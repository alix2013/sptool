//Author: LiXue An(anlixue@cn.ibm.com)
//Description: all commands entry point
package main

import (
	"flag"
	"fmt"
	"os"
	"sptool/actlog"

	// "sptool/clienttrace"
	// "sptool/device"
	// "sptool/webserver"
	"strings"
)

func runCommand() {
	flag.Usage = appUsage

	if len(os.Args) == 1 {
		flag.Usage()
		os.Exit(1)
	}
	//command for actlog, split cmds to sub module
	actlog.RunCommand()

	switch strings.ToUpper(os.Args[1]) {
	case "VERSION", "V":
		printVersion()
	case "HELP", "H":
		printHelp()
	default:
		// fmt.Println("Params:", os.Args[0:])
		fmt.Printf("%s is not valid command.\n", os.Args[1])
		flag.Usage()
		os.Exit(1)
	}
}

// //old version
// func runCommand_old() {
// 	//for device cmd setup
// 	flag.Usage = appUsage
// 	// redefineCommand := flag.NewFlagSet("redefine", flag.ExitOnError)
// 	// redefineCommandInputFileFlag := redefineCommand.String("f", "", "Generate redefine Library/Drive/Path cmd")
// 	// redefineCommandOutputFileFlag := redefineCommand.String("o", "", "(Optional)output file, default is stdout")
// 	// // isDebugFlag := redefineCommand.Bool("v", false, "Verbose output")
// 	// // redefineCommand.Usage = deviceUsage //overwrite default usage for subcmd

// 	// cmpSlotsLibvCommand := flag.NewFlagSet("cmpSlotsLibv", flag.ExitOnError)
// 	// cmpSlotsLibvCommandLIbvInputFileFlag := cmpSlotsLibvCommand.String("l", "", "admin cmd 'q libv' output file name")
// 	// cmpSlotsLibvCommandSlotsInputFileFlag := cmpSlotsLibvCommand.String("s", "", "admin cmd 'show slots library' output file name")

// 	//for splitactlog cmd
// 	splitActlogCommand := flag.NewFlagSet("SplitActlog", flag.ExitOnError)
// 	splitActlogInputFileFlag := splitActlogCommand.String("f", "", "actlog file name")
// 	splitActlogCommandEnableByHourFlag := splitActlogCommand.Bool("t", false, "(Optional) Generate one file per hour")
// 	// splitActlogCommand.Usage = splitActlogUsage

// 	//for filter actlog by session id  cmd
// 	filterActlogBySSCommand := flag.NewFlagSet("FilterActlogbySession", flag.ExitOnError)
// 	filterActlogBySSCommandInputFileFlag := filterActlogBySSCommand.String("f", "", `actlog file name `)
// 	filterActlogBySSCommandSessionIDFlag := filterActlogBySSCommand.String("s", "", "session id ")
// 	filterActlogBySSCommandOutputFileFlag := filterActlogBySSCommand.String("o", "", `(Optional) Output file name`)
// 	// filterActlogBySSCommand.Usage = filterActlogBySSUsage

// 	//for filter actlog by session id list cmd
// 	filterActlogBySessionsCommand := flag.NewFlagSet("FilterActlogbySessionIDs", flag.ExitOnError)
// 	filterActlogBySessionsCommandInputFileFlag := filterActlogBySessionsCommand.String("f", "", `actlog file name  `)
// 	filterActlogBySessionsCommandSessionIDFlag := filterActlogBySessionsCommand.String("s", "", `session ID list,multiple IDs separate by comma ","`)
// 	filterActlogBySessionsCommandOutputFileFlag := filterActlogBySessionsCommand.String("o", "", `(Optional) Output file name `)
// 	// filterActlogBySSCommand.Usage = filterActlogBySSUsage

