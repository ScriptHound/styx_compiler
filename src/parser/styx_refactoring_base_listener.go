// Code generated from Styx_refactoring.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Styx_refactoring

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseStyx_refactoringListener is a complete listener for a parse tree produced by Styx_refactoringParser.
type BaseStyx_refactoringListener struct{}

var _ Styx_refactoringListener = &BaseStyx_refactoringListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseStyx_refactoringListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseStyx_refactoringListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseStyx_refactoringListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseStyx_refactoringListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseStyx_refactoringListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseStyx_refactoringListener) ExitProg(ctx *ProgContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseStyx_refactoringListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseStyx_refactoringListener) ExitStatement(ctx *StatementContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseStyx_refactoringListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseStyx_refactoringListener) ExitExpression(ctx *ExpressionContext) {}

// EnterMultiplyExpr is called when production multiplyExpr is entered.
func (s *BaseStyx_refactoringListener) EnterMultiplyExpr(ctx *MultiplyExprContext) {}

// ExitMultiplyExpr is called when production multiplyExpr is exited.
func (s *BaseStyx_refactoringListener) ExitMultiplyExpr(ctx *MultiplyExprContext) {}

// EnterStatementsBlock is called when production statementsBlock is entered.
func (s *BaseStyx_refactoringListener) EnterStatementsBlock(ctx *StatementsBlockContext) {}

// ExitStatementsBlock is called when production statementsBlock is exited.
func (s *BaseStyx_refactoringListener) ExitStatementsBlock(ctx *StatementsBlockContext) {}

// EnterFunctionArgs is called when production functionArgs is entered.
func (s *BaseStyx_refactoringListener) EnterFunctionArgs(ctx *FunctionArgsContext) {}

// ExitFunctionArgs is called when production functionArgs is exited.
func (s *BaseStyx_refactoringListener) ExitFunctionArgs(ctx *FunctionArgsContext) {}

// EnterFunctionDefinition is called when production functionDefinition is entered.
func (s *BaseStyx_refactoringListener) EnterFunctionDefinition(ctx *FunctionDefinitionContext) {}

// ExitFunctionDefinition is called when production functionDefinition is exited.
func (s *BaseStyx_refactoringListener) ExitFunctionDefinition(ctx *FunctionDefinitionContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseStyx_refactoringListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseStyx_refactoringListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseStyx_refactoringListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseStyx_refactoringListener) ExitAssignment(ctx *AssignmentContext) {}
