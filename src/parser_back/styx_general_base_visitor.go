// Code generated from styx_general.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // styx_general

import "github.com/antlr/antlr4/runtime/Go/antlr"

type Basestyx_generalVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *Basestyx_generalVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basestyx_generalVisitor) VisitMultiplicativeExpr(ctx *MultiplicativeExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basestyx_generalVisitor) VisitAdditiveExpr(ctx *AdditiveExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basestyx_generalVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basestyx_generalVisitor) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basestyx_generalVisitor) VisitExpressionsBlock(ctx *ExpressionsBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basestyx_generalVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basestyx_generalVisitor) VisitProcedure(ctx *ProcedureContext) interface{} {
	return v.VisitChildren(ctx)
}
