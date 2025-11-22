package server

import (
	"apparently-typing/internal/views/apphealth"
	"apparently-typing/internal/views/apps"
	"apparently-typing/internal/views/blog"
	"apparently-typing/internal/views/home"
	"apparently-typing/internal/views/tech"
	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	// Register routes
	fileServer := http.FileServer(http.FS(Files))
	mux.Handle("/assets/", fileServer)

	apps := apps.NewHandler()
	apphealth := apphealth.NewHandler()
	blog := blog.NewHandler()
	home := home.NewHandler()
	tech := tech.NewHandler()
	mux.Handle("/", home)
	mux.Handle("/apps", apps)
	mux.Handle("/apps/{id}", apps)
	mux.Handle("/apphealth/{id}", apphealth)
	mux.Handle("/blog", blog)
	mux.Handle("/blog/{id}", blog)
	mux.Handle("/tech", tech)
	// Wrap the mux with CORS middleware
	return s.corsMiddleware(mux)
}

func (s *Server) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// w.Header().Set("Access-Control-Allow-Origin", "*") // Replace "*" with specific origins if needed
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "false") // Set to "true" if credentials are required

		// Handle preflight OPTIONS requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Proceed with the next handler
		next.ServeHTTP(w, r)
	})
}
