package blog

import (
	"apparently-typing/static"
	"net/http"
	"path"
	"strconv"
	"strings"

	"time"

	"github.com/a-h/templ"
)

type BlogPost struct {
	Date    time.Time
	Title   string
	Content string
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
		}

		showpost(post_id, w, r)
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
		titles[i] = post.Name()
	}
	templ.Handler(List("Blog Posts", titles)).ServeHTTP(w, r)
}

func showpost(id string, w http.ResponseWriter, r *http.Request) {
	index, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Malformed post ID", http.StatusBadRequest)
		return
	}
	posts, err := static.BlogFiles.ReadDir("blog")

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	if index > len(posts) || index < 0 {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}
	filename := posts[index].Name()
	filecontent, err := static.BlogFiles.ReadFile(path.Join("blog", filename))
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	datetitle := strings.TrimSuffix(filename, ".md")
	filedate, filetitle, found := strings.Cut(datetitle, "-")
	if !found {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	date, err := time.Parse("YYYYMMDD", filedate)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	var post = BlogPost{
		Date:    date,
		Title:   filetitle,
		Content: string(filecontent),
	}
	templ.Handler(Post(post)).ServeHTTP(w, r)
}
