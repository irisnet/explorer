all: get_deps build

get_deps:
	rm -rf ./vendor
	go mod download
	go mod vendor

build:
	go build -o build/irisplorer explorer.go

irisplorer:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/irisplorer explorer.go