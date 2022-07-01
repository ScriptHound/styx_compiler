// Code generated from styx.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // styx

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasestyxListener is a complete listener for a parse tree produced by styxParser.
type BasestyxListener struct{}

var _ styxListener = &BasestyxListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasestyxListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasestyxListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasestyxListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasestyxListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterR is called when production r is entered.
func (s *BasestyxListener) EnterR(ctx *RContext) {}

// ExitR is called when production r is exited.
func (s *BasestyxListener) ExitR(ctx *RContext) {}
