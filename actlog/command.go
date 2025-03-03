//Author: LiXue An(anlixue@cn.ibm.com)
//Description: actlog commands entry
package actlog

import (
	"fmt"
	"os"
	"runtime"
	"sptool/app"
	"sptool/cmdlog"
	"strings"
)

//FilterActlogBySessionCommand for command params
type ViewActlogCommand struct {
	inputFileName string
}

//FilterActlogBySessionCommand for command params
type FilterActlogBySessionCommand struct {
	inputFileName  string
	sessionID      string
	outputFileName string
}

//FilterActlogBySessionIDsCommand for command params
type FilterActlogBySessionIDsCommand struct {
	inputFileName  string
	sessionIDs     string
	outputFileName string
}

type FilterActlogByProcessIDsCommand struct {
	inputFileName  string
	processIDs     string
	outputFileName string
}

//
type FilterActlogByMessageTypeCommand struct {
	inputFileName   string
	messageTypeList string
	outputFileName  string
	invertMatch     bool
}

//FilterActlogByKWCommand filter by keyword
type FilterActlogByKWCommand struct {
	inputFileName  string
	keyword        string
	outputFileName string
}

//FilterActlogByKWExtCommand extension for ignoreCase and invert match
type FilterActlogByKWExtCommand struct {
	inputFileName  string
	keyword        string
	ignoreCae      bool
	outputFileName string
	invertMatch    bool
}

//FilterActlogByNodesCommand filter by nodes
type FilterActlogByNodesCommand struct {
	inputFileName  string
	nodes          []string
	outputFileName string
}

//SplitActlogCommand : implement cmd
type SplitActlogCommand struct {
	inputFileName string
	byHour        bool
}

//ConvertActlog2HTMLCommand for convert to html
type ConvertActlog2HTMLCommand struct {
	inputFileName  string
	outputFileName string
}

//NewSplitActlogCommand : return  NewSplitActlogCommand object
func NewSplitActlogCommand(inputFile string, byHour bool) *SplitActlogCommand {
	return &SplitActlogCommand{inputFileName: inputFile, byHour: byHour}
}

//SplitActlog split by date
func (cmd *SplitActlogCommand) SplitActlog() {
	LoadConfigFromOSEnv("SPTOOL_ACTLOG_")
	app.Debug("Start split actlog for file:", fmt.Sprintf("%+v", cmd))
	app.Debug("Actlog config:", fmt.Sprintf("%+v", config))
	if cmd.byHour {
		splitActlogByHour(cmd.inputFileName)
	} else {
		splitActlogByDate(cmd.inputFileName)
	}
	cmdlog.Info("sa " + fmt.Sprintf("%+v", *cmd))
}

//NewFilterActlogBySessionCommand for cmd execution
func NewFilterActlogBySessionCommand(inputFile string, sessionID string, outputFile string) *FilterActlogBySessionCommand {
	return &FilterActlogBySessionCommand{inputFileName: inputFile, sessionID: sessionID, outputFileName: outputFile}
}

//NewFilterActlogBySessionIDsCommand for cmd execution
func NewFilterActlogBySessionIDsCommand(inputFile string, sessionIDs string, outputFile string) *FilterActlogBySessionIDsCommand {
	return &FilterActlogBySessionIDsCommand{inputFileName: inputFile, sessionIDs: sessionIDs, outputFileName: outputFile}
}

//NewFilterActlogByProcessIDsCommand for cmd execution
func NewFilterActlogByProcessIDsCommand(inputFile string, processIDs string, outputFile string) *FilterActlogByProcessIDsCommand {
	return &FilterActlogByProcessIDsCommand{inputFileName: inputFile, processIDs: processIDs, outputFileName: outputFile}
}

//NewFilterActlogByMessageTypeCommand
func NewFilterActlogByMessageTypeCommand(inputFile string, messageTypeList string, outputFile string, invertMatch bool) *FilterActlogByMessageTypeCommand {
	return &FilterActlogByMessageTypeCommand{inputFileName: inputFile, messageTypeList: messageTypeList, outputFileName: outputFile, invertMatch: invertMatch}
}

