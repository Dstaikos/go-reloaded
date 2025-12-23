#!/bin/bash

# Go Text Processing Tool - Test Runner
# This script runs all unit tests with coverage reporting

echo "🧪 Running Go Text Processing Tool Tests..."
echo "=========================================="

# Change to project root directory
cd ..

# Run all tests with verbose output and coverage
echo "📋 Running unit tests..."
go test -v ./pkg/... -cover

echo ""
echo "📊 Generating detailed coverage report..."
go test ./pkg/... -coverprofile=coverage.out

echo ""
echo "🎯 Coverage by function:"
go tool cover -func=coverage.out

echo ""
echo "📈 To view HTML coverage report, run:"
echo "go tool cover -html=coverage.out"

echo ""
echo "✅ Test run complete!"