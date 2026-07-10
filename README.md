# The Good Café

A café ordering app built with two microservices and a Vue frontend.

## Architecture

```
menu-service (Go + SQLite) ──── :8080
ai-service (Python/FastAPI + Mistral) ──── :8001
frontend (Vue 3 + daisyUI + Tailwind CSS) ──── :80
```

- **Menu Service** — CRUD API for menu items (Go, Gin, GORM, SQLite)
- **AI Service** — Chat endpoint that fetches the live menu and uses it as context for Mistral AI responses (Python, FastAPI)
- **Frontend** — Vue 3 SPA with daisyUI components and Tailwind CSS

## Quick Start (Docker)

```sh
cp .env.example .env
# edit .env with your Mistral API key
docker compose up --build
# → http://localhost
```

## Local Dev

### Prerequisites

- [Go 1.22+](https://go.dev/dl/)
- [Python 3.13+](https://www.python.org/downloads/) + [uv](https://docs.astral.sh/uv/)
- [Bun](https://bun.sh/)
- A [Mistral API key](https://console.mistral.ai/)

### Menu Service

```sh
cd menu-service
go run main.go
# → running on :8080
```

### AI Service

```sh
cd ai-service
export MISTRAL_API_KEY=your-key-here
uv run fastapi dev main.py
# → running on :8001
```

### Frontend

```sh
cd frontend
bun install
bun dev
# → http://localhost:5173
```
