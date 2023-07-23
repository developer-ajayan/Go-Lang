#!/bin/bash

set -e

if [ -n "$POSTGRES_DB" ]; then
    echo "Creating database: $POSTGRES_DB"
    psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
        CREATE DATABASE IF NOT EXISTS $POSTGRES_DB;
    EOSQL
fi