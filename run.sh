#!/bin/bash

# Install packages
go mod download

export PORT=3001

# Build and run the server
go run main.go