// 	//for message type filter
// 	filterActlogByMessageTypeCommand := flag.NewFlagSet("FilterActlogbyMessageType", flag.ExitOnError)
// 	filterActlogByMessageTypeCommandInputFileFlag := filterActlogByMessageTypeCommand.String("f", "", `actlog file name `)
// 	filterActlogByMessageTypeCommandMessageTypeFlag := filterActlogByMessageTypeCommand.String("t", "", `message type:[I,W,E,S],I:Information,W:Warning,E:Error,S:Severe,multiple values separate by comma "," `)
// 	filterActlogByMessageTypeCommandOutputFileFlag := filterActlogByMessageTypeCommand.String("o", "", `(Optional) Output file name`)

// 	//filer by keyword
// 	filterActlogByKWCommand := flag.NewFlagSet("FilterActlogbyKeyword", flag.ExitOnError)
// 	filterActlogByKWCommandInputFileFlag := filterActlogByKWCommand.String("f", "", "actlog file name")
// 	filterActlogByKWCommandKeywordFlag := filterActlogByKWCommand.String("k", "", "keyword")
// 	filterActlogByKWCommandOutputFileFlag := filterActlogByKWCommand.String("o", "", "(Optional) Output file name")
// 	// filterActlogByKWCommand.Usage = filterActlogByKWUsage

// 	filterActlogByKWExtCommand := flag.NewFlagSet("FilterActlogbyKW", flag.ExitOnError)
// 	filterActlogByKWExtCommandInputFileFlag := filterActlogByKWExtCommand.String("f", "", "actlog file name")
// 	filterActlogByKWExtCommandKeywordFlag := filterActlogByKWExtCommand.String("k", "", "keyword")
// 	filterActlogByKWExtCommandOutputFileFlag := filterActlogByKWExtCommand.String("o", "", "(Optional) Output file name")
// 	filterActlogByKWExtCommandIgnoreCaseFlag := filterActlogByKWExtCommand.Bool("i", false, "(Optional) ignore case")
// 	filterActlogByKWExtCommandInvertMatchFlag := filterActlogByKWExtCommand.Bool("v", false, "(Optional) invert match(not match specified keyword)")
// 	//filter by node name
// 	filterActlogByNodesCommand := flag.NewFlagSet("FilterActlogbyNode", flag.ExitOnError)
// 	filterActlogByNodesCommandInputFileFlag := filterActlogByNodesCommand.String("f", "", "actlog file name")
// 	filterActlogByNodesCommandNodesFlag := filterActlogByNodesCommand.String("n", "", "node name, multiple nodes separate by ','")
// 	filterActlogByNodesCommandOutputFileFlag := filterActlogByNodesCommand.String("o", "", "(Optional) Output file ")
// 	// filterActlogByKWCommand.Usage = filterActlogByKWUsage

// 	//convert actlog to html
// 	convertActlog2HTMLCommand := flag.NewFlagSet("ConvertActlogToHtml", flag.ExitOnError)
// 	convertActlog2HTMLCommandInputFileFlag := convertActlog2HTMLCommand.String("f", "", `actlog file name `)
// 	convertActlog2HTMLCommandOutputFileFlag := convertActlog2HTMLCommand.String("o", "", `(Optional) Output file `)
// 	// convertActlog2HTMLCommand.Usage = convertActlog2HTMLUsage

// 	viewActlogCommand := flag.NewFlagSet("ViewActlog", flag.ExitOnError)
// 	viewActlogCommandInputFileFlag := viewActlogCommand.String("f", "", "actlog file name")
// 	// convertActlog2HTMLCommandOutputFileFlag := convertActlog2HTMLCommand.String("o", "", "(Optional) Output file ")

// 	//web ui for actlog
// 	// webuiActlogCommand := flag.NewFlagSet("WebUiforActlog", flag.ExitOnError)
// 	// webuiActlogCommandInputFileFlag := webuiActlogCommand.String("f", "", "actlog file name")
// 	// webuiActlogCommandPort := webuiActlogCommand.Int("p", 0, "web server tcp port,listening on a random port by default")

