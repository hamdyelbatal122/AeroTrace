BINARY := aerotrace
ENTRY := foo.sky

.PHONY: build test vet fmt check run clean

build:
	go build -o $(BINARY) .

test:
	go test ./...

vet:
	go vet ./...

fmt:
	gofmt -w *.go

check:
	test -z "$(gofmt -l *.go)"
	go test ./...
	go vet ./...

run:
	go run . $(ENTRY)

clean:
	rm -f $(BINARY) *.out *.test
