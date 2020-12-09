# BankApp built in Golang

A Golang API that tracks transactions and bank accounts

## Installation

You can install the bank app by cloning the repo and using `go install`

```bash
git clone https://github.com/ezeoleaf/verbose-octo-happiness
cd verbose-octo-happiness
go install
```

You can run it on the go
```bash
git clone https://github.com/ezeoleaf/verbose-octo-happiness
cd verbose-octo-happiness
go run .
```

Or you can use the following make targets

Build && Run
```bash
git clone https://github.com/ezeoleaf/verbose-octo-happiness
cd verbose-octo-happiness
make build
./bin/main
```

Run
```bash
git clone https://github.com/ezeoleaf/verbose-octo-happiness
cd verbose-octo-happiness
make run
```

## Usage

### Using prebuilt app

Just run the built app inside the bin folder that suits your OS

### Using Docker

For running the application using Docker you can follow these steps:
```bash
docker-composer build && docker-compose up -d
```

### Using Make

To run the application using make targets, first you could try the help section.

`make help`

As a response you will see the entire targets available.

```
Usage: 

  build     build the application and place the built app in the bin folder
  run       runs the application
  test      runs the tests with the cover flag
  install   installs the app to $GOPATH/bin
  compile   compiles the application for multiple environments and place the output executables under the bin folder
  help      prints this help message
```

### Using Go commands

Inside the repo folder:

To run the application:

&nbsp;&nbsp;`go run .`

### Documentation

After running the application you can read the API documentation in:

`http://localhost:1323/swagger/index.html`
