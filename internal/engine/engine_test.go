package engine

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/rainmana/tinythoughts/internal/types"
)

func TestNew(t *testing.T) {
	eng := New("test-session")
	if eng.session.SessionID != "test-session" {
		t.Errorf("Expected session ID 'test-session', got '%s'", eng.session.SessionID)
	}
	if len(eng.frameworks) != 0 {
		t.Errorf("Expected 0 frameworks, got %d", len(eng.frameworks))
	}
}

func TestSequentialThinking(t *testing.T) {
	eng := New("test-session")
	data := types.ThoughtData{
		Thought:           "Test thought",
		ThoughtNumber:     1,
		TotalThoughts:     1,
		NextThoughtNeeded: false,
	}

	result := eng.SequentialThinking(data)
	if result.Operation != "sequential_thinking" {
		t.Errorf("Expected operation 'sequential_thinking', got '%s'", result.Operation)
	}
	if result.Status != "success" {
		t.Errorf("Expected status 'success', got '%s'", result.Status)
	}
	if len(eng.session.Thoughts) != 1 {
		t.Errorf("Expected 1 thought in session, got %d", len(eng.session.Thoughts))
	}
}

func TestMentalModel(t *testing.T) {
	eng := New("test-session")
	data := types.MentalModelData{
		ModelName:  "first_principles",
		Problem:    "Test problem",
		Steps:      []string{"Step 1"},
		Reasoning:  "Test reasoning",
		Conclusion: "Test conclusion",
	}

	result := eng.MentalModel(data)
	if result.Operation != "mental_model" {
		t.Errorf("Expected operation 'mental_model', got '%s'", result.Operation)
	}
	if len(eng.session.MentalModels) != 1 {
		t.Errorf("Expected 1 mental model in session, got %d", len(eng.session.MentalModels))
	}
}

func TestDebuggingApproach(t *testing.T) {
	eng := New("test-session")
	data := types.DebuggingApproachData{
		ApproachName: "binary_search",
		Issue:        "Test issue",
		Steps:        []string{"Step 1"},
		Findings:     "Test findings",
		Resolution:   "Test resolution",
	}

	result := eng.DebuggingApproach(data)
	if result.Operation != "debugging_approach" {
		t.Errorf("Expected operation 'debugging_approach', got '%s'", result.Operation)
	}
	if len(eng.session.DebuggingSessions) != 1 {
		t.Errorf("Expected 1 debugging session, got %d", len(eng.session.DebuggingSessions))
	}
}

func TestLoadFrameworks(t *testing.T) {
	tmpDir := t.TempDir()
	yamlContent := `name: test_framework
description: Test framework
type: mental_model
steps:
  - name: step1
    description: Test step
    prompt: Test prompt
prompts:
  conclusion: Test conclusion
`
	err := os.WriteFile(filepath.Join(tmpDir, "test.yaml"), []byte(yamlContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test framework: %v", err)
	}

	eng := New("test-session")
	err = eng.LoadFrameworks(tmpDir)
	if err != nil {
		t.Fatalf("LoadFrameworks failed: %v", err)
	}

	if len(eng.frameworks) != 1 {
		t.Errorf("Expected 1 framework, got %d", len(eng.frameworks))
	}

	fw, err := eng.GetFramework("test_framework")
	if err != nil {
		t.Errorf("GetFramework failed: %v", err)
	}
	if fw.Name != "test_framework" {
		t.Errorf("Expected framework name 'test_framework', got '%s'", fw.Name)
	}
}

func TestSaveLoadSession(t *testing.T) {
	tmpFile := filepath.Join(t.TempDir(), "session.json")

	eng := New("test-session")
	eng.SequentialThinking(types.ThoughtData{
		Thought:           "Test",
		ThoughtNumber:     1,
		TotalThoughts:     1,
		NextThoughtNeeded: false,
	})

	err := eng.SaveSession(tmpFile)
	if err != nil {
		t.Fatalf("SaveSession failed: %v", err)
	}

	eng2 := New("test-session")
	err = eng2.LoadSession(tmpFile)
	if err != nil {
		t.Fatalf("LoadSession failed: %v", err)
	}

	if len(eng2.session.Thoughts) != 1 {
		t.Errorf("Expected 1 thought after loading, got %d", len(eng2.session.Thoughts))
	}
}

func TestListFrameworks(t *testing.T) {
	eng := New("test-session")
	eng.frameworks["fw1"] = &types.Framework{Name: "fw1"}
	eng.frameworks["fw2"] = &types.Framework{Name: "fw2"}

	names := eng.ListFrameworks()
	if len(names) != 2 {
		t.Errorf("Expected 2 frameworks, got %d", len(names))
	}
}

func TestGetFrameworkNotFound(t *testing.T) {
	eng := New("test-session")
	_, err := eng.GetFramework("nonexistent")
	if err == nil {
		t.Error("Expected error for nonexistent framework, got nil")
	}
}
