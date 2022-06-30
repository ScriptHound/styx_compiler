package main

import (
	"fmt"

	"parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	is := antlr.NewInputStream("Hello world")
	lexer := parser.NewCalcLexer(is)

	for {
		token := lexer.NextToken()
		if token.GetTokenType() == antlr.TokenEOF {
			break
		}

		fmt.Printf("%s (%q)\n",
			lexer.SymbolicNames[token.GetTokenType()], token.GetText())
	}
}
