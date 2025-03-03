package terminalcolor

var terminalStyle = struct {
	severeStart string
	errorStart  string
	warnStart   string
	infoStart   string
	end         string
}{
	severeStart: "\033[37;1;41m", // White text bold with red background ,
	errorStart:  "\033[1;31m",    // bold red text
	warnStart:   "\033[1;33m",    //Yellow bold text
	infoStart:   "\033[;32m",     //Green Text
	end:         "\033[0m",
}

func InfoColor(line string) string {
	return terminalStyle.infoStart + line + terminalStyle.end
}

func WarnColor(line string) string {
	return terminalStyle.warnStart + line + terminalStyle.end
}

func ErrorColor(line string) string {
	return terminalStyle.errorStart + line + terminalStyle.end
}

func SevereColor(line string) string {
	return terminalStyle.severeStart + line + terminalStyle.end
}
