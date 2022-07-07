// Code generated from Styx_refactoring.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Styx_refactoring

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type Styx_refactoringParser struct {
	*antlr.BaseParser
}

var styx_refactoringParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func styx_refactoringParserInit() {
	staticData := &styx_refactoringParserStaticData
	staticData.literalNames = []string{
		"", "'+'", "'-'", "'*'", "'/'", "'{'", "'}'", "'('", "')'", "'proc'",
		"'='", "", "", "';'", "','",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "INT", "ID", "SEMICOLON",
		"COMMA", "WS",
	}
	staticData.ruleNames = []string{
		"prog", "statement", "expression", "multiplyExpr", "statementsBlock",
		"functionArgs", "functionDefinition", "functionCall", "assignment",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 15, 90, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 4, 0, 25, 8, 0, 11, 0, 12, 0, 26, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 1, 3, 1, 34, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		4, 1, 4, 1, 4, 3, 4, 47, 8, 4, 1, 4, 1, 4, 5, 4, 51, 8, 4, 10, 4, 12, 4,
		54, 9, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 64, 8,
		5, 10, 5, 12, 5, 67, 9, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3,
		8, 88, 8, 8, 1, 8, 0, 0, 9, 0, 2, 4, 6, 8, 10, 12, 14, 16, 0, 3, 1, 0,
		1, 2, 1, 0, 11, 12, 1, 0, 3, 4, 90, 0, 24, 1, 0, 0, 0, 2, 33, 1, 0, 0,
		0, 4, 35, 1, 0, 0, 0, 6, 39, 1, 0, 0, 0, 8, 43, 1, 0, 0, 0, 10, 57, 1,
		0, 0, 0, 12, 70, 1, 0, 0, 0, 14, 75, 1, 0, 0, 0, 16, 87, 1, 0, 0, 0, 18,
		19, 3, 2, 1, 0, 19, 20, 5, 13, 0, 0, 20, 25, 1, 0, 0, 0, 21, 22, 3, 4,
		2, 0, 22, 23, 5, 13, 0, 0, 23, 25, 1, 0, 0, 0, 24, 18, 1, 0, 0, 0, 24,
		21, 1, 0, 0, 0, 25, 26, 1, 0, 0, 0, 26, 24, 1, 0, 0, 0, 26, 27, 1, 0, 0,
		0, 27, 28, 1, 0, 0, 0, 28, 29, 5, 0, 0, 1, 29, 1, 1, 0, 0, 0, 30, 34, 3,
		12, 6, 0, 31, 34, 3, 14, 7, 0, 32, 34, 3, 16, 8, 0, 33, 30, 1, 0, 0, 0,
		33, 31, 1, 0, 0, 0, 33, 32, 1, 0, 0, 0, 34, 3, 1, 0, 0, 0, 35, 36, 3, 6,
		3, 0, 36, 37, 7, 0, 0, 0, 37, 38, 3, 6, 3, 0, 38, 5, 1, 0, 0, 0, 39, 40,
		7, 1, 0, 0, 40, 41, 7, 2, 0, 0, 41, 42, 7, 1, 0, 0, 42, 7, 1, 0, 0, 0,
		43, 52, 5, 5, 0, 0, 44, 47, 3, 2, 1, 0, 45, 47, 3, 4, 2, 0, 46, 44, 1,
		0, 0, 0, 46, 45, 1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 49, 5, 13, 0, 0, 49,
		51, 1, 0, 0, 0, 50, 46, 1, 0, 0, 0, 51, 54, 1, 0, 0, 0, 52, 50, 1, 0, 0,
		0, 52, 53, 1, 0, 0, 0, 53, 55, 1, 0, 0, 0, 54, 52, 1, 0, 0, 0, 55, 56,
		5, 6, 0, 0, 56, 9, 1, 0, 0, 0, 57, 65, 5, 7, 0, 0, 58, 59, 5, 12, 0, 0,
		59, 64, 5, 14, 0, 0, 60, 61, 3, 4, 2, 0, 61, 62, 5, 14, 0, 0, 62, 64, 1,
		0, 0, 0, 63, 58, 1, 0, 0, 0, 63, 60, 1, 0, 0, 0, 64, 67, 1, 0, 0, 0, 65,
		63, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 68, 1, 0, 0, 0, 67, 65, 1, 0, 0,
		0, 68, 69, 5, 8, 0, 0, 69, 11, 1, 0, 0, 0, 70, 71, 5, 9, 0, 0, 71, 72,
		5, 12, 0, 0, 72, 73, 3, 10, 5, 0, 73, 74, 3, 8, 4, 0, 74, 13, 1, 0, 0,
		0, 75, 76, 5, 12, 0, 0, 76, 77, 3, 10, 5, 0, 77, 15, 1, 0, 0, 0, 78, 79,
		5, 12, 0, 0, 79, 80, 5, 10, 0, 0, 80, 88, 3, 4, 2, 0, 81, 82, 5, 12, 0,
		0, 82, 83, 5, 10, 0, 0, 83, 88, 5, 12, 0, 0, 84, 85, 5, 12, 0, 0, 85, 86,
		5, 10, 0, 0, 86, 88, 5, 11, 0, 0, 87, 78, 1, 0, 0, 0, 87, 81, 1, 0, 0,
		0, 87, 84, 1, 0, 0, 0, 88, 17, 1, 0, 0, 0, 8, 24, 26, 33, 46, 52, 63, 65,
		87,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// Styx_refactoringParserInit initializes any static state used to implement Styx_refactoringParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewStyx_refactoringParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Styx_refactoringParserInit() {
	staticData := &styx_refactoringParserStaticData
	staticData.once.Do(styx_refactoringParserInit)
}

// NewStyx_refactoringParser produces a new parser instance for the optional input antlr.TokenStream.
func NewStyx_refactoringParser(input antlr.TokenStream) *Styx_refactoringParser {
	Styx_refactoringParserInit()
	this := new(Styx_refactoringParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &styx_refactoringParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Styx_refactoring.g4"

	return this
}

// Styx_refactoringParser tokens.
const (
	Styx_refactoringParserEOF       = antlr.TokenEOF
	Styx_refactoringParserT__0      = 1
	Styx_refactoringParserT__1      = 2
	Styx_refactoringParserT__2      = 3
	Styx_refactoringParserT__3      = 4
	Styx_refactoringParserT__4      = 5
	Styx_refactoringParserT__5      = 6
	Styx_refactoringParserT__6      = 7
	Styx_refactoringParserT__7      = 8
	Styx_refactoringParserT__8      = 9
	Styx_refactoringParserT__9      = 10
	Styx_refactoringParserINT       = 11
	Styx_refactoringParserID        = 12
	Styx_refactoringParserSEMICOLON = 13
	Styx_refactoringParserCOMMA     = 14
	Styx_refactoringParserWS        = 15
)

// Styx_refactoringParser rules.
const (
	Styx_refactoringParserRULE_prog               = 0
	Styx_refactoringParserRULE_statement          = 1
	Styx_refactoringParserRULE_expression         = 2
	Styx_refactoringParserRULE_multiplyExpr       = 3
	Styx_refactoringParserRULE_statementsBlock    = 4
	Styx_refactoringParserRULE_functionArgs       = 5
	Styx_refactoringParserRULE_functionDefinition = 6
	Styx_refactoringParserRULE_functionCall       = 7
	Styx_refactoringParserRULE_assignment         = 8
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Styx_refactoringParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Styx_refactoringParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) EOF() antlr.TerminalNode {
	return s.GetToken(Styx_refactoringParserEOF, 0)
}

func (s *ProgContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(Styx_refactoringParserSEMICOLON)
}

func (s *ProgContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(Styx_refactoringParserSEMICOLON, i)
}

func (s *ProgContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Styx_refactoringListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Styx_refactoringListener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *Styx_refactoringParser) Prog() (localctx IProgContext) {
	this := p
	_ = this

	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Styx_refactoringParserRULE_prog)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(24)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<Styx_refactoringParserT__8)|(1<<Styx_refactoringParserINT)|(1<<Styx_refactoringParserID))) != 0) {
		p.SetState(24)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(18)
				p.Statement()
			}
			{
				p.SetState(19)
				p.Match(Styx_refactoringParserSEMICOLON)
			}

		case 2:
			{
				p.SetState(21)
				p.Expression()
			}
			{
				p.SetState(22)
				p.Match(Styx_refactoringParserSEMICOLON)
			}

		}

		p.SetState(26)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(28)
		p.Match(Styx_refactoringParserEOF)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Styx_refactoringParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Styx_refactoringParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) FunctionDefinition() IFunctionDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionDefinitionContext)
}

