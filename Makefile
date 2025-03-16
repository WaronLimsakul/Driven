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

build: # build for deploy
	make tw-build
	make templ-generate
	go build -ldflags "-X main.Environment=production" -o ./bin/$(APP_NAME) ./cmd/main.go
	# check the last line later
