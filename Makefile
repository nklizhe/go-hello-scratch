all: hello

hello: *.go
	go build -o bin/hello .

linux: *.go
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/linux/hello .

docker: linux
	docker build -t hello-scratch -f Dockerfile .

.PHONY = linux docker