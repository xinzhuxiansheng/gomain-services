export tag=v1.0
root:
	export ROOT=github.com/xinzhuxiansheng/golang

build:
	echo "building httpserver binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64/httpserver .

release: build
	echo "building httpserver container"
	docker build -t xinzhuxiansheng/httpserver:${tag} .

push: release
	echo "pushing xinzhuxiansheng/httpserver"
	docker push xinzhuxiansheng/httpserver:${tag}
