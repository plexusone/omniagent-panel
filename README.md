# OmniAgent Panel

[![Go CI][go-ci-svg]][go-ci-url]
[![Go Lint][go-lint-svg]][go-lint-url]
[![Go SAST][go-sast-svg]][go-sast-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![Visualization][viz-svg]][viz-url]
[![License][license-svg]][license-url]

 [go-ci-svg]: https://github.com/plexusone/omniagent-panel/actions/workflows/go-ci.yaml/badge.svg?branch=main
 [go-ci-url]: https://github.com/plexusone/omniagent-panel/actions/workflows/go-ci.yaml
 [go-lint-svg]: https://github.com/plexusone/omniagent-panel/actions/workflows/go-lint.yaml/badge.svg?branch=main
 [go-lint-url]: https://github.com/plexusone/omniagent-panel/actions/workflows/go-lint.yaml
 [go-sast-svg]: https://github.com/plexusone/omniagent-panel/actions/workflows/go-sast-codeql.yaml/badge.svg?branch=main
 [go-sast-url]: https://github.com/plexusone/omniagent-panel/actions/workflows/go-sast-codeql.yaml
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/plexusone/omniagent-panel
 [docs-godoc-url]: https://pkg.go.dev/github.com/plexusone/omniagent-panel
 [docs-mkdoc-svg]: https://img.shields.io/badge/Go-dev%20guide-blue.svg
 [docs-mkdoc-url]: https://plexusone.dev/omniagent-panel
 [viz-svg]: https://img.shields.io/badge/Go-visualizaton-blue.svg
 [viz-url]: https://mango-dune-07a8b7110.1.azurestaticapps.net/?repo=plexusone%2Fomniagent-panel
 [loc-svg]: https://tokei.rs/b1/github/plexusone/omniagent-panel
 [repo-url]: https://github.com/plexusone/omniagent-panel
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/plexusone/omniagent-panel/blob/main/LICENSE

Multi-agent panel discussion system powered by LiveKit, LLMs, and optional AI avatars.

## Features

- 🎙️ **Panel Discussions** - Multiple AI panelists engage in dynamic discussions on any topic
- 🔄 **Dual Modes** - Human moderator mode or fully automated AI moderator mode
- 🔊 **Real-time Audio** - LiveKit integration for low-latency audio streaming
- 🎭 **AI Avatars** - Optional HeyGen avatar support for visual representation
- 🧠 **Flexible LLMs** - Support for Anthropic Claude and OpenAI models
- 🗣️ **Multiple TTS Providers** - OpenAI, ElevenLabs, and more via omnivoice
- ☸️ **Kubernetes Ready** - Dockerfile with health endpoints for K8s deployment

## Installation

```bash
go install -tags opus github.com/plexusone/omniagent-panel/cmd/panel@latest
```

Or build from source:

```bash
git clone https://github.com/plexusone/omniagent-panel.git
cd omniagent-panel
go build -tags opus -o panel ./cmd/panel
```

## Quick Start

### Environment Variables

```bash
# Required: LiveKit configuration
export LIVEKIT_URL="wss://your-project.livekit.cloud"
export LIVEKIT_API_KEY="your-api-key"
export LIVEKIT_API_SECRET="your-api-secret"

# Required: LLM configuration
export LLM_PROVIDER="anthropic"  # or "openai"
export LLM_MODEL="claude-sonnet-4-20250514"
export ANTHROPIC_API_KEY="your-key"

# Required: TTS configuration
export TTS_PROVIDER="openai"
export OPENAI_API_KEY="your-key"

# Panel configuration
export PANEL_TOPIC="The future of AI agents"
export PANEL_SIZE=3  # 2-4 panelists
export PANEL_MODE=auto  # "human" or "auto"
```

### Run the Panel

```bash
# Auto mode (AI moderator)
export PANEL_MODE=auto
./panel

# Human mode (you moderate via LiveKit Meet)
export PANEL_MODE=human
export STT_PROVIDER="deepgram"
export DEEPGRAM_API_KEY="your-key"
./panel
```

## Configuration

### Panel Settings

