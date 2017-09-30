package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// TODO: 修改为当前目录
const parentDir = "/Users/alfred/Desktop"

const (
	makefileName    = "Makefile"
	readmeName      = "README.md"
	startScriptName = "start.sh"
	stopScriptName  = "stop.sh"
)

// Create the project root directory. Project full path will be
// returned if success.
func createRootDir(projectName string) (string, error) {
	if projectName == "" {
		return "", errors.New("project name can't be empty")
	}

	projectPath := parentDir + string(os.PathSeparator) + projectName
	return projectPath, os.Mkdir(projectPath, 0755)
}

const makefileTemplate = `GOOS=linux
GOARCH=amd64

DEPS=\
# github.com/go-sql-driver/mysql

build:
	@echo "Build program for $(GOOS) $(GOARCH)"
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o {{projectName}}

	@echo "Package program in dist directory"
	mkdir dist && mkdir dist/{{projectName}}
	cp *.sh dist/{{projectName}}
	cp ./{{projectName}} dist/{{projectName}}
	cp ./*.json dist/{{projectName}}
	cd dist; zip -r {{projectName}}-$(GOOS)-$(GOARCH).zip {{projectName}}
	rm -rf dist/{{projectName}}

clean: {{projectName}}
	rm -rf dist
	rm {{projectName}}

dep:
	@echo "Get dependency"
	for DEP in $(DEPS); do \
		go get $$DEP; \
	done
`

// Create the Makefile. Makefile path will be returned if success.
func createMakefile(projectPath string, projectName string) (string, error) {
	if projectPath == "" {
		return "", errors.New("project path can't be empty")
	}

	makefilePath := projectPath + string(os.PathSeparator) + makefileName
	f, err := os.Create(makefilePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	content := makefileTemplate
	content = strings.Replace(content, "{{projectName}}", projectName, -1)
	_, err = f.WriteString(content)
	return makefilePath, err
}

// Create the README.md. Readme file path will be returned if success.
func createReadme(projectPath string, projectName string) (string, error) {
	if projectPath == "" {
		return "", errors.New("project path can't be empty")
	}

	readmePath := projectPath + string(os.PathSeparator) + readmeName
	f, err := os.Create(readmePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("# %s\n", projectName))
	return readmePath, err
}

const startScriptTemplate = `#!/usr/bin/env bash

cd "$(dirname "$0")"
chmod u+x ./{{projectName}}
nohup ./{{projectName}} ./conf.json & echo $! > .pid`

const stopScriptTemplate = `#!/usr/bin/env bash
cat .pid | xargs kill -15`

// Create the start.sh and stop.sh.
func createRunScripts(projectPath string, projectName string) error {
	if projectPath == "" {
		return errors.New("project path can't be empty")
	}

	startPath := projectPath + string(os.PathSeparator) + startScriptName
	startFile, err := os.Create(startPath)
	if err != nil {
		return err
	}
	defer startFile.Close()

	content := startScriptTemplate
	content = strings.Replace(content, "{{projectName}}", projectName, -1)
	_, err = startFile.WriteString(content)
	if err != nil {
		return err
	}

	stopPath := projectPath + string(os.PathSeparator) + stopScriptName
	stopFile, err := os.Create(stopPath)
	if err != nil {
		return err
	}
	defer stopFile.Close()

	_, err = stopFile.WriteString(stopScriptTemplate)
	return err
}
