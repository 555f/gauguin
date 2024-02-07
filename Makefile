VERSION = snapshot

.PHONY: build

default: build

build:
	go mod download
	go build -o ./build/gg ./cmd/gg

install:
	go install ./cmd/gg

chglog:
	git-chglog -o CHANGELOG.md

check:
	go vet ./...
	go test -v ./...

css:
	npx tailwindcss -i ./internal/plugin/http/files/tailwind.css -o ./internal/plugin/http/files/style.css  -c ./internal/plugin/http/files/tailwind.config.js --watch

gen-examples: gen-examples-rest-service-echo gen-examples-rest-service-chi

gen-examples-rest-service-echo:
	 go run cmd/gg/main.go run  --config examples/rest-service-echo/gg.yaml

gen-examples-rest-service-chi:
	 go run cmd/gg/main.go run  --config examples/rest-service-chi/gg.yaml	 