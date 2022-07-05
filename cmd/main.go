package main

import (
	parser "MTUCI_studying_practice/src/parser"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	content, err := ioutil.ReadFile("example.styx")
	string_content := string(content)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string_content)

	is := antlr.NewInputStream(string_content)

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
