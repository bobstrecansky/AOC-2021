#!/bin/sh
ROOT=$(git rev-parse --show-toplevel)

MOUNTOPTION=""
if [ "$(uname -s)" = "Darwin" ]; then
    MOUNTOPTION=":delegated"
fi

dockerwrap() {
    IMAGE=$1
    shift
    docker run --rm \
        -u $(id -u):$(id -g) \
        -v "$ROOT":"$ROOT"${MOUNTOPTION} \
        -w "$ROOT" \
        -e GO111MODULE=on \
        -e GOCACHE="$ROOT/.cache" \
        -e GOLANGCI_LINT_CACHE="$ROOT/.cache" \
        -e CGO_ENABLED=0 \
        -e GOOS=linux \
        -e GOARCH=amd64 \
        "$IMAGE" \
        "$@"
}
