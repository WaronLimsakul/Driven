staticcheck:
	staticcheck ./...

test:
	go test -race -v timout 30s ./...

# monitor for dev
tw-watch:
	tailwind -i ./static/css/input.css -o ./static/css/style.css --watch

# build for production
tw-build:
	tailwind -i ./static/css/input.css -o ./static/css/style.min.css --minify

templ-watch:
	templ generate --watch

templ-generate:
	templ generate

dev:
	go build -o ./tmp/main ./cmd/main.go && air

full-dev:
	tmux kill-session -t driven-dev || true
	tmux new-session -d -s driven-dev
	tmux split-window -h -t driven-dev
	tmux split-window -v -t driven-dev:0.0
	tmux send-keys -t driven-dev:0.0 'make tw-watch' C-m
	tmux send-keys -t driven-dev:0.1 'make templ-watch' C-m
	tmux send-keys -t driven-dev:0.2 'make dev' C-m
	tmux attach-session -t driven-dev

build: # build for deploy
	make tw-build
	make templ-generate
	go build -ldflags "-X main.Environment=production" -o ./bin/$(APP_NAME) ./cmd/main.go
	# check the last line later