func (s *StatementContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *StatementContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Styx_refactoringListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Styx_refactoringListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *Styx_refactoringParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Styx_refactoringParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(30)
			p.FunctionDefinition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(31)
			p.FunctionCall()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(32)
			p.Assignment()
		}

	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Styx_refactoringParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Styx_refactoringParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) AllMultiplyExpr() []IMultiplyExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMultiplyExprContext); ok {
			len++
		}
	}

	tst := make([]IMultiplyExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMultiplyExprContext); ok {
			tst[i] = t.(IMultiplyExprContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) MultiplyExpr(i int) IMultiplyExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiplyExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiplyExprContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Styx_refactoringListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Styx_refactoringListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *Styx_refactoringParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, Styx_refactoringParserRULE_expression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(35)
		p.MultiplyExpr()
	}
	{
		p.SetState(36)
		_la = p.GetTokenStream().LA(1)

		if !(_la == Styx_refactoringParserT__0 || _la == Styx_refactoringParserT__1) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(37)
		p.MultiplyExpr()
	}

	return localctx
}

// IMultiplyExprContext is an interface to support dynamic dispatch.
type IMultiplyExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiplyExprContext differentiates from other interfaces.
	IsMultiplyExprContext()
}

type MultiplyExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplyExprContext() *MultiplyExprContext {
	var p = new(MultiplyExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Styx_refactoringParserRULE_multiplyExpr
	return p
}

func (*MultiplyExprContext) IsMultiplyExprContext() {}

func NewMultiplyExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplyExprContext {
	var p = new(MultiplyExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Styx_refactoringParserRULE_multiplyExpr

	return p
}

func (s *MultiplyExprContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplyExprContext) AllINT() []antlr.TerminalNode {
	return s.GetTokens(Styx_refactoringParserINT)
}

func (s *MultiplyExprContext) INT(i int) antlr.TerminalNode {
	return s.GetToken(Styx_refactoringParserINT, i)
}

func (s *MultiplyExprContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(Styx_refactoringParserID)
}

func (s *MultiplyExprContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(Styx_refactoringParserID, i)
}

func (s *MultiplyExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplyExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplyExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Styx_refactoringListener); ok {
		listenerT.EnterMultiplyExpr(s)
	}
}

func (s *MultiplyExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Styx_refactoringListener); ok {
		listenerT.ExitMultiplyExpr(s)
	}
}

