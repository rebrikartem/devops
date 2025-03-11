#!/bin/sh

if [ -z "$DATABASE_URL" ]; then
  echo "Error: DATABASE_URL is not set."
  exit 1
fi

if [ "$1" = "migrate" ]; then
  echo "Running migrations..."
  /app/main migrate
fi

# Запуск приложения
echo "Starting application on port ${APP_PORT}..."
exec /app/main