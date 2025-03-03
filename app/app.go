//Author: LiXue An(anlixue@cn.ibm.com)
//Description: package app for gloabal debug and log file configure

package app

import (
	"fmt"
	"log"
	"os"
	"strings"
)

//Config for global config
type appconfig struct {
	//Debug global
	debug         bool
	logFile       string
	viewerOptions string
	// //LineStartwith  regexp for start line
	// LineStartwith string
	// //HTMLInfo  regexp for Info
	// HTMLInfo string
	// //HTMLWarn regexp for warning
	// HTMLWarn string
	// //HTMLError regexp for error
	// HTMLError string
	// //HTMLSeverity regexp for severity
	// HTMLSeverity string
	// //DateFormat for date field format, for filter by time
	// DateFormat string
}

//Config instance
var Config appconfig

func init() {
	// Config.LineStartwith = `^\d+/\d+/\d+`
	// Config.HTMLInfo = "info"
	// Config.HTMLWarn = "warn|warning"
	// Config.HTMLError = `\serror\s|\serr\s`
	// Config.HTMLSeverity = `\sserv\s`
	// Config.DateFormat = "mm/dd/yyyy"
	// Config.viewerOptions = "-N -R -I"
	Config.viewerOptions = "-N -R -I -C -j 5 -# 10"
}

//LoadConfigFromOSEnv set config values from OS env
func LoadConfigFromOSEnv(envPrefix string) {

	key := strings.ToUpper(envPrefix + "Debug")
	if v := os.Getenv(key); v != "" {
		if strings.ToUpper(v) == "TRUE" || strings.ToUpper(v) == "ON" || strings.ToUpper(v) == "1" {
			Config.debug = true
		}
	}

	key = strings.ToUpper(envPrefix + "LOG")
	if v := os.Getenv(key); v != "" {
		Config.logFile = v
	}

	key = strings.ToUpper(envPrefix + "VIEWER_OPTIONS")
	if v := os.Getenv(key); v != "" {
		Config.viewerOptions = v
		//if no options set, set -R as default for show with color
		if strings.TrimSpace(Config.viewerOptions) == "" {
			Config.viewerOptions = "-R"
		}
	}

	// key = strings.ToUpper(envPrefix + "LineStartwith")
	// if v := os.Getenv(key); v != "" {
	// 	Config.LineStartwith = v
	// }

	// key = strings.ToUpper(envPrefix + "HTMLInfo")
	// if v := os.Getenv(key); v != "" {
	// 	Config.HTMLInfo = v
	// }

	// key = strings.ToUpper(envPrefix + "HTMLWarn")
	// if v := os.Getenv(key); v != "" {
	// 	Config.HTMLWarn = v
	// }

	// key = strings.ToUpper(envPrefix + "HTMLError")
	// if v := os.Getenv(key); v != "" {
	// 	Config.HTMLError = v
	// }

	// key = strings.ToUpper(envPrefix + "HTMLSeverity")
	// if v := os.Getenv(key); v != "" {
	// 	Config.HTMLSeverity = v
	// }

	// key = strings.ToUpper(envPrefix + "DateFormat")
	// if v := os.Getenv(key); v != "" {
	// 	Config.DateFormat = v
	// }

}

//GetConfig global config
func GetConfig() string {
	//return fmt.Sprintf("Debug:%v, LineStartwith:%v, HTMLInfo:%s, HTMLWarn:%s, HTMLSeverity:%s, DateFormat:%s", Config.Debug, Config.LineStartwith, Config.HTMLInfo, Config.HTMLWarn, Config.HTMLSeverity, Config.DateFormat)
	return fmt.Sprintf("%+v", Config)
}

//Debug print debug info to log file
func Debug(v ...interface{}) {
	if Config.debug {
		log.Println("DEBUG==>", v)
	}
}

//IsDebug
func IsDebug() bool {
	return Config.debug
}

//GetViewerOptions
func GetViewerOptions() string {
	return Config.viewerOptions
}

//Info print debug info to log file
func Info(v ...interface{}) {
	if Config.debug {
		log.Println("Info==>", v)
	}
}

//ConfigLogFile if set SPTOOL_LOG env, write log to this file
func ConfigLogFile() *os.File {
	if Config.logFile != "" {
		f, err := os.OpenFile(Config.logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			fmt.Printf("Error while opening log file: %+v \n", err)
		}
		// defer f.Close()
		log.SetOutput(f)
		return f
	} else {
		return nil
	}
}
