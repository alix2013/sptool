package actlog

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"sptool/app"
	"sptool/terminalcolor"
	"sptool/util"
	"time"
)

func viewActlog(actlogFileName string) {
	user, _ := user.Current()

	// outFile := actlogFileName + "-sptool-viewer-" + user.Username + ".tmp"
	outFile := "/tmp/.sptool-viewer-" + user.Username + "-" + time.Now().Format("150405") + ".tmp"
	if actlogFileName == "-" {
		// outFile = "sptool-viewer-" + user.Username + ".tmp"
		outFile = "/tmp/.sptool-viewer-" + user.Username + "-" + time.Now().Format("150405") + ".tmp"
	}
	var outLines []string
	lines := getFormatedActlog(actlogFileName)

	for _, line := range lines {
		//if util.IsMatchRegexp(`AN.\d{4}S`, line) || util.IsMatchRegexp(`AN.\d{4}D`, line) || util.IsMatchRegexp(`AN.\d{4}K`, line) {
		if util.IsMatchRegexp(config.severe, line) {
			outLines = append(outLines, terminalcolor.SevereColor(line))
			continue
		}
		//if util.IsMatchRegexp(`AN.\d{4}E`, line) {
		if util.IsMatchRegexp(config.error, line) {
			outLines = append(outLines, terminalcolor.ErrorColor(line))
			continue
		}
		//if util.IsMatchRegexp(`AN.\d{4}W`, line) {
		if util.IsMatchRegexp(config.warn, line) {
			outLines = append(outLines, terminalcolor.WarnColor(line))
			continue
		}
		outLines = append(outLines, terminalcolor.InfoColor(line))
	}

	if len(outLines) <= 0 {
		fmt.Println("Output is empty! please check input file format")
		return
	}
	err := util.WriteLinesToFile(outLines, outFile)
	if err != nil {
		fmt.Println("Write file error", err)
		return
	}

	//call external cmd
	viewerOptions := app.GetViewerOptions()
	cmd := exec.Command("less", viewerOptions, outFile)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	//	clean up tmp file
	if !app.IsDebug() {
		err = os.Remove(outFile)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
