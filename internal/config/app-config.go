package config

// ApplicationConfig - struct to hold values of env and app container
type ApplicationConfig struct {
	envValues *envConfig
}

// NewApplicationConfig - loads config values from environment and initialises config
func NewApplicationConfig() *ApplicationConfig {
	envValues := newEnvironmentConfig()

	return &ApplicationConfig{
		envValues: envValues,
	}
}

//NewMockApplicationConfig - returns new mock application config
func NewMockApplicationConfig() *ApplicationConfig {
	return &ApplicationConfig{}
}

//ServiceName returns the name of service
func (cfg *ApplicationConfig) ServiceName() string {
	return cfg.envValues.ServiceName
}

//ServerPort returns the port no to listen for requests
func (cfg *ApplicationConfig) ServerPort() int {
	return cfg.envValues.ServerPort
}
