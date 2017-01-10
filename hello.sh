#!/usr/bin/env bash

MESSAGE="hello"

if [[ $1 ]]; then
    MESSAGE="$MESSAGE $1"
fi

curl -H "Content-Type: application/json" -X POST -d "{\"message\":\"$MESSAGE\"}" http://localhost:3000/message
