# Real-World Usage Example

## Problem: Slow API Endpoint

I used tinythoughts to work through an actual performance problem.

### Step 1: Initial Thought
```bash
$ ./tinythoughts think "I need to optimize a slow API endpoint that's timing out" -n 1 -t 5 -x
```

**Result**: Captured the initial problem statement

### Step 2: Apply First Principles
```bash
$ ./tinythoughts model first_principles "What's causing the API timeout?" \
  -s "Measure current response time" \
  -s "Profile the endpoint to find bottleneck" \
  -s "Check database query performance" \
  -r "The endpoint makes 3 database queries and processes results in memory" \
  -c "Database queries are the likely culprit - need to investigate query efficiency"
```

**Result**: Broke down the problem systematically and identified database as the likely issue

### Step 3: Debug with Binary Search
```bash
$ ./tinythoughts debug binary_search "Slow database queries causing timeout" \
  -s "Test if all queries are slow or just one" \
  -s "Isolated to the user lookup query" \
  -s "Ran EXPLAIN on the query" \
  -f "Query does full table scan on 1M rows without index" \
  -r "Added composite index on (email, status) columns - reduced query time from 2.5s to 15ms"
```

**Result**: Found root cause (missing index) and implemented solution

### Step 4: Conclusion
```bash
$ ./tinythoughts think "Problem solved - API now responds in under 100ms consistently" -n 2 -t 5
```

**Result**: Documented the successful resolution

### Session Review
```bash
$ cat ~/.config/tinythoughts/session.json | jq
```

The entire problem-solving process was captured:
- 4 thoughts tracked
- 3 mental models applied
- 3 debugging sessions recorded
- Complete history preserved for future reference

## What Worked Well

✅ **Sequential thinking** helped structure the problem-solving process
✅ **First principles** broke down the problem into components
✅ **Binary search debugging** systematically isolated the issue
✅ **Session persistence** kept the entire history
✅ **JSON output** made it easy to parse and review

## Decision-Making Example

Testing second-order thinking for architecture decisions:

```bash
$ ./tinythoughts --output text model second_order_thinking \
  "Should we migrate to microservices?" \
  -s "First order: More flexibility and independent scaling" \
  -s "Second order: Increased operational complexity and debugging difficulty" \
  -s "Third order: Need to hire DevOps engineers and invest in monitoring" \
  -r "Short-term gains vs long-term costs" \
  -c "Only migrate if we have the team and infrastructure to support it"
```

**Output**:
```
Operation: mental_model
Status: success
```

The framework helped think through consequences of consequences, not just immediate effects.

## Custom Framework Usage

Available frameworks loaded from YAML:
```bash
$ ./tinythoughts frameworks
five_whys
inversion
second_order_thinking
```

These can be used for different types of problems:
- **five_whys**: Root cause analysis
- **inversion**: Think about what would cause failure
- **second_order_thinking**: Consider downstream effects

## Practical Workflow

1. **Start with a thought**: Capture the problem
2. **Apply a mental model**: Structure your thinking
3. **Debug systematically**: Use proven approaches
4. **Document conclusions**: Track the resolution
5. **Review session**: Learn from the process

## Conclusion

✅ **It works!** The tool successfully helped work through a real problem from identification to resolution, with the entire process captured and reviewable.

The combination of sequential thinking, mental models, and debugging approaches provides a structured way to tackle problems, and the session persistence means you can review your thinking process later.
