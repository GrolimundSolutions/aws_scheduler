BINARY=aws_scheduler




.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY}_amd64 ./cmd/scheduler/

.PHONY: lint
lint:
	CGO_ENABLED=0 GOGC=40 golangci-lint run --timeout 5m

.PHONY: clean
clean:
	rm -f ${BINARY}_amd64