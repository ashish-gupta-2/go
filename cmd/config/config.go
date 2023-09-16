package config

import (
	"github.com/spf13/viper"
)

// Constant declarations for package level usage.
const (
	envDBEndpoint = "APP_DB_ENDPOINT"
	envGRPCPort   = "APP_GRPC_PORT"
	envHTTPPort   = "APP_HTTP_PORT"
)

// Config represent all application configuration.
type Config struct {
	DBEndpoint string
	GRPCPort   uint16
	HTTPPort   uint16
}

// NewConfig creates a new instance of app config.
func NewConfig() Config {
	viper.AutomaticEnv()
	setDefaultAppConfig()

	vpr := viper.GetViper()
	return Config{
		DBEndpoint: vpr.GetString(envDBEndpoint),
		GRPCPort:   uint16(vpr.GetUint(envGRPCPort)),
		HTTPPort:   uint16(vpr.GetUint(envHTTPPort)),
	}
}

func setDefaultAppConfig() {
	viper.SetDefault(envDBEndpoint, "postgres://postgres:postgres@0.0.0.0:5432/postgres?sslmode=disable")
	viper.SetDefault(envGRPCPort, 8181)
}
