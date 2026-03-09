package types

// ThoughtData represents a single thought in sequential thinking
type ThoughtData struct {
	Thought           string `json:"thought"`
	ThoughtNumber     int    `json:"thoughtNumber"`
	TotalThoughts     int    `json:"totalThoughts"`
	NextThoughtNeeded bool   `json:"nextThoughtNeeded"`
	IsRevision        bool   `json:"isRevision,omitempty"`
	RevisesThought    int    `json:"revisesThought,omitempty"`
	BranchFromThought int    `json:"branchFromThought,omitempty"`
	BranchID          string `json:"branchId,omitempty"`
	NeedsMoreThoughts bool   `json:"needsMoreThoughts,omitempty"`
}

// MentalModelData represents application of a mental model
type MentalModelData struct {
	ModelName  string   `json:"modelName"`
	Problem    string   `json:"problem"`
	Steps      []string `json:"steps"`
	Reasoning  string   `json:"reasoning"`
	Conclusion string   `json:"conclusion"`
}

// DebuggingApproachData represents a debugging session
type DebuggingApproachData struct {
	ApproachName string   `json:"approachName"`
	Issue        string   `json:"issue"`
	Steps        []string `json:"steps"`
	Findings     string   `json:"findings"`
	Resolution   string   `json:"resolution"`
}

// Framework represents a custom thinking framework loaded from YAML
type Framework struct {
	Name        string            `yaml:"name" json:"name"`
	Description string            `yaml:"description" json:"description"`
	Type        string            `yaml:"type" json:"type"` // sequential, mental_model, debugging, etc.
	Steps       []FrameworkStep   `yaml:"steps" json:"steps"`
	Prompts     map[string]string `yaml:"prompts" json:"prompts"`
}

// FrameworkStep represents a step in a custom framework
type FrameworkStep struct {
	Name        string `yaml:"name" json:"name"`
	Description string `yaml:"description" json:"description"`
	Prompt      string `yaml:"prompt" json:"prompt"`
}

// SessionState tracks thinking session state
type SessionState struct {
	SessionID     string                  `json:"sessionId"`
	Thoughts      []ThoughtData           `json:"thoughts"`
	MentalModels  []MentalModelData       `json:"mentalModels"`
	DebuggingSessions []DebuggingApproachData `json:"debuggingSessions"`
}

// Result represents the output of a thinking operation
type Result struct {
	Operation string      `json:"operation"`
	Status    string      `json:"status"`
	Data      interface{} `json:"data"`
	Session   *SessionState `json:"session,omitempty"`
}
