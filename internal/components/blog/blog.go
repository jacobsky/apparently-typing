package blog

import (
	"apparently-typing/static"
	"bytes"
	"context"
	"io"
	"log/slog"
	"net/http"
	"path"
	"strconv"
	"strings"

	"time"

	"github.com/a-h/templ"
	"github.com/yuin/goldmark"
)

const (
	YYYYMMDD = "20060102"
)

type BlogPost struct {
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
		post_id := r.PathValue("id")
		if post_id == "" {
			list(w, r)
			return
		}
		index, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "Malformed post ID", http.StatusBadRequest)
			return
		}
		showpost(index, w, r)
		// Check

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func list(w http.ResponseWriter, r *http.Request) {
	posts, err := static.BlogFiles.ReadDir("blog")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	titles := make([]string, len(posts))
	for i, post := range posts {
		title := strings.Split(post.Name(), "-")[1]
		titles[i] = strings.TrimSuffix(title, ".md")
	}
	templ.Handler(List("Blog Posts", titles)).ServeHTTP(w, r)
}

func unsafeRenderMarkdown(html string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		_, err = io.WriteString(w, html)
		return
	})
}

func showpost(index int, w http.ResponseWriter, r *http.Request) {
	posts, err := static.BlogFiles.ReadDir("blog")

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		slog.Error("Read Blog Files Error: ", "error", err)
		return
	}
	if index >= len(posts) || index < 0 {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}
	filename := posts[index].Name()
	filecontent, err := static.BlogFiles.ReadFile(path.Join("blog", filename))
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		slog.Error("Read Files Content Error: ", "error", err)
		return
	}

	datetitle := strings.TrimSuffix(filename, ".md")
	filedate, filetitle, found := strings.Cut(datetitle, "-")
	if !found {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		slog.Error("Parsing filename error", "error", err)
		return
	}
	t, err := time.Parse(YYYYMMDD, filedate)

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		slog.Error("Time converstaion error", "error", err)
		return
	}

	var buf bytes.Buffer
	if err := goldmark.Convert(filecontent, &buf); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		slog.Error("Markdown parsing error", "error", err)
	}
	htmlcontent := unsafeRenderMarkdown(buf.String())
	var post = BlogPost{
		Date:    t,
		Title:   filetitle,
		Content: htmlcontent,
	}
	templ.Handler(Post(post)).ServeHTTP(w, r)
}
