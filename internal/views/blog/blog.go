package blog

import (
	"apparently-typing/static"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"path"
	"strconv"
	"strings"

	"time"

	"github.com/a-h/templ"
	"github.com/starfederation/datastar-go/datastar"
	"github.com/yuin/goldmark"
)

const (
	YYYYMMDD = "20060102"
)

type BlogPost struct {
	ID      int
	Date    time.Time
	Title   string
	Content templ.Component
}
type Handler struct{}

func NewHandler() http.Handler {
	return &Handler{}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

		switch r.PathValue("id") {
		// continous reading for all the posts.
		case "all":
			posts, err := static.BlogFiles.ReadDir("blog")
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
			index := len(posts) - 1
			post, _, _ := getpost(index)
			// If it can scroll, add the element that will allow autoscroll, otherwise, standard element
			if post.ID > 0 {
				templ.Handler(Post(post, PostScroll(post.ID-1))).ServeHTTP(w, r)
			} else {
				templ.Handler(Post(post, Nav(post.ID+1, -1))).ServeHTTP(w, r)
			}
		// Shows only the latest post
		case "latest":
			posts, err := static.BlogFiles.ReadDir("blog")
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
			index := len(posts) - 1
			showpost(index, w, r)
		case "":
			showindex(w, r)
		default:
			index, err := strconv.Atoi(r.PathValue("id"))
			// If failed, we just return index with a redirect
			if err != nil {
				http.Redirect(w, r, "/blog/", http.StatusMovedPermanently)
				return
			}
			if r.URL.Query().Has("continuous") {
				scroll(w, r)
			} else {
				showpost(index, w, r)
			}
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func scroll(w http.ResponseWriter, r *http.Request) {
	sse := datastar.NewSSE(w, r)
	index, err := strconv.Atoi(r.PathValue("id"))
	// If failed, we'll just let the connection close without updating anything.
	if err != nil {
		_ = sse.ConsoleError(fmt.Errorf("internal server error. path value not valid integer: %w", err))
		slog.Error("Datastar-Request emitted", "error", err)
		return
	}

	post, _, err := getpost(index)
	if err != nil {
		_ = sse.ConsoleError(fmt.Errorf("fetch post error: %w", err))
		slog.Error("Internal Server Error", "error", err)
		return
	}

	err = sse.PatchElementTempl(PostFrag(post), datastar.WithSelectorID("post_nav"), datastar.WithModeBefore())
	if err != nil {
		_ = sse.ConsoleError(fmt.Errorf("something went wrong with patching %w", err))
		slog.Error("Internal Server Error", "error", err)
		return
	}

	if index > 0 {
		err = sse.PatchElementTempl(PostScroll(index-1), datastar.WithSelectorID("post_nav"))
		if err != nil {
			_ = sse.ConsoleError(fmt.Errorf("something went wrong with patching %w", err))
			slog.Error("Internal Server Error", "error", err)
		}

	}
}

func showindex(w http.ResponseWriter, r *http.Request) {
	posts, err := static.BlogFiles.ReadDir("blog")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	titles := make([]BlogPost, 0, len(posts))
	for i, post := range posts {
		datetitle := strings.TrimSuffix(post.Name(), ".md")
		date, title, found := strings.Cut(datetitle, "-")
		if found {
			t, err := time.Parse(YYYYMMDD, date)
			if err != nil {
				slog.Error("List Posts", "error", err)
				continue
			}
			titles = append(titles, BlogPost{ID: i, Title: title, Date: t})
		} else {
			slog.Error("Blogpost name malformed", "title", datetitle)
		}
	}
	templ.Handler(Index("Blog Posts", titles)).ServeHTTP(w, r)
}

func unsafeRenderMarkdown(html string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		_, err = io.WriteString(w, html)
		return
	})
}

func showpost(index int, w http.ResponseWriter, r *http.Request) {
	post, total, err := getpost(index)
	if errors.Is(err, ErrPostNotFound) {
		http.Redirect(w, r, "/blog/", http.StatusMovedPermanently)
		return
	} else if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	// Conditionally allow for a naviation based on whether there are additional available
	next, prev := -1, -1
	if post.ID+1 < total {
		next = post.ID + 1
	}

	if post.ID-1 > 0 {
		prev = post.ID - 1
	}
	templ.Handler(Post(post, Nav(next, prev))).ServeHTTP(w, r)
}

var ErrPostNotFound = errors.New("blog post not found")

func getpost(index int) (*BlogPost, int, error) {
	posts, err := static.BlogFiles.ReadDir("blog")

	if err != nil {
		slog.Error("Read Blog Directory Error: ", "error", err)
		return nil, 0, fmt.Errorf("read blog directory error %w", err)
	}

	totalPosts := len(posts)
	if index >= totalPosts || index < 0 {
		return nil, 0, ErrPostNotFound
	}

	filename := posts[index].Name()
	filecontent, err := static.BlogFiles.ReadFile(path.Join("blog", filename))
	if err != nil {
		slog.Error("Read Files Content Error: ", "error", err)
		return nil, 0, fmt.Errorf("read files content error %w", err)
	}

	datetitle := strings.TrimSuffix(filename, ".md")
	filedate, filetitle, found := strings.Cut(datetitle, "-")
	if !found {
		slog.Error("Parsing filename error", "error", err)
		return nil, 0, fmt.Errorf("parsing filename error: %w", err)
	}
	t, err := time.Parse(YYYYMMDD, filedate)

	if err != nil {
		slog.Error("Time conversation error", "error", err)
		return nil, 0, fmt.Errorf("time conversion error %w", err)
	}

	var buf bytes.Buffer
	if err := goldmark.Convert(filecontent, &buf); err != nil {
		slog.Error("Goldmark parsing error", "error", err)
		return nil, 0, fmt.Errorf("goldmark parsing error %w", err)
	}
	htmlcontent := unsafeRenderMarkdown(buf.String())
	var post = &BlogPost{
		ID:      index,
		Date:    t,
		Title:   filetitle,
		Content: htmlcontent,
	}
	return post, totalPosts, nil
}
