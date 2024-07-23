#!/bin/bash

# Test file for Dockerfile

# Test 1: Check if the image builds successfully
echo "Test 1: Building Docker image"
if docker build -t my-go-app .; then
    echo "Test 1: Passed - Image built successfully"
else
    echo "Test 1: Failed - Image build failed"
    exit 1
fi

# Test 2: Check if the container runs
echo "Test 2: Running container"
if docker run -d --name test-container -p 8080:8080 my-go-app; then
    echo "Test 2: Passed - Container started successfully"
else
    echo "Test 2: Failed - Container failed to start"
    docker rm -f test-container
    exit 1
fi

# Test 3: Check if the application is responding
echo "Test 3: Checking application response"
sleep 5  # Give the application a moment to start up
if curl -s -o /dev/null -w "%{http_code}" http://localhost:8080 | grep -q "200"; then
    echo "Test 3: Passed - Application is responding"
else
    echo "Test 3: Failed - Application is not responding"
    docker rm -f test-container
    exit 1
fi

# Clean up
docker rm -f test-container

echo "All tests passed successfully"