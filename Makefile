GO_MODULE := $(shell awk '/module/{print $$2; exit}' go.mod)
GO_BUILDMETA = github.com/avakarev/dotfiles-cli/internal/buildmeta

lint:
	@echo ">> Running revive..."
	@revive -config .revive.toml -formatter friendly ./...
	@echo ">> Running staticcheck..."
	@staticcheck ./...

vet:
	@echo ">> Vetting..."
	@go vet ./...

sec:
	@echo ">> Auditing..."
	@gosec -quiet -tests ./...

test:
	@echo ">> Running tests..."
	@go test -v -race ./...
.PHONY: test

setup-ci:
	@go install github.com/mgechev/revive@latest
	@go install github.com/securego/gosec/v2/cmd/gosec@latest
	@go install honnef.co/go/tools/cmd/staticcheck@latest

ci: lint vet sec test

build:
	@echo ">> Building ./bin/dotfiles..."
	@CGO_ENABLED=0 go build -o ./bin/dotfiles ./cmd
	@echo "   Done!"

release:
	BUILDMETA=${GO_BUILDMETA} goreleaser release --clean

release-dryrun:
	BUILDMETA=${GO_BUILDMETA} goreleaser release --snapshot --skip-publish --skip-sign --clean

release-build:
	BUILDMETA=${GO_BUILDMETA} goreleaser build --snapshot --single-target --clean
