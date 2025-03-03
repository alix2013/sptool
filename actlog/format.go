//Author: LiXue An(anlixue@cn.ibm.com)
//Description: format actlog to slice
package actlog

import (
	"bufio"
	"fmt"
	"os"
	"sptool/util"
	"strings"
)

func getFormatedActlog(inputFileName string) []string {
	formatedLines := []string{}
	// file, err := os.Open(inputFileName)
	// if err != nil {
	// 	fmt.Println("Error when open file ---> ", err)
	// 	file.Close()
	// 	os.Exit(1)

	// }
	// defer file.Close()
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
		//if len(words) > 3 && util.IsMatchRegexp(`^\d{2}/\d{2}/\d{2}|^\d{4}-\d{2}-\d{2}`, words[0]) {
		// if len(words) > 3 && util.IsMatchRegexp(config.lineStart, words[0]) {
		if len(words) > 3 && isMatchLineStart(words[0]) {
			formatedLines = append(formatedLines, line)

		} else {
			//append current line to last
			if len(formatedLines) > 0 {
				lastIndex := len(formatedLines)
				formatedLines[lastIndex-1] = strings.TrimSpace(formatedLines[lastIndex-1]) + " " + strings.TrimSpace(line)
			}
		}
	}
	return formatedLines
}
