package main

import (
	// "github.com/davecgh/go-spew/spew"
	// "github.com/robertkrimen/otto/parser"
	// "cs4215/goophy/parser"
	"cs4215/goophy/parser"
	"fmt"
	"os"

	antlr "github.com/antlr/antlr4/runtime/Go/antlr/v4"
	// "github.com/davecgh/go-spew/spew"
)

func main() {
	// 	filename := "" // A filename is optional
	// 	src := `
	//     // Sample xyzzy example
	//     (function(){
	//         if (3.14159 > 0) {
	//             console.log("Hello, World.");
	//             return;
	//         }

	//         var xyzzy = NaN;
	//         console.log("Nothing happens.");
	//         return xyzzy;
	//     })();
	// `
	// Parse some JavaScript, yielding a *ast.Program and/or an ErrorList
	// parser = JavaScriptParserBase
	// program, _ := parser.ParseFile(nil, filename, src, 0)
	// spew.Dump(program)
	Parse("1 + 1;2+3")
}

type TreeShapeListener struct {
	*parser.BaseGrammarListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

// func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
// 	fmt.Println(ctx.GetText())
// 	// fmt.Println("CHILD COUNT IS HERE")
// 	// fmt.Println(ctx.GetChildCount())
// }

func (s *TreeShapeListener) VisitTerminal(node antlr.TerminalNode) {
	fmt.Println(node.GetText())
	// fmt.Println("CHILD COUNT IS HERE")
	// fmt.Println(ctx.GetChildCount())
}

func Parse(program string) {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewGrammarLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewGrammarParser(stream)
	p.BuildParseTrees = true
	// tree := p.
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), p.Prog())
}
