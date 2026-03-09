package types

import "testing"

func TestThoughtData(t *testing.T) {
	thought := ThoughtData{
		Thought:           "Test thought",
		ThoughtNumber:     1,
		TotalThoughts:     3,
		NextThoughtNeeded: true,
	}

	if thought.Thought != "Test thought" {
		t.Errorf("Expected 'Test thought', got '%s'", thought.Thought)
	}
	if thought.ThoughtNumber != 1 {
		t.Errorf("Expected thought number 1, got %d", thought.ThoughtNumber)
	}
}

func TestMentalModelData(t *testing.T) {
	model := MentalModelData{
		ModelName:  "first_principles",
		Problem:    "Test problem",
		Steps:      []string{"Step 1", "Step 2"},
		Reasoning:  "Test reasoning",
		Conclusion: "Test conclusion",
	}

	if model.ModelName != "first_principles" {
		t.Errorf("Expected 'first_principles', got '%s'", model.ModelName)
	}
	if len(model.Steps) != 2 {
		t.Errorf("Expected 2 steps, got %d", len(model.Steps))
	}
}

func TestDebuggingApproachData(t *testing.T) {
	debug := DebuggingApproachData{
		ApproachName: "binary_search",
		Issue:        "Test issue",
		Steps:        []string{"Step 1"},
		Findings:     "Test findings",
		Resolution:   "Test resolution",
	}

	if debug.ApproachName != "binary_search" {
		t.Errorf("Expected 'binary_search', got '%s'", debug.ApproachName)
	}
}

func TestFramework(t *testing.T) {
	fw := Framework{
		Name:        "test_framework",
		Description: "Test description",
		Type:        "mental_model",
		Steps: []FrameworkStep{
			{Name: "step1", Description: "Test step", Prompt: "Test prompt"},
		},
		Prompts: map[string]string{"conclusion": "Test conclusion"},
	}

	if fw.Name != "test_framework" {
		t.Errorf("Expected 'test_framework', got '%s'", fw.Name)
	}
	if len(fw.Steps) != 1 {
		t.Errorf("Expected 1 step, got %d", len(fw.Steps))
	}
}

func TestSessionState(t *testing.T) {
	session := SessionState{
		SessionID:         "test-session",
		Thoughts:          []ThoughtData{},
		MentalModels:      []MentalModelData{},
		DebuggingSessions: []DebuggingApproachData{},
	}

	if session.SessionID != "test-session" {
		t.Errorf("Expected 'test-session', got '%s'", session.SessionID)
	}
	if len(session.Thoughts) != 0 {
		t.Errorf("Expected 0 thoughts, got %d", len(session.Thoughts))
	}
}
