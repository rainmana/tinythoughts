# Framework Selection Guide

## When to Use Each Framework (and When NOT To)

This guide shows you which framework to use for different types of decisions, with real examples of good and bad use cases.

---

## 1. Regret Minimization (Jeff Bezos)

### ✅ USE FOR:
- Life-changing decisions
- Career pivots
- Major commitments
- Irreversible choices

### ❌ DON'T USE FOR:
- Trivial daily choices
- Reversible decisions
- Low-stakes options

### Good Example:
```bash
tinythoughts model regret_minimization "Should I move to another country for my dream job?"
# At 80, you'll care about experiences and growth
# Conclusion: Move - regret of not trying outweighs challenges
```

### Bad Example:
```bash
tinythoughts model regret_minimization "Should I buy the blue or red shirt?"
# At 80, you won't remember or care about shirt color
# Conclusion: Framework overkill - just pick one
```

**Key Insight**: If you won't remember it at 80, don't use this framework.

---

## 2. Eisenhower Matrix

### ✅ USE FOR:
- Daily task prioritization
- Workload management
- Time management
- Deciding what to delegate/eliminate

### ❌ DON'T USE FOR:
- Strategic business decisions
- Long-term planning
- Decisions without urgency component

### Good Example:
```bash
tinythoughts model eisenhower_matrix "Manage my workday"
# Urgent+Important: Production outage
# Important not urgent: Refactor technical debt
# Conclusion: Fix outage now, schedule refactor
```

### Bad Example:
```bash
tinythoughts model eisenhower_matrix "Should we pivot our business model?"
# Business pivots aren't about urgency vs importance
# Conclusion: Wrong framework - use SWOT or strategic analysis
```

**Key Insight**: Only works when urgency is a meaningful dimension.

---

## 3. Pre-Mortem Analysis

### ✅ USE FOR:
- Before starting projects
- Launch planning
- Risk identification
- Preventing failure

### ❌ DON'T USE FOR:
- Analyzing past failures (use postmortem)
- Ongoing projects (too late)
- Low-risk activities

### Good Example:
```bash
tinythoughts model premortem "Launch mobile app"
# Failure: App crashes, 1-star reviews
# Causes: Didn't test on older devices, poor onboarding
# Conclusion: Test on 5+ device types before launch
```

### Bad Example:
```bash
tinythoughts model premortem "Why did our last product fail?"
# Failure already happened
# Conclusion: Use postmortem for past, premortem for future
```

**Key Insight**: Premortem is preventive, not diagnostic.

---

## 4. Circle of Competence (Warren Buffett)

### ✅ USE FOR:
- Investment decisions
- Hiring for expertise
- Knowing when to get help
- Avoiding overconfidence