// 	// //web ui for actlog
// 	// webuiOCCommand := flag.NewFlagSet("WebUiforOcopscntrlog", flag.ExitOnError)
// 	// webuiOCCommandInputFileFlag := webuiOCCommand.String("f", "", "actlog file name")
// 	// webuiOCCommandPort := webuiOCCommand.Int("p", 0, "web server tcp port,listening on a random port by default")

// 	// //filtr client trace by regExp
// 	// filterCTByKWCommand := flag.NewFlagSet("FilterClienttracebyKeyword", flag.ExitOnError)
// 	// filterCTByKWCommandInputFileFlag := filterCTByKWCommand.String("f", "", "client/api trace file name")
// 	// filterCTByKWCommandKeywordFlag := filterCTByKWCommand.String("k", "", "keyword")
// 	// filterCTByKWCommandOutputFileFlag := filterCTByKWCommand.String("o", "", "(Optional) Output file ")
// 	// filterCTByKWCommandCaseInSensitiveFlag := filterCTByKWCommand.Bool("i", false, "(Optional)keyword case insensitive")

// 	// //filtr client trace by time
// 	// filterCTByTimeCommand := flag.NewFlagSet("FilterClienttracebyTime", flag.ExitOnError)
// 	// filterCTByTimeCommandInputFileFlag := filterCTByTimeCommand.String("f", "", "client/api trace file name")
// 	// filterCTByTimeCommandBeginTimeFlag := filterCTByTimeCommand.String("b", "", "begin time, yyyy-mm-ddThh:mm:ss")
// 	// filterCTByTimeCommandEndTimeFlag := filterCTByTimeCommand.String("e", "", "end time, yyyy-mm-ddThh:mm:ss")
// 	// filterCTByTimeCommandOutputFileFlag := filterCTByTimeCommand.String("o", "", "(Optional) Output file ")

// 	// //split by time
// 	// splitCTByTimeCommand := flag.NewFlagSet("SplitClienttracebyTime", flag.ExitOnError)
// 	// splitCTByTimeCommandInputFileFlag := splitCTByTimeCommand.String("f", "", "client/api trace file name")
// 	// splitCTByTimeCommandMinutesFlag := splitCTByTimeCommand.Int("m", 60, "minutes")

// 	// //convert client trace to html
// 	// convertCT2HTMLCommand := flag.NewFlagSet("ConvertClientTracetoHtml", flag.ExitOnError)
// 	// convertCT2HTMLCommandInputFileFlag := convertCT2HTMLCommand.String("f", "", "client/api trace file name")
// 	// convertCT2HTMLCommandOutputFileFlag := convertCT2HTMLCommand.String("o", "", "(Optional) Output file ")

// 	if len(os.Args) == 1 {
// 		// fmt.Println(usageStr)
// 		flag.Usage()
// 		return
// 	}

// 	switch strings.ToUpper(os.Args[1]) {
// 	case "SPLITACTLOG", "SA":
// 		splitActlogCommand.Parse(os.Args[2:])
// 		if splitActlogCommand.Parsed() {
// 			if *splitActlogInputFileFlag == "" {
// 				splitActlogCommand.Usage()
// 				return
// 			}
// 			cmd := actlog.NewSplitActlogCommand(*splitActlogInputFileFlag, *splitActlogCommandEnableByHourFlag)
// 			cmd.SplitActlog()

// 		}

// 	case "FILTERACTLOGBYSESSION", "FAS":
// 		filterActlogBySSCommand.Parse(os.Args[2:])
// 		if filterActlogBySSCommand.Parsed() {
// 			if *filterActlogBySSCommandInputFileFlag == "" || *filterActlogBySSCommandSessionIDFlag == "" {
// 				filterActlogBySSCommand.Usage()
// 				return
// 			}

// 			cmd := actlog.NewFilterActlogBySessionCommand(*filterActlogBySSCommandInputFileFlag, *filterActlogBySSCommandSessionIDFlag, *filterActlogBySSCommandOutputFileFlag)
// 			cmd.FilterActlogBySessionID()
// 			// fmt.Println("Args:", os.Args[2:])
// 		}

