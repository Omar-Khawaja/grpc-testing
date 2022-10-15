#!/usr/bin/env bash

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./hello.proto

# if you have mockery installed, run the following to generate mocks
# mockery --all -inpkg

