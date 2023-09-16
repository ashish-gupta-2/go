BASE_DIR = $(shell pwd)
API = $(BASE_DIR)/cmd/main.go
CIURL = https://ci.ashish.com
PIPELINE = test
TEAMNAME = self
OUTPUT_DIR = $(BASE_DIR)/out


all: vend lint build test

vend:
	go mod tidy
	go mod vendor

lint:
	golangci-lint run ./...

clean:
	rm -rf $(OUTPUT_DIR)

build: clean
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
	go build -ldflags=$(LDFLAGS_API) -o $(OUTPUT_DIR)/bin/goapi $(API)

docker-build: build
	# docker login docker.hub.com
	docker build -t go-api -f Dockerfile .
	docker tag go-api ashish.com/go-api/goapi:test
	docker push ashish.com/go-api/goapi:tes
	# docker run go-api execute -j $(ARGS)

test:
	ginkgo -r -v -race -trace -cover --label-filter="unit"

cover: clean
	mkdir -p $(OUTPUT_DIR)/cover
	ginkgo -cover -outputdir=$(OUTPUT_DIR)/cover ./...

run-api:
	go run $(API)

mockgen:
	mockgen -destination=mocks/mock_storage.go -package=mocks -source=pkg/storage/storage.go
	mockgen -destination=mocks/mock_service.go -package=mocks -source=pkg/services/service.go