// 	case "FILTERACTLOGBYSESSIONIDS", "FASS":
// 		filterActlogBySessionsCommand.Parse(os.Args[2:])
// 		if filterActlogBySessionsCommand.Parsed() {
// 			if *filterActlogBySessionsCommandInputFileFlag == "" || *filterActlogBySessionsCommandSessionIDFlag == "" {
// 				filterActlogBySessionsCommand.Usage()
// 				return
// 			}

// 			cmd := actlog.NewFilterActlogBySessionIDsCommand(*filterActlogBySessionsCommandInputFileFlag, *filterActlogBySessionsCommandSessionIDFlag, *filterActlogBySessionsCommandOutputFileFlag)
// 			cmd.FilterActlogBySessionIDs()
// 			// fmt.Println("Args:", os.Args[2:])
// 		}
// 	case "FILTERACTLOGBYMESSAGETYPE", "FAT":
// 		filterActlogByMessageTypeCommand.Parse(os.Args[2:])
// 		if filterActlogByMessageTypeCommand.Parsed() {
// 			if *filterActlogByMessageTypeCommandInputFileFlag == "" || *filterActlogByMessageTypeCommandMessageTypeFlag == "" {
// 				filterActlogByMessageTypeCommand.Usage()
// 				return
// 			}

// 			cmd := actlog.NewFilterActlogByMessageTypeCommand(*filterActlogByMessageTypeCommandInputFileFlag, *filterActlogByMessageTypeCommandMessageTypeFlag, *filterActlogByMessageTypeCommandOutputFileFlag)
// 			cmd.FilterActlogByMessageType()
// 			// fmt.Println("Args:", os.Args[2:])
// 		}

// 	case "FILTERACTLOGBYKEYWORD", "FAK":
// 		filterActlogByKWCommand.Parse(os.Args[2:])
// 		if filterActlogByKWCommand.Parsed() {
// 			if *filterActlogByKWCommandInputFileFlag == "" || *filterActlogByKWCommandKeywordFlag == "" {
// 				filterActlogByKWCommand.Usage()
// 				return
// 			}

// 			cmd := actlog.NewFilterActlogByKWCommand(*filterActlogByKWCommandInputFileFlag, *filterActlogByKWCommandKeywordFlag, *filterActlogByKWCommandOutputFileFlag)
// 			cmd.FilterActlogByKeyword()
// 			// fmt.Println("Args:", os.Args[2:])
// 		}

// 	case "FILTERACTLOGBYREGEXP", "FAR":
// 		filterActlogByKWExtCommand.Parse(os.Args[2:])
// 		if filterActlogByKWExtCommand.Parsed() {
// 			if *filterActlogByKWExtCommandInputFileFlag == "" || *filterActlogByKWExtCommandKeywordFlag == "" {
// 				filterActlogByKWExtCommand.Usage()
// 				return
// 			}

// 			cmd := actlog.NewFilterActlogByKWExtCommand(*filterActlogByKWExtCommandInputFileFlag, *filterActlogByKWExtCommandKeywordFlag, *filterActlogByKWExtCommandOutputFileFlag, *filterActlogByKWExtCommandIgnoreCaseFlag, *filterActlogByKWExtCommandInvertMatchFlag)
// 			cmd.FilterActlogByKeywordExt()
// 			// fmt.Println("Args:", os.Args[2:])
// 		}

