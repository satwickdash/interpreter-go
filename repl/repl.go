package repl

import (
	"bufio"
	"fmt"
	"interpreter-go/lexer"
	"interpreter-go/token"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {

	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scannedLine := scanner.Scan()
		if !scannedLine {
			return
		}

		line := scanner.Text()
		lex := lexer.New(line)
		for tok := lex.NextToken(); tok.Type != token.EOF; tok = lex.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
