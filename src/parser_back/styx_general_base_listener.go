// Code generated from styx_general.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // styx_general

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Basestyx_generalListener is a complete listener for a parse tree produced by styx_generalParser.
type Basestyx_generalListener struct{}

var _ styx_generalListener = &Basestyx_generalListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *Basestyx_generalListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *Basestyx_generalListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *Basestyx_generalListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *Basestyx_generalListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpr is called when production expr is entered.
func (s *Basestyx_generalListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *Basestyx_generalListener) ExitExpr(ctx *ExprContext) {}
