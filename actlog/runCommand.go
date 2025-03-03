package actlog

import (
	"flag"
	"os"
	"strings"
)

func RunCommand() {

	switch strings.ToUpper(os.Args[1]) {
	case "SPLITACTLOG", "SA":
		splitActlogCommand := flag.NewFlagSet("SplitActlog", flag.ExitOnError)
		splitActlogInputFileFlag := splitActlogCommand.String("f", "", "actlog file name")
		splitActlogCommandEnableByHourFlag := splitActlogCommand.Bool("t", false, "(Optional) Generate one file per hour")
		// splitActlogCommand.Usage = splitActlogUsage

		splitActlogCommand.Parse(os.Args[2:])
		if splitActlogCommand.Parsed() {
			if *splitActlogInputFileFlag == "" {
				splitActlogCommand.Usage()
				os.Exit(1)
			}
			cmd := NewSplitActlogCommand(*splitActlogInputFileFlag, *splitActlogCommandEnableByHourFlag)
			cmd.SplitActlog()
		}
		os.Exit(0)
	case "FILTERACTLOGBYSESSION", "FAS-": //replaced by new FAS
		//for filter actlog by session id  cmd
		filterActlogBySSCommand := flag.NewFlagSet("FilterActlogbySession", flag.ExitOnError)
		filterActlogBySSCommandInputFileFlag := filterActlogBySSCommand.String("f", "", `actlog file name `)
		filterActlogBySSCommandSessionIDFlag := filterActlogBySSCommand.String("s", "", "session id ")
		filterActlogBySSCommandOutputFileFlag := filterActlogBySSCommand.String("o", "", `(Optional) Output file name`)
		// filterActlogBySSCommand.Usage = filterActlogBySSUsage

		filterActlogBySSCommand.Parse(os.Args[2:])
		if filterActlogBySSCommand.Parsed() {
			if *filterActlogBySSCommandInputFileFlag == "" || *filterActlogBySSCommandSessionIDFlag == "" {
				filterActlogBySSCommand.Usage()
				os.Exit(1)
			}

			cmd := NewFilterActlogBySessionCommand(*filterActlogBySSCommandInputFileFlag, *filterActlogBySSCommandSessionIDFlag, *filterActlogBySSCommandOutputFileFlag)
			cmd.FilterActlogBySessionID()
			// fmt.Println("Args:", os.Args[2:])
		}
		os.Exit(0)
	case "FILTERACTLOGBYSESSIONS", "FAS":
		//for filter actlog by session id list cmd
		filterActlogBySessionsCommand := flag.NewFlagSet("FilterActlogbySessions", flag.ExitOnError)
		filterActlogBySessionsCommandInputFileFlag := filterActlogBySessionsCommand.String("f", "", `actlog file name  `)
		filterActlogBySessionsCommandSessionIDFlag := filterActlogBySessionsCommand.String("s", "", `session ID list,multiple IDs separate by comma ","`)
		filterActlogBySessionsCommandOutputFileFlag := filterActlogBySessionsCommand.String("o", "", `(Optional) Output file name `)
		// filterActlogBySSCommand.Usage = filterActlogBySSUsage

		filterActlogBySessionsCommand.Parse(os.Args[2:])
		if filterActlogBySessionsCommand.Parsed() {
			if *filterActlogBySessionsCommandInputFileFlag == "" || *filterActlogBySessionsCommandSessionIDFlag == "" {
				filterActlogBySessionsCommand.Usage()
				os.Exit(1)
			}

			cmd := NewFilterActlogBySessionIDsCommand(*filterActlogBySessionsCommandInputFileFlag, *filterActlogBySessionsCommandSessionIDFlag, *filterActlogBySessionsCommandOutputFileFlag)
			cmd.FilterActlogBySessionIDs()
			// fmt.Println("Args:", os.Args[2:])
		}
		os.Exit(0)
	case "FILTERACTLOGBYPROCESSES", "FAP":
		//for filter actlog by session id list cmd
		filterActlogByProcessesCommand := flag.NewFlagSet("FilterActlogbyProcesses", flag.ExitOnError)
		filterActlogByProcessesCommandInputFileFlag := filterActlogByProcessesCommand.String("f", "", `actlog file name  `)
		filterActlogByProcessesCommandProcessIDFlag := filterActlogByProcessesCommand.String("p", "", `process ID list,multiple IDs separate by comma ","`)
		filterActlogByProcessesCommandOutputFileFlag := filterActlogByProcessesCommand.String("o", "", `(Optional) Output file name `)
		// filterActlogBySSCommand.Usage = filterActlogBySSUsage

		filterActlogByProcessesCommand.Parse(os.Args[2:])
		if filterActlogByProcessesCommand.Parsed() {
			if *filterActlogByProcessesCommandInputFileFlag == "" || *filterActlogByProcessesCommandProcessIDFlag == "" {
				filterActlogByProcessesCommand.Usage()
				os.Exit(1)
			}

			cmd := NewFilterActlogByProcessIDsCommand(*filterActlogByProcessesCommandInputFileFlag, *filterActlogByProcessesCommandProcessIDFlag, *filterActlogByProcessesCommandOutputFileFlag)
			cmd.FilterActlogByProcessIDs()
			// fmt.Println("Args:", os.Args[2:])
		}
		os.Exit(0)

	case "FILTERACTLOGBYMESSAGETYPE", "FAT":
		//for message type filter
		filterActlogByMessageTypeCommand := flag.NewFlagSet("FilterActlogbyMessageType", flag.ExitOnError)
		filterActlogByMessageTypeCommandInputFileFlag := filterActlogByMessageTypeCommand.String("f", "", `actlog file name `)
		filterActlogByMessageTypeCommandMessageTypeFlag := filterActlogByMessageTypeCommand.String("t", "", `message type:[I,W,E,S],I:Information,W:Warning,E:Error,S:Severe,multiple values separate by comma "," `)
		filterActlogByMessageTypeCommandOutputFileFlag := filterActlogByMessageTypeCommand.String("o", "", `(Optional) Output file name`)
		filterActlogByMessageTypeCommandInvertMatchFlag := filterActlogByMessageTypeCommand.Bool("v", false, `(Optional) Invert match(not match specified type), for example: -t "I" -v equals -t "W,E,S" `)

		filterActlogByMessageTypeCommand.Parse(os.Args[2:])
		if filterActlogByMessageTypeCommand.Parsed() {
			if *filterActlogByMessageTypeCommandInputFileFlag == "" || *filterActlogByMessageTypeCommandMessageTypeFlag == "" {
				filterActlogByMessageTypeCommand.Usage()
				os.Exit(1)
			}

			cmd := NewFilterActlogByMessageTypeCommand(*filterActlogByMessageTypeCommandInputFileFlag, *filterActlogByMessageTypeCommandMessageTypeFlag, *filterActlogByMessageTypeCommandOutputFileFlag, *filterActlogByMessageTypeCommandInvertMatchFlag)
			cmd.FilterActlogByMessageType()
			// fmt.Println("Args:", os.Args[2:])
		}
		os.Exit(0)
	case "FILTERACTLOGBYKEYWORD", "FAK-":
		//filer by keyword
		filterActlogByKWCommand := flag.NewFlagSet("FilterActlogbyKeyword", flag.ExitOnError)
		filterActlogByKWCommandInputFileFlag := filterActlogByKWCommand.String("f", "", "actlog file name")
		filterActlogByKWCommandKeywordFlag := filterActlogByKWCommand.String("k", "", "keyword, the value is case sensitive")
		filterActlogByKWCommandOutputFileFlag := filterActlogByKWCommand.String("o", "", "(Optional) Output file name")
		// filterActlogByKWCommand.Usage = filterActlogByKWUsage

		filterActlogByKWCommand.Parse(os.Args[2:])
		if filterActlogByKWCommand.Parsed() {
			if *filterActlogByKWCommandInputFileFlag == "" || *filterActlogByKWCommandKeywordFlag == "" {
				filterActlogByKWCommand.Usage()
				os.Exit(1)
			}
			cmd := NewFilterActlogByKWCommand(*filterActlogByKWCommandInputFileFlag, *filterActlogByKWCommandKeywordFlag, *filterActlogByKWCommandOutputFileFlag)
			cmd.FilterActlogByKeyword()
			// fmt.Println("Args:", os.Args[2:])
		}
		os.Exit(0)
	case "FILTERACTLOGBYREGEXP", "FAR":
		filterActlogByKWExtCommand := flag.NewFlagSet("FilterActlogByRegExp", flag.ExitOnError)
		filterActlogByKWExtCommandInputFileFlag := filterActlogByKWExtCommand.String("f", "", "actlog file name")
		filterActlogByKWExtCommandKeywordFlag := filterActlogByKWExtCommand.String("e", "", "regular expression")
		filterActlogByKWExtCommandOutputFileFlag := filterActlogByKWExtCommand.String("o", "", "(Optional) Output file name")
		filterActlogByKWExtCommandIgnoreCaseFlag := filterActlogByKWExtCommand.Bool("i", false, "(Optional) ignore case")
		filterActlogByKWExtCommandInvertMatchFlag := filterActlogByKWExtCommand.Bool("v", false, "(Optional) Invert match(not match specified keyword")

		filterActlogByKWExtCommand.Parse(os.Args[2:])
		if filterActlogByKWExtCommand.Parsed() {
			if *filterActlogByKWExtCommandInputFileFlag == "" || *filterActlogByKWExtCommandKeywordFlag == "" {
				filterActlogByKWExtCommand.Usage()
				os.Exit(1)
			}
			cmd := NewFilterActlogByKWExtCommand(*filterActlogByKWExtCommandInputFileFlag, *filterActlogByKWExtCommandKeywordFlag, *filterActlogByKWExtCommandOutputFileFlag, *filterActlogByKWExtCommandIgnoreCaseFlag, *filterActlogByKWExtCommandInvertMatchFlag)
			cmd.FilterActlogByKeywordExt()
			// fmt.Println("Args:", os.Args[2:])
		}
		os.Exit(0)
	case "FILTERACTLOGBYNODE", "FAN-":
		//filter by node name
		filterActlogByNodesCommand := flag.NewFlagSet("FilterActlogbyNode", flag.ExitOnError)
		filterActlogByNodesCommandInputFileFlag := filterActlogByNodesCommand.String("f", "", "actlog file name")
		filterActlogByNodesCommandNodesFlag := filterActlogByNodesCommand.String("n", "", "node name, multiple nodes separate by ','")
		filterActlogByNodesCommandOutputFileFlag := filterActlogByNodesCommand.String("o", "", "(Optional) Output file ")
		// filterActlogByKWCommand.Usage = filterActlogByKWUsage

		filterActlogByNodesCommand.Parse(os.Args[2:])
		if filterActlogByNodesCommand.Parsed() {
			if *filterActlogByNodesCommandInputFileFlag == "" || *filterActlogByNodesCommandNodesFlag == "" {
				filterActlogByNodesCommand.Usage()
				os.Exit(1)
			}
			nodes := strings.Split(*filterActlogByNodesCommandNodesFlag, ",")
			//
			cmd := NewFilterActlogByNodesCommand(*filterActlogByNodesCommandInputFileFlag, nodes, *filterActlogByNodesCommandOutputFileFlag)
			cmd.FilterActlogByNodes()
			// fmt.Println("Args:", os.Args[2:])
		}
		os.Exit(0)
	case "CONVERTACTLOGTOHTML", "CAH":
		//convert actlog to html
		convertActlog2HTMLCommand := flag.NewFlagSet("ConvertActlogToHtml", flag.ExitOnError)
		convertActlog2HTMLCommandInputFileFlag := convertActlog2HTMLCommand.String("f", "", `actlog file name `)
		convertActlog2HTMLCommandOutputFileFlag := convertActlog2HTMLCommand.String("o", "", `(Optional) Output file `)
		// convertActlog2HTMLCommand.Usage = convertActlog2HTMLUsage

		convertActlog2HTMLCommand.Parse(os.Args[2:])
		if convertActlog2HTMLCommand.Parsed() {
			if *convertActlog2HTMLCommandInputFileFlag == "" {
				convertActlog2HTMLCommand.Usage()
				os.Exit(1)
			}

			cmd := NewConvertActlog2HTMLCommand(*convertActlog2HTMLCommandInputFileFlag, *convertActlog2HTMLCommandOutputFileFlag)
			cmd.ConvertActlog2HTML()
			// fmt.Println("Args:", os.Args[2:])
		}
		os.Exit(0)
	case "VIEWACTLOG", "VA":
		viewActlogCommand := flag.NewFlagSet("ViewActlog", flag.ExitOnError)
		viewActlogCommandInputFileFlag := viewActlogCommand.String("f", "", "actlog file name")

		viewActlogCommand.Parse(os.Args[2:])
		if viewActlogCommand.Parsed() {
			if *viewActlogCommandInputFileFlag == "" {
				viewActlogCommand.Usage()
				os.Exit(1)
			}
			cmd := NewViewActlogCommand(*viewActlogCommandInputFileFlag)
			cmd.ViewActlog()
			// fmt.Println("Args:", os.Args[2:])
		}
		os.Exit(0)
	}
}
