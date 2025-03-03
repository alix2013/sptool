//Author: LiXue An(anlixue@cn.ibm.com)
//Description: filter actlog by message type "I,W,E,S"

package actlog

import (
	"fmt"
	"sptool/util"
	"strings"
	"time"
)

//getFTOutfileName
func getFTOutfileName(inputFileName string, outputFileName string, messageTypeList string, invertMatch bool) string {
	outfile := outputFileName
	if outputFileName == "" {
		if inputFileName == "-" {
			if invertMatch { //add invert tag to filename
				outfile = "sptool-filter-messagetype-" + messageTypeList + "-v-" + time.Now().Format("150405") + ".txt"
			} else {
				outfile = "sptool-filter-messagetype-" + messageTypeList + "-" + time.Now().Format("150405") + ".txt"
			}
		} else {
			if invertMatch { //add invert tag to filename
				outfile = inputFileName + "-sptool-filter-messagetype-" + messageTypeList + "-v-" + time.Now().Format("150405") + ".txt"
			} else {
				outfile = inputFileName + "-sptool-filter-messagetype-" + messageTypeList + "-" + time.Now().Format("150405") + ".txt"
			}
		}
	}
	return outfile
}

//filterByMessageType
func filterByMessageType(inputFileName string, outputFileName string, messageTypeList string, invertMatch bool) {
	outfile := getFTOutfileName(inputFileName, outputFileName, messageTypeList, invertMatch)

	resultList := getFormatedActlogWithFilter(inputFileName, messageTypeList, customFilterForMessageType, invertMatch)
	if len(resultList) > 0 {
		if outfile != "-" && strings.ToUpper(outfile) != "STDOUT" {
			fmt.Println("Output to file --->  ", outfile)
		}
		err := util.WriteLinesToFile(resultList, outfile)
		if err != nil {
			fmt.Println("Output error", err)
		}
	} else {
		fmt.Println("Output is empty! please check input file format")
	}

}

//customFilterForMessageType  filter last line
func customFilterForMessageType(lines *[]string, messageList string, options ...interface{}) {
	// fmt.Println("----", options)
	if len(*lines) == 0 {
		return
	}
	index := len(*lines) - 1 //get last line index
	isInvertMatch := options[0].(bool)

	isMatch := false
	messageSlice := strings.Split(messageList, ",")
	for _, kw := range messageSlice {
		if isMatch {
			break //exit if at lease one session match
		} else {
			line := (*lines)[index]
			// fmt.Println(kw, line)
			switch strings.ToUpper(kw) {
			case "S":
				if isSevere(line) {
					isMatch = true
					break
				}
			case "E":
				if isError(line) {
					isMatch = true
					break
				}
			case "W":
				if isWarn(line) {
					isMatch = true
					break
				}
			case "I":
				if isInfo(line) {
					isMatch = true
					break
				}
			}
		}
	}
	if isInvertMatch {
		isMatch = !isMatch
	}
	//remove this line
	// *lines = append((*lines)[:index], (*lines)[(index+1):]...)
	if !isMatch {
		*lines = (*lines)[:index]
	}

}
