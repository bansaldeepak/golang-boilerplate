#!/bin/sh -e
set -x

# Set environment variables
# export APP_ENVIRONMENT=migration

# Get migration name from command line argument or environment variable
MIGRATION_NAME=${1:-$APP_NAME}

# Check if MIGRATION_NAME is set
if [ -z "$MIGRATION_NAME" ]; then
  echo "Error: Migration name not provided and PROJECT_NAME environment variable is not set."
  exit 1
fi

# Generate a new migration file with datetime prefix
DATETIME=$(date +"%Y%m%d%H%M%S")
migrate create -ext sql -dir migrations -seq "${DATETIME}_${MIGRATION_NAME}"
