# go-worker-service

Tiny Go service with one HTTP API and a background worker.

## Quickstart
```bash
make run
# in another terminal
curl -s localhost:8080/health
curl -s -X POST localhost:8080/tasks -d '{"title":"demo"}' -H 'content-type: application/json' | jq
curl -s localhost:8080/tasks | jq
```

## What it shows
- HTTP server with `/health`, `/tasks` (GET/POST)
- Background worker processing a queue (channel-based)
- Unit tests, Dockerfile, CI (Go 1.22, race detector, coverage artifact)
