
.PHONY: test
test: vet
	go test -v -coverprofile coverage.txt  ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: goimports
goimports:
	goimports