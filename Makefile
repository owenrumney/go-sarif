
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
	go fmt ./...