package blog

import (
	"net/http"

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
		// Check
		var post = BlogPost{
			Date:    time.Now(),
			Title:   "This is a title",
			Content: "This is the markdown content",
		}
		templ.Handler(Post(post)).ServeHTTP(w, r)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
