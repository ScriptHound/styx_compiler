package main

import (
	"fmt"

	parser "MTUCI_studying_practice/src/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	is := antlr.NewInputStream("1+5-4*7")

	lexer := parser.Newstyx_generalLexer(is)

	for {
		token := lexer.NextToken()
		if token.GetTokenType() == antlr.TokenEOF {
			break
		}

		fmt.Printf("%s (%q)\n",
			lexer.SymbolicNames[token.GetTokenType()], token.GetText())
	}
}
