package cmdlog

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"runtime"
)

var accessLogger *log.Logger

//if SPTOOL_DISABLE_CMDLOG have value, no output to cmdlog
//defualt cmd log file is /tmp/sptcmd.log
//do not write logs to windows platform
func init() {
	if runtime.GOOS != "windows" { //AIX, linux or macOS
		if v := os.Getenv("SPTOOL_DISABLE_CMDLOG"); v != "" {
			return
		}

		cmdlogfile := "/tmp/sptcmds.log"
		if v := os.Getenv("SPTOOL_CMDLOG"); v != "" {
			cmdlogfile = v
		}

		file, err := os.OpenFile(cmdlogfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println("Failed to open log file:", err)
		}
		accessLogger = log.New(file, "", log.Ldate|log.Ltime)
		// change file permission again, not sure why OpenFile not set permission correctly
		err = os.Chmod(cmdlogfile, 0666)
		if err != nil {
			fmt.Println("Failed to operate log file:", err)
		}
	}
}

//Info write logs to file
func Info(s string) {
	if accessLogger != nil {
		user, _ := user.Current()
		path, _ := os.Getwd()
		// if err != nil {
		// 	log.Println(err)
		// }
		// fmt.Println(path)
		//log current user and current directory
		accessLogger.Println(user.Username + " " + path + " " + s)
	}
}
