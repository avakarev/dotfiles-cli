PKGS = $(shell go list ./... 2> /dev/null)
TEST_PKGS = $(shell go list ./internal/... ./pkg/... 2> /dev/null)

print-%: ; @echo $*=$($*)

tidy:
	@echo ">> Tidying..."
	@go mod tidy

fmt:
	@echo ">> Formatting..."
	@go fmt $(PKGS)

vet:
	@echo ">> Vetting..."
	@go vet ${PKGS}

lint:
	@echo ">> Linting..."
	@golint $(PKGS)

sec:
	@echo ">> Auditing..."
	@gosec -quiet ./...

test:
	@echo ">> Running tests..."
	@go test -v -race ${TEST_PKGS}
.PHONY: test

setup-ci:
	@go get -u golang.org/x/lint/golint
	@go get -u github.com/securego/gosec/cmd/gosec

ci: lint vet sec test

build:
	@echo ">> Building ./bin/dotfiles ..."
	@CGO_ENABLED=0 go build -o ./bin/dotfiles ./cmd
	@echo "   Done!"
