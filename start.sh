#!/bin/bash -e

case $1 in
  "run")
    docker-compose build whistleblower
    docker-compose run --rm whistleblower migrate:run
    docker-compose up whistleblower
    ;;
  "up")
    docker-compose up whistleblower
    ;;
  *)
    echo "usage: $0 [run|up]"
    exit 1
    ;;
esac