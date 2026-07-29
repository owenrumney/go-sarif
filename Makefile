
.PHONY: test
test: vet
	go test -v  ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: goimports
goimports:
	goimports

.PHONY: fmt
fmt:
	gofmt -w -s ./..

.PHONY: lint
lint:
	which golangci-lint || go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.12.2
	golangci-lint run