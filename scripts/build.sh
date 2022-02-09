cd ..
docker run --rm -it -v "$GOPATH":/usr/src/server -w /usr/src/server -e GOPATH=/usr/src/server --name 1.17-alpine-cur golang:1.17-alpine sh -c 'cd src/github.com/deepch/Server && go build -ldflags="-s -w" -o build/server-linux-amd64-docker *.go'