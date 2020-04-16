.PHONY: build

fmt:
	@gofmt -l -w -s .
	@echo "[OK] Ran gofmt"

lint:
	@golint ./...
	@echo "[OK] Ran golint"

tidy:
	@go mod tidy -v
	@echo "[OK] Updated go.mod"

build-static:
	@pkger -o ./internal/dictionary
	@echo "[OK] Updated static files"

security:
	@gosec -quiet ./...
	@echo "[OK] Ran gosec"

vet:
	@go vet
	@echo "[OK] Ran go vet"

test:
	@go test -vet=off ./...
	@echo "[OK] Tests passed"

validate: lint vet tidy security test

build-full: validate fmt build-static build

build:
	@go build
	@echo "[OK] Built binary"

run: build
	@./text_linter
