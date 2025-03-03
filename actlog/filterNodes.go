//Author: LiXue An(anlixue@cn.ibm.com)
//Description: filter aclog by nodes name
package actlog

import (
	"fmt"
	"regexp"
	"sptool/util"
	"strings"
)

//====================================================================================
//filter actlog by nodes
//input filename, output filename and nodes list separate by ","
// 1. format actlog to line
// 2. filter line if match fixed message:ANR0406I "started by node NODENAME", save sessionID, or match Node name, save session
// 3. filter actlog by saved session ID
// 4. output to file
//filterByNodes
func filterByNodes(inputFileName string, outputFileName string, nodes []string) {
	outfile := outputFileName
	if outputFileName == "" {
		outfile = inputFileName + "-nodes-" + strings.Join(nodes, "-") + ".txt"
	}
	//get formated line
	actlogList := getFormatedActlog(inputFileName)
	//filter by nodes
	nodesList := getActLogByNodes(actlogList, nodes)
	if len(nodesList) > 0 {
		fmt.Println("Output to file --->  ", outfile)
		err := util.WriteLinesToFile(nodesList, outfile)
		if err != nil {
			fmt.Println("Write file error", err)
		}
	} else {
		fmt.Println("Output is empty! please check input file format and nodes name")
	}
}

// getActLogByNodes, if line match nodename or related session, return the list
// example:
// 02/06/20 10:28:57 ANR0406I Session 1248242 started for node NODE_TEST (DB2/AIX64) (SSL 10.0.10.161:46978). (SESSION: 1248242)
// func getActLogByNodes(lines []string, nodes []string) []string {
// 	retList := []string{}
// 	//get session list related this node
// 	sessionMap := make(map[string]string)
// 	for _, line := range lines {
// 		//找到node开始时候的session ID
// 		sessionID := getFirstSessionIDbyNodes(line, nodes)
// 		if sessionID != "" { //get first session id, save in map and add line to list
// 			if _, exists := sessionMap[sessionID]; !exists {
// 				sessionMap[sessionID] = sessionID
// 			}
// 		} else {
// 			//如果node开始的session信息不存在，按 node名称匹配，获取 (SESSION: xxxx)信息
// 			ssID := getMatchedSessionByNodes(line, nodes)
// 			if ssID != "" {
// 				// fmt.Println("Debug find node sessionID -->", ssID)
// 				if _, exists := sessionMap[ssID]; !exists {
// 					sessionMap[ssID] = ssID
// 				}
// 			}
// 		}
// 		// fmt.Println("Debug sessionID:", sessionMap)
// 		//if match session ID then append to list
// 		if isMatchSession(line, sessionMap) {
// 			retList = append(retList, line)
// 		}
// 	}
// 	return retList
// }

func getActLogByNodes(lines []string, nodes []string) []string {
	retList := []string{}
	//get session list related this node
	sessionMap := make(map[string]string)
	for _, line := range lines {
		//找到node开始时候的session ID
		if sessionID := getFirstSessionIDbyNodes(line, nodes); sessionID != "" {
			//get first session id, save in map and add line to list
			if _, exists := sessionMap[sessionID]; !exists {
				sessionMap[sessionID] = sessionID
				retList = append(retList, line)
				continue
			}
		}
		//如果node开始的session信息不存在，按 node名称匹配，获取 (SESSION: xxxx) sessionID, save in map
		if ssID := getMatchedSessionByNodes(line, nodes); ssID != "" {
			// if ssID != "" {
			// fmt.Println("Debug find node sessionID -->", ssID)
			if _, exists := sessionMap[ssID]; !exists {
				sessionMap[ssID] = ssID
				retList = append(retList, line)
				continue
			}
		}
		// fmt.Println("Debug sessionID:", sessionMap)
		//if match session ID then append to list
		if isMatchSession(line, sessionMap) {
			retList = append(retList, line)
		}
	}
	return retList
}

//getMatchedNodeSessionID get sessionID
//example line, match this, return session ID
// 02/06/20 10:28:57 ANR0406I Session 1248242 started for node NODE_TEST (DB2/AIX64) (SSL 10.0.10.161:46978). (SESSION: 1248242)
func getFirstSessionIDbyNodes(line string, nodes []string) string {
	fields := strings.Fields(line)
	if len(fields) >= 8 && fields[2] == "ANR0406I" && isValueInList(nodes, fields[8]) {
		return fields[4]
	}
	return ""
}

//isMatchSession
// if line contains SESSION: xxx, return true, sessionMap key is sessionID
func isMatchSession(line string, sessionMap map[string]string) bool {
	for k := range sessionMap {
		if strings.Contains(line, "SESSION: "+k) {
			return true
		}
	}
	return false
}

// slice list contains value return true
//isValueInList
func isValueInList(slice []string, val string) bool {
	for _, item := range slice {
		if strings.ToUpper(item) == strings.ToUpper(val) {
			return true
		}
	}
	return false
}

//getMatchedSessionByNodes if line contains node name, return the session ID
// for some actlog have no "started session info but line contains node info"
// match format:  "node nodename ", "node nodename," "node nodename:"
// keyword SESSION: 1234
func getMatchedSessionByNodes(line string, nodes []string) string {
	for _, node := range nodes {
		// fmt.Println("sessionID in map:", k)
		//avoid match wrong nodename,nodename with prefix " " and suffix " " or suffix ":"
		if strings.Contains(line, "node "+node+" ") ||
			strings.Contains(line, "node "+node+":") ||
			strings.Contains(line, "node "+node+",") {
			regexp, _ := regexp.Compile(`SESSION: \d*`)
			sessionStr := regexp.FindString(line) //SESSION: 1234
			// fmt.Println("Debug session:", sessionStr)
			ss := strings.ReplaceAll(sessionStr, ":", " ")
			fields := strings.Fields(ss)
			// fmt.Println("Debug:",fields)

			if len(fields) >= 2 {
				return fields[1]
			}
		}
	}
	return ""
}
