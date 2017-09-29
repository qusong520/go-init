package main

import (
	"flag"
	"fmt"
)

// arguments to create the project
type Args struct {
	ProjectName string // name of the project to be created

	GenerateMakefile   bool // generate Makefile
	GenerateReadme     bool // generate README.md
	GenerateRunScripts bool // generate start/stop scripts

	InteractiveMode bool // launch the interactive mode to get the args
}

func main() {
	args := parseArgs()

	if args.InteractiveMode {
		launchInteractiveMode(&args)
	}

	fmt.Printf("%#v\n", args)
}

// parsing the args
func parseArgs() (args Args) {
	flag.BoolVar(&args.GenerateMakefile, "m", true, "generate the Makefile")
	flag.BoolVar(&args.GenerateReadme, "r", true, "generate the README.md")
	flag.BoolVar(&args.GenerateRunScripts, "s", true, "generate the scripts to start/stop the program")
	flag.Parse()

	if flag.NArg() > 0 {
		args.ProjectName = flag.Arg(0)
	} else {
		args.InteractiveMode = true
	}
	return
}

func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
