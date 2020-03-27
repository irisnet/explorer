all: get_vendor_deps build

get_vendor_deps:
	go mod download
	go mod vendor

build:
	go build -o build/irisplorer explorer.go

irisplorer:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/irisplorer explorer.go