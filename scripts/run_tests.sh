#!/bin/bash

# Enable Go modules
export GO111MODULE=on
# export ENVIRONMENT=test

# Load environment variables from .env.test
if [ -f .env.test ]; then
    echo $(cat .env.test | xargs)
    export $(cat .env.test | xargs)
fi

# Change to the src directory
cd src

# Default coverage threshold
COVERAGE_THRESHOLD=50

# Parse command line arguments
while [[ "$#" -gt 0 ]]; do
    case $1 in
        --coverage)
            COVERAGE=true
            if [[ "$2" =~ ^[0-9]+$ ]]; then
                COVERAGE_THRESHOLD="$2"
                shift
            fi
            ;;
    esac
    shift
done

# Run tests with verbose output and coverage if requested
if [ "$COVERAGE" == true ]; then
    go test -v -coverprofile=coverage.out ./...
else
    go test -v ./...
fi

# Check if tests passed
if [ $? -eq 0 ]; then
    echo "All tests passed!"
else
    echo "Some tests failed."
    exit 1
fi

# Generate coverage report if requested
if [ "$COVERAGE" == true ]; then
    go tool cover -html=coverage.out -o coverage.html
    echo "Coverage report generated: coverage.html"

    # Check if coverage is less than the threshold
    COVERAGE_PERCENT=$(go tool cover -func=coverage.out | grep total: | awk '{print substr($3, 1, length($3)-1)}')
    if (( $(echo "$COVERAGE_PERCENT < $COVERAGE_THRESHOLD" | bc -l) )); then
        echo "Coverage is less than $COVERAGE_THRESHOLD%: $COVERAGE_PERCENT%"
        exit 1
    fi
fi