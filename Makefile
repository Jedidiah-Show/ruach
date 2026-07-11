APP_NAME=ruach

.PHONY: build
build:
	go build -o bin/$(APP_NAME) ./apps/cli

.PHONY: run
run:
	go run ./apps/cli

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: test
test:
	go test ./...

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: clean
clean:
	rm -rf bin