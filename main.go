//Author: LiXue An(anlixue@cn.ibm.com)
//Description: main entry point, load OS environments and run cmd

package main

import (
	"sptool/app"
)

func main() {
	//get running program folder name for log file location instead of current folder
	// dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	// if err != nil {
	// 	// log.Fatalf("Error filepath dir: %v", err)
	// 	fmt.Printf("Error filepath dir: %v \n", err)
	// }

	// //create log file if it not exist
	// // f, err := os.OpenFile("sptool.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// f, err := os.OpenFile(dir+"/sptool.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	// log.Fatalf("Error opening file: %v", err)
	// 	fmt.Printf("Error opening file: %v \n", err)
	// }

	//if not call this function, debug info output to stdout
	app.LoadConfigFromOSEnv("SPTOOL_")
	f := app.ConfigLogFile()
	if f != nil {
		defer f.Close()
	}
	// if err != nil {
	// 	log.Fatalf("Error opening file: %v", err)
	// }
	// if f != nil {
	// 	defer f.Close()
	// 	log.SetOutput(f)
	// }
	// log.SetFlags(log.LstdFlags | log.Lshortfile)
	// appconfig.PrintConfig()
	// app.Debug(app.GetConfig())
	runCommand()
}

// func configLogFile() *os.File {
// 	if app.Config.LogFile != "" {
// 		f, err := os.OpenFile(app.Config.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
// 		if err != nil {
// 			fmt.Printf("Error while opening log file: %v \n", err)
// 		}
// 		// defer f.Close()
// 		log.SetOutput(f)
// 		return f
// 	} else {
// 		return nil
// 	}
// }
