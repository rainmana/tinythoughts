#!/bin/bash
set -e

echo "=========================================="
echo "  Comprehensive Framework Testing"
echo "=========================================="
echo ""

cd ~/Development/tinythoughts

echo "1. REGRET MINIMIZATION"
echo "----------------------"
echo "✓ GOOD USE: Life-changing decision"
./tinythoughts --output text model regret_minimization "Should I move to another country for my dream job?" \
  -s "At 80: Will value experiences and growth over comfort" \
  -s "Regret if yes: Might miss family, cultural adjustment hard" \
  -s "Regret if no: Never experienced living abroad, always wondered" \
  -c "Move - regret of not trying outweighs adjustment challenges"
echo ""

echo "✗ BAD USE: Trivial decision"
./tinythoughts --output text model regret_minimization "Should I buy the blue or red shirt?" \
  -s "At 80: Won't remember or care about shirt color" \
  -s "Regret if blue: None" \
  -s "Regret if red: None" \
  -c "Framework overkill - just pick one, doesn't matter at 80"
echo ""
echo "LESSON: Use for irreversible life decisions, not trivial choices"
echo ""

echo "2. EISENHOWER MATRIX"
echo "--------------------"
echo "✓ GOOD USE: Task prioritization"
./tinythoughts --output text model eisenhower_matrix "Manage my workday" \
  -s "Urgent+Important: Production outage" \
  -s "Important not urgent: Refactor technical debt" \
  -s "Urgent not important: Respond to Slack messages" \
  -s "Neither: Reorganize desktop icons" \
  -c "Fix outage now, schedule refactor, batch Slack, skip desktop"
echo ""

echo "✗ BAD USE: Strategic decision"
./tinythoughts --output text model eisenhower_matrix "Should we pivot our business model?" \
  -s "Urgent+Important: ???" \
  -s "Important not urgent: ???" \
  -c "Framework doesn't fit - business pivot isn't about urgency, it's about strategy"
echo ""
echo "LESSON: Use for task management, not strategic decisions"
echo ""

echo "3. PREMORTEM"
echo "------------"
echo "✓ GOOD USE: Project planning"
./tinythoughts --output text model premortem "Launch mobile app" \
  -s "Failure: App crashes, 1-star reviews, users uninstall" \
  -s "Causes: Didn't test on older devices, poor onboarding, bugs" \
  -s "Mitigations: Test on 5+ device types, add tutorial, QA sprint" \
  -c "Critical: Device testing and onboarding before launch"
echo ""

echo "✗ BAD USE: Past decision"
./tinythoughts --output text model premortem "Why did our last product fail?" \
  -s "Failure: Already happened" \
  -c "Framework is for FUTURE planning, not analyzing past failures - use postmortem instead"
echo ""
echo "LESSON: Use before projects start, not after they fail"
echo ""

echo "4. CIRCLE OF COMPETENCE"
echo "-----------------------"
echo "✓ GOOD USE: Investment decision"
./tinythoughts --output text model circle_of_competence "Invest in biotech startup" \
  -s "Inside: Software, startups" \
  -s "Edge: Basic biology" \
  -s "Outside: Drug development, FDA approval, clinical trials" \
  -s "Decision: Way outside my circle - need expert or pass" \
  -c "Pass on investment - too far outside competence"
echo ""

echo "✗ BAD USE: Learning new skill"
./tinythoughts --output text model circle_of_competence "Should I learn Python?" \
  -s "Inside: JavaScript" \
  -s "Outside: Python" \
  -c "Framework doesn't help - learning is about EXPANDING your circle, not staying in it"
echo ""
echo "LESSON: Use for decisions requiring expertise, not for learning"
echo ""

echo "5. SIX THINKING HATS"
echo "--------------------"
echo "✓ GOOD USE: Complex team decision"
./tinythoughts --output text model six_thinking_hats "Adopt remote-first policy" \
  -s "White: 30% want remote, 20% want office, 50% hybrid" \
  -s "Red: Managers anxious about productivity, employees excited" \
  -s "Black: Communication gaps, culture erosion, timezone issues" \
  -s "Yellow: Talent access, cost savings, flexibility" \
  -s "Green: Hybrid model, async-first communication, quarterly meetups" \
  -s "Blue: Hybrid with strong async culture is the path" \
  -c "Implement hybrid with async-first practices"
echo ""

echo "✗ BAD USE: Simple yes/no decision"
./tinythoughts --output text model six_thinking_hats "Should I upgrade my laptop?" \
  -s "White: Current laptop is 5 years old" \
  -s "Red: Want new laptop" \
  -c "Framework overkill - simple cost/benefit analysis would suffice"
echo ""
echo "LESSON: Use for complex problems needing multiple perspectives, not simple decisions"
echo ""

echo "6. 10/10/10 RULE"
echo "----------------"
echo "✓ GOOD USE: Emotional decision"
./tinythoughts --output text model ten_ten_ten "Should I quit after bad performance review?" \
  -s "10 minutes: Angry, want to quit immediately" \
  -s "10 months: Might regret hasty decision, could have improved" \
  -s "10 years: Won't matter much, but career gap might" \
  -c "Don't quit in anger - wait, improve, then decide rationally"
echo ""

echo "✗ BAD USE: Time-insensitive decision"
./tinythoughts --output text model ten_ten_ten "Which database should we use?" \
  -s "10 minutes: Doesn't matter" \
  -s "10 months: Still using it" \
  -s "10 years: Still using it" \
  -c "Framework doesn't help - technical decisions aren't about time perspective"
echo ""
echo "LESSON: Use for emotional/impulse decisions, not technical choices"
echo ""

echo "7. SWOT ANALYSIS"
echo "----------------"
echo "✓ GOOD USE: Strategic planning"
./tinythoughts --output text model swot_analysis "Expand to enterprise market" \
  -s "Strengths: Strong product, good reputation" \
  -s "Weaknesses: No enterprise sales team, limited support" \
  -s "Opportunities: Growing enterprise demand, weak competitors" \
  -s "Threats: Established players, long sales cycles" \
  -c "Hire enterprise sales lead, build support team, target mid-market first"
echo ""

echo "✗ BAD USE: Personal decision"
./tinythoughts --output text model swot_analysis "Should I ask someone on a date?" \
  -s "Strengths: Confident, good conversationalist" \
  -s "Weaknesses: Nervous" \
  -c "Framework doesn't fit - personal relationships aren't strategic business decisions"
echo ""
echo "LESSON: Use for business/competitive strategy, not personal life"
echo ""

echo "=========================================="
echo "  Summary: When to Use Each Framework"
echo "=========================================="
echo ""
echo "Regret Minimization: Life decisions (NOT trivial choices)"
echo "Eisenhower Matrix: Task management (NOT strategy)"
echo "Pre-Mortem: Future planning (NOT past analysis)"
echo "Circle of Competence: Expertise check (NOT learning)"
echo "Six Thinking Hats: Complex problems (NOT simple yes/no)"
echo "10/10/10 Rule: Emotional decisions (NOT technical)"
echo "SWOT Analysis: Business strategy (NOT personal life)"
echo ""
