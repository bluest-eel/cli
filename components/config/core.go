package config

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/geomyidia/zylog/logger"
	log "github.com/sirupsen/logrus"
	cfg "github.com/spf13/viper"
)

// Configuration related constants
const (
	AppName         string = "bluest-eel"
	ConfigDir       string = "configs"
	ConfigFile      string = "client"
	ConfigType      string = "yaml"
	ConfigReadError string = "Fatal error config file"
)

var (
	_, b, _, _    = runtime.Caller(0)
	componentsDir = filepath.Dir(filepath.Dir(b))
	projectDir    = filepath.Dir(componentsDir)
	configsDir    = filepath.Join(projectDir, ConfigDir)
)

func init() {
	cfg.AddConfigPath(ConfigDir)
	cfg.SetConfigName(ConfigFile)
	cfg.SetConfigType(ConfigType)
	cfg.SetEnvPrefix(AppName)

	cfg.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	cfg.Set("Verbose", true)
	cfg.AutomaticEnv()
	cfg.AddConfigPath(configsDir)
	cfg.AddConfigPath("/")                   // support for Docker
	cfg.AddConfigPath("/etc/bluest-eel/cli") // support for bare-metal deploys

	err := cfg.ReadInConfig()
	if err != nil {
		log.Panicf("%s: %s", ConfigReadError, err)
	}
	// log.Infof("Env vars: %v", os.Environ())
}

// MessagingServerConfig ...
type MessagingServerConfig struct {
	Host string
	Port int
}

// GRPCDConfig ...
type GRPCDConfig struct {
	Host string
	Port int
}

// Config ...
type Config struct {
	MessagingServer *MessagingServerConfig
	GRPCD           *GRPCDConfig
	Logging         *logger.ZyLogOptions
}

// NewConfig is a constructor that creates the full coniguration data structure
// for use by our application(s) and client(s) as an in-memory copy of the
// config data (saving from having to make repeated and somewhat expensive
// calls to the viper library).
//
// Note that Viper does provide both the AllSettings() and Unmarshall()
// functions, but these require that you have a struct defined that will be
// used to dump the Viper config data into. We've already got that set up, so
// there's no real benefit to switching.
//
// Furthermore, in our case, we're utilizing structs from other libraries to
// be used when setting those up (see how we initialize the logging component
// in ./components/logging.go, Setup).
func NewConfig() *Config {
	return &Config{
		GRPCD: &GRPCDConfig{
			Host: cfg.GetString("grpc.host"),
			Port: cfg.GetInt("grpc.port"),
		},
		Logging: &logger.ZyLogOptions{
			Colored:      cfg.GetBool("logging.colored"),
			Level:        cfg.GetString("logging.level"),
			Output:       cfg.GetString("logging.output"),
			ReportCaller: cfg.GetBool("logging.report-caller"),
		},
		MessagingServer: &MessagingServerConfig{
			Host: cfg.GetString("messaging-server.host"),
			Port: cfg.GetInt("messaging-server.port"),
		},
	}
}

// GRPCConnectionString ...
func (c *Config) GRPCConnectionString() string {
	return fmt.Sprintf("%s:%d", c.GRPCD.Host, c.GRPCD.Port)
}
