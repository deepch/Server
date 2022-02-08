APP=server
SERVER_FLAGS ?= -config config.json
PLATFORMS=linux windows darwin
ARCHITECTURES=amd64 arm
VERSION=1.0.1
BUILD=`git rev-parse HEAD`
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

P="\\033[34m[+]\\033[0m"

#generate api proto and api doc swagger
gen:
	protoc -I ./api/proto \
	  --gogofaster_out=:api/grpc \
	  --go_out ./api/proto --go_opt paths=source_relative \
	  --go-grpc_out ./api/proto \
	  --go-grpc_opt paths=source_relative \
	  --grpc-gateway_out ./api/proto \
	  --doc_out=markdown,grpc-api-spec.md:docs \
	  --grpc-gateway_opt paths=source_relative \
	  --swagger_out=logtostderr=true:docs \
	  ./api/proto/server/server.proto
	  cp ./docs/server/server.swagger.json ./assets/docs/server.swagger.json

release: gen build_darwin_amd64 build_darwin_arm64 build_windows_amd64 build_windows_386 build_linux_amd64 build_linux_386 build_linux_arm_5 build_linux_arm_6 build_linux_arm_7 build_linux_mipsle_softfloat build_linux_mipsle_hardfloat build_linux_mips_softfloat build_linux_mips_hardfloat

build:
	@echo "$(P) build"
	GO111MODULE=on go build *.go

build_darwin_amd64:
	@echo "$(P) build mac os 64 bit"
	GO111MODULE=on GOOS=darwin GOARCH=amd64 go build -o release/server-darwin-amd64 *.go

build_darwin_arm64:
	@echo "$(P) build mac os M1"
	GO111MODULE=on GOOS=darwin GOARCH=arm64 go build -o release/server-darwin-arm64 *.go

build_windows_amd64:
	@echo "$(P) build windows 64 bit"
	GO111MODULE=on GOOS=windows GOARCH=amd64 go build -o release/server-windows-amd64 *.go

build_windows_386:
	@echo "$(P) build windows 32 bit"
	GO111MODULE=on GOOS=windows GOARCH=386 go build -o release/server-windows-386 *.go

build_linux_amd64:
	@echo "$(P) build linux 64 bit"
	GO111MODULE=on GOOS=linux GOARCH=amd64 go build -o release/server-linux-amd64 *.go

build_linux_386:
	@echo "$(P) build linux 32 bit"
	GO111MODULE=on GOOS=linux GOARCH=386 go build -o release/server-linux-386 *.go

build_linux_arm_5:
	@echo "$(P) build linux arm 5"
	GO111MODULE=on GOOS=linux GOARCH=arm GOARM=5 go build -o release/server-linux-arm-5 *.go

build_linux_arm_6:
	@echo "$(P) build linux arm 6"
	GO111MODULE=on GOOS=linux GOARCH=arm GOARM=6 go build -o release/server-linux-arm-6 *.go

build_linux_arm_7:
	@echo "$(P) build linux arm 7"
	GO111MODULE=on GOOS=linux GOARCH=arm GOARM=7 go build -o release/server-linux-arm-6 *.go

build_linux_mipsle_softfloat:
	@echo "$(P) build linux mipsle softfloat"
	GO111MODULE=on CFLAGS=-msoft-float CGO_CFLAGS=-msoft-float CGO_LDFLAGS=-msoft-float GOOS=linux GOARCH=mipsle GOMIPS=softfloat go build -o release/server-linux-mipsle-softfloat *.go

build_linux_mipsle_hardfloat:
	@echo "$(P) build linux mipsle hardfloat"
	GO111MODULE=on CFLAGS=-mhard-float CGO_CFLAGS=-mhard-float CGO_LDFLAGS=-mhard-float GOOS=linux GOARCH=mipsle GOMIPS=hardfloat go build -o release/server-linux-mipsle-hardfloat *.go

build_linux_mips_softfloat:
	@echo "$(P) build linux mips softfloat"
	GO111MODULE=on CFLAGS=-msoft-float CGO_CFLAGS=-msoft-float CGO_LDFLAGS=-msoft-float GOOS=linux GOARCH=mips GOMIPS=softfloat go build -o release/server-linux-mips-softfloat *.go

build_linux_mips_hardfloat:
	@echo "$(P) build linux mips hardfloat"
	GO111MODULE=on CFLAGS=-mhard-float CGO_CFLAGS=-mhard-float CGO_LDFLAGS=-mhard-float GOOS=linux GOARCH=mips GOMIPS=hardfloat go build -o release/server-linux-mips-hardfloat *.go

run:
	@echo "$(P) run"
	GO111MODULE=on go run *.go

serve:
	@$(MAKE) server

server:
	@echo "$(P) server $(SERVER_FLAGS)"
	./${APP} $(SERVER_FLAGS)

lint:
	@echo "$(P) lint"
	go vet

.NOTPARALLEL:

.PHONY: build run server test lint
