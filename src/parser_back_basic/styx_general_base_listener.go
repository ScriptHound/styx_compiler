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

// EnterAssignment is called when production assignment is entered.
func (s *Basestyx_generalListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *Basestyx_generalListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterMultiplicativeExpr is called when production multiplicativeExpr is entered.
func (s *Basestyx_generalListener) EnterMultiplicativeExpr(ctx *MultiplicativeExprContext) {}

// ExitMultiplicativeExpr is called when production multiplicativeExpr is exited.
func (s *Basestyx_generalListener) ExitMultiplicativeExpr(ctx *MultiplicativeExprContext) {}

// EnterAdditiveExpr is called when production additiveExpr is entered.
func (s *Basestyx_generalListener) EnterAdditiveExpr(ctx *AdditiveExprContext) {}

// ExitAdditiveExpr is called when production additiveExpr is exited.
func (s *Basestyx_generalListener) ExitAdditiveExpr(ctx *AdditiveExprContext) {}
