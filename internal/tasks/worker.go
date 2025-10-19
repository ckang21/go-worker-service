package tasks

import (
	"log"
	"time"

	"github.com/ckang21/go-worker-service/internal/queue"
)

type Worker struct {
	q      *queue.Queue[Job]
	store  *Store
	logger *log.Logger
	stop   chan struct{}
}

func NewWorker(q *queue.Queue[Job], store *Store, logger *log.Logger) *Worker {
	return &Worker{
		q:      q,
		store:  store,
		logger: logger,
		stop:   make(chan struct{}),
	}
}

func (w *Worker) Run() {
	for {
		select {
		case <-w.stop:
			w.logger.Println("worker stopped")
			return
		case job, ok := <-w.q.Chan():
			if !ok {
				w.logger.Println("queue closed; worker exit")
				return
			}
			w.process(job)
		}
	}
}

func (w *Worker) process(job Job) {
	w.store.UpdateStatus(job.TaskID, StatusRunning)
	// Simulate work
	time.Sleep(500 * time.Millisecond)
	w.store.UpdateStatus(job.TaskID, StatusDone)
	w.logger.Printf("processed task=%s", job.TaskID)
}

func (w *Worker) Stop() {
	close(w.stop)
}
