# tinythoughts

A lightweight CLI tool for systematic thinking, mental models, and debugging approaches. Built from the clear-thought MCP server, tinythoughts makes powerful thinking frameworks easily accessible from the command line.

## Features

- **Sequential Thinking**: Track and organize thoughts in a structured sequence
- **Mental Models**: Apply proven mental models like first principles, opportunity cost, and more
- **Debugging Approaches**: Systematic debugging with binary search, root cause analysis, etc.
- **Custom Frameworks**: Load your own thinking frameworks from YAML files
- **Session Management**: Automatically saves your thinking sessions
- **Environment Configuration**: Configure via YAML files or environment variables

## Installation

### From Source

```bash
git clone https://github.com/waldzellai/tinythoughts
cd tinythoughts
go build -o tinythoughts ./cmd/tinythoughts
sudo mv tinythoughts /usr/local/bin/
```

### Using Go Install

```bash
go install github.com/waldzellai/tinythoughts/cmd/tinythoughts@latest
```

## Quick Start

```bash
# Sequential thinking
tinythoughts think "First, I need to understand the problem" -n 1 -t 3 -x

# Apply a mental model
tinythoughts model first_principles "How should I architect this system?" \
  -s "Break down into components" \
  -s "Identify core requirements" \
  -r "Start from fundamental needs" \
  -c "Build from the ground up"

# Debugging approach
tinythoughts debug binary_search "Performance issue in API" \
  -s "Check if issue is in frontend or backend" \
  -s "Narrow down to specific endpoint" \
  -f "Database query is slow" \
  -r "Added index to improve query performance"

# List available frameworks
tinythoughts frameworks
```

## Configuration

### Config File

Create `~/.config/tinythoughts/tinythoughts.yaml`:

```yaml
frameworks_dir: ~/.config/tinythoughts/frameworks
session_file: ~/.config/tinythoughts/session.json
output_format: json
```

### Environment Variables

All configuration can be set via environment variables with the `TINYTHOUGHTS_` prefix:

```bash
export TINYTHOUGHTS_FRAMEWORKS_DIR=~/my-frameworks
export TINYTHOUGHTS_SESSION_FILE=~/my-session.json
export TINYTHOUGHTS_OUTPUT_FORMAT=text
```

### Command-Line Flags

```bash
tinythoughts --config /path/to/config.yaml --output json think "My thought"
```

## Custom Frameworks

Create custom thinking frameworks as YAML files in your frameworks directory.

### Framework Structure

```yaml
name: framework_name
description: Brief description of the framework
type: mental_model  # or debugging, sequential
steps:
  - name: step_name
    description: What this step does
    prompt: Question to guide thinking
prompts:
  conclusion: Final synthesis question
```

### Example: Five Whys

```yaml
name: five_whys
description: Root cause analysis by asking "why" five times
type: debugging
steps:
  - name: identify_problem
    description: State the problem clearly
    prompt: What is the problem you're trying to solve?
  - name: first_why
    description: Ask why the problem occurs
    prompt: Why does this problem occur?
  - name: second_why
    description: Ask why the first answer is true
    prompt: Why is that the case?
  - name: third_why
    description: Continue drilling down
    prompt: Why does that happen?
  - name: fourth_why
    description: Get closer to root cause
    prompt: What causes that?
  - name: fifth_why
    description: Identify the root cause
    prompt: What is the fundamental reason?
prompts:
  conclusion: "Based on the five whys, what is the root cause?"
```

Place this in `~/.config/tinythoughts/frameworks/five_whys.yaml` and it will be automatically loaded.

## Commands

### think

Sequential thinking with thought tracking.

```bash
tinythoughts think [thought] [flags]

Flags:
  -n, --number int    Thought number (default 1)
  -t, --total int     Total thoughts (default 1)
  -x, --next          Next thought needed
```

### model

Apply mental models to problems.

```bash
tinythoughts model [model-name] [problem] [flags]

Flags:
  -s, --steps stringSlice      Steps taken
  -r, --reasoning string       Reasoning process
  -c, --conclusion string      Conclusion

Built-in models:
  - first_principles
  - opportunity_cost
  - error_propagation
  - rubber_duck
  - pareto_principle
  - occams_razor
```

### debug

Apply debugging approaches.

