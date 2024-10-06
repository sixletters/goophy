package main

import (
	"cs4215/goophy/pkg/compiler"
	"cs4215/goophy/pkg/lexer"
	"cs4215/goophy/pkg/machine"
	"cs4215/goophy/pkg/parser"
	"fmt"
	"unsafe"
)

// Exported function to be called from JavaScript
//
//export WasmRunGo
func WasmRunGo(inputPtr *byte, length uint32) *byte {
	// Convert the memory pointer and length into a Go string
	inputBytes := unsafe.Slice(inputPtr, length)
	input := string(inputBytes)

	// Run the compiler process
	l := lexer.NewLexer(input)
	p := parser.New(l)
	program := p.ParseProgram()
	instrs := compiler.NewCompiler().Compile(*program)
	mach := machine.NewMachine().Init()
	res := mach.Run(instrs)

	// Convert the result to a string and return a pointer to it
	result := []byte(fmt.Sprintf("%v", res))
	return &result[0]
}

func main() {
	// TinyGo requires a main function, but it's not used when compiling to WASM
}
