## build: build the application and place the built app in the bin folder
build:
	go build -o bin/main .

## run: runs the application
run:
	go run .

## test: runs the tests with the cover flag
test:
	go test . --cover

## install: installs the app to $GOPATH/bin
install:
	go install

## compile: compiles the application for multiple environments and place the output executables under the bin folder
compile:
	# 32-Bit Systems
	# FreeBDS
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 .
	# MacOS
	GOOS=darwin GOARCH=386 go build -o bin/main-darwin-386 .
	# Linux
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 .
	# Windows
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 .
	# 64-Bit
	# FreeBDS
	GOOS=freebsd GOARCH=amd64 go build -o bin/main-freebsd-amd64 .
	# MacOS
	GOOS=darwin GOARCH=amd64 go build -o bin/main-darwin-amd64 .
	# Linux
	GOOS=linux GOARCH=amd64 go build -o bin/main-linux-amd64 .
	# Windows
	GOOS=windows GOARCH=amd64 go build -o bin/main-windows-amd64 .

## help: prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
