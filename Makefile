

all: cli

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

.PHONY: third-party
third-party:
	rm -rf api/google/api api/protoc-gen-swagger/options
	mkdir -p api/google/api api/protoc-gen-swagger/options
	wget https://raw.githubusercontent.com/grpc-ecosystem/grpc-gateway/v1/protoc-gen-swagger/options/annotations.proto -O api/protoc-gen-swagger/options/annotations.proto
	wget https://raw.githubusercontent.com/grpc-ecosystem/grpc-gateway/v1/protoc-gen-swagger/options/openapiv2.proto -O api/protoc-gen-swagger/options/openapiv2.proto
	wget https://raw.githubusercontent.com/googleapis/api-common-protos/master/google/api/annotations.proto -O api/google/api/annotations.proto
	wget https://raw.githubusercontent.com/googleapis/api-common-protos/master/google/api/http.proto -O api/google/api/http.proto

.PHONY: protoc
protoc: third-party
	docker run --rm --user $$(id -u):$$(id -g) -v ${PWD}:/work metalstack/builder protoc -I api --go_out=plugins=grpc:api api/lightbits/api/duros/v1/*.proto
	docker run --rm --user $$(id -u):$$(id -g) -v ${PWD}:/work metalstack/builder protoc -I api --go_out=plugins=grpc:api api/lightbits/api/duros/v2/*.proto

.PHONY: protoc-ci
protoc-ci: third-party
	protoc -I api --go_out=plugins=grpc:api api/lightbits/api/duros/v1/*.proto
	protoc -I api --go_out=plugins=grpc:api api/lightbits/api/duros/v2/*.proto
