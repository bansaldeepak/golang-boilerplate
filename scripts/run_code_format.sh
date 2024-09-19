#!/bin/sh -e
set -x

# Ensure golint is installed
# if ! command -v golint &> /dev/null
# then
#     go install golang.org/x/lint/golint@latest
# fi

# Change to the src directory
# cd src

# Function to check formatting
check_format() {
    gofmt -l . | tee /dev/tty
    goimports -l . | tee /dev/tty
    golines -m 80 -l . | tee /dev/tty
    golint ./... | tee /dev/tty
}

# Function to fix formatting
fix_format() {
    gofmt -s -w .
    goimports -w .
    golines -w --max-len=80 .
}

# Check if the --fix argument is passed
if [ "$1" = "--fix" ]; then
    fix_format
else
    check_format
fi