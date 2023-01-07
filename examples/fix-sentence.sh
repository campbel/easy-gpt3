#!/bin/bash

# Usage
# echo "What is your problem man" | ./examples/fix-sentence.sh

MESSAGE=$(cat /dev/stdin)

read -r -d '' PROMPT << EOM
Please rewrite the following sentence so it sounds smart.

$MESSAGE
EOM

echo "$PROMPT" | ezgpt3