
.PHONY: test
test: vet
	go test -v  ./... && cd v2 && go test -v  ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: goimports
goimports:
	goimports