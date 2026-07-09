# Café de Microservices

A café ordering app built with three microservices and a Vue frontend.

## Architecture

```
menu-service (Go/Gin + SQLite) ──► :8080
ai-service (Python/FastAPI + Mistral) ──► :8000
frontend (Vue 3 + daisyUI + Tailwind CSS) ──► :5173
```

- **Menu Service** — CRUD API for menu items (Go, Gin, GORM, SQLite)
- **AI Service** — Chat endpoint that fetches the live menu from the menu service and uses it as context for Mistral AI responses (Python, FastAPI)
- **Frontend** — Vue 3 SPA with daisyUI components and Tailwind CSS

## Prerequisites

- [Go 1.22+](https://go.dev/dl/)
- [Python 3.13+](https://www.python.org/downloads/) + [uv](https://docs.astral.sh/uv/)
- [Bun](https://bun.sh/) (or Node 22+)
- A [Mistral API key](https://console.mistral.ai/)

## Run

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
# → running on :8000
```

Set `MENU_SERVICE_URL` if the menu service isn't on localhost:8080.

### Frontend

```sh
cd frontend
bun install
bun dev
# → http://localhost:5173
```
