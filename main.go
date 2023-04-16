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
	"strconv"
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
	numCores := int64(4)
	if len(os.Args) == 3 {
		numCores, err = strconv.ParseInt(os.Args[2], 10, 64)
		if err != nil {
			panic("Invalid number of cores given")
		}
	}
	l := lexer.NewLexer(string(data))
	p := parser.New(l)
	program := p.ParseProgram()
	instrs := compiler.NewCompiler().Compile(*program)
	for i, ints := range instrs {
		fmt.Printf("%d ", i)
		fmt.Println(ints)
	}
	machine := machine.NewMachine().Init().WithCores(int(numCores))
	machine.Run(instrs)
	//repl.Start(os.Stdin, os.Stdout)

}
