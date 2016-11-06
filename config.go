package GoConfig

import (
	"fmt"

	"github.com/spf13/viper"
)

var allowedConfigTypes = map[string]bool{
	"json": true,
	"yaml": true,
	"toml": true,
}

type ConfigOptions struct {
	configType string
	configFile string
	configPath string
}

func Init(options *ConfigOptions) error {
	if !isConfigTypeAllowed(options.configType) {
		return fmt.Errorf("Config type of %s not allowed", options.configType)
	}
	viper.WatchConfig()
	setConfigOptions(options)
	return viper.ReadInConfig()
}

func isConfigTypeAllowed(configType string) bool {
	if _, ok := allowedConfigTypes[configType]; !ok {
		return false
	}
	return true
}

func setConfigOptions(options *ConfigOptions) {
	viper.SetConfigType(options.configType)
	viper.SetConfigName(options.configFile)
	viper.AddConfigPath(options.configPath)
}

func GetConfigStringValue(key string) string {
	return viper.GetString(key)
}

func GetConfigIntValue(key string) int {
	return viper.GetInt(key)
}

func GetConfigFloatValue(key string) float64 {
	return viper.GetFloat64(key)
}

func GetConfigBoolValue(key string) bool {
	return viper.GetBool(key)
}

func GetConfigMapValue(key string) map[string]interface{} {
	return viper.GetStringMap(key)
}

func HasKey(key string) bool {
	return viper.IsSet(key)
}
