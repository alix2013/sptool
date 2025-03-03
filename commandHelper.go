//Author: LiXue An(anlixue@cn.ibm.com)
//Description: command helper for usage printing

package main

import "fmt"

const versionInfo = `Spectrum Protect Support Tools 
Version: 1.8.0

Version history: 
V1.0.0	Initial version,Split actlog by date(one file per day) and time(24 files per day)
V1.1.0  Convert actlog to HTML file
V1.1.1  Display Information,Warning,Error and Severe type messages of actlog with different style in HTML file
V1.1.2	Support Linux(x86_64), Windows(x86_64), AIX 7.2 on Power8(ppc64)
V1.2.0	Filter actlog by session ID
V1.2.1	Filter actlog by multiple session IDs
V1.3.0  Filter actlog by message type("I":Information,"W":Warning,"E":Error,"S":Severe)
V1.3.1	Filter actlog by message type support invert-match(exclude specified type)
V1.4.0	Filter actlog by process IDs
V1.4.1	Support Linux on Power(ppc64le),Linux on zOS(s390x)
V1.5.0	Filter actlog by keyword
V1.6.0	Filter actlog by regular expression
V1.7.0	View actlog on Unix terminal,display Info,Warn,Error and Severe message with different color
V1.8.0  Support pipeline("-":STDIN or STDOUT)
`

//Report bugs or enhence feature: LiXue An(anlixue@cn.ibm.com)
//V1.8.0  Support OS pipeline, "-" as standard input and output
// V1.1.1  Support various date format of actlog
//,i.e sptool fat -t W,E,S -f qact.txt -o - | ./sptool va -f -

const appUsageStr = `Usage: sptool <command> [<args>]

Commands for Spectrum Protect Actlog:
sa,  SplitActlog               Split actlog file by day or hour
cah, ConvertActlogToHtml       Convert actlog file to html file
fas, FilterActlogBySessions    Filter actlog by sessions 
fap, FilterActlogByProcesses   Filter actlog by processes
fat, FilterActlogByMessageType Filter actlog by message type
far, FilterActlogByRegExp      Filter actlog by regular expression
va,  ViewActlog                View actlog on Unix terminal

Other Commands:
v,   Version                   Print version information
h,   Help                      Print Help

Report bugs or enhence feature: LiXue An(anlixue@cn.ibm.com)
`

//fak, FilterActlogByKeyword     Filter actlog by keyword
//fas, FilterActlogBySession     Filter actlog by session ID
// fak	FilterActlogbyKeyword,Filter actlog by keyword
// wua	WebUiviewActlog,Start Web UI to view actlog

// rd	RedefineDevice,Generate re-define Library/Drive/Path cmds
//
//wua	WebUiforActlog, start Web UI to view actlog
//wuops	WebUiviewocOPSlog, start Web UI to view operation center tsm_opscntr log file
// h	Help,more examples

const helpStr = `
Name:
	sptool - Spectrum Protect Support Tool

Description:
	sptool is a simple utility for IBM Spectrum Protect(TSM) support,
It can run on Windows(x86_64),AIX(V7.2 on Power8),Linux(x86_64,ppc64le,s390x),MacOS(x86_64)

Usage:
	sptool <command> [<args>]
* Each command has short name and long name, for example: sa" equals "SplitActlog".
* Each arg may be required or optional, if it's an optional and not specified value, the tool will assign a default value.
* If an arg need multiple values, separate them by comma

Examples:
- Split spectrum protect/TSM actlog file 
sptool sa -f qactlog.txt or sptool SplitActlog -f qactlog.txt

- Convert spectrum protect/TSM actlog to HTML file:
sptool cah -f qactlog.txt 

- Filter spectrum protect/TSM actlog by session IDs
sptool fas -f qactlog.txt -s 123,234,345

- Filter spectrum protect/TSM actlog by process IDs
sptool fap -f qactlog.txt -p 123,234,345

- Filter spectrum protect/TSM actlog by message type
sptool fat -f qactlog.txt -t W,E,S
sptool fat -f qactlog.txt -t I -v (exclude information messsages, equals to -t W,E,S )

- Filter spectrum protect/TSM actlog by regular express 
sptool far -f qactlog.txt -e "drive[0-1]|library" -i
Case insensitive match express "drive[0-1]|library"

sptool far -f qactlog.txt -e "ANR2034E" -v (exclude expression ANR2034E)

- View actlog in Unix terminal, high light warning,error and severe messages
sptool va -f qactlog.txt

- Unix pipeline 
sptool fat -f qactlog.txt -t W,E,S -o - | sptool cah -f -

`

