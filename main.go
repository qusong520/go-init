package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	exitCodeCreateRootFail       = 1
	exitCodeCreateMakefileFail   = 2
	exitCodeCreateReadmeFail     = 3
	exitCodeCreateRunScriptsFail = 4
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

	// create the root directory of the project
	projectPath, err := createRootDir(args.ProjectName)
	if err != nil {
		fmt.Println(err)
		os.Exit(exitCodeCreateRootFail)
	}

	// Makefile
	if args.GenerateMakefile {
		_, err := createMakefile(projectPath, args.ProjectName)
		if err != nil {
			fmt.Println(err)
			os.Exit(exitCodeCreateMakefileFail)
		}
	}

	// Readme
	if args.GenerateReadme {
		_, err := createReadme(projectPath, args.ProjectName)
		if err != nil {
			fmt.Println(err)
			os.Exit(exitCodeCreateReadmeFail)
		}
	}

	// Run scripts
	if args.GenerateRunScripts {
		err := createRunScripts(projectPath, args.ProjectName)
		if err != nil {
			fmt.Println(err)
			os.Exit(exitCodeCreateRunScriptsFail)
		}
	}

	fmt.Printf("Project created: %s\n", args.ProjectName)
}

// Parsing the args
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
