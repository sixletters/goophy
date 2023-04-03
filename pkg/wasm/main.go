package main

import (
	"cs4215/goophy/pkg/compiler"
	"cs4215/goophy/pkg/lexer"
	"cs4215/goophy/pkg/machine"
	"cs4215/goophy/pkg/parser"
	"syscall/js"
)

func main() {
	ch := make(chan struct{}, 0)
	js.Global().Set("WasmRunGo", js.FuncOf(runGo))
	<-ch
}

func runGo(this js.Value, args []js.Value) interface{} {
	input := args[0].String()
	l := lexer.NewLexer(input)
	p := parser.New(l)
	program := p.ParseProgram()
	instrs := compiler.Compile(*program)
	machine := machine.NewMachine().Init()
	res := machine.Run(instrs)
	return res
}
