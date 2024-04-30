all: lint build test

build:
	@echo "Building..."
	go build ./...

lint: vet staticcheck misspell

vet:
	@echo "Vetting..."
	go vet ./...

staticcheck:
	@[ -x "$(shell which staticcheck)" ] || go install honnef.co/go/tools/cmd/staticcheck@master
	staticcheck ./...

test:
	 go test -cover ./...

start:
	 go run main.go

misspell:
	@[ -x "$(shell which misspell)" ] || go install github.com/client9/misspell/cmd/misspell
	find . -name '*.go' -not -path './vendor/*' | xargs misspell -error