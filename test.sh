#!/usr/bin/env bash

echo ">> Testing API.."
pushd api
    ./bin/rails test
popd

echo
echo ">> Testing Worker.."
pushd worker
    go test
popd
