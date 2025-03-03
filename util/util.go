//Author: LiXue An(anlixue@cn.ibm.com)
//Description: for common functions

package util

import (
	"fmt"
	"regexp"
	"runtime"
	"strings"
)

//IsMatchRegexp if line match regExp return true
func IsMatchRegexp(regExp, line string) bool {
	match, _ := regexp.MatchString(regExp, line)
	if match {
		return true
	}
	return false
}

//IsMatchRegexpCaseInsensitive if line match regExp return true
func IsMatchRegexpCaseInsensitive(regExp, line string) bool {
	pattern := "(?i)" + regExp
	match, _ := regexp.MatchString(pattern, line)
	if match {
		return true
	}
	return false
}

//Where return
func Where(depthList ...int) string {
	var depth int
	if depthList == nil {
		depth = 1
	} else {
		depth = depthList[0]
	}
	// function, file, line, _ := runtime.Caller(depth)
	_, file, line, _ := runtime.Caller(depth)
	// return fmt.Sprintf("File: %s  Function: %s Line: %d", chopPath(file), runtime.FuncForPC(function).Name(), line)
	return fmt.Sprintf("%s-%d-->", strings.ReplaceAll(chopPath(file), ".go", ""), line)
}

// return the source filename after the last slash
func chopPath(original string) string {
	i := strings.LastIndex(original, "/")
	if i == -1 {
		return original
	}
	return original[i+1:]
}

//Debug print debug info
// func Debug(v ...interface{}) {
// 	// log.SetFlags(log.LstdFlags | log.Lshortfile)
// 	// fmt.Println("==Debug==>", v)
// 	if app.Config.Debug {
// 		log.Println("DEBUG==>", v)
// 	}
// }
