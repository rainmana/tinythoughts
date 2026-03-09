# tinythoughts Architecture

## Overview

tinythoughts is a standalone Go CLI binary that provides systematic thinking frameworks. It's derived from the clear-thought MCP server but redesigned for direct command-line use without the MCP protocol overhead.

## Design Principles

1. **Minimal**: Only essential code, no unnecessary abstractions
2. **Portable**: Single binary with no runtime dependencies
3. **Extensible**: Custom frameworks via YAML files
4. **Configurable**: Environment variables, config files, and CLI flags
5. **Stateful**: Persistent session tracking across invocations

## Architecture

```
┌─────────────────────────────────────────────────────────┐
│                    CLI Layer (Cobra)                     │
│  Commands: think, model, debug, frameworks               │
└────────────────────┬────────────────────────────────────┘
                     │
┌────────────────────▼────────────────────────────────────┐
│              Configuration (Viper)                       │
│  Sources: YAML files, env vars, CLI flags               │
└────────────────────┬────────────────────────────────────┘
                     │
┌────────────────────▼────────────────────────────────────┐
│                  Engine Layer                            │
│  - Sequential thinking logic                             │
│  - Mental model application                              │
│  - Debugging approaches                                  │
│  - Framework loader (YAML)                               │
│  - Session state management                              │
└────────────────────┬────────────────────────────────────┘
                     │
┌────────────────────▼────────────────────────────────────┐
│                  Storage Layer                           │
│  - Session persistence (JSON)                            │
│  - Framework definitions (YAML)                          │
└─────────────────────────────────────────────────────────┘
```

## Directory Structure

```
tinythoughts/
├── cmd/tinythoughts/          # CLI entry point
│   └── main.go               # Cobra commands, flag binding
├── internal/
│   ├── types/                # Core data structures
│   │   └── types.go         # ThoughtData, MentalModelData, etc.
│   ├── config/               # Configuration management
│   │   └── config.go        # Viper integration
│   └── engine/               # Core logic
│       └── engine.go        # Thinking frameworks, session state
├── frameworks/               # Example YAML frameworks
│   ├── five_whys.yaml
│   ├── inversion.yaml
│   └── second_order_thinking.yaml
├── README.md                 # User documentation
├── ARCHITECTURE.md           # This file
├── Makefile                  # Build automation
└── go.mod                    # Go dependencies
```

## Key Components

### 1. CLI Layer (cmd/tinythoughts/main.go)

- **Root Command**: Entry point, loads config and initializes engine
- **Subcommands**: think, model, debug, frameworks
- **Flag Binding**: Connects CLI flags to Viper configuration
- **Output Formatting**: JSON or text output

### 2. Configuration (internal/config/config.go)

- **Viper Integration**: Unified config from multiple sources
- **Precedence**: CLI flags > env vars > config file > defaults
- **Environment Variables**: TINYTHOUGHTS_* prefix
- **Config File**: ~/.config/tinythoughts/tinythoughts.yaml

### 3. Engine (internal/engine/engine.go)

- **Session State**: Tracks thoughts, models, debugging sessions
- **Framework Loader**: Reads YAML files from frameworks directory
- **Core Operations**: Sequential thinking, mental models, debugging
- **Persistence**: Saves/loads session state to JSON

### 4. Types (internal/types/types.go)

- **ThoughtData**: Sequential thinking structure
- **MentalModelData**: Mental model application
- **DebuggingApproachData**: Debugging session
- **Framework**: Custom framework definition
- **SessionState**: Complete session state
- **Result**: Operation output structure

## Data Flow

### Command Execution Flow

```
User Input (CLI)
    ↓
Cobra Command Handler
    ↓
Parse Flags & Arguments
    ↓
Create Data Structure (ThoughtData, etc.)
    ↓
Engine Operation (SequentialThinking, etc.)
    ↓
Update Session State
    ↓
Format Output (JSON/Text)
    ↓
Save Session to Disk
    ↓
Display to User
```

### Configuration Loading Flow

```
Application Start
    ↓
Load Defaults (Viper)
    ↓
Read Config File (if exists)
    ↓
Read Environment Variables
    ↓
Apply CLI Flags
    ↓
Final Configuration
```

