#!/bin/bash

# Check if .env file exists
if [ ! -f .env ]; then
  echo "No .env file found. Creating from sample..."
  cp .env.sample .env
  echo "Please update the .env file with your actual configuration values."
  exit 1
fi

# Build and run the application
echo "Building and starting the reminder app..."
go build -o reminder
./reminder