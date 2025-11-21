package apps

import (
	"context"
	"log/slog"
	"net/http"
	"sync"
	"time"

	"github.com/a-h/templ"
)

const (
	JPASSISTANT = "jpassistant"
	HIRE        = "hire"
	CHECKPOINT  = "checkpoint"
)

var apps = map[string]string{
	JPASSISTANT: "https://joushu.apparently-typ.ing/healthcheck",
	HIRE:        "https://hire.apparently-typ.ing/healthcheck",
	CHECKPOINT:  "https://checkpoint.apparently-typ.ing/healthcheck",
}

var AppHealth = newHealthMap(
	map[string]bool{
		JPASSISTANT: false,
		HIRE:        false,
		CHECKPOINT:  false,
	},
)

// Simple concurrent lock to ensure that the health status endpoints are
// read safe.
type healthMap struct {
	rw   sync.RWMutex
	data map[string]bool
}

func newHealthMap(data map[string]bool) *healthMap {
	return &healthMap{
		data: data,
	}
}

func (hm *healthMap) Read(key string) bool {
	hm.rw.RLock()
	defer hm.rw.RUnlock()
	return hm.data[key]
}

func (hm *healthMap) Write(key string, val bool) {
	hm.rw.Lock()
	defer hm.rw.Unlock()
	hm.data[key] = val
}

type Handler struct{}

func NewHandler() http.Handler {
	h := &Handler{}
	go h.serve(context.Background())
	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return

	}
	// Parameter is passed to check specific endpoint's state
	appname := r.URL.Query().Get("app")
	templ.Handler(Indicator(appname, AppHealth.Read(appname))).ServeHTTP(w, r)

}

func checkapp(endpoint string) bool {
	resp, err := http.Get(apps[endpoint])
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}

func (h *Handler) serve(ctx context.Context) {
	ticker := time.NewTicker(60 * time.Second)
	slog.Info("Health Check Service Started")
	var wg sync.WaitGroup
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			slog.Debug("Health Heartbeat Check")
			for key := range apps {
				key := key
				wg.Go(
					func() {
						health := checkapp(key)
						AppHealth.Write(key, health)
						slog.Debug("App Health", "app", key, "healthy", health)
					})
			}
			wg.Wait()
		}
	}
}
