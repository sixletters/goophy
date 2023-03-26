package main

import (
	"cs4215/goophy/pkg/compiler"
	"cs4215/goophy/pkg/lexer"
	"cs4215/goophy/pkg/machine"
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
<<<<<<< Updated upstream
	fmt.Println(compiler.Compile(*program))
	instrs := compiler.Compile(*program)
	val := machine.Run(instrs)
	fmt.Println(val)
	//thread pool / when u need a concurrent thread create one
	//simple concurrency control by next tues, looking forward then we see how much of go we cna implement
	// val := run(instrs)
	// fmt.Println(compiler.Compile(*program))
	// fmt.Println(program.String())
	// fmt.Printf("Feel free to type in commands\n")
	// repl.Start(os.Stdin, os.Stdout)
=======
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
>>>>>>> Stashed changes
}
