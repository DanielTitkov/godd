# godd

Simple dd replica in Go

## Usage

`godd -i <input file path> -o <output file path> -b <buffer size>`

Buffer size is optional. Default is 1024 bytes. 

## Preparation

### Install golangci-lint

If you don't have golangci-lint installed.

`make setup`

### Run all the tests

`make test`

### Run linters

`make lint`

### Install godd

Note the before intallation tests and linters will be employed.

`make install`




