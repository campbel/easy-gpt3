#!/bin/bash

CHANGES=$(git diff --staged --ignore-all-space)

read -r -d '' PROMPT << EOM
The following is a diff for recent changes to this repository.
Write a commit message for the changes:

$CHANGES
EOM

echo "$PROMPT"  | go run ./cmd/ezgpt3/main.go