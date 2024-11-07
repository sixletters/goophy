package main

import (
	"cs4215/goophy/pkg/compiler"
	"cs4215/goophy/pkg/lexer"
	"cs4215/goophy/pkg/machine"
	"cs4215/goophy/pkg/parser"
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("starting goophy")
	js.Global().Set("runGoophy", runGoophyFunc())
	<-make(chan struct{})
}

func runGoophyFunc() js.Func {
	goophyRunFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "Invalid no of aruguments passed"
		}
		inputCode := args[0].String()
		l := lexer.NewLexer(inputCode)
		p := parser.New(l)
		program := p.ParseProgram()
		instrs := compiler.NewCompiler().Compile(*program)
		fmt.Println(instrs)
		mach := machine.NewMachine().Init()
		res := mach.Run(instrs)
		fmt.Println(res)
		// Convert the result to a string and return a pointer to it
		return fmt.Sprintf("%v", res)
	})
	return goophyRunFunc
}
