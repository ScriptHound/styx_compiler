package main

import (
	"fmt"

	parser "MTUCI_studying_practice/src/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	is := antlr.NewInputStream("Hello world")

	lexer := parser.NewstyxLexer(is)

	for {
		token := lexer.NextToken()
		if token.GetTokenType() == antlr.TokenEOF {
			break
		}

		fmt.Printf("%s (%q)\n",
			lexer.SymbolicNames[token.GetTokenType()], token.GetText())
	}
}
