.PHONY: all deps templ css css-watch build run dev

SHELL      := /bin/bash
export PATH := $(HOME)/go/bin:$(PATH)

APP_NAME   := MeGrugMeLikeSimple
GO_FILES   := ./cmd/server/...
GO_FLAGS   := -ldflags="-s -w"

all: build

# ─── Dependencies ──────────────────────────────────────────
deps:
	go mod tidy
	npm install

# ─── Templ ──────────────────────────────────────────────────
templ:
	templ generate

templ-watch:
	templ generate --watch

# ─── Tailwind CSS ──────────────────────────────────────────
css:
	npx @tailwindcss/cli -i ./static/css/input.css -o ./static/css/output.css --minify

css-watch:
	npx @tailwindcss/cli -i ./static/css/input.css -o ./static/css/output.css --watch

# ─── Build ──────────────────────────────────────────────────
build: templ css
	go build $(GO_FLAGS) -o bin/$(APP_NAME) $(GO_FILES)

# ─── Run ────────────────────────────────────────────────────
run: templ
	go run $(GO_FILES)

# ─── Dev (run in separate terminals or use tmux) ──────────
dev:
	@echo "Run these in separate terminals:"
	@echo "  make templ-watch    # Templ hot-reload"
	@echo "  make css-watch      # Tailwind watcher"
	@echo "  make run            # Go server on :8080"

# ─── Clean ──────────────────────────────────────────────────
clean:
	rm -rf bin/ static/css/output.css
