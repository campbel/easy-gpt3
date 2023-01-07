#!/bin/bash

# Usage:
# ./fix-file.sh <file>

DOC=$(cat "$1")

read -r -d '' PROMPT << EOM
Please rewrite the document so it sounds smart.

$DOC
EOM

echo "$PROMPT" | ezgpt3 -tokens 1200