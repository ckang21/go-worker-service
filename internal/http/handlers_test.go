package http

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ckang21/go-worker-service/internal/queue"
	"github.com/ckang21/go-worker-service/internal/tasks"
)

func TestHealth(t *testing.T) {
	store := tasks.NewStore()
	q := queue.New[tasks.Job](8)
	mux := http.NewServeMux()
	RegisterRoutes(mux, store, q)

	req := httptest.NewRequest("GET", "/health", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != 200 || rec.Body.String() != "ok" {
		t.Fatalf("health failed: code=%d body=%q", rec.Code, rec.Body.String())
	}
}

func TestCreateTask(t *testing.T) {
	store := tasks.NewStore()
	q := queue.New[tasks.Job](8)
	mux := http.NewServeMux()
	RegisterRoutes(mux, store, q)

	body := bytes.NewBufferString(`{"title":"demo"}`)
	req := httptest.NewRequest("POST", "/tasks", body)
	req.Header.Set("content-type", "application/json")
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusAccepted {
		t.Fatalf("expected 202; got %d", rec.Code)
	}
	if len(store.List()) != 1 {
		t.Fatalf("task not stored")
	}
}
