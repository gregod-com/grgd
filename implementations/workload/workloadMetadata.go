package workload

import (
	I "github.com/gregod-com/interfaces"
)

// WorkloadMetadata ...
type WorkloadMetadata struct {
	Name             string `yaml:"name`
	Active           bool   `yaml:"active"`
	Env              string `yaml:"env"`
	PathToDescriptor string `yaml:"path"`
}

// CreateWorkloadMetadata ...
func CreateWorkloadMetadata() I.IWorkloadMetadata {
	return &WorkloadMetadata{}
}

// GetName ...
func (meta *WorkloadMetadata) GetName() string {
	return meta.Name
}

// GetActive ...
func (meta *WorkloadMetadata) GetActive() (string, string, bool) {
	// log.Println("called getact in dcWL MEtatadata for " + meta.GetName())

	if meta.Active {
		return "true", "true", true
	}
	return "false", "false", false
}

// ToggleActive ...
func (meta *WorkloadMetadata) ToggleActive() error {
	meta.Active = !meta.Active
	return nil
}

// GetEnvAsString ...
func (meta *WorkloadMetadata) GetEnvAsString() string {
	return meta.Env
}

// GetEnvAsEmoji ...
func (meta *WorkloadMetadata) GetEnvAsEmoji() string {
	if meta.Env == "prod" {
		return "🐳"
	}
	return "🚧"
}

// GetPathToPodDescriptor ...
func (meta *WorkloadMetadata) GetPathToPodDescriptor() string {
	return meta.PathToDescriptor
}
