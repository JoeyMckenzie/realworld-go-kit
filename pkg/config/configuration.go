package config

const development = "development"
const production = "production"
const docker = "docker"

type ServiceConfig struct {
	Environment string `hcl:"environment"`
	Port        int    `hcl:"port"`
	Timeout     int    `hcl:"timeout"`
}

type Configuration struct {
	Service ServiceConfig `hcl:"service,block"`
}

func (c *Configuration) IsValidEnvironment() bool {
	return c.Service.Environment == production || c.Service.Environment == development || c.Service.Environment == docker
}

func (c *Configuration) IsProduction() bool {
	return c.Service.Environment == production
}

func (c *Configuration) IsDocker() bool {
	return c.Service.Environment == docker
}
