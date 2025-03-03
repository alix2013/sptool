//Author: LiXue An(anlixue@cn.ibm.com)
//Description: filter actlog by session ID

package actlog

import (
	"fmt"
	"sptool/util"
	"strings"
	"time"
)

func filterBySessionID(inputFileName string, outputFileName string, sessionID string) {
	outfile := outputFileName
	if outputFileName == "" {
		outfile = inputFileName + "-sptool-filter-session-" + sessionID + "-" + time.Now().Format("150405") + ".txt"
		if inputFileName == "-" {
			outfile = "sptool-filter-session-" + sessionID + "-" + time.Now().Format("150405") + ".txt"
		}
	}

	// actlogList := getFormatedActlog(inputFileName)
	// sessionList := getBySessionID(actlogList, sessionID)
	sessionList := getFormatedActlogWithFilter(inputFileName, sessionID, customFilterForSessionID)
	if len(sessionList) > 0 {
		if outfile != "-" && strings.ToUpper(outfile) != "STDOUT" {
			fmt.Println("Output to file --->  ", outfile)
		}
		err := util.WriteLinesToFile(sessionList, outfile)
		if err != nil {
			fmt.Println("Output file error", err)
		}
	} else {
		fmt.Println("Output is empty! please check input file format")
	}

}

//customFilterForSessionID  filter last line
func customFilterForSessionID(lines *[]string, kw string, options ...interface{}) {
	if len(*lines) == 0 {
		return
	}
	index := len(*lines) - 1
	if !strings.Contains((*lines)[index], "SESSION: "+kw) {
		//remove this line
		// *lines = append((*lines)[:index], (*lines)[(index+1):]...)
		*lines = (*lines)[:index]
	}
}

func filterBySessionIDList(inputFileName string, outputFileName string, sessionIDList string) {
	outfile := outputFileName
	if outputFileName == "" {
		// outfile = inputFileName + "-ssid-" + sessionID + ".txt"
		if inputFileName == "-" {
			outfile = "sptool-filter-session-" + sessionIDList + "-" + time.Now().Format("150405") + ".txt"
		} else {
			outfile = inputFileName + "-sptool-filter-session-" + sessionIDList + "-" + time.Now().Format("150405") + ".txt"
		}
	}

	// actlogList := getFormatedActlog(inputFileName)
	// sessionList := getBySessionID(actlogList, sessionID)
	//sessionList := getFormatedActlogWithFilterSessionList(inputFileName, sessionIDList, customFilterForSessionIDList)
	sessionList := getFormatedActlogWithFilter(inputFileName, sessionIDList, customFilterForSessionIDList)
	if len(sessionList) > 0 {
		if outfile != "-" && strings.ToUpper(outfile) != "STDOUT" {
			fmt.Println("Output to file --->  ", outfile)
		}
		err := util.WriteLinesToFile(sessionList, outfile)
		if err != nil {
			fmt.Println("Output error", err)
		}
	} else {
		fmt.Println("Output is empty! please check input file format")
	}

}

//customFilterForSessionIDList  filter last line
func customFilterForSessionIDList(lines *[]string, sessionList string, options ...interface{}) {
	if len(*lines) == 0 {
		return
	}
	index := len(*lines) - 1 //get last line index

	isMatch := false
	sessionSlice := strings.Split(sessionList, ",")

	for _, kw := range sessionSlice {
		if isMatch {
			break //exit if at lease one session match
		} else {
			if strings.Contains((*lines)[index], "SESSION: "+strings.TrimSpace(kw)) {
				isMatch = true
			}
		}
	}
	//remove this line
	// *lines = append((*lines)[:index], (*lines)[(index+1):]...)
	if !isMatch {
		*lines = (*lines)[:index]
	}

}
