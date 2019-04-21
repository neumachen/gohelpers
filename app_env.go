package paratils

// AppEnv is used to determine the environment of the application
type AppEnv interface {
	String() string
	IsDevelopment() bool
	IsTesting() bool
	IsStaging() bool
	IsProduction() bool
}

type appEnv string

// String ...
func (a appEnv) String() string {
	return string(a)
}

// IsDevelopment ...
func (a appEnv) IsDevelopment() bool {
	return a == DevelopmentEnv
}

// IsTesting ...
func (a appEnv) IsTesting() bool {
	return a == TestingEnv
}

// IsStaging ...
func (a appEnv) IsStaging() bool {
	return a == StagingEnv
}

// IsProductio ...
func (a appEnv) IsProduction() bool {
	return a == ProductionEnv
}

// DevelopmentEnv ...
const DevelopmentEnv appEnv = "development"

// TestingEnv ...
const TestingEnv appEnv = "testing"

// StagingEnv ...
const StagingEnv appEnv = "staging"

// ProductionEnv ...
const ProductionEnv appEnv = "production"
