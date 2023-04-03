package main

import (
	"cs4215/goophy/pkg/compiler"
	"cs4215/goophy/pkg/lexer"
	"cs4215/goophy/pkg/machine"
	"cs4215/goophy/pkg/parser"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
)

func main() {
	_, err := user.Current()
	if err != nil {
		panic(err)
	}
	if len(os.Args) < 2 {
		panic("There is no input file name given")
	}
	fileName := os.Args[1]
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	l := lexer.NewLexer(string(data))
	p := parser.New(l)
	program := p.ParseProgram()
	instrs := compiler.Compile(*program)
	machine := machine.NewMachine().Init()
	machine.Run(instrs)
	// fmt.Println(res)
	//repl.Start(os.Stdin, os.Stdout)

}
