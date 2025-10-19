package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ckang21/go-worker-service/internal/queue"
	"github.com/ckang21/go-worker-service/internal/tasks"
)

type createTaskReq struct {
	Title string `json:"title"`
}

func RegisterRoutes(mux *http.ServeMux, store *tasks.Store, q *queue.Queue[tasks.Job]) {
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	mux.HandleFunc("GET /tasks", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, store.List())
	})

	mux.HandleFunc("POST /tasks", func(w http.ResponseWriter, r *http.Request) {
		var req createTaskReq
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Title == "" {
			http.Error(w, "invalid body", http.StatusBadRequest)
			return
		}

		task := tasks.Task{
			ID:        tasks.NewID(),
			Title:     req.Title,
			Status:    tasks.StatusQueued,
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		}
		store.Create(task)

		// Enqueue work
		q.Push(tasks.Job{TaskID: task.ID})

		writeJSON(w, http.StatusAccepted, task)
	})
}

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}
