#!/bin/bash
echo "Building Windows EXE..."

GOOS=windows GOARCH=amd64 CGO_ENABLED=1 \
CC=x86_64-w64-mingw32-gcc \
go build -o DualMonitorClient.exe main.go

echo "Done: DualMonitorClient.exe"
