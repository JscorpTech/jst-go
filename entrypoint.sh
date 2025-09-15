#!/bin/bash
set -e
while ! nc -z db 5432; do
  sleep 2
  echo "Waiting postgress...."
done

air --build.cmd "go build -o bin/api ./cmd/api/main.go" --build.bin "./bin/api"
exit $?
