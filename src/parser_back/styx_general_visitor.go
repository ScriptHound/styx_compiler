// Code generated from styx_general.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // styx_general

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by styx_generalParser.
type styx_generalVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by styx_generalParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by styx_generalParser#multiplicativeExpr.
	VisitMultiplicativeExpr(ctx *MultiplicativeExprContext) interface{}

	// Visit a parse tree produced by styx_generalParser#additiveExpr.
	VisitAdditiveExpr(ctx *AdditiveExprContext) interface{}

	// Visit a parse tree produced by styx_generalParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by styx_generalParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by styx_generalParser#expressionsBlock.
	VisitExpressionsBlock(ctx *ExpressionsBlockContext) interface{}

	// Visit a parse tree produced by styx_generalParser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by styx_generalParser#procedure.
	VisitProcedure(ctx *ProcedureContext) interface{}
}
