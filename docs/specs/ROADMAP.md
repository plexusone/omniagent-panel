# OmniAgent Panel Roadmap

## v0.1.0 (Current Release Target)

### High Priority

- [x] **Dockerfile** - Container image build for K8s deployment
  - Multi-stage build with opus support
  - Minimal runtime image
  - Health check command

- [x] **Health Endpoints** - HTTP endpoints for K8s probes
  - `GET /health` - Liveness probe
  - `GET /ready` - Readiness probe
  - Health server runs alongside panel

- [x] **README.md** - Project documentation
  - Features overview
  - Installation instructions
  - Environment variables reference
  - Usage examples for human and auto modes
  - Avatar configuration guide

- [x] **CHANGELOG** - Release history
  - CHANGELOG.json in structured-changelog format
  - CHANGELOG.md generated via schangelog

### Medium Priority

- [ ] **Tests** - Unit and integration tests
  - coordinator_test.go
  - panelist_test.go
  - transcript_test.go
  - Integration tests with mock LiveKit

- [ ] **Additional Avatar Providers** - Beyond HeyGen
  - Tavus integration via omniavatar
  - bitHuman integration via omniavatar
  - Provider selection via AVATAR_PROVIDER env var

### Low Priority

- [ ] **MkDocs Documentation** - Full documentation site
  - Getting started guide
  - Configuration reference
  - Architecture overview
  - Avatar provider comparison

- [ ] **Observability** - Metrics and tracing
  - Prometheus metrics endpoint
  - OpenTelemetry tracing
  - Structured logging with slog

## v0.2.0 (Future)

- [ ] **Web UI** - Browser-based control panel
- [ ] **Persistent Transcripts** - Save discussions to database
- [ ] **Custom Personas** - YAML-based persona definitions
- [ ] **Multi-room Support** - Run multiple panels simultaneously
