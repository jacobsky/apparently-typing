package apphealth

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"sync"
	"time"

	"github.com/starfederation/datastar-go/datastar"
)

const (
	JPASSISTANT = "jpassistant"
	HIRE        = "hire"
	CHECKPOINT  = "checkpoint"
)

var AppURLs = map[string]string{
	JPASSISTANT: "https://joushu.apparently-typ.ing/",
	HIRE:        "https://hire.apparently-typ.ing/",
	CHECKPOINT:  "https://checkpoint.apparently-typ.ing/",
}

var appHealthEndpoints = map[string]string{
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
	if r.Header.Get("Datastar-Request") == "" {
		slog.Warn("Endpoint queried externally")
		http.Error(w, "Request not allowed", http.StatusBadRequest)
		return
	}

	slog.Info("Received health check query")
	appname := r.PathValue("id")
	sse := datastar.NewSSE(w, r)
	if appname == "" {
		slog.Error("appname is empty")
		_ = sse.ConsoleError(fmt.Errorf("Bad request Appname is empty"))
		return
	}
	err := sse.PatchElementTempl(Indicator(appname))
	if err != nil {
		slog.Error("datastar error", "error", err)
	}
}

func checkapp(endpoint string) bool {
	resp, err := http.Get(appHealthEndpoints[endpoint])
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
			for key := range appHealthEndpoints {
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
