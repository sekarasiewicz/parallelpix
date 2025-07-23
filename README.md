# ParallelPix - Concurrent Image-Processing Microservice

This document serves as both the Product Requirements (PRD) and the `README.md` for the ParallelPix service: a Go-based, Fiber-driven microservice for asynchronous image uploads, processing, and retrieval.

---

## 1. Overview

**ParallelPix** enables clients to upload images via REST, have them processed (resize, filter, watermark) in the background, and query status/results. It balances high concurrency, resilience, and observability.

## 2. Goals

* **Reliability**: Safe processing under load, graceful shutdowns.
* **Scalability**: Handle thousands of jobs concurrently via worker pool.
* **Extensibility**: Plug-in storage drivers, filters, and middleware.
* **Observability**: Metrics (Prometheus), tracing, and structured logs.

## 3. Key Features

1. **REST API** for image upload, status, and download.
2. **Background Worker Pool** using goroutines, channels, and context.
3. **Storage Abstraction**: Local filesystem & S3/MinIO via `Storage` interface.
4. **Config & Middleware** via Fiber (logging, CORS, JWT auth).
5. **Database**: PostgreSQL + GORM migrations for `ImageJob` schema.
6. **Metrics & Profiling**: `/metrics` (Prometheus) & `/debug/pprof`.
7. **Testing**: Unit & integration tests with mocks.

## 4. User Stories

* **As a client**, I can upload an image and immediately receive a job ID.
* **As a client**, I can poll job status until processing completes.
* **As a client**, I can download the processed image once ready.
* **As an operator**, I can monitor service health and performance metrics.

## 5. Architecture

```
[Client] --> [Fiber REST API] --> [DB: ImageJob]  \
                                      → [Channel Queue] → [Worker Pool] → [Storage]
                                            ↘                           ↘
                                             → [Redis Pub/Sub] (optional)
```

* **API Layer:** Fiber app with middleware chain.
* **DB Layer:** GORM-managed Postgres for job records.
* **Queue & Workers:** Channel-driven, N goroutines processing jobs.
* **Storage:** Pluggable via interface; Local & S3/MinIO implementations.

## 6. API Endpoints

| Method | Path               | Description                        |
| ------ | ------------------ | ---------------------------------- |
| POST   | `/v1/upload`       | Upload image, returns `{job_id}`   |
| GET    | `/v1/status/:id`   | Get job status & metadata          |
| GET    | `/v1/download/:id` | Download processed image (if done) |
| GET    | `/metrics`         | Prometheus metrics                 |
| GET    | `/debug/pprof/...` | pprof endpoints for profiling      |

### Request / Response Examples

**Upload**
`POST /v1/upload` (multipart/form-data, field `file`)
**Response**: `200 OK` `{ "job_id": "abc123" }`

**Status**
`GET /v1/status/abc123`
**Response**: `200 OK` `{ "job_id": "abc123", "status": "processing", "submitted_at": "..." }`

**Download**
`GET /v1/download/abc123`
**Response**: `200 OK` (image/png)

## 7. Data Model

```go
type ImageJob struct {
    ID         uuid.UUID  `gorm:"type:uuid;primaryKey"`
    Status     string     `gorm:"index"` // queued|processing|done|error
    InputPath  string
    OutputPath string
    CreatedAt  time.Time
    UpdatedAt  time.Time
}
```

## 8. Tech Stack

* **Language**: Go 1.21+
* **Web**: Fiber 2.x
* **ORM**: GORM + PostgreSQL
* **Storage**: AWS SDK / MinIO
* **Queue**: Go channels & goroutines
* **Config**: Viper or env
* **Logging**: Zap
* **Metrics**: Prometheus / OpenTelemetry
* **Testing**: `testing`, `httptest`, `testify/mock`

## 9. Workflow

1. Client uploads → API writes `ImageJob{Status: queued}` → sends job to channel.
2. Worker picks job → updates status to `processing` → downloads input → processes image → saves via `Storage` → updates job to `done`.
3. Client polls status until `done`, then downloads via `/download`.

## 10. Next Steps

* Define image transformations and plugin interface.
* Implement authentication (JWT).
* Add CI pipeline with linting, tests, and security scans.

---

Created by **ParallelPix** team – ready to ship!
