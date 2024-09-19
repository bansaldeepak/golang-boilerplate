#!/bin/sh -e
set -x

# Set (or export) necessary environment variables, if they aren't set elsewhere
# Example:
# export POSTGRES_SCHEMA=your_schema
# export POSTGRES_USER=your_user
# export POSTGRES_DATABASE=your_database
# export POSTGRES_PASSWORD=your_password

# Load environment variables from .env
if [ -f .env ]; then
    export $(cat .env | xargs)
fi

# Create the SQL script using the POSTGRES_SCHEMA environment variable
echo "CREATE SCHEMA IF NOT EXISTS $POSTGRES_SCHEMA;" > /tmp/init.sql
echo "ALTER USER $POSTGRES_USER SET search_path TO $POSTGRES_SCHEMA;" >> /tmp/init.sql

# Run the script with psql to create the schema and set the search path
PGPASSWORD=$POSTGRES_PASSWORD psql -h localhost -p 5432 -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DATABASE" < /tmp/init.sql

# Run migrations
migrate -path migrations -database "postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@localhost:5432/$POSTGRES_DATABASE?sslmode=disable" up