//Author: LiXue An(anlixue@cn.ibm.com)
//Description: filter actlog by regular expression, support invert-match and ignore case

package actlog

import (
	"encoding/base64"
	"fmt"
	"sptool/util"
	"time"
)

//filter by keyword extenssion
func filterByKeyWordRegExpExt(inputFileName string, outputFileName string, regExp string, ignoreCase bool, invertMatch bool) {
	fmt.Println("Filter actlog by regular expression:", regExp, "Ignore Case?:", ignoreCase, "Invert Match?:", invertMatch)
	outfile := outputFileName
	// if outputFileName == "" {
	// 	// outfile = inputFileName + "-regExp" + ".txt"
	// 	outfile = inputFileName + "-" + base64.StdEncoding.EncodeToString([]byte(regExp)) + ".txt"
	// }
	if outputFileName == "" {
		outfile = inputFileName + "-sptool-filter-regexp-" + base64.StdEncoding.EncodeToString([]byte(regExp)) + "-" + time.Now().Format("150405") + ".txt"
		if inputFileName == "-" {
			outfile = "sptool-filter-regexp-" + base64.StdEncoding.EncodeToString([]byte(regExp)) + "-" + time.Now().Format("150405") + ".txt"
		}
	}
	// actlogList := getFormatedActlog(inputFileName
	// sessionList := getByRegExp(actlogList, regExp)
	filteredList := getFormatedActlogWithFilter(inputFileName, regExp, customFilterForRegExpExt, ignoreCase, invertMatch)
	if len(filteredList) > 0 {
		if outfile != "-" {
			fmt.Println("Output to file --->  ", outfile)
		}
		err := util.WriteLinesToFile(filteredList, outfile)
		if err != nil {
			fmt.Println("Write file error", err)
		}
	} else {
		fmt.Println("Output is empty! please check input file format")
	}
}

func customFilterForRegExpExt(lines *[]string, kw string, options ...interface{}) {
	if len(*lines) == 0 {
		return
	}
	ignoreCase := options[0].(bool)
	invertMatch := options[1].(bool)
	index := len(*lines) - 1

	matchFunc := util.IsMatchRegexp
	if ignoreCase {
		matchFunc = util.IsMatchRegexpCaseInsensitive
	}
	if !invertMatch { //if not match, remove it
		if !matchFunc(kw, (*lines)[index]) {
			//remove this line
			// *lines = append((*lines)[:index], (*lines)[(index+1):]...)
			*lines = (*lines)[:index]
		}
	} else { //if match, remove it
		if matchFunc(kw, (*lines)[index]) {
			*lines = (*lines)[:index]
		}
	}
}

// func customFilterForRegExpExt(lines *[]string, kw string, options ...interface{}) {
// 	if len(*lines) == 0 {
// 		return
// 	}
// 	// fmt.Println(len(options))
// 	// fmt.Println("--->", (*lines)[0])
// 	// var slice []bool = []bool{false, false}
// 	// // var slice []interface{}
// 	// // slice := make([]bool, len(options))
// 	// for i, v := range options {
// 	// 	slice[i] = v.(bool)
// 	// 	// fmt.Println("a")
// 	// 	// fmt.Println(i, v)
// 	// 	// fmt.Println("b")
// 	// }
// 	// ignoreCase := slice[0]
// 	// invertMatch := slice[1]
// 	ignoreCase := options[0].(bool)
// 	invertMatch := options[1].(bool)

// 	index := len(*lines) - 1
// 	if !ignoreCase {
// 		if !invertMatch {
// 			if !util.IsMatchRegexp(kw, (*lines)[index]) {
// 				//remove this line
// 				// *lines = append((*lines)[:index], (*lines)[(index+1):]...)
// 				*lines = (*lines)[:index]
// 			}
// 		} else {
// 			if util.IsMatchRegexp(kw, (*lines)[index]) {
// 				//remove this line
// 				// *lines = append((*lines)[:index], (*lines)[(index+1):]...)
// 				*lines = (*lines)[:index]
// 			}
// 		}
// 	} else { //for ignore case
// 		if !invertMatch {
// 			if !util.IsMatchRegexpCaseInsensitive(kw, (*lines)[index]) {
// 				//remove this line
// 				// *lines = append((*lines)[:index], (*lines)[(index+1):]...)
// 				*lines = (*lines)[:index]
// 			}
// 		} else {
// 			if util.IsMatchRegexpCaseInsensitive(kw, (*lines)[index]) {
// 				//remove this line
// 				// *lines = append((*lines)[:index], (*lines)[(index+1):]...)
// 				*lines = (*lines)[:index]
// 			}
// 		}
// 	}

// }