//FilterActlogBySessionID execute cmd
func (cmd *FilterActlogBySessionCommand) FilterActlogBySessionID() {
	LoadConfigFromOSEnv("SPTOOL_ACTLOG_")
	app.Debug("Start filter actlog by session:", fmt.Sprintf("%+v", cmd))
	app.Debug("Actlog config:", fmt.Sprintf("%+v", config))
	filterBySessionID(cmd.inputFileName, cmd.outputFileName, cmd.sessionID)
	app.Debug("End filter actlog by session")

	cmdlog.Info("fas " + fmt.Sprintf("%+v", *cmd))
}

//FilterActlogBySessionID execute cmd
func (cmd *FilterActlogBySessionIDsCommand) FilterActlogBySessionIDs() {
	LoadConfigFromOSEnv("SPTOOL_ACTLOG_")
	app.Debug("Start filter actlog by session id list:", fmt.Sprintf("%+v", cmd))
	app.Debug("Actlog config:", fmt.Sprintf("%+v", config))
	filterBySessionIDList(cmd.inputFileName, cmd.outputFileName, cmd.sessionIDs)
	app.Debug("End filter actlog by session id list")
	cmdlog.Info("fass " + fmt.Sprintf("%+v", *cmd))
}

//FilterActlogByProcessIDs execute cmd
func (cmd *FilterActlogByProcessIDsCommand) FilterActlogByProcessIDs() {
	LoadConfigFromOSEnv("SPTOOL_ACTLOG_")
	app.Debug("Start filter actlog by process id list:", fmt.Sprintf("%+v", cmd))
	app.Debug("Actlog config:", fmt.Sprintf("%+v", config))
	filterByProcessIDList(cmd.inputFileName, cmd.outputFileName, cmd.processIDs)
	app.Debug("End filter actlog by process id list")
	cmdlog.Info("fap " + fmt.Sprintf("%+v", *cmd))
}

//FilterActlogBySessionID execute cmd
func (cmd *FilterActlogByMessageTypeCommand) FilterActlogByMessageType() {
	if !validateMessageType(cmd.messageTypeList) {
		fmt.Println(`Invalid message type, valide value are "I","W","E","S" or specify multiple value separate by comma, for example "W,E"`)
		os.Exit(1)
	}
	LoadConfigFromOSEnv("SPTOOL_ACTLOG_")
	app.Debug("Start filter actlog by message type list:", fmt.Sprintf("%+v", cmd))
	app.Debug("Actlog config:", fmt.Sprintf("%+v", config))
	filterByMessageType(cmd.inputFileName, cmd.outputFileName, cmd.messageTypeList, cmd.invertMatch)
	app.Debug("End filter actlog by by message type list")
	cmdlog.Info("fat " + fmt.Sprintf("%+v", *cmd))
}

func validateMessageType(messageType string) bool {
	messageSlice := strings.Split(messageType, ",")
	for _, kw := range messageSlice {
		if strings.ToUpper(kw) != "I" && strings.ToUpper(kw) != "W" && strings.ToUpper(kw) != "E" && strings.ToUpper(kw) != "S" {
			return false
		}
	}
	return true
}

//NewFilterActlogByKWCommand for cmd execution
func NewFilterActlogByKWCommand(inputFile string, keyword string, outputFile string) *FilterActlogByKWCommand {
	return &FilterActlogByKWCommand{inputFileName: inputFile, keyword: keyword, outputFileName: outputFile}
}

//NewFilterActlogByKWExtCommand for cmd execution
func NewFilterActlogByKWExtCommand(inputFile string, keyword string, outputFile string, ignoreCase bool, invertMatch bool) *FilterActlogByKWExtCommand {
	return &FilterActlogByKWExtCommand{inputFileName: inputFile, keyword: keyword, outputFileName: outputFile, ignoreCae: ignoreCase, invertMatch: invertMatch}
}

