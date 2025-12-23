# Testing Documentation

## Overview

This document describes the comprehensive unit testing suite for the Go Text Processing Tool. The test suite ensures reliability, correctness, and maintainability of all core functionalities.

## Test Structure

### Unit Tests

Each core function has dedicated unit tests:

- **`hexconv_test.go`** - Tests hex/binary to decimal conversions
- **`UpLowCap_test.go`** - Tests case transformations and modifiers  
- **`anconv_test.go`** - Tests article conversions (a/an)
- **`Punctuation_test.go`** - Tests punctuation and quote processing
- **`pipeline_test.go`** - Tests complete pipeline integration

### Test Categories

#### 1. Basic Functionality Tests
- Simple conversions and transformations
- Standard use cases
- Expected behavior validation

#### 2. Edge Case Tests  
- Boundary conditions
- Invalid inputs
- Empty/null scenarios
- Overflow conditions

#### 3. Integration Tests
- Complete pipeline processing
- Multi-stage transformations
- Order-dependent operations

#### 4. Error Handling Tests
- Invalid modifier formats
- Malformed input handling
- Graceful degradation

## Running Tests

### Command Line

```bash
# Run all tests
go test ./pkg/...

# Run tests with verbose output
go test -v ./pkg/...

# Run tests with coverage
go test -cover ./pkg/...

# Generate coverage report
go test ./pkg/... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Using Test Scripts

**Linux/Mac:**
```bash
chmod +x test.sh
./test.sh
```

**Windows:**
```cmd
test.bat
```

## Test Coverage Goals

- **Minimum Coverage**: 85%
- **Critical Functions**: 95%+
- **Edge Cases**: 100% of identified scenarios

## Test Data

### HexBin Function Tests
- Valid hex conversions (1E → 30)
- Valid binary conversions (10 → 2)  
- Large number handling
- Invalid format handling
- Multiple conversions in one string

### UpLowCap Function Tests
- Basic case transformations
- Modifier counts (up, 2)
- Article conversions with modifiers
- Word boundary detection
- Alphanumeric sequence handling

### AnConvert Function Tests
- Basic article conversions (a → an)
- Vowel detection (a, e, i, o, u, h)
- Case preservation
- Multiple article scenarios
- Quoted word handling

### FixPunctuation Function Tests
- Punctuation spacing rules
- Quote pair processing
- Nested quote handling
- Contraction preservation
- Complex punctuation sequences

### Pipeline Integration Tests
- Complete processing workflow
- Order dependency validation
- Multi-feature interactions
- Real-world scenarios

## Continuous Integration

Tests should be run:
- Before every commit
- On pull requests
- In CI/CD pipeline
- Before releases

## Adding New Tests

When adding new functionality:

1. Create corresponding test cases
2. Include edge cases
3. Test error conditions
4. Verify integration impact
5. Update coverage goals

## Test Maintenance

- Review tests quarterly
- Update for new requirements
- Remove obsolete tests
- Optimize slow tests
- Document test rationale