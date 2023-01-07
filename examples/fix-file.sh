#!/bin/bash

DOC=$(cat "$1")

read -r -d '' PROMPT << EOM
Please rewrite the document so it sounds smart.

$DOC
EOM

echo "$PROMPT" | go run ./cmd/ezgpt3/main.go -tokens 1200