func (p *Styx_refactoringParser) MultiplyExpr() (localctx IMultiplyExprContext) {
	this := p
	_ = this

	localctx = NewMultiplyExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, Styx_refactoringParserRULE_multiplyExpr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(39)
		_la = p.GetTokenStream().LA(1)

		if !(_la == Styx_refactoringParserINT || _la == Styx_refactoringParserID) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(40)
		_la = p.GetTokenStream().LA(1)

		if !(_la == Styx_refactoringParserT__2 || _la == Styx_refactoringParserT__3) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(41)
		_la = p.GetTokenStream().LA(1)

		if !(_la == Styx_refactoringParserINT || _la == Styx_refactoringParserID) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IStatementsBlockContext is an interface to support dynamic dispatch.
type IStatementsBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementsBlockContext differentiates from other interfaces.
	IsStatementsBlockContext()
}

type StatementsBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementsBlockContext() *StatementsBlockContext {
	var p = new(StatementsBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Styx_refactoringParserRULE_statementsBlock
	return p
}

func (*StatementsBlockContext) IsStatementsBlockContext() {}

func NewStatementsBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementsBlockContext {
	var p = new(StatementsBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Styx_refactoringParserRULE_statementsBlock

	return p
}

func (s *StatementsBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementsBlockContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(Styx_refactoringParserSEMICOLON)
}

func (s *StatementsBlockContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(Styx_refactoringParserSEMICOLON, i)
}

func (s *StatementsBlockContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *StatementsBlockContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementsBlockContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *StatementsBlockContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StatementsBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementsBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementsBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Styx_refactoringListener); ok {
		listenerT.EnterStatementsBlock(s)
	}
}

func (s *StatementsBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Styx_refactoringListener); ok {
		listenerT.ExitStatementsBlock(s)
	}
}

