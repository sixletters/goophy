package main

import (
	"cs4215/goophy/pkg/lexer"
	"cs4215/goophy/pkg/parser"
	"cs4215/goophy/pkg/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {

	_, err := user.Current()
	if err != nil {
		panic(err)
	}
	input := `
	1;
	let x = "12";
	`
	l := lexer.NewLexer(input)
	p := parser.New(l)
	program := p.ParseProgram()
	fmt.Println(program.String())
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
