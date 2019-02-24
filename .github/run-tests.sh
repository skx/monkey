#!/bin/sh

# init modules
go mod init

# Run golang tests
go test ./...

# Build a binary and run it.
go build -ldflags "-X main.version=test-code" -o "monkey"
./monkey -eval 'puts("ok - stdlib\n"); exit(0);'