func (p *Styx_refactoringParser) StatementsBlock() (localctx IStatementsBlockContext) {
	this := p
	_ = this

	localctx = NewStatementsBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, Styx_refactoringParserRULE_statementsBlock)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(43)
		p.Match(Styx_refactoringParserT__4)
	}
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<Styx_refactoringParserT__8)|(1<<Styx_refactoringParserINT)|(1<<Styx_refactoringParserID))) != 0 {
		p.SetState(46)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(44)
				p.Statement()
			}

		case 2:
			{
				p.SetState(45)
				p.Expression()
			}

		}
		{
			p.SetState(48)
			p.Match(Styx_refactoringParserSEMICOLON)
		}

		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(55)
		p.Match(Styx_refactoringParserT__5)
	}

	return localctx
}

// IFunctionArgsContext is an interface to support dynamic dispatch.
type IFunctionArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionArgsContext differentiates from other interfaces.
	IsFunctionArgsContext()
}

type FunctionArgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionArgsContext() *FunctionArgsContext {
	var p = new(FunctionArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Styx_refactoringParserRULE_functionArgs
	return p
}

func (*FunctionArgsContext) IsFunctionArgsContext() {}

func NewFunctionArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionArgsContext {
	var p = new(FunctionArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Styx_refactoringParserRULE_functionArgs

	return p
}

func (s *FunctionArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionArgsContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(Styx_refactoringParserID)
}

func (s *FunctionArgsContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(Styx_refactoringParserID, i)
}

func (s *FunctionArgsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(Styx_refactoringParserCOMMA)
}

func (s *FunctionArgsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(Styx_refactoringParserCOMMA, i)
}

func (s *FunctionArgsContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *FunctionArgsContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FunctionArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Styx_refactoringListener); ok {
		listenerT.EnterFunctionArgs(s)
	}
}

func (s *FunctionArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Styx_refactoringListener); ok {
		listenerT.ExitFunctionArgs(s)
	}
}

func (p *Styx_refactoringParser) FunctionArgs() (localctx IFunctionArgsContext) {
	this := p
	_ = this

	localctx = NewFunctionArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, Styx_refactoringParserRULE_functionArgs)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(57)
		p.Match(Styx_refactoringParserT__6)
	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == Styx_refactoringParserINT || _la == Styx_refactoringParserID {
		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(58)
				p.Match(Styx_refactoringParserID)
			}
			{
				p.SetState(59)
				p.Match(Styx_refactoringParserCOMMA)
			}

		case 2:
			{
				p.SetState(60)
				p.Expression()
			}
			{
				p.SetState(61)
				p.Match(Styx_refactoringParserCOMMA)
			}

		}

		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(68)
		p.Match(Styx_refactoringParserT__7)
	}

	return localctx
}

