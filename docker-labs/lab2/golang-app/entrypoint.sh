#!/bin/sh

if [ -z "$DATABASE_URL" ]; then
  echo "Error: DATABASE_URL is not set."
  exit 1
fi

echo "Running migrations..."
/app/main migrate

echo "Starting application on port ${APP_PORT}..."
exec /app/main