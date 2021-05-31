package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type ExecuteSuperBatch struct {
	parameters map[string]string
}

func (batch ExecuteSuperBatch) execute(filename string) {
	if _, err := os.Stat(filename); err != nil {
		content, err := ioutil.ReadFile(filename)
		if err != nil {
			// log.Fatal(err)
			fmt.Println(err)
		}
		fmt.Print(content)
	} else {
		fmt.Println(err)
	}
}

func performCommand(command string, parameters map[string]string) {
	// perform commands from the command line
	if len(command) == 0 {
		fmt.Print("HELP")
	} else {
		if command == "help" {
			fmt.Print("Help")
		} else {
			ExecuteSuperBatch{parameters: parameters}.execute(command)
		}
	}
}

// The argument parser type used to
// parse command line arguments
type ArgumentParser struct {
	// The command line arguments
	arguments []string
	// The length of the command line
	// arguments
	length int
}

// The function that parser the command line arguments
// and seperates them into commands and parameters
// Example syntax - <exe> <commad> <param1>=<value1>
func (parser ArgumentParser) parse() (string, map[string]string) {
	var command string = ""
	var parameters map[string]string = make(map[string]string)
	for index := 0; index < parser.length; index++ {
		var currentParserArgument string = parser.arguments[index]
		// The first elemenet in the arguments is
		// considered as the command to run
		if index == 0 {
			command = currentParserArgument
			continue
		}

		var statement []string = strings.Split(currentParserArgument, "=")
		key, value := statement[0], ""
		if len(statement) > 1 {
			value = statement[1]
		}

		parameters[key] = value
	}
	return command, parameters
}

func main() {
	var arguments []string = os.Args[1:]
	var parser = ArgumentParser{arguments: arguments, length: len(arguments)}
	command, parameters := parser.parse()
	performCommand(command, parameters)
}
