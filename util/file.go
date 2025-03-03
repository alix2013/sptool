//Author: LiXue An(anlixue@cn.ibm.com)
//Description: util package for common functions

package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//WriteLinesToFile write the lines to the given file
func WriteLinesToFile(lines []string, fileName string) error {
	if fileName == "-" || strings.ToUpper(fileName) == "STDOUT" {
		WriteLinesToStdout(lines)
		return nil
	}
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func WriteLinesToStdout(lines []string) {
	for _, line := range lines {
		fmt.Println(line)
	}
}

func OpenFileOrStdin(inputFileName string) (*os.File, error) {
	var file *os.File
	if inputFileName == "-" {
		file = os.Stdin
		return file, nil
	} else {
		return os.Open(inputFileName)
	}
}
