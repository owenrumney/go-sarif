
.PHONY: test
test: vet
	go test -v -covermode=atomic -coverpkg ./... -coverprofile coverage.txt ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: goimports
goimports:
	goimports