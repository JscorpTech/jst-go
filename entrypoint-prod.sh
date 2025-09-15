#!/bin/bash
set -e

while ! nc -z db 5432; do
  sleep 2
  echo "Waiting postgress...."
done

./bin/api

exit $?
