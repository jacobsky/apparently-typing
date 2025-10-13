package blog

import (
	"apparently-typing/static"
	"log/slog"
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
		titles[i] = post.Name()
	}
	templ.Handler(List("Blog Posts", titles)).ServeHTTP(w, r)
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
	year, err := strconv.Atoi(filedate[0:4])
	month, err2 := strconv.Atoi(filedate[5:6])
	day, err3 := strconv.Atoi(filedate[7:8])
	if err != nil || err2 != nil || err3 != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		slog.Error("Time converstaion error", "error", err, "error", err2, "error", err3)
		return
	}

	var post = BlogPost{
		Date:    time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC),
		Title:   filetitle,
		Content: string(filecontent),
	}
	templ.Handler(Post(post)).ServeHTTP(w, r)
}
