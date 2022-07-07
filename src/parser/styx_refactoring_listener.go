// Code generated from Styx_refactoring.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Styx_refactoring

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Styx_refactoringListener is a complete listener for a parse tree produced by Styx_refactoringParser.
type Styx_refactoringListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMultiplyExpr is called when entering the multiplyExpr production.
	EnterMultiplyExpr(c *MultiplyExprContext)

	// EnterStatementsBlock is called when entering the statementsBlock production.
	EnterStatementsBlock(c *StatementsBlockContext)

	// EnterFunctionArgs is called when entering the functionArgs production.
	EnterFunctionArgs(c *FunctionArgsContext)

	// EnterFunctionDefinition is called when entering the functionDefinition production.
	EnterFunctionDefinition(c *FunctionDefinitionContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMultiplyExpr is called when exiting the multiplyExpr production.
	ExitMultiplyExpr(c *MultiplyExprContext)

	// ExitStatementsBlock is called when exiting the statementsBlock production.
	ExitStatementsBlock(c *StatementsBlockContext)

	// ExitFunctionArgs is called when exiting the functionArgs production.
	ExitFunctionArgs(c *FunctionArgsContext)

	// ExitFunctionDefinition is called when exiting the functionDefinition production.
	ExitFunctionDefinition(c *FunctionDefinitionContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)
}
