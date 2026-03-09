# LLM Usage Guide for tinythoughts

This guide helps LLMs (and AI assistants) effectively use tinythoughts to help users with decision-making and problem-solving.

## Quick Discovery

### List all available frameworks
```bash
tinythoughts frameworks -v
```

This shows all frameworks with descriptions, helping you choose the right one.

### Get detailed information about a framework
```bash
tinythoughts describe [framework-name]
```

This shows:
- Framework type
- Description
- All steps with prompts
- Conclusion question

## Decision Flow for LLMs

### 1. Understand the User's Problem

Ask yourself:
- What type of decision is this?
- What's the key dimension? (time, risk, emotion, strategy, expertise)
- How complex is it?

### 2. Choose the Right Framework

Use this mapping:

| User Says | Framework to Use |
|-----------|------------------|
| "Should I quit my job?" | regret_minimization |
| "Help me prioritize tasks" | eisenhower_matrix |
| "What could go wrong?" | premortem |
| "I don't know if I understand this" | circle_of_competence |
| "Let's look at this from all angles" | six_thinking_hats |
| "Should I send this angry email?" | ten_ten_ten |
| "Analyze our competitive position" | swot_analysis |
| "Why did this fail?" | five_whys |
| "What if we did the opposite?" | inversion |
| "What are the long-term effects?" | second_order_thinking |

### 3. Verify Framework Appropriateness

Before using a framework, check:
```bash
tinythoughts describe [framework-name]
```

Read the steps and prompts. If they don't match the user's problem, choose a different framework.

### 4. Execute the Framework

```bash
tinythoughts model [framework-name] "[problem]" \
  -s "[step 1 answer]" \
  -s "[step 2 answer]" \
  -r "[reasoning]" \
  -c "[conclusion]"
```

### 5. Interpret Results

The output includes:
- The framework applied
- Steps taken
- Reasoning
- Conclusion
- Session history

Use this to provide structured advice to the user.

## Example Workflow

### User: "I'm thinking about starting a company but I'm scared"

**Step 1**: Identify this as a life-changing decision with fear/emotion

**Step 2**: Choose framework
```bash
tinythoughts frameworks -v | grep -i regret
# Output: regret_minimization - Jeff Bezos's framework - minimize regret when you're 80 years old
```

**Step 3**: Verify appropriateness
```bash
tinythoughts describe regret_minimization
# Read the steps - yes, this fits!
```

**Step 4**: Execute
```bash
tinythoughts model regret_minimization "Should I quit my job to start a company?" \
  -s "At 80, I'll value taking risks and pursuing dreams over playing it safe" \
  -s "Regret if yes: Might fail, lose savings, career setback" \
  -s "Regret if no: Never tried, always wondered 'what if', played it safe" \
  -s "Compare: Regret of not trying feels worse than trying and failing" \
  -c "Start the company - you'll regret not trying more than trying and failing"
```

**Step 5**: Present to user
"Based on the Regret Minimization framework (used by Jeff Bezos for major decisions), when you're 80 years old, you'll likely regret not trying more than trying and failing. The fear of failure is temporary, but the regret of not pursuing your dreams could last a lifetime."

## Common Patterns

### Pattern 1: User is Overwhelmed
```bash
# Use Eisenhower Matrix to prioritize
tinythoughts model eisenhower_matrix "Help me prioritize" \
  -s "Urgent+Important: [critical tasks]" \
  -s "Important not urgent: [strategic work]" \
  -s "Urgent not important: [interruptions]" \
  -s "Neither: [time wasters]" \
  -c "Focus on urgent+important, schedule important work, delegate/eliminate rest"
```

### Pattern 2: User is Making Emotional Decision
```bash
# Use 10/10/10 to add perspective
tinythoughts model ten_ten_ten "Should I [emotional decision]?" \
  -s "10 minutes: [immediate feeling]" \
  -s "10 months: [medium-term impact]" \
  -s "10 years: [long-term perspective]" \
  -c "[decision based on time perspective]"
```

