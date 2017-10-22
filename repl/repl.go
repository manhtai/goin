package repl

import (
	"bufio"
	"fmt"
	"github.com/manhtai/goin/lexer"
	"github.com/manhtai/goin/token"
	"io"
)

// PROMPT is our prompt characters
const PROMPT = ">>> "

// Start read text from stdin and print token to stdout
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
