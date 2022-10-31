BINARY=aws_scheduler
# Build values
VERSION=`git describe --tags --always`
BUILD=`date +%FT%T%z`
MODULE_PATH="github.com/GrolimundSolutions/aws_scheduler/cmd/scheduler/schedulermain"

# Setup the -ldflags option for go build here
LDFLAGS=-ldflags "-w -s -X ${MODULE_PATH}.Version=${VERSION} -X ${MODULE_PATH}.Build=${BUILD}"



.PHONY: build
build:
	@echo "Version: ${VERSION}"
	@echo "Build: ${BUILD}"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ./${BINARY} ./cmd/scheduler/

.PHONY: lint
lint:
	CGO_ENABLED=0 GOGC=40 golangci-lint run --timeout 4m

run:
	@echo "Version: ${VERSION}"
	@echo "Build: ${BUILD}"
	go run ${LDFLAGS} ./cmd/scheduler

tests:
	go test -v -short ./...

.PHONY: clean
clean:
	rm -f ${BINARY}_amd64

# make release bump=major  // 1.0.0
# make release bump=minor  // 0.1.0
# make release			   // 0.0.1
release:
	$(eval v := $(shell git ls-remote --refs --tags | cut --delimiter='/' --fields=3 | tr '-' '~' | sort --version-sort | tail --lines=1))
ifeq ($(bump), major)
	$(eval f := 1)
else ifeq ($(bump), minor)
	$(eval f := 2)
else
	$(eval f := 3)
endif
	$(eval OLD_VERSION = $(shell echo "${v}"))
	$(eval NEW_VERSION = $(shell echo $(v) | awk -F. -v OFS=. -v f=$(f) '{ $$f++ } 1'))
	@echo "OLD_VERSION: ${OLD_VERSION}"
	@echo "NEW_VERSION: ${NEW_VERSION}"
	git tag $(NEW_VERSION) && git push --tags
