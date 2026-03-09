# New Thinking Frameworks

This document describes the 7 new thinking frameworks added to tinythoughts.

## Decision-Making Frameworks

### 1. Regret Minimization (Jeff Bezos)

**When to use**: Major life decisions, career changes, risk-taking

**Key question**: "When I'm 80, will I regret not trying this?"

**Example**:
```bash
tinythoughts model regret_minimization "Should I quit my job to start a company?" \
  -s "At 80: Will care about taking risks and pursuing dreams" \
  -s "Regret if yes: Might fail, lose savings" \
  -s "Regret if no: Never tried, always wondered 'what if'" \
  -c "Take the leap - you'll regret not trying more than failing"
```

**Best for**: Irreversible decisions, entrepreneurship, career pivots

---

### 2. 10/10/10 Rule

**When to use**: Decisions where time perspective matters

**Key question**: "How will this feel in 10 minutes, 10 months, and 10 years?"

**Example**:
```bash
tinythoughts model ten_ten_ten "Should I send this angry email?" \
  -s "10 minutes: Feel satisfied" \
  -s "10 months: Regret damaging relationship" \
  -s "10 years: Won't matter, but reputation damage might" \
  -c "Don't send - short-term satisfaction, long-term cost"
```

**Best for**: Emotional decisions, relationship conflicts, impulse control

---

### 3. Eisenhower Matrix

**When to use**: Task prioritization, time management

**Key question**: "Is this urgent? Is this important?"

**Example**:
```bash
tinythoughts model eisenhower_matrix "Prioritize my tasks" \
  -s "Urgent+Important: Fix production bug" \
  -s "Important not urgent: Plan architecture" \
  -s "Urgent not important: Answer emails" \
  -s "Neither: Browse social media" \
  -c "Do bug now, schedule architecture, delegate emails, eliminate social media"
```

**Best for**: Daily planning, workload management, saying no

---

## Risk & Planning Frameworks

### 4. Pre-Mortem Analysis

**When to use**: Before starting projects, launching products

**Key question**: "If this failed, what would have caused it?"

**Example**:
```bash
tinythoughts model premortem "Launch new product feature" \
  -s "Imagine failure: Users hate it, churn increased" \
  -s "Causes: No user testing, poor onboarding, performance issues" \
  -s "Mitigations: Run user testing, add onboarding, optimize performance" \
  -c "Critical risks: 1) No user testing 2) Poor onboarding 3) Performance"
```

**Best for**: Project planning, risk identification, preventing failure

---

### 5. SWOT Analysis

**When to use**: Strategic planning, competitive analysis

**Key question**: "What are our strengths, weaknesses, opportunities, threats?"

**Example**:
```bash
tinythoughts model swot_analysis "Enter new market" \
  -s "Strengths: Strong brand, technical expertise" \
  -s "Weaknesses: No market presence, limited budget" \
  -s "Opportunities: Growing market, weak competitors" \
  -s "Threats: Regulatory changes, established players" \
  -c "Use brand strength to capture growth opportunity before regulations tighten"
```

**Best for**: Business strategy, market entry, competitive positioning

---

## Thinking & Analysis Frameworks

### 6. Six Thinking Hats (Edward de Bono)

**When to use**: Complex problems needing multiple perspectives

**Key question**: "What does each perspective reveal?"

**Example**:
```bash
tinythoughts model six_thinking_hats "Adopt AI coding assistant" \
  -s "White (Facts): Costs $20/user/month, 30% productivity gain" \
  -s "Red (Emotion): Developers excited, managers worried about dependency" \
  -s "Black (Risks): Code quality concerns, security risks, vendor lock-in" \
  -s "Yellow (Benefits): Faster development, fewer bugs, happier team" \
  -s "Green (Creative): Start with pilot team, measure impact, expand gradually" \
  -s "Blue (Overview): Benefits outweigh risks if we pilot first" \
  -c "Run 3-month pilot with one team, measure results, then decide"
```

**Best for**: Team decisions, complex problems, avoiding groupthink

---

### 7. Circle of Competence (Warren Buffett)

**When to use**: Evaluating decisions outside your expertise

**Key question**: "Do I truly understand this, or am I outside my circle?"

**Example**:
```bash
tinythoughts model circle_of_competence "Invest in cryptocurrency startup" \
  -s "Inside circle: Software development, startups" \
  -s "Edge of circle: Basic blockchain concepts" \
  -s "Outside circle: Crypto economics, regulatory landscape, tokenomics" \
  -s "Decision: Outside my circle - need expert help or pass" \
  -c "Either bring in crypto expert advisor, or pass on this investment"
```

**Best for**: Investment decisions, hiring, knowing when to get help

---

## Quick Reference

| Framework | Best For | Time to Use |
|-----------|----------|-------------|
| Regret Minimization | Life decisions | 5-10 min |
| 10/10/10 Rule | Emotional decisions | 2-5 min |
| Eisenhower Matrix | Daily prioritization | 5-10 min |
| Pre-Mortem | Project planning | 15-30 min |
| SWOT Analysis | Strategic planning | 20-45 min |
| Six Thinking Hats | Complex problems | 15-30 min |
| Circle of Competence | Expertise check | 5-10 min |

## Combining Frameworks

These frameworks work well together:

1. **Pre-Mortem + SWOT**: Identify risks, then analyze strategic position
2. **Circle of Competence + Regret Minimization**: Check expertise, then decide if worth learning
3. **Eisenhower Matrix + 10/10/10**: Prioritize tasks, then validate with time perspective
4. **Six Thinking Hats + Pre-Mortem**: Multiple perspectives on what could go wrong

## Tips for Effective Use

1. **Be honest**: Frameworks only work if you're truthful with yourself
2. **Write it down**: Use tinythoughts to capture your reasoning
3. **Review later**: Check your session.json to see how decisions played out
4. **Combine frameworks**: Use multiple perspectives on important decisions
5. **Share with team**: Use frameworks to structure team discussions

## Creating Your Own

See the YAML files in `frameworks/` for examples. The structure is:

```yaml
name: your_framework
description: What it does
type: mental_model
steps:
  - name: step_name
    description: What to consider
    prompt: Question to ask
prompts:
  conclusion: "Final synthesis question"
```

Add your custom frameworks to `~/.config/tinythoughts/frameworks/` and they'll be automatically loaded.