### Pattern 3: User Needs Risk Analysis
```bash
# Use Pre-Mortem before starting
tinythoughts model premortem "Launch [project]" \
  -s "Imagine failure: [what went wrong]" \
  -s "Causes: [specific reasons]" \
  -s "Mitigations: [preventive actions]" \
  -c "Top 3 risks to address: [list]"
```

### Pattern 4: User Outside Their Expertise
```bash
# Use Circle of Competence
tinythoughts model circle_of_competence "Should I [decision outside expertise]?" \
  -s "Inside circle: [what they know]" \
  -s "Outside circle: [what they don't know]" \
  -s "Decision: [stay in circle or get help]" \
  -c "[recommendation with expert help if needed]"
```

## Red Flags: Wrong Framework

If you see these, choose a different framework:

- Using **regret_minimization** for "which shirt to buy"
- Using **eisenhower_matrix** for "should we pivot the business"
- Using **premortem** for "why did this already fail"
- Using **circle_of_competence** for "should I learn X"
- Using **six_thinking_hats** for simple yes/no
- Using **ten_ten_ten** for technical architecture
- Using **swot_analysis** for personal relationships

## Multi-Framework Approach

For complex decisions, use multiple frameworks:

```bash
# 1. Check expertise
tinythoughts model circle_of_competence "[decision]" ...

# 2. If outside expertise, get help. If inside, continue:
tinythoughts model regret_minimization "[decision]" ...

# 3. Identify risks
tinythoughts model premortem "[decision]" ...

# 4. Make final decision with all perspectives
```

## Output Formats

### JSON (default)
```bash
tinythoughts model [framework] "[problem]" ...
# Returns structured JSON with full session history
```

### Text (human-readable)
```bash
tinythoughts --output text model [framework] "[problem]" ...
# Returns simple text output
```

Use JSON when you need to parse results, text when presenting to users.

## Session Persistence

All decisions are saved to `~/.config/tinythoughts/session.json`

You can reference past decisions:
```bash
cat ~/.config/tinythoughts/session.json | jq '.mentalModels[-1]'
# Shows the last mental model used
```

## Best Practices for LLMs

1. **Always verify framework appropriateness** with `describe` before using
2. **Match framework complexity to problem complexity**
3. **Use verbose mode** (`-v`) to see descriptions
4. **Combine frameworks** for complex decisions
5. **Explain your framework choice** to the user
6. **Show the reasoning**, not just the conclusion
7. **Reference the framework by name** so users can learn it

## Example: Full LLM Response

**User**: "I'm not sure if I should take this new job offer"

**LLM Response**:
"This is a significant career decision. Let me help you think through it using the Regret Minimization framework, which Jeff Bezos used for major life decisions.

[Runs tinythoughts command]

Based on this framework, here's the analysis:

When you're 80 years old looking back:
- If you take the job and it doesn't work out, you'll know you tried
- If you don't take it, you might always wonder "what if"

The key question is: which regret would be harder to live with?

From the framework's perspective: [conclusion]

Would you like to also analyze this from a risk perspective using a Pre-Mortem, or check if this role is within your circle of competence?"

## Quick Reference Commands

```bash
# Discovery
tinythoughts frameworks -v              # List all with descriptions
tinythoughts describe [framework]       # Detailed framework info

# Execution
tinythoughts model [framework] "[problem]" -s "[step]" -r "[reasoning]" -c "[conclusion]"
tinythoughts think "[thought]" -n 1 -t 3 -x
tinythoughts debug [approach] "[issue]" -s "[step]" -f "[findings]" -r "[resolution]"

# Output
tinythoughts --output text [command]    # Human-readable
tinythoughts --output json [command]    # Structured (default)
```

## Error Handling

If a framework doesn't exist:
```bash
tinythoughts describe nonexistent
# Error: Framework not found: nonexistent
```

Check available frameworks first:
```bash
tinythoughts frameworks
```

## Summary

1. Use `frameworks -v` to discover
2. Use `describe` to verify
3. Execute with appropriate framework
4. Explain your reasoning to the user
5. Offer to use additional frameworks if needed

The goal is to help users make better decisions through structured thinking, not to force every decision through a framework.
