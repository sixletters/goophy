package main

import (
	"cs4215/goophy/pkg/compiler"
	"cs4215/goophy/pkg/lexer"
	"cs4215/goophy/pkg/machine"
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
	let z = 5;
	let z = 10;
	let urmom = fn(x,y){
		x + y;
	}
	if (z == 5){
		let z = 10;
		if (z == 10){
			let z = 48;
		}
	}
	z;
	`
	// "+" | "-" | "*" | "/" | "<" | ">" | "==" | "!="
	// "!" | "-"
	// let add = fn(x, y) {
	// 	x < y;
	//    };
	// let five = 5;
	// let ten =10;
	// let result = add(five, ten);
	// result;
	l := lexer.NewLexer(input)
	p := parser.New(l)
	program := p.ParseProgram()
	instrs := compiler.Compile(*program)
	fmt.Println(compiler.Compile(*program))
	res := machine.Run(instrs)
	fmt.Println(res)
	// fmt.Println(instrs)
	// fmt.Println(program.String())
	// fmt.Println(compiler.Compile(*program))
	// fmt.Printf("Feel free to type in commands\n")
	//repl.Start(os.Stdin, os.Stdout)
}
