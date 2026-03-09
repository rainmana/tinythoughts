package engine

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/waldzellai/tinythoughts/internal/types"
	"gopkg.in/yaml.v3"
)

type Engine struct {
	session    *types.SessionState
	frameworks map[string]*types.Framework
}

func New(sessionID string) *Engine {
	return &Engine{
		session: &types.SessionState{
			SessionID:         sessionID,
			Thoughts:          []types.ThoughtData{},
			MentalModels:      []types.MentalModelData{},
			DebuggingSessions: []types.DebuggingApproachData{},
		},
		frameworks: make(map[string]*types.Framework),
	}
}

func (e *Engine) LoadFrameworks(dir string) error {
	expanded := os.ExpandEnv(dir)
	entries, err := os.ReadDir(expanded)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".yaml" {
			continue
		}

		path := filepath.Join(expanded, entry.Name())
		data, err := os.ReadFile(path)
		if err != nil {
			continue
		}

		var fw types.Framework
		if err := yaml.Unmarshal(data, &fw); err != nil {
			continue
		}

		e.frameworks[fw.Name] = &fw
	}

	return nil
}

func (e *Engine) SequentialThinking(data types.ThoughtData) types.Result {
	e.session.Thoughts = append(e.session.Thoughts, data)
	
	return types.Result{
		Operation: "sequential_thinking",
		Status:    "success",
		Data:      data,
		Session:   e.session,
	}
}

func (e *Engine) MentalModel(data types.MentalModelData) types.Result {
	e.session.MentalModels = append(e.session.MentalModels, data)
	
	return types.Result{
		Operation: "mental_model",
		Status:    "success",
		Data:      data,
		Session:   e.session,
	}
}

func (e *Engine) DebuggingApproach(data types.DebuggingApproachData) types.Result {
	e.session.DebuggingSessions = append(e.session.DebuggingSessions, data)
	
	return types.Result{
		Operation: "debugging_approach",
		Status:    "success",
		Data:      data,
		Session:   e.session,
	}
}

func (e *Engine) SaveSession(path string) error {
	expanded := os.ExpandEnv(path)
	dir := filepath.Dir(expanded)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(e.session, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(expanded, data, 0644)
}

func (e *Engine) LoadSession(path string) error {
	expanded := os.ExpandEnv(path)
	data, err := os.ReadFile(expanded)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	return json.Unmarshal(data, e.session)
}

func (e *Engine) GetFramework(name string) (*types.Framework, error) {
	fw, ok := e.frameworks[name]
	if !ok {
		return nil, fmt.Errorf("framework not found: %s", name)
	}
	return fw, nil
}

func (e *Engine) ListFrameworks() []string {
	names := make([]string, 0, len(e.frameworks))
	for name := range e.frameworks {
		names = append(names, name)
	}
	return names
}