//application usage output
func appUsage() {
	fmt.Println(appUsageStr)
}

//version info
func printVersion() {
	fmt.Println(versionInfo)
}

func printHelp() {
	// `
	// Examples:
	// # Split actlog file by day
	// sptool sa -f qactlog.txt

	// # Split actlog file by day + hour
	// sptool sa -f qactlog.txt -t

	// # Filter actlog file by session id
	// sptool fas -f qactlog.txt -s 1234

	// # Filter actlog file by keyword or regular expression
	// sptool fak -f qactlog.txt -k "NODE1"
	// sptool fak -f qactlog.txt -k "ANR\d{4}E"
	// sptool fak -f qactlog.txt -k "DRIVE[0-4]"
	//`
	fmt.Println(helpStr)
}

//usage for device cmd
// func deviceUsage() {
// 	const usageString = `
// Usage of redefineDevice|rd:
// -f string
// 	dsmamdc> q path f=d > qpath.txt, i.e -f qpath.txt
// -o string
// 	(Optional) generated cmd lines output file name, default is stdout
// `
// 	fmt.Println(usageString)
// }

// // for single redefine device application
// func main() {

// 	// isDebug := false
// 	inputFileName := flag.String("f", "", "dsmadmc> q path f=d > qpath.txt output file,  i.e: -f qpath.txt")
// 	outputFileName := flag.String("o", "", "output file name")
// 	verboseFlag := flag.Bool("v", false, "Verbose output")
// 	flag.Usage = deviceUsage
// 	flag.Parse()
// 	if *inputFileName == "" {
// 		// flag.PrintDefaults()
// 		flag.Usage()
// 		os.Exit(1)
// 	}

// 	cmd := device.NewRedefineCommand(*inputFileName, *outputFileName, *verboseFlag)
// 	cmd.GenRedefineCMD()

// }

// //for split actlog
// func main() {
// 	inputFileName := flag.String("f", "", "dsmadmc> q actlog > qactlog.txt,  i.e: -f qact.txt, default split by date")
// 	// outputFileName := flag.String("o", "", "output file name")
// 	byHourFlag := flag.Bool("t", false, "split actlog by hour")
// 	// flag.Usage = deviceUsage
// 	flag.Parse()
// 	if *inputFileName == "" {
// 		flag.Usage()
// 		os.Exit(1)
// 	}
// 	if *byHourFlag {
// 		actlog.SplitActlogByHour(*inputFileName)
// 	} else {
// 		actlog.SplitActlog(*inputFileName)
// 	}
// 	//actlog.SplitActlog(*inputFileName, *byHourFlag)
// 	// actlog.SplitActlog(*inputFileName)

// }

// func splitActlogUsage() {
// 	const usageString = `
// Usage of splitActlog|sa:
// -f string
// 	dsmamdc> q actlog > qact.txt, i.e -f qact.txt
// -t
// 	(Optional) split by hour
// `
// 	fmt.Println(usageString)

// }

// func filterActlogBySSUsage() {
// 	const usageString = `
// Usage of filterActlogBySessionId|fas:
// -f string
// 	dsmamdc> q actlog > qact.txt, i.e -f qact.txt
// -s string
// 	session id
// -o string
// 	(Optional) output file name
// `
// 	fmt.Println(usageString)

// }

// func filterActlogByKWUsage() {
// 	const usageString = `
// Usage of filterActlogBySessionId|fas:
// -f string
// 	dsmamdc> q actlog > qact.txt, i.e -f qact.txt
// -k string
// 	keyword	or regular expression
// -o string
// 	(Optional) output file name
// `
// 	fmt.Println(usageString)

// }

// func convertActlog2HTMLUsage() {
// 	const usageString = `
// Usage of convertActlog2HTML|cah:
// -f string
// 	dsmamdc> q actlog > qact.txt, i.e -f qact.txt
// -o string
// 	(Optional) output file name
// `
// 	fmt.Println(usageString)

// }
