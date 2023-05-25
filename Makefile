export GO111MODULE = on
export COMPOSE_HTTP_TIMEOUT ?= 480



.PHONY: test benchmark run build update-deps lint test-extra 

GOPATH := $(shell go env GOPATH)
LINTERS := \
	github.com/golang/lint/golint \
	github.com/kisielk/errcheck \
	honnef.co/go/tools/cmd/staticcheck \
	honnef.co/go/tools/cmd/unused

PACKAGES = $(shell go list ./... | grep -v /vendor/)

lint:
	env GO111MODULE=off go fmt ./...
	env GO111MODULE=on go vet -mod=vendor ./...
	gofmt -w ./pkg/

install-build-deps:
	go install -v $(LINTERS)

test: lint benchmark
	mkdir -p builds
	env GO111MODULE=on go test -mod=vendor -race -coverprofile=${UNIT_COVERAGE_OUTPUT} ./...

random-test:
	env GO111MODULE=on go test -v -count=1 -mod=vendor -race -run=TestRandomMethod ./...

test-extra: test install-build-deps
	$(GOPATH)/bin/errcheck ./...
	$(GOPATH)/bin/staticcheck ./...
	$(GOPATH)/bin/unused ./...
	for pkg in $$(go list ./... |grep -v /vendor/); do $(GOPATH)/bin/golint $$pkg; done

benchmark:
	env GO111MODULE=on go test -mod=vendor -v -run=NOTHING_USE_MAKE_TEST_INSTEAD -benchmem=true -bench=. ./...

run:
	env GO111MODULE=on go build -mod=vendor
# 	env ./go_sample_stuffs


build:
	env GOOS=linux CGO_ENABLED=0 GO111MODULE=on /usr/local/go/bin/go build -mod=vendor -o builds/go_sample_stuffs main.go

build-atl-upload:
	env GOOS=linux CGO_ENABLED=0 GO111MODULE=on /usr/local/go/bin/go build -mod=vendor -o builds/send_atl_upload_notif kafkaprem/atlupload/main/send_atl_upload_msg.go

docker-build:
	docker build --rm -t go_sample_stuffs .

modtidy:
	go mod tidy

modvendor:
	go mod vendor

modall: modtidy modvendor