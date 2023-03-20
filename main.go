package main

import (
	"cs4215/goophy/pkg/compiler"
	"cs4215/goophy/pkg/lexer"
	"cs4215/goophy/pkg/parser"
	"fmt"
	"os/user"
)

// "cs4215/goophy/parser"

// "github.com/davecgh/go-spew/spew"

// func main() {
// 	filename := "" // A filename is optional
// 	src := `
// 	    // Sample xyzzy example
// 	    (function(){
// 	        if (3.14159 > 0) {
// 	            console.log("Hello, World.");
// 	            return;
// 	        }

// 	        var xyzzy = NaN;
// 	        console.log("Nothing happens.");
// 	        return xyzzy;
// 	    })();
// 	`
// 	// Parse some JavaScript, yielding a *ast.Program and/or an ErrorList
// 	// parser = JavaScriptParserBase
// 	node, _ := testparser.ParseFile(nil, "", `if (abc > 1) {}`, 0)
// 	node.
// 	// spew.Dump(node)
// 	fmt.Println(node.Comments)
// 	for _, f := range node.DeclarationList {
// 		fn, ok := f.(*ast.FunctionDeclaration)
// 		if !ok {
// 			fmt.Println("SKIPS")
// 			continue
// 		}
// 		fmt.Println(fn.Function.Name.Name)
// 	}
// 	// Parse("1 + 1;2+3")
// }

// type TreeShapeListener struct {
// 	*parser.BaseGrammarListener
// }

// func NewTreeShapeListener() *TreeShapeListener {
// 	return new(TreeShapeListener)
// }

// func (s *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
// 	fmt.Println(ctx.)
// 	ctx.
// 	// fmt.Println("CHILD COUNT IS HERE")
// 	// fmt.Println(ctx.GetChildCount())
// }

// func (s *TreeShapeListener) VisitTerminal(node antlr.TerminalNode) {
// 	fmt.Println(node.GetText())
// 	fmt.Println(node.GetPayload())
// 	fmt.Println(node.)
// 	// fmt.Println("CHILD COUNT IS HERE")
// 	// fmt.Println(ctx.GetChildCount())
// }

// func Parse(program string) {
// 	input, _ := antlr.NewFileStream(os.Args[1])
// 	lexer := parser.NewGrammarLexer(input)
// 	stream := antlr.NewCommonTokenStream(lexer, 0)
// 	p := parser.NewGrammarParser(stream)
// 	parser.
// 		p.BuildParseTrees = true
// 	// tree := p.
// 	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), p.Prog())
// }

func main() {
	_, err := user.Current()
	if err != nil {
		panic(err)
	}
	input := `
	let x = 10 - 5;
	1 + 1;
	`
	l := lexer.NewLexer(input)
	p := parser.New(l)
	program := p.ParseProgram()
	fmt.Println(compiler.Compile(*program))
	fmt.Println(program.String())
	// fmt.Printf("Feel free to type in commands\n")
	// repl.Start(os.Stdin, os.Stdout)
}
