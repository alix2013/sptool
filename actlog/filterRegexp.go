//Author: LiXue An(anlixue@cn.ibm.com)
//Description: filter actlog by regular expression

package actlog

import (
	"encoding/base64"
	"fmt"
	"sptool/util"
	"time"
)

func filterByKeyWordRegExp(inputFileName string, outputFileName string, regExp string) {
	fmt.Println("filter by keyword:", regExp)
	outfile := outputFileName
	// if outputFileName == "" {
	// 	outfile = inputFileName + "-" + base64.StdEncoding.EncodeToString([]byte(regExp)) + ".txt"
	// }

	if outputFileName == "" {
		outfile = inputFileName + "-sptool-filter-keyword-" + base64.StdEncoding.EncodeToString([]byte(regExp)) + "-" + time.Now().Format("150405") + ".txt"
		if inputFileName == "-" {
			outfile = "sptool-filter-keyword-" + base64.StdEncoding.EncodeToString([]byte(regExp)) + "-" + time.Now().Format("150405") + ".txt"
		}
	}
	// actlogList := getFormatedActlog(inputFileName
	// sessionList := getByRegExp(actlogList, regExp)
	filteredList := getFormatedActlogWithFilter(inputFileName, regExp, customFilterForRegExp)
	if len(filteredList) > 0 {
		fmt.Println("Output to file --->  ", outfile)
		err := util.WriteLinesToFile(filteredList, outfile)
		if err != nil {
			fmt.Println("Write file error", err)
		}
	} else {
		fmt.Println("Output is empty! please check input file format")
	}
}

func customFilterForRegExp(lines *[]string, kw string, options ...interface{}) {
	if len(*lines) == 0 {
		return
	}
	index := len(*lines) - 1
	if !util.IsMatchRegexp(kw, (*lines)[index]) {
		//remove this line
		// *lines = append((*lines)[:index], (*lines)[(index+1):]...)
		*lines = (*lines)[:index]
	}
}
