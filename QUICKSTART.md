# tinythoughts Quick Start Guide

Get up and running with tinythoughts in 5 minutes.

## Installation

```bash
cd ~/Development/tinythoughts
make build
make install
```

Or manually:

```bash
go build -o tinythoughts ./cmd/tinythoughts
sudo mv tinythoughts /usr/local/bin/
```

## Setup Configuration

```bash
make setup-config
```

This creates:
- `~/.config/tinythoughts/frameworks/` with example frameworks
- `~/.config/tinythoughts/tinythoughts.yaml` with default config

## Your First Commands

### 1. Sequential Thinking

```bash
tinythoughts think "I need to solve the database performance issue" -n 1 -t 3 -x
```

Output:
```json
{
  "operation": "sequential_thinking",
  "status": "success",
  "data": {
    "thought": "I need to solve the database performance issue",
    "thoughtNumber": 1,
    "totalThoughts": 3,
    "nextThoughtNeeded": true
  }
}
```

### 2. Apply a Mental Model

```bash
tinythoughts model first_principles "How to scale our API?" \
  -s "Break down the request flow" \
  -s "Identify bottlenecks" \
  -r "Database is the constraint" \
  -c "Add read replicas and caching"
```

### 3. Debugging Approach

```bash
tinythoughts debug binary_search "API timeout on /users endpoint" \
  -s "Check if all endpoints are slow" \
  -s "Isolate to /users endpoint" \
  -s "Profile the database query" \
  -f "N+1 query problem" \
  -r "Added eager loading"
```

### 4. List Available Frameworks

```bash
tinythoughts frameworks
```

Output:
```
five_whys
inversion
second_order_thinking
```

## Using Custom Frameworks

### Create a Framework

```bash
cat > ~/.config/tinythoughts/frameworks/premortem.yaml << 'YAML'
name: premortem
description: Imagine failure and work backwards
type: mental_model
steps:
  - name: imagine_failure
    description: Project has failed
    prompt: It's 6 months from now. The project failed. What happened?
  - name: identify_causes
    description: List failure reasons
    prompt: What were the top 3 reasons it failed?
  - name: create_mitigations
    description: Design preventive measures
    prompt: What can you do now to prevent each failure?
prompts:
  conclusion: "What are the most critical risks?"
YAML
```

### Use Your Framework

```bash
tinythoughts model premortem "New product launch" \
  -s "Imagined complete failure" \
  -s "No user testing was done" \
  -r "Need early feedback" \
  -c "Schedule user testing sessions"
```

## Configuration Options

### Via Environment Variables

```bash
export TINYTHOUGHTS_OUTPUT_FORMAT=text
export TINYTHOUGHTS_FRAMEWORKS_DIR=~/my-frameworks
tinythoughts think "Test" -n 1 -t 1
```

### Via Config File

Edit `~/.config/tinythoughts/tinythoughts.yaml`:

```yaml
frameworks_dir: ~/my-frameworks
session_file: ~/my-session.json
output_format: text
```

### Via CLI Flags

```bash
tinythoughts --output text think "Test" -n 1 -t 1
```

## Session Persistence

Sessions are automatically saved to `~/.config/tinythoughts/session.json`.

View your session:

```bash
cat ~/.config/tinythoughts/session.json | jq
```

Example session:

```json
{
  "sessionId": "cli-session-1234567890",
  "thoughts": [
    {
      "thought": "First thought",
      "thoughtNumber": 1,
      "totalThoughts": 3,
      "nextThoughtNeeded": true
    }
  ],
  "mentalModels": [
    {
      "modelName": "first_principles",
      "problem": "How to scale?",
      "steps": ["Break down", "Identify bottlenecks"],
      "reasoning": "Start from fundamentals",
      "conclusion": "Focus on database"
    }
  ],
  "debuggingSessions": []
}
```

## Common Workflows

### Problem-Solving Session

```bash
# 1. Initial thought
tinythoughts think "Need to improve API performance" -n 1 -t 5 -x

# 2. Apply first principles
tinythoughts model first_principles "What's causing slow API?" \
  -s "Measure current performance" \
  -s "Identify bottlenecks" \
  -r "Database queries are slow" \
  -c "Optimize queries and add caching"

# 3. Debug with binary search
tinythoughts debug binary_search "Slow database queries" \
  -s "Check if all queries are slow" \
  -s "Isolate specific query" \
  -f "JOIN on unindexed column" \
  -r "Added composite index"

# 4. Continue thinking
tinythoughts think "Index improved performance by 10x" -n 2 -t 5 -x

# 5. Review session
cat ~/.config/tinythoughts/session.json | jq
```

### Decision Making

```bash
# Use inversion
tinythoughts model inversion "Should we migrate to microservices?" \
  -s "What would guarantee failure?" \
  -s "Identified: poor service boundaries, no monitoring" \
  -r "Need clear boundaries and observability" \
  -c "Start with one service, build monitoring first"

# Use second-order thinking
tinythoughts model second_order_thinking "Migrate to microservices" \
  -s "First order: more flexibility" \
  -s "Second order: increased complexity" \
  -s "Third order: need more DevOps resources" \
  -r "Consider long-term costs" \
  -c "Evaluate if benefits outweigh complexity"
```

## Tips

1. **Use text output for quick checks**: `--output text`
2. **Chain commands in scripts**: Output is JSON by default
3. **Create custom frameworks**: YAML files in frameworks directory
4. **Review sessions**: Check `session.json` to see your thinking process
5. **Use environment variables**: Set defaults for your workflow

## Next Steps

- Read the full [README.md](README.md) for detailed documentation
- Check [ARCHITECTURE.md](ARCHITECTURE.md) to understand the design
- Create custom frameworks for your specific thinking patterns
- Integrate into your development workflow

## Getting Help

```bash
tinythoughts --help
tinythoughts think --help
tinythoughts model --help
tinythoughts debug --help
```

## Troubleshooting

### Command not found

```bash
# Check if installed
which tinythoughts

# If not, install
cd ~/Development/tinythoughts
make install
```

### No frameworks showing

```bash
# Setup frameworks
make setup-config

# Or manually
mkdir -p ~/.config/tinythoughts/frameworks
cp frameworks/*.yaml ~/.config/tinythoughts/frameworks/
```

### Session not persisting

```bash
# Check session file location
cat ~/.config/tinythoughts/tinythoughts.yaml

# Ensure directory exists
mkdir -p ~/.config/tinythoughts
```
