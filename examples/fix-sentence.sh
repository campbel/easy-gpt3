#!/bin/bash

# Example
# echo "What is your problem man" | ./examples/fix-paragraph.sh
# What is the issue you are facing, friend?

MESSAGE=$(cat /dev/stdin)

read -r -d '' PROMPT << EOM
Please rewrite the following sentence so it sounds smart.

$MESSAGE
EOM

echo "$PROMPT" | go run ./cmd/ezgpt3/main.go