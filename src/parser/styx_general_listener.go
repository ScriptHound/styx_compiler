// Code generated from styx_general.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // styx_general

import "github.com/antlr/antlr4/runtime/Go/antlr"

// styx_generalListener is a complete listener for a parse tree produced by styx_generalParser.
type styx_generalListener interface {
	antlr.ParseTreeListener

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)
}
