package config

type (
	AppConfig struct {
		Env string
	}

	AllowedEnvs struct {
		Environments []string
		Default      string
	}
)

const DevEnv = "development"
const StagingEnv = "staging"
const ProductionEnv = "production"
const TestEnv = "test"

var allowedEnvs = &AllowedEnvs{
	Environments: []string{DevEnv, TestEnv, StagingEnv, ProductionEnv},
	Default:      DevEnv,
}

func New(env string) *AppConfig {
	return &AppConfig{
		Env: validateEnv(env),
	}
}

func validateEnv(env string) string {
	for _, e := range allowedEnvs.Environments {
		if env == e {
			return e
		}
	}
	return allowedEnvs.Default
}
