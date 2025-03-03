//Author: LiXue An(anlixue@cn.ibm.com)
//Description: actlog helper functions within package

package actlog

import "sptool/util"

func isSevere(line string) bool {
	if util.IsMatchRegexp(config.severe, line) {
		return true
	} else {
		return false
	}
}

func isError(line string) bool {
	if util.IsMatchRegexp(config.error, line) {
		return true
	} else {
		return false
	}
}

func isWarn(line string) bool {
	if util.IsMatchRegexp(config.warn, line) {
		return true
	} else {
		return false
	}
}

func isInfo(line string) bool {
	if isWarn(line) || isError(line) || isSevere(line) {
		return false
	} else {
		return true
	}
	// if util.IsMatchRegexp(config.info, line) {
	// 	return true
	// } else {
	// 	return false
	// }
}

func isMatchLineStart(word string) bool {
	if util.IsMatchRegexp(config.lineStart, word) {
		return true
	} else {
		return false
	}
}