// 	case "FILTERACTLOGBYNODE", "FAN":
// 		filterActlogByNodesCommand.Parse(os.Args[2:])
// 		if filterActlogByNodesCommand.Parsed() {
// 			if *filterActlogByNodesCommandInputFileFlag == "" || *filterActlogByNodesCommandNodesFlag == "" {
// 				filterActlogByNodesCommand.Usage()
// 				return
// 			}
// 			nodes := strings.Split(*filterActlogByNodesCommandNodesFlag, ",")
// 			//
// 			cmd := actlog.NewFilterActlogByNodesCommand(*filterActlogByNodesCommandInputFileFlag, nodes, *filterActlogByNodesCommandOutputFileFlag)
// 			cmd.FilterActlogByNodes()
// 			// fmt.Println("Args:", os.Args[2:])
// 		}
// 	case "CONVERTACTLOGTOHTML", "CAH":
// 		convertActlog2HTMLCommand.Parse(os.Args[2:])
// 		if convertActlog2HTMLCommand.Parsed() {
// 			if *convertActlog2HTMLCommandInputFileFlag == "" {
// 				convertActlog2HTMLCommand.Usage()
// 				return
// 			}

// 			cmd := actlog.NewConvertActlog2HTMLCommand(*convertActlog2HTMLCommandInputFileFlag, *convertActlog2HTMLCommandOutputFileFlag)
// 			cmd.ConvertActlog2HTML()
// 			// fmt.Println("Args:", os.Args[2:])
// 		}
// 	case "VIEWACTLOG", "VA":
// 		viewActlogCommand.Parse(os.Args[2:])
// 		if viewActlogCommand.Parsed() {
// 			if *viewActlogCommandInputFileFlag == "" {
// 				viewActlogCommand.Usage()
// 				return
// 			}

// 			cmd := actlog.NewViewActlogCommand(*viewActlogCommandInputFileFlag)
// 			cmd.ViewActlog()
// 			// fmt.Println("Args:", os.Args[2:])
// 		}

// 	// case "WebUiforActlog", "WUA":
// 	// 	webuiActlogCommand.Parse(os.Args[2:])
// 	// 	if webuiActlogCommand.Parsed() {
// 	// 		if *webuiActlogCommandInputFileFlag == "" {
// 	// 			webuiActlogCommand.Usage()
// 	// 			return
// 	// 		}
// 	// 		cmd := webserver.NewWebUICommand(*webuiActlogCommandInputFileFlag, *webuiActlogCommandPort)
// 	// 		cmd.StartActlogServer()
// 	// 	}

// 	// case "WebUiOc", "WUO":
// 	// 	webuiOCCommand.Parse(os.Args[2:])
// 	// 	if webuiOCCommand.Parsed() {
// 	// 		if *webuiOCCommandInputFileFlag == "" {
// 	// 			webuiOCCommand.Usage()
// 	// 			return
// 	// 		}
// 	// 		cmd := webserver.NewWebUICommand(*webuiOCCommandInputFileFlag, *webuiOCCommandPort)
// 	// 		cmd.StartOCOpslogServer()
// 	// 	}
// 	// 	//===============client trace=============================
// 	// case "FIlTERCLIENTTRACEBYKEYWORD", "FCK":
// 	// 	filterCTByKWCommand.Parse(os.Args[2:])
// 	// 	if filterCTByKWCommand.Parsed() {
// 	// 		if *filterCTByKWCommandInputFileFlag == "" {
// 	// 			filterCTByKWCommand.Usage()
// 	// 			return
// 	// 		}
// 	// 		if *filterCTByKWCommandKeywordFlag == "" {
// 	// 			filterCTByKWCommand.Usage()
// 	// 			return
// 	// 		}

// 	// 		cmd := clienttrace.NewFilterCTByKWCommand(*filterCTByKWCommandInputFileFlag, *filterCTByKWCommandKeywordFlag, *filterCTByKWCommandOutputFileFlag, *filterCTByKWCommandCaseInSensitiveFlag)
// 	// 		cmd.FilterCTByKeyword()
// 	// 	}

// 	// case "FIlTERCLIENTTRACEBYTIME", "FCT":
// 	// 	filterCTByTimeCommand.Parse(os.Args[2:])
// 	// 	if filterCTByTimeCommand.Parsed() {
// 	// 		if *filterCTByTimeCommandInputFileFlag == "" {
// 	// 			filterCTByTimeCommand.Usage()
// 	// 			return
// 	// 		}
// 	// 		if *filterCTByTimeCommandBeginTimeFlag == "" || *filterCTByTimeCommandEndTimeFlag == "" {
// 	// 			filterCTByTimeCommand.Usage()
// 	// 			return
// 	// 		}

