#!/bin/bash
set -e

echo "=== Functional Testing ==="

# Clean test environment
rm -f /tmp/test-session.json
export TINYTHOUGHTS_SESSION_FILE=/tmp/test-session.json
export TINYTHOUGHTS_OUTPUT_FORMAT=json

echo "✓ Test 1: Sequential thinking"
./tinythoughts think "Test thought" -n 1 -t 1 | jq -e '.operation == "sequential_thinking"' > /dev/null

echo "✓ Test 2: Mental model"
./tinythoughts model first_principles "Test problem" -s "Step 1" -r "Reasoning" -c "Conclusion" | jq -e '.operation == "mental_model"' > /dev/null

echo "✓ Test 3: Debugging approach"
./tinythoughts debug binary_search "Test issue" -s "Step 1" -f "Findings" -r "Resolution" | jq -e '.operation == "debugging_approach"' > /dev/null

echo "✓ Test 4: Session persistence"
test -f /tmp/test-session.json || (echo "Session file not created" && exit 1)

echo "✓ Test 5: Session contains data"
jq -e '.thoughts | length > 0' /tmp/test-session.json > /dev/null

echo "✓ Test 6: Text output format"
export TINYTHOUGHTS_OUTPUT_FORMAT=text
OUTPUT=$(./tinythoughts think "Test" -n 1 -t 1)
echo "$OUTPUT" | grep -q "Operation: sequential_thinking" || (echo "Text output failed" && exit 1)

echo "✓ Test 7: Frameworks listing"
mkdir -p /tmp/test-frameworks
cat > /tmp/test-frameworks/test.yaml << 'YAML'
name: test_framework
description: Test
type: mental_model
steps:
  - name: step1
    description: Test step
    prompt: Test prompt
YAML
export TINYTHOUGHTS_FRAMEWORKS_DIR=/tmp/test-frameworks
./tinythoughts frameworks | grep -q "test_framework" || (echo "Framework loading failed" && exit 1)

echo ""
echo "=== All functional tests passed! ==="
