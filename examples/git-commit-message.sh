#!/bin/bash

# Example (run on this repo)
# ./examples/git-commit-message.sh
# Commit: Add support for OpenAI GPT-3 Completions API

CHANGES=$(git diff --staged --ignore-all-space)

read -r -d '' PROMPT << EOM
The following is a diff of the changes you are about to commit.
Please write a commit message that describes the changes you are making.
The commit message should be written in the imperative mood, as if you were
telling someone to do something. For example, "Fix bug" and not "Fixed bug"
or "Fixes bug." It should also be less than 70 characters long.

Changes:
$CHANGES
EOM

echo "$PROMPT"  | go run ./cmd/ezgpt3/main.go -tokens 70