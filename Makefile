

all: protoc mocks cli

.PHONY:
cli: 
	go build -o bin/cli cmd/cli/cli.go 

# clean generated code
.PHONY: clean
clean:
	rm -f api/v1/*pb.go

.PHONY:
test:
	CGO_ENABLED=1 go test ./... -coverprofile=coverage.out -covermode=atomic && go tool cover -func=coverage.out

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...

.PHONY: protoc
protoc: 
	rm -rf api/*
	make -C proto protoc

.PHONY: protoc-ci
protoc-ci: third-party
	protoc -I api --go_out=plugins=grpc:api api/lightbits/api/duros/v2/*.proto

.PHONY: mocks
mocks:
	rm -rf api/duros/v2/mocks/*
	docker run --user $$(id -u):$$(id -g) --rm -w /work -v ${PWD}:/work vektra/mockery:v2.21.1 -r --all --keeptree --dir api/duros/v2 --output api/duros/v2/mocks 
