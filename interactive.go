package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	projectNamePrint = "Enter the name of the project: "
	makefilePrint    = "Makefile will be created? (y/n, default: y): "
	readmePrint      = "README.md will be created? (y/n, default: y): "
	runScriptsPrint  = "start.sh and stop.sh will be created? (y/n, default: y): "
)

func launchInteractiveMode(args *Args) {
	fmt.Printf("Create golang project.\n")

	reader := bufio.NewReader(os.Stdin)

	// Read project name
	args.ProjectName = readNotEmptyString(projectNamePrint, reader)

	// Read generate related flags
	args.GenerateMakefile = readYesOrNo(makefilePrint, reader)
	args.GenerateReadme = readYesOrNo(readmePrint, reader)
	args.GenerateRunScripts = readYesOrNo(runScriptsPrint, reader)
}

// read from the reader and get the first not empty string
func readNotEmptyString(print string, reader *bufio.Reader) (result string) {
	for result == "" {
		fmt.Printf(print)
		result, _ = reader.ReadString('\n')

		result = fetchFirst(result)
	}
	return
}

// read from the reader and get the first not empty string as yes or no.
// y/Y/yes/YES return true, n/N/No/NO return false. Empty also treated as y.
func readYesOrNo(print string, reader *bufio.Reader) bool {
	s := ""
	for s == "" {
		fmt.Printf(print)
		s, _ = reader.ReadString('\n')
		s = fetchFirst(s)

		if sUpper := strings.ToUpper(s); sUpper != "Y" && sUpper != "YES" &&
			sUpper != "N" && sUpper != "NO" {
			if s == "" { // treat empty as y
				s = "y"
			} else {
				s = ""
			}
		}
	}

	sUpper := strings.ToUpper(s)
	return sUpper == "Y" || sUpper == "YES"
}

func fetchFirst(s string) string {
	fields := strings.Fields(s)
	if len(fields) > 0 {
		return fields[0]
	} else {
		return ""
	}
}
