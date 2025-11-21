package apps

import (
	"log/slog"
	"net/http"
	"sync"
	"time"

	"github.com/a-h/templ"
)

const (
	jpassistant = "jpassistant"
	hire        = "hire"
	checkpoint  = "checkpoint"
)

var apps = map[string]string{
	jpassistant: "https://joushu.apparently-typ.ing/healthcheck",
	hire:        "https://hire.apparently-typ.ing/healthcheck",
	checkpoint:  "https://checkpoint.apparently-typ.ing/healthcheck",
}

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

var AppHealth = newHealthMap(
	map[string]bool{
		jpassistant: false,
		hire:        false,
		checkpoint:  false,
	},
)

func NewHandler() http.Handler {
	h := &Handler{}
	go h.serve()
	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Parameter is passed to check specific endpoint's state
		appname := r.URL.Query().Get("app")
		templ.Handler(Indicator(appname, AppHealth.Read(appname))).ServeHTTP(w, r)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func checkapp(endpoint string) bool {
	resp, err := http.Get(apps[endpoint])
	if err != nil {
		return false
	}
	if resp.StatusCode == http.StatusOK {
		return true
	}
	return false
}

func (h *Handler) serve() {
	slog.Info("Health Check Service Started")
	var wg sync.WaitGroup
	for {
		slog.Debug("Health Heartbeat Check")
		for key := range apps {
			wg.Go(
				func() {
					health := checkapp(key)
					AppHealth.Write(key, health)
					slog.Debug("App Health", key, health)
				})
		}
		wg.Wait()
		time.Sleep(60 * time.Second)
	}
}
