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

// EnterExpression is called when production expression is entered.
func (s *Basestyx_generalListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *Basestyx_generalListener) ExitExpression(ctx *ExpressionContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *Basestyx_generalListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *Basestyx_generalListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterExpressionsBlock is called when production expressionsBlock is entered.
func (s *Basestyx_generalListener) EnterExpressionsBlock(ctx *ExpressionsBlockContext) {}

// ExitExpressionsBlock is called when production expressionsBlock is exited.
func (s *Basestyx_generalListener) ExitExpressionsBlock(ctx *ExpressionsBlockContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *Basestyx_generalListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *Basestyx_generalListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterProcedure is called when production procedure is entered.
func (s *Basestyx_generalListener) EnterProcedure(ctx *ProcedureContext) {}

// ExitProcedure is called when production procedure is exited.
func (s *Basestyx_generalListener) ExitProcedure(ctx *ProcedureContext) {}
