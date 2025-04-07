.PHONY: dev
dev:
		go build -o ./tmp/main ./cmd/app/main.go && air

.PHONY: build
build:
		go build -ldflags "-X main.Environment=production" -o ./bin/bike-guide-api ./cmd/app/main.go

.PHONY: vet
vet:
		go vet ./...

.PHONY: staticcheck
staticcheck:
		staticcheck ./...

.PHONY: test
test:
		go test -race -v -timeout 30s ./..