```bash
tinythoughts debug [approach] [issue] [flags]

Flags:
  -s, --steps stringSlice      Steps taken
  -f, --findings string        Findings discovered
  -r, --resolution string      How it was resolved

Built-in approaches:
  - binary_search
  - reverse_engineering
  - divide_conquer
  - backtracking
  - cause_elimination
  - program_slicing
  - log_analysis
  - static_analysis
  - root_cause_analysis
  - delta_debugging
  - fuzzing
  - incremental_testing
```

### frameworks

List all available frameworks (built-in and custom).

```bash
tinythoughts frameworks
```

## Output Formats

### JSON (default)

```json
{
  "operation": "mental_model",
  "status": "success",
  "data": {
    "modelName": "first_principles",
    "problem": "How to scale the system?",
    "steps": ["Break down components", "Identify bottlenecks"],
    "reasoning": "Start from fundamentals",
    "conclusion": "Focus on database optimization"
  },
  "session": {
    "sessionId": "cli-session-1234567890",
    "thoughts": [],
    "mentalModels": [...]
  }
}
```

### Text

```bash
tinythoughts --output text model first_principles "My problem"

Operation: mental_model
Status: success
```

## Session Management

Sessions are automatically saved to `~/.config/tinythoughts/session.json` after each command. This allows you to:

- Track your thinking process over time
- Review previous thoughts and models
- Build on earlier work
- Maintain context across CLI invocations

## Examples

### Problem-Solving Workflow

```bash
# Start with sequential thinking
tinythoughts think "Need to improve API performance" -n 1 -t 5 -x

# Apply first principles
tinythoughts model first_principles "What's causing slow API?" \
  -s "Measure current performance" \
  -s "Identify bottlenecks" \
  -s "Test hypotheses" \
  -r "Database queries are the issue" \
  -c "Need to optimize queries and add caching"

# Debug with binary search
tinythoughts debug binary_search "Slow database queries" \
  -s "Check if all queries are slow" \
  -s "Isolate specific query" \
  -f "JOIN on unindexed column" \
  -r "Added composite index"

# Continue thinking
tinythoughts think "Index improved performance by 10x" -n 2 -t 5 -x
```

### Using Custom Frameworks

```bash
# Create a custom framework
mkdir -p ~/.config/tinythoughts/frameworks
cat > ~/.config/tinythoughts/frameworks/premortem.yaml << 'EOF'
name: premortem
description: Imagine the project failed and work backwards
type: mental_model
steps:
  - name: imagine_failure
    description: Project has failed spectacularly
    prompt: Imagine it's 6 months from now and the project failed. What happened?
  - name: identify_causes
    description: List reasons for failure
    prompt: What were the top 3 reasons it failed?
  - name: assess_likelihood
    description: Evaluate probability
    prompt: How likely is each failure mode?
  - name: create_mitigations
    description: Design preventive measures
    prompt: What can you do now to prevent each failure mode?
prompts:
  conclusion: "What are the most critical risks to address?"
EOF

# Use it
tinythoughts model premortem "New product launch" \
  -s "Imagined complete failure" \
  -s "Identified lack of user testing as key risk" \
  -r "Need early user feedback" \
  -c "Schedule user testing sessions before launch"
```

## Architecture

tinythoughts is built with:

- **Cobra**: Command-line interface framework
- **Viper**: Configuration management (files, env vars, flags)
- **YAML**: Custom framework definitions
- **JSON**: Session state persistence

The architecture is intentionally minimal:

```
tinythoughts/
├── cmd/tinythoughts/     # CLI entry point
├── internal/
│   ├── types/           # Core data structures
│   ├── config/          # Configuration loading
│   └── engine/          # Thinking framework logic
└── frameworks/          # Example YAML frameworks
```

## Comparison to MCP Server

tinythoughts is derived from the clear-thought MCP server but optimized for CLI use:

| Feature | MCP Server | tinythoughts |
|---------|-----------|--------------|
| Transport | stdio/HTTP/SSE | Direct CLI |
| Integration | AI assistants | Command line |
| Configuration | Server config | YAML + env vars |
| Custom Frameworks | Not supported | YAML files |
| Session State | In-memory | Persistent JSON |
| Output | MCP protocol | JSON/text |

## Contributing

Contributions welcome! Please:

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

Apache License 2.0 - see LICENSE file for details.

## Acknowledgments

- Based on the [clear-thought MCP server](https://github.com/waldzellai/waldzell-mcp/tree/main/servers/server-clear-thought)
- Mental models inspired by [James Clear's guide](https://jamesclear.com/mental-models)
- Built with [Cobra](https://github.com/spf13/cobra) and [Viper](https://github.com/spf13/viper)
