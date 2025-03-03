//Author: LiXue An(anlixue@cn.ibm.com)
//Description: convert actlog to html file
package actlog

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"time"
)

const htmlTemplacte = `
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta
      name="viewport"
      content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0"
    />
  <title>sptool generated actlog </title>
<style>
	.warning {
		color: orange;
		#background-color: #EEEEEE;
		width: 90%;
		margin: 2px auto 2px 20px;
		border: 1px solid #eee;
		box-shadow: 0 2px 3px #ccc;
		padding: 3px;
		/*text-align: center*/
	  }
	  
	  .error {
		/* color: white; */
		/* background-color: #48D1CC ; */
		color: white;
		background-color: #d62d20; 
		width: 90%;
		margin: 2px auto 2px 40px;
		border: 1px solid #eee;
		box-shadow: 0 2px 3px #ccc;
		padding: 3px;
	  }
	  
	  .severe {
		/* color: 	#8B0000;  800000*/
		color: white;
		/* background-color: #000000; */
		background-color: #B00020;
		/* #40E0D0 ; */
		width: 90%;
		margin: 2px auto 2px 40px;
		border: 1px solid #eee;
		box-shadow: 0 2px 3px #ccc;
		padding: 3px;
	  }
	  
	  .info {
		color: green;
		width: 90%;
		/* top right bottom left  */
		margin: 8px auto 8px 20px;
		/*
		  border: 1px solid #eee;
		  box-shadow: 0 2px 3px #ccc;
		  padding: 16px;
		   */
	  }	
	  </style>
</head>
<body>
<script>
//for hide
	function hideInfo() {
		var infoElements=document.getElementsByClassName('info')
    	for ( var i=0;i< infoElements.length ; i++ ) {
          infoElements[i].style.display='none'
        }
}
function displayInfo() {
	var infoElements=document.getElementsByClassName('info')
	for ( var i=0;i< infoElements.length; i++ ) {
	  infoElements[i].style.display='block'
	}
  }

</script>

{{ range .}}
<div class="{{.CSSClass}}"> {{.Line}}</div>
{{ end }}
</body>
</html>
`

type htmlLineElement struct {
	CSSClass string
	Line     string
}

func convert2HTML(actlogFileName string, outputFileName string) {
	outFile := outputFileName
	if outputFileName == "" {
		outFile = actlogFileName + "-sptool-convert-" + time.Now().Format("150405") + ".html"
		if actlogFileName == "-" {
			outFile = "sptool-convert-" + time.Now().Format("150405") + ".html"
		}
	}

	htmlLines := []htmlLineElement{}
	lines := getFormatedActlog(actlogFileName)

	for _, line := range lines {
		//if util.IsMatchRegexp(`AN.\d{4}S`, line) || util.IsMatchRegexp(`AN.\d{4}D`, line) || util.IsMatchRegexp(`AN.\d{4}K`, line) {
		//if util.IsMatchRegexp(config.severe, line) {
		if isSevere(line) {
			htmlLines = append(htmlLines, htmlLineElement{CSSClass: "severe", Line: line})
			continue
		}
		//if util.IsMatchRegexp(`AN.\d{4}E`, line) {
		// if util.IsMatchRegexp(config.error, line) {
		if isError(line) {
			htmlLines = append(htmlLines, htmlLineElement{CSSClass: "error", Line: line})
			continue
		}
		//if util.IsMatchRegexp(`AN.\d{4}W`, line) {
		// if util.IsMatchRegexp(config.warn, line) {
		if isWarn(line) {
			htmlLines = append(htmlLines, htmlLineElement{CSSClass: "warning", Line: line})
			continue
		}
		htmlLines = append(htmlLines, htmlLineElement{CSSClass: "info", Line: line})
	}

	if len(htmlLines) <= 0 {
		fmt.Println("Output is empty! please check input file format")
		return
	}

	//write to html file
	t := template.Must(template.New("html-tmpl").Parse(htmlTemplacte))

	file, err := os.Create(outFile)
	defer file.Close()
	if err != nil {
		fmt.Println("Write file error", outFile, err)
	}
	w := bufio.NewWriter(file)
	defer w.Flush()
	// errs := t.Execute(os.Stdout, htmlLines)
	fmt.Println("Output to file --->", outFile)
	errs := t.Execute(w, htmlLines)
	if errs != nil {
		fmt.Println("Generate html error", err)
	}
}
