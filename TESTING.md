# Testing Documentation

## Test Coverage Summary

```
Package                                          Coverage
-------------------------------------------------------
internal/config                                  100.0%
internal/engine                                   77.1%
internal/types                                   N/A (no logic)
cmd/tinythoughts                                  0.0% (CLI entry point)
-------------------------------------------------------
Overall Core Logic Coverage                      ~88.5%
```

## Unit Tests

### Config Package Tests (`internal/config/config_test.go`)

- ✓ `TestLoad` - Verifies basic configuration loading
- ✓ `TestLoadWithEnvVars` - Tests environment variable precedence
- ✓ `TestLoadDefaults` - Validates default values

**Coverage**: 100% of statements

### Engine Package Tests (`internal/engine/engine_test.go`)

- ✓ `TestNew` - Engine initialization
- ✓ `TestSequentialThinking` - Sequential thinking operation
- ✓ `TestMentalModel` - Mental model application
- ✓ `TestDebuggingApproach` - Debugging approach execution
- ✓ `TestLoadFrameworks` - YAML framework loading
- ✓ `TestSaveLoadSession` - Session persistence
- ✓ `TestListFrameworks` - Framework listing
- ✓ `TestGetFrameworkNotFound` - Error handling

**Coverage**: 77.1% of statements

### Types Package Tests (`internal/types/types_test.go`)

- ✓ `TestThoughtData` - ThoughtData structure validation
- ✓ `TestMentalModelData` - MentalModelData structure validation
- ✓ `TestDebuggingApproachData` - DebuggingApproachData structure validation
- ✓ `TestFramework` - Framework structure validation
- ✓ `TestSessionState` - SessionState structure validation

**Coverage**: N/A (pure data structures)

## Functional Tests

Comprehensive end-to-end tests in `test_functionality.sh`:

1. ✓ Sequential thinking command execution
2. ✓ Mental model command execution
3. ✓ Debugging approach command execution
4. ✓ Session file creation and persistence
5. ✓ Session data integrity
6. ✓ Text output format
7. ✓ Framework loading from YAML files

## Running Tests

### All Unit Tests
```bash
make test
```

### Verbose Output
```bash
make test-verbose
```

### Coverage Report
```bash
make test-coverage
# Opens coverage.html in browser
```

### Functional Tests
```bash
./test_functionality.sh
```

## Test Results

### Unit Tests
```
=== RUN   TestLoad
--- PASS: TestLoad (0.00s)
=== RUN   TestLoadWithEnvVars
--- PASS: TestLoadWithEnvVars (0.00s)
=== RUN   TestLoadDefaults
--- PASS: TestLoadDefaults (0.00s)
PASS
ok      github.com/waldzellai/tinythoughts/internal/config      0.172s

=== RUN   TestNew
--- PASS: TestNew (0.00s)
=== RUN   TestSequentialThinking
--- PASS: TestSequentialThinking (0.00s)
=== RUN   TestMentalModel
--- PASS: TestMentalModel (0.00s)
=== RUN   TestDebuggingApproach
--- PASS: TestDebuggingApproach (0.00s)
=== RUN   TestLoadFrameworks
--- PASS: TestLoadFrameworks (0.00s)
=== RUN   TestSaveLoadSession
--- PASS: TestSaveLoadSession (0.00s)
=== RUN   TestListFrameworks
--- PASS: TestListFrameworks (0.00s)
=== RUN   TestGetFrameworkNotFound
--- PASS: TestGetFrameworkNotFound (0.00s)
PASS
ok      github.com/waldzellai/tinythoughts/internal/engine      0.431s

=== RUN   TestThoughtData
--- PASS: TestThoughtData (0.00s)
=== RUN   TestMentalModelData
--- PASS: TestMentalModelData (0.00s)
=== RUN   TestDebuggingApproachData
--- PASS: TestDebuggingApproachData (0.00s)
=== RUN   TestFramework
--- PASS: TestFramework (0.00s)
=== RUN   TestSessionState
--- PASS: TestSessionState (0.00s)
PASS
ok      github.com/waldzellai/tinythoughts/internal/types       0.303s
```

### Functional Tests
```
=== Functional Testing ===
✓ Test 1: Sequential thinking
✓ Test 2: Mental model
✓ Test 3: Debugging approach
✓ Test 4: Session persistence
✓ Test 5: Session contains data
✓ Test 6: Text output format
✓ Test 7: Frameworks listing

=== All functional tests passed! ===
```

## What's Tested

### Core Functionality
- ✓ Engine initialization and state management
- ✓ Sequential thinking operations
- ✓ Mental model applications
- ✓ Debugging approach execution
- ✓ Session persistence (save/load)
- ✓ Framework loading from YAML
- ✓ Configuration management (files, env vars, defaults)

### Data Structures
- ✓ All type definitions validated
- ✓ JSON serialization/deserialization
- ✓ YAML parsing for frameworks

### Configuration
- ✓ Default values
- ✓ Environment variable overrides
- ✓ Config file loading
- ✓ Precedence rules

### Error Handling
- ✓ Missing framework errors
- ✓ File I/O errors
- ✓ Invalid YAML handling

## What's Not Tested

### CLI Layer (cmd/tinythoughts)
- Command-line argument parsing (handled by Cobra)
- Flag binding (handled by Viper)
- Output formatting logic

**Rationale**: These are integration points with well-tested libraries (Cobra/Viper). Functional tests cover end-to-end CLI behavior.

### Edge Cases Not Covered
- Concurrent access to session files
- Very large session files (>1GB)
- Malformed YAML with complex nesting
- Network filesystem edge cases

## Test Maintenance

### Adding New Tests

1. Create `*_test.go` file in the same package
2. Follow naming convention: `Test<FunctionName>`
3. Use `t.TempDir()` for file operations
4. Clean up with `defer` or `t.Cleanup()`

Example:
```go
func TestNewFeature(t *testing.T) {
    // Setup
    eng := New("test-session")
    
    // Execute
    result := eng.NewFeature()
    
    // Assert
    if result.Status != "success" {
        t.Errorf("Expected success, got %s", result.Status)
    }
}
```

### Running Specific Tests
```bash
go test ./internal/engine -run TestSequentialThinking
go test ./internal/config -v
```

### Benchmarking
```bash
go test ./... -bench=.
```

## Continuous Integration

Recommended CI pipeline:

```yaml
- name: Test
  run: |
    go test ./... -v
    go test ./... -cover
    ./test_functionality.sh
```

## Test Philosophy

1. **Unit tests** cover individual functions and methods
2. **Functional tests** verify end-to-end CLI behavior
3. **Integration tests** (future) would test with real config files
4. **Focus on core logic**, not library wrappers

The test suite ensures:
- Core thinking operations work correctly
- Session state persists accurately
- Configuration loads from all sources
- Framework loading handles YAML correctly
- Error conditions are handled gracefully
