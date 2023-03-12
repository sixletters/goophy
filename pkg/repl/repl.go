package repl

import (
	"bufio"
	lexer "cs4215/goophy/pkg/lexer"
	token "cs4215/goophy/pkg/token"
	"fmt"
	"io"
)

// "bufio"
//    "fmt"
//    "io"
//    "cs"

const prompt = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.NewLexer(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
