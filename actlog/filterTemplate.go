//Author: LiXue An(anlixue@cn.ibm.com)
//Description: filter actlog template function, implement customized filter function to filter each line, remove last line if not match

package actlog

import (
	"bufio"
	"fmt"
	"os"
	"sptool/util"
	"strings"
)

/////////////////////////////////version 2
//performance tuning, just process one time when read file
//
//just process one time when read file, no need second time process list
//func getFormatedActlogWithFilter(inputFileName string, searchKey string, customFilter func(lines *[]string, index int, kw string)) []string {
func getFormatedActlogWithFilter(inputFileName string, searchKey string, customFilter func(lines *[]string, kw string, options ...interface{}), options ...interface{}) []string {
	formatedLines := []string{}
	// file, err := os.Open(inputFileName)
	// defer file.Close()
	// // defer fmt.Println("---test---")
	// if err != nil {
	// 	fmt.Println("Error when open file ---> ", err)
	// 	file.Close()
	// 	os.Exit(1)
	// }
	file, err := util.OpenFileOrStdin(inputFileName)
	if err != nil {
		fmt.Println("Error when open file ---> ", err)
		file.Close()
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		//if len(words) > 3 && util.IsMatchRegexp(`^\d{2}/\d{2}/\d{2}`, words[0]) {
		// some actlog date format maybe yyyy-mm-dd
		// if len(words) > 3 && util.IsMatchRegexp(`^\d{2}`, words[0]) {
		// if len(words) > 3 && util.IsMatchRegexp(`^\d{2}/\d{2}/\d{2}|^\d{4}-\d{2}-\d{2}`, words[0]) {
		//if len(words) > 3 && util.IsMatchRegexp(config.lineStart, words[0]) {
		if len(words) > 3 && isMatchLineStart(words[0]) {
			// formatedLines = append(formatedLines, line)

			// if len(formatedLines) >= 2 {
			// 	customFilter(&formatedLines, len(formatedLines)-2, searchKey)
			// }

			// if len(formatedLines) >= 1 {
			customFilter(&formatedLines, searchKey, options...)
			// }
			formatedLines = append(formatedLines, line)
		} else {
			//append current line to last
			if len(formatedLines) > 0 {
				lastIndex := len(formatedLines)
				formatedLines[lastIndex-1] = strings.TrimSpace(formatedLines[lastIndex-1]) + " " + strings.TrimSpace(line)
			}
		}
	}

	//filter last line
	customFilter(&formatedLines, searchKey, options...)
	return formatedLines
}
