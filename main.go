package main

import (
	"cs4215/goophy/pkg/compiler"
	"cs4215/goophy/pkg/lexer"
	"cs4215/goophy/pkg/parser"
	"fmt"
	"os/user"
)

func main() {

	_, err := user.Current()
	if err != nil {
		panic(err)
	}
	input := `
	let add = fn(x, y) {
		x + y;
	   };	   	
	let result = add(five, ten);
	`
	l := lexer.NewLexer(input)
	p := parser.New(l)
	program := p.ParseProgram()
	fmt.Println(program.String())
	fmt.Println(compiler.Compile(*program))
	fmt.Printf("Feel free to type in commands\n")
	//repl.Start(os.Stdin, os.Stdout)
}
