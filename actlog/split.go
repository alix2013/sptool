//Author: LiXue An(anlixue@cn.ibm.com)
//Description: split actlog by date or time(hours)
//actlog date format might be MM/dd/YY or MM/dd/YYYY

package actlog

import (
	"bufio"
	"fmt"
	"os"
	"sptool/util"
	"strings"
	"time"
)

// func isMatchKeyWord(regExp, word string) bool {
// 	match, _ := regexp.MatchString(regExp, word)
// 	if match {
// 		return true
// 	}
// 	return false
// }

//SplitActlog for split act log by date
func splitActlogByDate(inputFileName string) {
	formatedLines := []string{}
	// file, err := os.Open(inputFileName)
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
	//for large line
	// buf := make([]byte, 0, 64*1024)
	// scanner.Buffer(buf, 1024*1024)
	logDate := ""
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		//if len(words) > 3 && util.IsMatchRegexp(`^\d{2}/\d{2}/\d{2}`, words[0]) {
		// some actlog date format maybe yyyy-mm-dd
		// if len(words) > 3 && util.IsMatchRegexp(`^\d{2}/\d{2}/\d{2}|^\d{4}-\d{2}-\d{2}`, words[0]) {
		// if len(words) > 3 && util.IsMatchRegexp(config.lineStart, words[0]) {
		if len(words) > 3 && isMatchLineStart(words[0]) {
			if logDate != words[0] && logDate != "" {
				//write to file
				myOutFileName := inputFileName + "-" + strings.Replace(logDate, "/", "-", -1) + "-sptool-split-" + time.Now().Format("150405") + ".txt"
				fmt.Println("Output to file ---> ", myOutFileName)
				err := util.WriteLinesToFile(formatedLines, myOutFileName)
				if err != nil {
					fmt.Println("Write file error", err)
				}
				//clean formatLines
				formatedLines = []string{}
			}
			formatedLines = append(formatedLines, line)
			logDate = words[0]
		} else {
			//append current line to last
			if len(formatedLines) > 0 {
				lastIndex := len(formatedLines)
				formatedLines[lastIndex-1] = strings.TrimSpace(formatedLines[lastIndex-1]) + " " + strings.TrimSpace(line)
			}
		}
	}
	//last batch write to file
	if len(formatedLines) > 0 {
		logDateStr := strings.Fields(formatedLines[0])[0]
		myOutFileName := inputFileName + "-" + strings.Replace(logDateStr, "/", "-", -1) + "-sptool-split-" + time.Now().Format("150405") + ".txt"
		fmt.Println("Output to file ---> ", myOutFileName)
		err := util.WriteLinesToFile(formatedLines, myOutFileName)
		if err != nil {
			fmt.Println("Write file error", err)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Input file format is invalid,", "Error:", err)
	}
}

//SplitActlogByHour for split by hour
func splitActlogByHour(inputFileName string) {
	formatedLines := []string{}
	// file, err := os.Open(inputFileName)
	// if err != nil {
	// 	fmt.Println("Error when open file ---> ", err)
	// 	file.Close()
	// 	os.Exit(1)
	// 	// return
	// }
	file, err := util.OpenFileOrStdin(inputFileName)
	if err != nil {
		fmt.Println("Error when open file ---> ", err)
		file.Close()
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	logDate := ""
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		//if len(words) > 3 && util.IsMatchRegexp(`^\d{2}/\d{2}/\d{2}`, words[0]) {
		// some actlog date format maybe yyyy-mm-dd
		//if len(words) > 3 && util.IsMatchRegexp(`^\d{2}/\d{2}/\d{2}|^\d{4}-\d{2}-\d{2}`, words[0]) {
		//if len(words) > 3 && util.IsMatchRegexp(config.lineStart, words[0]) {
		if len(words) > 3 && isMatchLineStart(words[0]) {
			hourStr := words[0] + "-" + words[1][:2] //hour string
			if logDate != hourStr && logDate != "" {
				//write to file
				myOutFileName := inputFileName + "-" + strings.Replace(logDate, "/", "-", -1) + "-sptool-split-" + time.Now().Format("150405") + ".txt"
				if inputFileName == "-" {
					myOutFileName = strings.Replace(logDate, "/", "-", -1) + "-sptool-split-" + time.Now().Format("150405") + ".txt"
				}
				fmt.Println("Output to file ---> ", myOutFileName)
				err := util.WriteLinesToFile(formatedLines, myOutFileName)
				if err != nil {
					fmt.Println("Write file error", err)
				}
				//clean formatLines
				formatedLines = []string{}
			}
			formatedLines = append(formatedLines, line)
			logDate = words[0] + "-" + words[1][:2]
		} else {
			//append current line to last line string
			if len(formatedLines) > 0 {
				lastIndex := len(formatedLines)
				formatedLines[lastIndex-1] = strings.TrimSpace(formatedLines[lastIndex-1]) + " " + strings.TrimSpace(line)
			}
		}
	}
	//last batch write to file
	if len(formatedLines) > 0 {
		words := strings.Fields(formatedLines[0])
		logDateStr := words[0] + "-" + words[1][:2]
		myOutFileName := inputFileName + "-" + strings.Replace(logDateStr, "/", "-", -1) + "-sptool-split-" + time.Now().Format("150405") + ".txt"
		if inputFileName == "-" {
			myOutFileName = strings.Replace(logDate, "/", "-", -1) + "-sptool-split-" + time.Now().Format("150405") + ".txt"
		}
		fmt.Println("Output to file ---> ", myOutFileName)
		err := util.WriteLinesToFile(formatedLines, myOutFileName)
		if err != nil {
			fmt.Println("Write file error", err)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

/*
//SplitActlog for split act log by date
func SplitActlog(inputFileName string, isByHour bool) {
	//line[0:8] for split by date, [0:13] for by hour

	lastIndex := 8
	if isByHour {
		lastIndex = 13
	}
	formatedLines := []string{}
	//open input file
	file, err := os.Open(inputFileName)
	if err != nil {
		// log.Fatal(err)
		fmt.Println("Error when open file ---> ", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	logDate := ""
	for scanner.Scan() {
		line := scanner.Text()
		// words := strings.Fields(line)

		if isMatchKeyWord(`^\d{2}/\d{2}/\d{2}`, line) {

			if logDate != line[0:lastIndex] && logDate != "" {
				//write to file
				myOutFileName := inputFileName + "-" + strings.Replace(logDate, "/", "-", -1) + ".txt"
				myOutFileName = strings.ReplaceAll(myOutFileName, " ", "-")
				fmt.Println("Output to file ---> ", myOutFileName)
				util.WriteLinesToFile(formatedLines, myOutFileName)
				//clean formatLines
				formatedLines = []string{}
			}

			formatedLines = append(formatedLines, line)
			logDate = line[0:lastIndex]
			// matchBegin = true

		} else {
			//append current line to last
			if len(formatedLines) > 0 {
				lastIndex := len(formatedLines)
				formatedLines[lastIndex-1] = strings.TrimSpace(formatedLines[lastIndex-1]) + " " + strings.TrimSpace(line)
			}
		}

	}

	//last write
	if len(formatedLines) > 0 {
		// logDateStr := strings.Fields(formatedLines[0])[0]
		logDateStr := formatedLines[0][0:lastIndex]
		myOutFileName := inputFileName + "-" + strings.Replace(logDateStr, "/", "-", -1) + ".txt"
		myOutFileName = strings.ReplaceAll(myOutFileName, " ", "-")
		fmt.Println("Output to file ---> ", myOutFileName)
		util.WriteLinesToFile(formatedLines, myOutFileName)

	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)

	}
}
*/
