all: build

build:
	go build -o bin/drone-gce cmd/drone-gce/*.go

build-linux:
	GO111MODULE=on GOOS=linux GOARCH=amd64 go build -o bin/drone-gce-linux-amd64 cmd/drone-gce/*.go

clean:
	rm -rf bin