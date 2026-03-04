BINARY := kubectl-lofi

.PHONY: build install clean release-snapshot

build:
	go build -o $(BINARY) .

install: build
	cp $(BINARY) $(shell go env GOPATH)/bin/

clean:
	rm -f $(BINARY)

release-snapshot:
	goreleaser release --snapshot --clean
