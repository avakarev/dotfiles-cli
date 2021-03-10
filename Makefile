PKGS = $(shell go list ./... 2> /dev/null)
TEST_PKGS = $(shell go list ./internal/... ./pkg/... 2> /dev/null)

GO_MODULE := $(shell awk '/module/{print $$2; exit}' go.mod)
GO_BUILDMETA = github.com/avakarev/dotfiles-cli/internal/buildmeta


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
	@revive -config .revive.toml -formatter friendly ./...

sec:
	@echo ">> Auditing..."
	@gosec -conf .gosec.json -quiet ./...

test:
	@echo ">> Running tests..."
	@go test -v -race ${TEST_PKGS}
.PHONY: test

setup-ci:
	@GO111MODULE=off go get -u github.com/myitcv/gobin
	@gobin github.com/mgechev/revive
	@gobin github.com/securego/gosec/v2/cmd/gosec

ci: lint vet sec test

build:
	@echo ">> Building ./bin/dotfiles ..."
	@CGO_ENABLED=0 go build -o ./bin/dotfiles ./cmd
	@echo "   Done!"

release:
	BUILDMETA=${GO_BUILDMETA} goreleaser release --rm-dist

release-dryrun:
	BUILDMETA=${GO_BUILDMETA} goreleaser build --rm-dist --snapshot
