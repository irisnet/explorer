all: get_vendor_deps install

get_vendor_deps:
	go get github.com/Masterminds/glide
	glide install
	
install:
	go install ./cmd/irisplorer

irisplorer:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/irisplorer ./cmd/irisplorer