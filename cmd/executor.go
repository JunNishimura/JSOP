package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/JunNishimura/jsop/evaluator"
	"github.com/JunNishimura/jsop/lexer"
	"github.com/JunNishimura/jsop/object"
	"github.com/JunNishimura/jsop/parser"
)

func Run() error {
	filePath, err := parseCmdArgs()
	if err != nil {
		return err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("fail to open file: %s", err)
	}
	defer file.Close()

	// read file at once
	bytes, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("fail to read file: %s", err)
	}

	// parse program
	l := lexer.New(string(bytes))
	p := parser.New(l)
	program, err := p.ParseProgram()
	if err != nil {
		return fmt.Errorf("fail to parse program: %s", err)
	}

	// evaluate program
	env := object.NewEnvironment()

	// define macros
	if err := evaluator.DefineMacros(program, env); err != nil {
		return fmt.Errorf("fail to define macros: %s", err)
	}

	// expand macros
	expanded := evaluator.ExpandMacros(program, env)

	// evaluate expanded program
	evaluated := evaluator.Eval(expanded, env)
	finalResult := extractFinalEvaluation(evaluated)

	fmt.Println(finalResult.Inspect())

	return nil
}

func parseCmdArgs() (string, error) {
	// check if the user has provided a file to run
	cmdArgs := os.Args
	if len(cmdArgs) != 2 {
		return "", errors.New("please specify a file to run. Usage: ./jsop <filename>")
	}

	// check if the file extension is valid
	filePath := cmdArgs[1]
	fileName, err := filepath.Abs(filePath)
	if err != nil {
		return "", fmt.Errorf("fail to get absolute path of file: %s", err)
	}
	if !isValidFileExtension(fileName) {
		return "", errors.New("invalid file extension. Please use .jsop or .jsop.json files")
	}

	return filePath, nil
}

func isValidFileExtension(fileName string) bool {
	lowerCaseFileName := strings.ToLower(fileName)

	if filepath.Ext(lowerCaseFileName) == ".jsop" {
		return true
	}

	fileLen := len(lowerCaseFileName)
	if fileLen > 10 && lowerCaseFileName[fileLen-10:] == ".jsop.json" {
		return true
	}

	return true
}

func extractFinalEvaluation(obj object.Object) object.Object {
	if array, ok := obj.(*object.Array); ok {
		return array.Elements[len(array.Elements)-1]
	}

	return obj
}
