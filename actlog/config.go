//Author: LiXue An(anlixue@cn.ibm.com)
//Description: actlog config, initial configure and load OS env and helper functions
package actlog

import (
	"os"
	"strings"
)

type actlogConfig struct {
	lineStart string //regexp for start line
	// info      string //regexp for Info, all non-mached as info
	warn   string //regexp for warning
	error  string //regexp for error
	severe string //regexp for severe
}

var config actlogConfig

func init() {
	//config.linestart = `^\d{2}/\d{2}/\d{2}|^\d{4}-\d{2}-\d{2}`
	config.lineStart = `^\d{2}/\d{2}/\d{2}`
	//config.lineStart = `^\d{1,4}[/-]\d{1,2}[/-]\d{1,4}`   //all date format
	// config.info = `AN.\d{4}I` // or not match others as info
	config.warn = `AN.\d{4}W`
	config.error = `AN.\d{4}E`
	config.severe = `AN.\d{4}[SDK]` //`AN.\d{4}S|AN.\d{4}D|AN.\d{4}K`

}

//LoadConfigFromOSEnv set config values from OS env
func LoadConfigFromOSEnv(envPrefix string) {
	var key string
	key = strings.ToUpper(envPrefix + "LineStart")
	if v := os.Getenv(key); v != "" {
		config.lineStart = v
	}
	// key = strings.ToUpper(envPrefix + "Info")
	// if v := os.Getenv(key); v != "" {
	// 	config.info = v
	// }
	key = strings.ToUpper(envPrefix + "Warn")
	if v := os.Getenv(key); v != "" {
		config.warn = v
	}
	key = strings.ToUpper(envPrefix + "Error")
	if v := os.Getenv(key); v != "" {
		config.error = v
	}
	key = strings.ToUpper(envPrefix + "Severe")
	if v := os.Getenv(key); v != "" {
		config.severe = v
	}
}
