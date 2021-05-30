package main

import (
	"os"
	"fmt"
	"strings"
)

type ArgumentParser struct {
	arguments []string
	length int
}

func (parser ArgumentParser) parse() (string, map[string]string) {
	var command string = ""
	var parameters map[string]string = make(map[string]string)
	for index := 0; index < parser.length; index++ {
		var currentParserArgument string = parser.arguments[index]
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
	var parser = ArgumentParser{arguments:arguments, length:len(arguments)}
	command, parameters := parser.parse()
	fmt.Println(command, parameters)
}
