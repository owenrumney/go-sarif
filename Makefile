
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
	which golangci-lint || go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.7
	golangci-lint run