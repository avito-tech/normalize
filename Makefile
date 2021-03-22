.PHONY: deps
deps:
	go mod tidy && go mod verify

.PHONY: build
build:
	go build -v .

.PHONY: lint
lint:
	golangci-lint run --new-from-rev=origin/master --config=.golangci.yml ./...

.PHONY: lint-full
lint-full:
	golangci-lint run --config=.golangci.yml ./...

.PHONY: fmt
fmt:
	go fmt  ./...
	goimports -w ./

.PHONY: test
test:
	go test -v -coverprofile="coverage.txt" -covermode=atomic -race -count 1 -timeout 20s ./...
