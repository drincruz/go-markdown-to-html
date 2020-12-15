#!/usr/bin/env bash

# go_build.sh

function go_build () {
  go build -o go-markdown-to-html cmd/main.go
}

go_build
