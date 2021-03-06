# Setup scripts
# build binary for Mac and set command into your GOPATH/bin
darwin:
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o bin/mac/ .

# build binary for Linux AMD64 architecture and set command into your GOPATH/bin
linux64:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/linux64/ .