### ❌ DON'T USE FOR:
- Learning new skills (that's expanding your circle)
- Areas you want to develop
- Growth opportunities

### Good Example:
```bash
tinythoughts model circle_of_competence "Invest in biotech startup"
# Inside: Software, startups
# Outside: Drug development, FDA approval, clinical trials
# Conclusion: Pass - too far outside competence
```

### Bad Example:
```bash
tinythoughts model circle_of_competence "Should I learn Python?"
# Learning is about EXPANDING your circle
# Conclusion: Wrong framework - just learn it
```

**Key Insight**: Use to avoid mistakes, not to avoid growth.

---

## 5. Six Thinking Hats (Edward de Bono)

### ✅ USE FOR:
- Complex team decisions
- Problems with multiple stakeholders
- Avoiding groupthink
- Comprehensive analysis

### ❌ DON'T USE FOR:
- Simple yes/no decisions
- Solo decisions
- Time-sensitive choices (too slow)

### Good Example:
```bash
tinythoughts model six_thinking_hats "Adopt remote-first policy"
# White: 30% want remote, 50% hybrid
# Red: Managers anxious, employees excited
# Black: Communication gaps, culture erosion
# Yellow: Talent access, cost savings
# Green: Hybrid model with async-first
# Conclusion: Implement hybrid with strong async culture
```

### Bad Example:
```bash
tinythoughts model six_thinking_hats "Should I upgrade my laptop?"
# Simple cost/benefit decision
# Conclusion: Framework overkill
```

**Key Insight**: Complexity of framework should match complexity of problem.

---

## 6. 10/10/10 Rule

### ✅ USE FOR:
- Emotional decisions
- Impulse control
- Relationship conflicts
- Anger management

### ❌ DON'T USE FOR:
- Technical decisions
- Time-insensitive choices
- Decisions where time perspective doesn't matter

### Good Example:
```bash
tinythoughts model ten_ten_ten "Should I quit after bad performance review?"
# 10 minutes: Angry, want to quit
# 10 months: Might regret hasty decision
# 10 years: Career gap might matter
# Conclusion: Don't quit in anger - wait and improve
```

### Bad Example:
```bash
tinythoughts model ten_ten_ten "Which database should we use?"
# Technical decisions aren't about time perspective
# Conclusion: Use technical evaluation framework instead
```

**Key Insight**: Best for cooling down emotional reactions.

---

## 7. SWOT Analysis

### ✅ USE FOR:
- Business strategy
- Competitive analysis
- Market entry decisions
- Strategic planning

### ❌ DON'T USE FOR:
- Personal life decisions
- Non-competitive situations
- Simple choices

### Good Example:
```bash
tinythoughts model swot_analysis "Expand to enterprise market"
# Strengths: Strong product, good reputation
# Weaknesses: No enterprise sales team
# Opportunities: Growing demand, weak competitors
# Threats: Established players, long sales cycles
# Conclusion: Hire sales lead, target mid-market first
```

### Bad Example:
```bash
tinythoughts model swot_analysis "Should I ask someone on a date?"
# Personal relationships aren't strategic business decisions
# Conclusion: Just ask them out
```

**Key Insight**: SWOT is for competitive/strategic contexts.

---

## Quick Decision Tree

```
Is it a life-changing decision?
├─ Yes → Regret Minimization
└─ No ↓

Is it about task prioritization?
├─ Yes → Eisenhower Matrix
└─ No ↓

Is it a future project with risks?
├─ Yes → Pre-Mortem
└─ No ↓

Are you outside your expertise?
├─ Yes → Circle of Competence
└─ No ↓

Is it complex with multiple perspectives?
├─ Yes → Six Thinking Hats
└─ No ↓

Are you making an emotional decision?
├─ Yes → 10/10/10 Rule
└─ No ↓

Is it a strategic business decision?
├─ Yes → SWOT Analysis
└─ No → Use simpler analysis
```

---

## Combining Frameworks

Some decisions benefit from multiple frameworks:

### Example: Starting a Company

1. **Regret Minimization**: Will I regret not trying?
2. **Circle of Competence**: Do I understand this market?
3. **Pre-Mortem**: What could cause failure?
4. **SWOT**: What's my competitive position?

### Example: Hiring Decision

1. **Circle of Competence**: Do I understand this role?
2. **10/10/10**: Am I rushing due to pressure?
3. **Pre-Mortem**: What if this hire fails?

---

## Red Flags: You're Using the Wrong Framework

- **Regret Minimization** for what to eat for lunch
- **Eisenhower Matrix** for strategic planning
- **Pre-Mortem** for analyzing past failures
- **Circle of Competence** to avoid learning
- **Six Thinking Hats** for simple yes/no
- **10/10/10** for technical architecture
- **SWOT** for personal relationships

---

## Framework Complexity vs Decision Complexity

```
Simple Decision → Simple Framework
Complex Decision → Complex Framework

Mismatch = Wasted Time or Poor Decision
```

**Simple frameworks**: 10/10/10, Eisenhower Matrix  
**Medium frameworks**: Regret Minimization, Circle of Competence, Pre-Mortem  
**Complex frameworks**: Six Thinking Hats, SWOT Analysis

---

## Testing Your Framework Choice

Ask yourself:
1. Does this framework match my decision type?
2. Will it reveal insights I wouldn't see otherwise?
3. Is the time investment worth it?
4. Am I using it because it's appropriate or because it's familiar?

If you're forcing a framework to fit, you're probably using the wrong one.

---

## Summary Table

| Framework | Best For | Avoid For | Time |
|-----------|----------|-----------|------|
| Regret Minimization | Life decisions | Trivial choices | 5-10 min |
| Eisenhower Matrix | Task management | Strategy | 5-10 min |
| Pre-Mortem | Future planning | Past analysis | 15-30 min |
| Circle of Competence | Expertise check | Learning | 5-10 min |
| Six Thinking Hats | Complex problems | Simple yes/no | 15-30 min |
| 10/10/10 Rule | Emotional decisions | Technical | 2-5 min |
| SWOT Analysis | Business strategy | Personal life | 20-45 min |

---

## The Meta-Framework

**Before choosing a framework, ask:**
1. What type of decision is this?
2. What's the key dimension? (time, expertise, risk, emotion, strategy)
3. How complex is it?
4. How much time do I have?

Then pick the framework that matches those dimensions.

**Remember**: No framework is better than the wrong framework.