//FilterActlogByKeyword execute cmd
func (cmd *FilterActlogByKWCommand) FilterActlogByKeyword() {
	LoadConfigFromOSEnv("SPTOOL_ACTLOG_")
	app.Debug("Start filter actlog by keyword:", fmt.Sprintf("%+v", cmd))
	app.Debug("Actlog config:", fmt.Sprintf("%+v", config))
	filterByKeyWordRegExp(cmd.inputFileName, cmd.outputFileName, cmd.keyword)
	app.Debug("End filter actlog by keyword")
	cmdlog.Info("fak " + fmt.Sprintf("%+v", *cmd))
}

//FilterActlogByKeyword execute cmd
func (cmd *FilterActlogByKWExtCommand) FilterActlogByKeywordExt() {
	// fmt.Printf("execute cmd: filterByKeyWordRegExpExt: %+v", cmd)
	LoadConfigFromOSEnv("SPTOOL_ACTLOG_")
	app.Debug("Start filter actlog by regular expression:", fmt.Sprintf("%+v", cmd))
	app.Debug("Actlog config:", fmt.Sprintf("%+v", config))
	filterByKeyWordRegExpExt(cmd.inputFileName, cmd.outputFileName, cmd.keyword, cmd.ignoreCae, cmd.invertMatch)
	app.Debug("End filter actlog by regular expression")
	cmdlog.Info("far " + fmt.Sprintf("%+v", *cmd))
}

//NewFilterActlogByNodesCommand for cmd execution
func NewFilterActlogByNodesCommand(inputFile string, nodes []string, outputFile string) *FilterActlogByNodesCommand {
	return &FilterActlogByNodesCommand{inputFileName: inputFile, nodes: nodes, outputFileName: outputFile}
}

// FilterActlogByNodes execute cmd
func (cmd *FilterActlogByNodesCommand) FilterActlogByNodes() {
	LoadConfigFromOSEnv("SPTOOL_ACTLOG_")
	app.Debug("Start filter actlog by nodes:", fmt.Sprintf("%+v", cmd))
	app.Debug("Actlog config:", fmt.Sprintf("%+v", config))

	fmt.Println("Filter actlog by nodes -->", cmd.nodes)
	filterByNodes(cmd.inputFileName, cmd.outputFileName, cmd.nodes)
	app.Debug("End filter actlog by nodes")
	// cmdlog.Info("fan " + fmt.Sprintf("%+v", *cmd))
}

//NewConvertActlog2HTMLCommand for cmd instance
func NewConvertActlog2HTMLCommand(inputFile string, outputFile string) *ConvertActlog2HTMLCommand {
	return &ConvertActlog2HTMLCommand{inputFileName: inputFile, outputFileName: outputFile}
}

//ConvertActlog2HTML convert actions
func (cmd *ConvertActlog2HTMLCommand) ConvertActlog2HTML() {
	LoadConfigFromOSEnv("SPTOOL_ACTLOG_")
	app.Debug("Start convert actlog to HTML:", fmt.Sprintf("%+v", cmd))
	app.Debug("Actlog config:", fmt.Sprintf("%+v", config))

	convert2HTML(cmd.inputFileName, cmd.outputFileName)
	app.Debug("End convert actlog to HTML ")
	cmdlog.Info("cah " + fmt.Sprintf("%+v", *cmd))
}

func NewViewActlogCommand(inputFile string) *ViewActlogCommand {
	return &ViewActlogCommand{inputFileName: inputFile}
}

func (cmd *ViewActlogCommand) ViewActlog() {
	if runtime.GOOS == "windows" {
		fmt.Println("The command 'va' does not support windows platform")
		os.Exit(1)
	}
	LoadConfigFromOSEnv("SPTOOL_ACTLOG_")
	app.Debug("Start view actlog:", fmt.Sprintf("%+v", cmd))
	app.Debug("Actlog config:", fmt.Sprintf("%+v", config))
	viewActlog(cmd.inputFileName)
	app.Debug("End view actlog")
	cmdlog.Info("va " + fmt.Sprintf("%+v", *cmd))
}