// 	// 		cmd := clienttrace.NewFilterCTByTimeCommand(*filterCTByTimeCommandInputFileFlag, *filterCTByTimeCommandBeginTimeFlag, *filterCTByTimeCommandEndTimeFlag, *filterCTByTimeCommandOutputFileFlag)
// 	// 		cmd.FilterCTByTime()
// 	// 	}

// 	// 	//split client trace
// 	// case "SPLITCLIENTTRACEBYTIME", "SCT":
// 	// 	splitCTByTimeCommand.Parse(os.Args[2:])
// 	// 	if splitCTByTimeCommand.Parsed() {
// 	// 		if *splitCTByTimeCommandInputFileFlag == "" {
// 	// 			splitCTByTimeCommand.Usage()
// 	// 			return
// 	// 		}

// 	// 		cmd := clienttrace.NewSplitCTByTimeCommand(*splitCTByTimeCommandInputFileFlag, *splitCTByTimeCommandMinutesFlag)
// 	// 		cmd.SplitByTime()
// 	// 	}

// 	// case "CONVERTCLIENTTRACE2HTML", "CCH":
// 	// 	convertCT2HTMLCommand.Parse(os.Args[2:])
// 	// 	if convertCT2HTMLCommand.Parsed() {
// 	// 		if *convertCT2HTMLCommandInputFileFlag == "" {
// 	// 			convertCT2HTMLCommand.Usage()
// 	// 			return
// 	// 		}

// 	// 		cmd := clienttrace.NewConvertClientTrace2HTMLCommand(*convertCT2HTMLCommandInputFileFlag, *convertCT2HTMLCommandOutputFileFlag)
// 	// 		cmd.ConvertClientTrace2HTML()
// 	// 		// fmt.Println("Args:", os.Args[2:])
// 	// 	}

// 	// //==============for device=========================
// 	// case "REDEFINEDEVICE", "RD":
// 	// 	redefineCommand.Parse(os.Args[2:])

// 	// 	if redefineCommand.Parsed() {
// 	// 		if *redefineCommandInputFileFlag == "" {
// 	// 			redefineCommand.Usage()
// 	// 			return
// 	// 		}
// 	// 		// cmd := device.NewRedefineCommand(*redefineCommandInputFileFlag, *redefineCommandOutputFileFlag, *isDebugFlag)
// 	// 		cmd := device.NewRedefineCommand(*redefineCommandInputFileFlag, *redefineCommandOutputFileFlag)
// 	// 		cmd.GenRedefineCMD()
// 	// 	}

// 	// case "CMPSLOTSLIBV", "CSL":
// 	// 	cmpSlotsLibvCommand.Parse(os.Args[2:])

// 	// 	if cmpSlotsLibvCommand.Parsed() {
// 	// 		if *cmpSlotsLibvCommandLIbvInputFileFlag == "" || *cmpSlotsLibvCommandSlotsInputFileFlag == "" {
// 	// 			cmpSlotsLibvCommand.Usage()
// 	// 			return
// 	// 		}
// 	// 		// cmd := device.NewRedefineCommand(*redefineCommandInputFileFlag, *redefineCommandOutputFileFlag, *isDebugFlag)
// 	// 		cmd := device.NewcmpSlotsLibvCommand(*cmpSlotsLibvCommandLIbvInputFileFlag, *cmpSlotsLibvCommandSlotsInputFileFlag, app.Config.Debug)

// 	// 		cmd.Compare()
// 	// 	}

// 	case "VERSION", "V":
// 		printVersion()

// 	case "HELP", "H":
// 		printHelp()

// 	default:
// 		// fmt.Println("Params:", os.Args[0:])
// 		fmt.Printf("%s is not valid command.\n", os.Args[1])

// 		flag.Usage()
// 		os.Exit(1)
// 	}
//}
