#!/bin/bash

# Script to compile and execute the Go application

# Define the Go source file and output binary
SOURCE_FILE="./cmd/main.go"
OUTPUT_BINARY="aws-cli-simulator"

# Step 1: Compile the Go application
echo "Compiling the Go application..."
go build -o $OUTPUT_BINARY $SOURCE_FILE

# Check if the compilation was successful
if [ $? -ne 0 ]; then
    echo "Compilation failed. Exiting."
    exit 1
fi

echo "Compilation successful. Binary created: $OUTPUT_BINARY"

# Step 2: Execute the binary
echo "Executing the application..."
./$OUTPUT_BINARY