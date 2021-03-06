package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/kelseyhightower/envconfig"
	"github.com/mesg-foundation/core/version"
	"github.com/mesg-foundation/core/x/xstrings"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
)

const envPrefix = "mesg"

var (
	instance *Config
	once     sync.Once
)

// ServiceConfig contains information related to services that the config knows about
type ServiceConfig struct {
	URL  string
	Env  map[string]string
	Sid  string
	Hash string
}

// Config contains all the configuration needed.
type Config struct {
	Server struct {
		Address string
	}

	Client struct {
		Address string
	}

	Log struct {
		Format      string
		ForceColors bool
		Level       string
	}

	Core struct {
		Image string
		Name  string
		Path  string

		Database struct {
			ServiceRelativePath   string
			ExecutionRelativePath string
		}
	}

	// You can add any attribute that will be a `ServiceConfig` type
	// 		example: `Foo ServiceConfig`
	// You can then access it with config.Service.Foo.Sid in order to access the Sid of this service.
	// The services in this struct are automatically started when Core starts based on the URL and Env.
	// Sid and Hash are populated only after the deployment.
	// You need to initialize these services in the `New` function
	// 		example: `c.Service.Foo = ServiceConfig{URL: "https://api.github.com/repos/mesg-foundation/service-ethereum-erc20/tarball/master"}`
	// Also add it in the `Services()` function
	// 		example: `return []*ServiceConfig{ &c.Foo }`
	Service struct {
	}

	Docker struct {
		Socket string
		Core   struct {
			Path string
		}
	}
}

// New creates a new config with default values.
func New() (*Config, error) {
	home, err := homedir.Dir()
	if err != nil {
		return nil, err
	}

	var c Config
	c.Server.Address = ":50052"
	c.Client.Address = "localhost:50052"
	c.Log.Format = "text"
	c.Log.Level = "info"
	c.Log.ForceColors = false
	c.Core.Image = "mesg/core:" + strings.Split(version.Version, " ")[0]
	c.Core.Name = "core"
	c.Core.Path = filepath.Join(home, ".mesg")
	c.Core.Database.ServiceRelativePath = filepath.Join("database", "services")
	c.Core.Database.ExecutionRelativePath = filepath.Join("database", "executions")
	c.Docker.Core.Path = "/mesg"
	c.Docker.Socket = "/var/run/docker.sock"
	return &c, nil
}

// Global returns a singleton of a Config after loaded ENV and validate the values.
func Global() (*Config, error) {
	var err error
	once.Do(func() {
		instance, err = New()
		if err != nil {
			return
		}
		if err = instance.Load(); err != nil {
			return
		}
		if err = instance.Prepare(); err != nil {
			return
		}
	})
	if err != nil {
		return nil, err
	}
	if err := instance.Validate(); err != nil {
		return nil, err
	}
	return instance, nil
}

// Load reads config from environmental variables.
func (c *Config) Load() error {
	envconfig.MustProcess(envPrefix, c)
	return nil
}

// Prepare setups local directories or any other required thing based on config
func (c *Config) Prepare() error {
	return os.MkdirAll(c.Core.Path, os.FileMode(0755))
}

// Validate checks values and return an error if any validation failed.
func (c *Config) Validate() error {
	if !xstrings.SliceContains([]string{"text", "json"}, c.Log.Format) {
		return fmt.Errorf("value %q is not an allowed", c.Log.Format)
	}
	if _, err := logrus.ParseLevel(c.Log.Level); err != nil {
		return err
	}
	return nil
}

// Services returns all services that the configuration package is aware of
func (c *Config) Services() []*ServiceConfig {
	return []*ServiceConfig{}
}

// DaemonEnv returns the needed environmental variable for the Daemon.
func (c *Config) DaemonEnv() map[string]string {
	return map[string]string{
		"MESG_SERVER_ADDRESS":  c.Server.Address,
		"MESG_LOG_FORMAT":      c.Log.Format,
		"MESG_LOG_LEVEL":       c.Log.Level,
		"MESG_LOG_FORCECOLORS": strconv.FormatBool(c.Log.ForceColors),
		"MESG_CORE_NAME":       c.Core.Name,
		"MESG_CORE_PATH":       c.Docker.Core.Path,
	}
}