// IFunctionDefinitionContext is an interface to support dynamic dispatch.
type IFunctionDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionDefinitionContext differentiates from other interfaces.
	IsFunctionDefinitionContext()
}

type FunctionDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDefinitionContext() *FunctionDefinitionContext {
	var p = new(FunctionDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Styx_refactoringParserRULE_functionDefinition
	return p
}

func (*FunctionDefinitionContext) IsFunctionDefinitionContext() {}

func NewFunctionDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDefinitionContext {
	var p = new(FunctionDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Styx_refactoringParserRULE_functionDefinition

	return p
}

func (s *FunctionDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDefinitionContext) ID() antlr.TerminalNode {
	return s.GetToken(Styx_refactoringParserID, 0)
}

func (s *FunctionDefinitionContext) FunctionArgs() IFunctionArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionArgsContext)
}

func (s *FunctionDefinitionContext) StatementsBlock() IStatementsBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementsBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementsBlockContext)
}

func (s *FunctionDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Styx_refactoringListener); ok {
		listenerT.EnterFunctionDefinition(s)
	}
}

func (s *FunctionDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Styx_refactoringListener); ok {
		listenerT.ExitFunctionDefinition(s)
	}
}

func (p *Styx_refactoringParser) FunctionDefinition() (localctx IFunctionDefinitionContext) {
	this := p
	_ = this

	localctx = NewFunctionDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, Styx_refactoringParserRULE_functionDefinition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(70)
		p.Match(Styx_refactoringParserT__8)
	}
	{
		p.SetState(71)
		p.Match(Styx_refactoringParserID)
	}
	{
		p.SetState(72)
		p.FunctionArgs()
	}
	{
		p.SetState(73)
		p.StatementsBlock()
	}

	return localctx
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Styx_refactoringParserRULE_functionCall
	return p
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Styx_refactoringParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) ID() antlr.TerminalNode {
	return s.GetToken(Styx_refactoringParserID, 0)
}

func (s *FunctionCallContext) FunctionArgs() IFunctionArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionArgsContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Styx_refactoringListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Styx_refactoringListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (p *Styx_refactoringParser) FunctionCall() (localctx IFunctionCallContext) {
	this := p
	_ = this

	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, Styx_refactoringParserRULE_functionCall)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(75)
		p.Match(Styx_refactoringParserID)
	}
	{
		p.SetState(76)
		p.FunctionArgs()
	}

	return localctx
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Styx_refactoringParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Styx_refactoringParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(Styx_refactoringParserID)
}

func (s *AssignmentContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(Styx_refactoringParserID, i)
}

func (s *AssignmentContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentContext) INT() antlr.TerminalNode {
	return s.GetToken(Styx_refactoringParserINT, 0)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Styx_refactoringListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Styx_refactoringListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (p *Styx_refactoringParser) Assignment() (localctx IAssignmentContext) {
	this := p
	_ = this

	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, Styx_refactoringParserRULE_assignment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(78)
			p.Match(Styx_refactoringParserID)
		}
		{
			p.SetState(79)
			p.Match(Styx_refactoringParserT__9)
		}
		{
			p.SetState(80)
			p.Expression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(81)
			p.Match(Styx_refactoringParserID)
		}
		{
			p.SetState(82)
			p.Match(Styx_refactoringParserT__9)
		}
		{
			p.SetState(83)
			p.Match(Styx_refactoringParserID)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(84)
			p.Match(Styx_refactoringParserID)
		}
		{
			p.SetState(85)
			p.Match(Styx_refactoringParserT__9)
		}
		{
			p.SetState(86)
			p.Match(Styx_refactoringParserINT)
		}

	}

	return localctx
}