### Framework Loading Flow

```
Engine Initialization
    ↓
Expand Framework Directory Path
    ↓
Read Directory Entries
    ↓
For Each .yaml File:
    ↓
    Parse YAML → Framework Struct
    ↓
    Store in frameworks map
    ↓
Frameworks Available for Use
```

## Key Design Decisions

### 1. Why Go?

- Single binary distribution
- Fast compilation and execution
- Excellent CLI library ecosystem (Cobra, Viper)
- Strong standard library for JSON/YAML
- Cross-platform support

### 2. Why Not MCP Protocol?

- MCP adds transport layer complexity
- CLI tools need direct, synchronous execution
- Simpler for AI to invoke via command line
- No need for stdio/HTTP/SSE overhead

### 3. Why YAML for Frameworks?

- Human-readable and editable
- Standard format for configuration
- Easy to version control
- Simple structure for framework definitions

### 4. Why Persistent Sessions?

- Maintain context across CLI invocations
- Track thinking process over time
- Enable review and reflection
- Support iterative problem-solving

### 5. Why Cobra + Viper?

- Industry standard for Go CLIs
- Excellent flag parsing and validation
- Unified configuration management
- Auto-generated help and completion

## Extension Points

### Adding New Commands

1. Define command in `cmd/tinythoughts/main.go`
2. Add flags with `cmd.Flags()`
3. Implement handler function
4. Call engine method
5. Output result and save session

### Adding New Framework Types

1. Define type in `internal/types/types.go`
2. Add engine method in `internal/engine/engine.go`
3. Update session state structure
4. Create CLI command
5. Document in README

### Custom Framework Format

```yaml
name: framework_name
description: Brief description
type: mental_model | debugging | sequential
steps:
  - name: step_name
    description: What this step does
    prompt: Question to guide thinking
prompts:
  conclusion: Final synthesis question
```

## Testing Strategy

### Manual Testing

```bash
# Build
make build

# Test each command
./tinythoughts think "Test thought" -n 1 -t 1
./tinythoughts model first_principles "Test problem"
./tinythoughts debug binary_search "Test issue"
./tinythoughts frameworks

# Test configuration
export TINYTHOUGHTS_OUTPUT_FORMAT=text
./tinythoughts think "Test" -n 1 -t 1

# Test session persistence
./tinythoughts think "First" -n 1 -t 2 -x
./tinythoughts think "Second" -n 2 -t 2
# Check ~/.config/tinythoughts/session.json
```

### Future: Automated Tests

- Unit tests for engine operations
- Integration tests for CLI commands
- Framework loading tests
- Configuration precedence tests

## Performance Characteristics

- **Startup Time**: ~10ms (Go binary)
- **Memory Usage**: ~10MB (minimal footprint)
- **Session File**: ~1KB per 10 operations
- **Framework Loading**: ~1ms per framework

## Comparison to MCP Server

| Aspect | MCP Server | tinythoughts |
|--------|-----------|--------------|
| Language | TypeScript | Go |
| Runtime | Node.js | Native binary |
| Transport | stdio/HTTP/SSE | Direct CLI |
| Configuration | Server config | YAML + env vars |
| Custom Frameworks | Not supported | YAML files |
| Session State | In-memory | Persistent JSON |
| Distribution | npm package | Single binary |
| Startup Time | ~100ms | ~10ms |
| Memory | ~50MB | ~10MB |

## Future Enhancements

1. **Interactive Mode**: REPL for multi-step thinking
2. **Framework Validation**: Schema validation for YAML
3. **Export Formats**: Markdown, HTML output
4. **Framework Templates**: Generate framework scaffolds
5. **Session Management**: List, load, delete sessions
6. **Collaboration**: Share frameworks and sessions
7. **AI Integration**: Direct LLM API calls for guided thinking
8. **Visualization**: Generate diagrams from thinking sessions

## Contributing

When contributing, maintain these principles:

1. **Minimal Code**: Only add what's necessary
2. **Clear Separation**: Keep layers distinct
3. **Backward Compatibility**: Don't break existing commands
4. **Documentation**: Update README and this file
5. **Testing**: Manually test all affected commands
