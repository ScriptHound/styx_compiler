// Code generated from styx_general.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // styx_general

import "github.com/antlr/antlr4/runtime/Go/antlr"

// styx_generalListener is a complete listener for a parse tree produced by styx_generalParser.
type styx_generalListener interface {
	antlr.ParseTreeListener

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterMultiplicativeExpr is called when entering the multiplicativeExpr production.
	EnterMultiplicativeExpr(c *MultiplicativeExprContext)

	// EnterAdditiveExpr is called when entering the additiveExpr production.
	EnterAdditiveExpr(c *AdditiveExprContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterExpressionsBlock is called when entering the expressionsBlock production.
	EnterExpressionsBlock(c *ExpressionsBlockContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterProcedure is called when entering the procedure production.
	EnterProcedure(c *ProcedureContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitMultiplicativeExpr is called when exiting the multiplicativeExpr production.
	ExitMultiplicativeExpr(c *MultiplicativeExprContext)

	// ExitAdditiveExpr is called when exiting the additiveExpr production.
	ExitAdditiveExpr(c *AdditiveExprContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitExpressionsBlock is called when exiting the expressionsBlock production.
	ExitExpressionsBlock(c *ExpressionsBlockContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitProcedure is called when exiting the procedure production.
	ExitProcedure(c *ProcedureContext)
}
