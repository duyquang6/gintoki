
PROJECT ?= "gintoki"
PROJECT_NAME := $(PROJECT)
BUILD_DATE := `date +%FT%T%z`

BUILD_FLAGS = -ldflags "-X $(PROJECT_NAME)/app.BuildDate=$(BUILD_DATE)\
						-X $(PROJECT_NAME)/app.BuildCommitHash=$(SHORT_SHA)\
						-X $(PROJECT_NAME)/app.VersionName=$(VERSION_NAME)\
						-X $(PROJECT_NAME)/app.VersionCode=$(BUILD_ID)"
prepare:
	@go get -u -v github.com/gogo/protobuf/proto
	@go get -u -v github.com/gogo/protobuf/jsonpb
	@go get -u -v github.com/gogo/protobuf/protoc-gen-gogo
	@go get -u -v github.com/gogo/protobuf/protoc-gen-gogofaster
	@go get -u -v github.com/gogo/protobuf/protoc-gen-gogofast
	@go get -u -v github.com/gogo/protobuf/protoc-gen-gogoslick
	@go get -u -v github.com/gogo/protobuf/gogoproto
	@go get -u -v github.com/swaggo/swag/cmd/swag

docs:
	@swag init -g cmd/gateway/routers/router.go -o cmd/gateway/docs

dep:
	@go mod download

clean:
	@go mod tidy

test:
	@./scripts/test.sh

build:
	@env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(PROJECT) cmd/server/main.go 

genproto:
	# @protoc -I=. -I=${GOPATH}/src -I=${GOPATH}/src/github.com/gogo/protobuf/protobuf --gofast_out=plugins=grpc:. ./domain/dto/proto/*.proto
	@cd application/handler/proto && protoc -I=. -I=${GOPATH}/src -I=${GOPATH}/src/googleapis -I=${GOPATH}/src/github.com/gogo/protobuf/protobuf --gofast_out=plugins=grpc:. --grpc-gateway_out=logtostderr=true:. *.proto
	@cd application/handler/proto && protoc -I=. -I=${GOPATH}/src -I=${GOPATH}/src/googleapis -I=${GOPATH}/src/github.com/gogo/protobuf/protobuf --gofast_out=plugins=grpc:. *.proto

genclientproto:	
	@cd cmd/client && python -m grpc_tools.protoc -I=. -I=${GOPATH}/src -I=${GOPATH}/src/github.com/gogo/protobuf/protobuf --python_out=. --grpc_python_out=. *.proto

lint:
	@golangci-lint run