| Variable | Description | Default |
|----------|-------------|---------|
| `PANEL_TOPIC` | Discussion topic | "The future of artificial intelligence" |
| `PANEL_SIZE` | Number of panelists (2-4) | 3 |
| `PANEL_MODE` | "human" or "auto" | "human" |
| `PANEL_QUESTIONS` | Comma-separated questions (auto mode) | Generated |
| `PANEL_ROUNDS` | Number of discussion rounds (auto mode) | 5 |

### Moderator Configuration (Auto Mode)

| Variable | Description | Default |
|----------|-------------|---------|
| `MODERATOR_NAME` | Moderator name | "Sam" |
| `MODERATOR_VOICE` | TTS voice | "shimmer" |
| `MODERATOR_PERSONALITY` | Moderator personality | "Engaging moderator..." |
| `MODERATOR_AVATAR_ID` | HeyGen avatar ID | None |

### Panelist Configuration

Override default panelists with indexed environment variables:

| Variable | Description |
|----------|-------------|
| `PANELIST_1_NAME` | First panelist name |
| `PANELIST_1_VOICE` | First panelist TTS voice |
| `PANELIST_1_PERSONALITY` | First panelist personality |
| `PANELIST_1_AVATAR_ID` | First panelist HeyGen avatar ID |

Repeat for `PANELIST_2_*`, `PANELIST_3_*`, etc.

### Avatar Configuration (Optional)

| Variable | Description | Default |
|----------|-------------|---------|
| `HEYGEN_API_KEY` | HeyGen API key | None |
| `HEYGEN_SANDBOX` | Use sandbox mode | "false" |
| `HEYGEN_VIDEO_QUALITY` | Video quality | "high" |

### Recording Configuration (Optional)

| Variable | Description | Default |
|----------|-------------|---------|
| `PANEL_RECORD` | Enable recording | "false" |
| `PANEL_RECORD_FORMAT` | Recording format | "mp4" |
| `PANEL_RECORD_LAYOUT` | Recording layout | "speaker" |
| `PANEL_RECORD_PATH` | Local file path | None |
| `PANEL_RECORD_S3_BUCKET` | S3 bucket | None |
| `PANEL_RECORD_S3_REGION` | S3 region | None |

## Health Endpoints

The panel exposes health endpoints for Kubernetes probes:

| Endpoint | Description |
|----------|-------------|
| `GET /health` | Liveness probe - returns 200 if service is alive |
| `GET /ready` | Readiness probe - returns 200 if ready to receive traffic |

Configure the health port with `--health-port` flag (default: 8080).

## Docker

Build the container:

```bash
docker build -t omniagent-panel .
```

Run with environment variables:

```bash
docker run -p 8080:8080 \
  -e LIVEKIT_URL="wss://..." \
  -e LIVEKIT_API_KEY="..." \
  -e LIVEKIT_API_SECRET="..." \
  -e ANTHROPIC_API_KEY="..." \
  -e OPENAI_API_KEY="..." \
  -e PANEL_MODE=auto \
  omniagent-panel
```

## Kubernetes Deployment

Deploy using agentkit with the Kubernetes provider:

```bash
export AGENTKIT_DEPLOY_PROVIDER=kubernetes
export KUBECONFIG=~/.kube/config
go run ./deploy/cmd
```

See [deploy/deploy.yaml](deploy/deploy.yaml) for configuration options.

## Architecture

```
omniagent-panel/
├── cmd/panel/          # Main application
│   ├── main.go         # Entry point, mode selection
│   ├── coordinator.go  # Turn-taking orchestration
│   ├── panelist.go     # Panelist agents
│   ├── moderator.go    # AI moderator (auto mode)
│   ├── health.go       # Health endpoints
│   └── ...
├── panel/              # Core panel types
├── deploy/             # Kubernetes deployment
│   ├── deploy.yaml     # Deployment configuration
│   └── cmd/main.go     # Deploy CLI
└── Dockerfile          # Container build
```

## Dependencies

- [omni-livekit](https://github.com/plexusone/omni-livekit) - LiveKit agent framework
- [omnivoice](https://github.com/plexusone/omnivoice) - Multi-provider TTS/STT
- [agentkit](https://github.com/plexusone/agentkit) - Deployment framework
- [agentkit-k8s-pulumi](https://github.com/plexusone/agentkit-k8s-pulumi) - Kubernetes provider

## License

MIT License - see [LICENSE](LICENSE) for details.
