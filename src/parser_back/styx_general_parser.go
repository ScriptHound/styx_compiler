// Code generated from styx_general.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // styx_general

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

type styx_generalParser struct {
	*antlr.BaseParser
}

var styx_generalParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func styx_generalParserInit() {
	staticData := &styx_generalParserStaticData
	staticData.literalNames = []string{
		"", "'='", "'*'", "'/'", "'+'", "'-'", "','", "'{'", "';'", "'}'", "'return'",
		"'proc'", "'('", "')'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "INT", "ID",
		"WS",
	}
	staticData.ruleNames = []string{
		"assignment", "multiplicativeExpr", "additiveExpr", "expression", "expressionList",
		"expressionsBlock", "returnStatement", "procedure",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 16, 106, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 3, 0, 26, 8, 0, 1, 1, 1, 1, 1, 1, 5, 1, 31, 8, 1,
		10, 1, 12, 1, 34, 9, 1, 1, 1, 1, 1, 1, 1, 5, 1, 39, 8, 1, 10, 1, 12, 1,
		42, 9, 1, 1, 1, 1, 1, 1, 1, 5, 1, 47, 8, 1, 10, 1, 12, 1, 50, 9, 1, 1,
		1, 1, 1, 1, 1, 5, 1, 55, 8, 1, 10, 1, 12, 1, 58, 9, 1, 3, 1, 60, 8, 1,
		1, 2, 1, 2, 1, 2, 5, 2, 65, 8, 2, 10, 2, 12, 2, 68, 9, 2, 1, 3, 1, 3, 3,
		3, 72, 8, 3, 1, 4, 1, 4, 1, 4, 5, 4, 77, 8, 4, 10, 4, 12, 4, 80, 9, 4,
		1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 86, 8, 5, 10, 5, 12, 5, 89, 9, 5, 1, 5, 1,
		5, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 97, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 0, 0, 8, 0, 2, 4, 6, 8, 10, 12, 14, 0, 2, 1, 0, 2, 3,
		1, 0, 4, 5, 111, 0, 25, 1, 0, 0, 0, 2, 59, 1, 0, 0, 0, 4, 61, 1, 0, 0,
		0, 6, 71, 1, 0, 0, 0, 8, 73, 1, 0, 0, 0, 10, 81, 1, 0, 0, 0, 12, 96, 1,
		0, 0, 0, 14, 98, 1, 0, 0, 0, 16, 17, 5, 15, 0, 0, 17, 18, 5, 1, 0, 0, 18,
		26, 5, 14, 0, 0, 19, 20, 5, 15, 0, 0, 20, 21, 5, 1, 0, 0, 21, 26, 5, 15,
		0, 0, 22, 23, 5, 15, 0, 0, 23, 24, 5, 1, 0, 0, 24, 26, 3, 4, 2, 0, 25,
		16, 1, 0, 0, 0, 25, 19, 1, 0, 0, 0, 25, 22, 1, 0, 0, 0, 26, 1, 1, 0, 0,
		0, 27, 32, 5, 14, 0, 0, 28, 29, 7, 0, 0, 0, 29, 31, 5, 14, 0, 0, 30, 28,
		1, 0, 0, 0, 31, 34, 1, 0, 0, 0, 32, 30, 1, 0, 0, 0, 32, 33, 1, 0, 0, 0,
		33, 60, 1, 0, 0, 0, 34, 32, 1, 0, 0, 0, 35, 40, 5, 15, 0, 0, 36, 37, 7,
		0, 0, 0, 37, 39, 5, 15, 0, 0, 38, 36, 1, 0, 0, 0, 39, 42, 1, 0, 0, 0, 40,
		38, 1, 0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 60, 1, 0, 0, 0, 42, 40, 1, 0, 0,
		0, 43, 48, 5, 14, 0, 0, 44, 45, 7, 0, 0, 0, 45, 47, 5, 15, 0, 0, 46, 44,
		1, 0, 0, 0, 47, 50, 1, 0, 0, 0, 48, 46, 1, 0, 0, 0, 48, 49, 1, 0, 0, 0,
		49, 60, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 51, 56, 5, 15, 0, 0, 52, 53, 7,
		0, 0, 0, 53, 55, 5, 14, 0, 0, 54, 52, 1, 0, 0, 0, 55, 58, 1, 0, 0, 0, 56,
		54, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 60, 1, 0, 0, 0, 58, 56, 1, 0, 0,
		0, 59, 27, 1, 0, 0, 0, 59, 35, 1, 0, 0, 0, 59, 43, 1, 0, 0, 0, 59, 51,
		1, 0, 0, 0, 60, 3, 1, 0, 0, 0, 61, 66, 3, 2, 1, 0, 62, 63, 7, 1, 0, 0,
		63, 65, 3, 2, 1, 0, 64, 62, 1, 0, 0, 0, 65, 68, 1, 0, 0, 0, 66, 64, 1,
		0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 5, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 69,
		72, 3, 0, 0, 0, 70, 72, 3, 4, 2, 0, 71, 69, 1, 0, 0, 0, 71, 70, 1, 0, 0,
		0, 72, 7, 1, 0, 0, 0, 73, 78, 3, 6, 3, 0, 74, 75, 5, 6, 0, 0, 75, 77, 3,
		6, 3, 0, 76, 74, 1, 0, 0, 0, 77, 80, 1, 0, 0, 0, 78, 76, 1, 0, 0, 0, 78,
		79, 1, 0, 0, 0, 79, 9, 1, 0, 0, 0, 80, 78, 1, 0, 0, 0, 81, 87, 5, 7, 0,
		0, 82, 83, 3, 6, 3, 0, 83, 84, 5, 8, 0, 0, 84, 86, 1, 0, 0, 0, 85, 82,
		1, 0, 0, 0, 86, 89, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0,
		88, 90, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 90, 91, 5, 9, 0, 0, 91, 11, 1,
		0, 0, 0, 92, 93, 5, 10, 0, 0, 93, 97, 5, 15, 0, 0, 94, 95, 5, 10, 0, 0,
		95, 97, 5, 14, 0, 0, 96, 92, 1, 0, 0, 0, 96, 94, 1, 0, 0, 0, 97, 13, 1,
		0, 0, 0, 98, 99, 5, 11, 0, 0, 99, 100, 5, 15, 0, 0, 100, 101, 5, 12, 0,
		0, 101, 102, 3, 8, 4, 0, 102, 103, 5, 13, 0, 0, 103, 104, 3, 10, 5, 0,
		104, 15, 1, 0, 0, 0, 11, 25, 32, 40, 48, 56, 59, 66, 71, 78, 87, 96,
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

// styx_generalParserInit initializes any static state used to implement styx_generalParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// Newstyx_generalParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Styx_generalParserInit() {
	staticData := &styx_generalParserStaticData
	staticData.once.Do(styx_generalParserInit)
}

// Newstyx_generalParser produces a new parser instance for the optional input antlr.TokenStream.
func Newstyx_generalParser(input antlr.TokenStream) *styx_generalParser {
	Styx_generalParserInit()
	this := new(styx_generalParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &styx_generalParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "styx_general.g4"

	return this
}

// styx_generalParser tokens.
const (
	styx_generalParserEOF   = antlr.TokenEOF
	styx_generalParserT__0  = 1
	styx_generalParserT__1  = 2
	styx_generalParserT__2  = 3
	styx_generalParserT__3  = 4
	styx_generalParserT__4  = 5
	styx_generalParserT__5  = 6
	styx_generalParserT__6  = 7
	styx_generalParserT__7  = 8
	styx_generalParserT__8  = 9
	styx_generalParserT__9  = 10
	styx_generalParserT__10 = 11
	styx_generalParserT__11 = 12
	styx_generalParserT__12 = 13
	styx_generalParserINT   = 14
	styx_generalParserID    = 15
	styx_generalParserWS    = 16
)

// styx_generalParser rules.
const (
	styx_generalParserRULE_assignment         = 0
	styx_generalParserRULE_multiplicativeExpr = 1
	styx_generalParserRULE_additiveExpr       = 2
	styx_generalParserRULE_expression         = 3
	styx_generalParserRULE_expressionList     = 4
	styx_generalParserRULE_expressionsBlock   = 5
	styx_generalParserRULE_returnStatement    = 6
	styx_generalParserRULE_procedure          = 7
)

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
	p.RuleIndex = styx_generalParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = styx_generalParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(styx_generalParserID)
}

func (s *AssignmentContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(styx_generalParserID, i)
}

func (s *AssignmentContext) INT() antlr.TerminalNode {
	return s.GetToken(styx_generalParserINT, 0)
}

func (s *AssignmentContext) AdditiveExpr() IAdditiveExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdditiveExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdditiveExprContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(styx_generalListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(styx_generalListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case styx_generalVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *styx_generalParser) Assignment() (localctx IAssignmentContext) {
	this := p
	_ = this

	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, styx_generalParserRULE_assignment)

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

	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(16)
			p.Match(styx_generalParserID)
		}
		{
			p.SetState(17)
			p.Match(styx_generalParserT__0)
		}
		{
			p.SetState(18)
			p.Match(styx_generalParserINT)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(19)
			p.Match(styx_generalParserID)
		}
		{
			p.SetState(20)
			p.Match(styx_generalParserT__0)
		}
		{
			p.SetState(21)
			p.Match(styx_generalParserID)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(22)
			p.Match(styx_generalParserID)
		}
		{
			p.SetState(23)
			p.Match(styx_generalParserT__0)
		}
		{
			p.SetState(24)
			p.AdditiveExpr()
		}

	}

	return localctx
}

// IMultiplicativeExprContext is an interface to support dynamic dispatch.
type IMultiplicativeExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiplicativeExprContext differentiates from other interfaces.
	IsMultiplicativeExprContext()
}

type MultiplicativeExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicativeExprContext() *MultiplicativeExprContext {
	var p = new(MultiplicativeExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = styx_generalParserRULE_multiplicativeExpr
	return p
}

func (*MultiplicativeExprContext) IsMultiplicativeExprContext() {}

func NewMultiplicativeExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicativeExprContext {
	var p = new(MultiplicativeExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = styx_generalParserRULE_multiplicativeExpr

	return p
}

func (s *MultiplicativeExprContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicativeExprContext) AllINT() []antlr.TerminalNode {
	return s.GetTokens(styx_generalParserINT)
}

func (s *MultiplicativeExprContext) INT(i int) antlr.TerminalNode {
	return s.GetToken(styx_generalParserINT, i)
}

func (s *MultiplicativeExprContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(styx_generalParserID)
}

func (s *MultiplicativeExprContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(styx_generalParserID, i)
}

func (s *MultiplicativeExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplicativeExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(styx_generalListener); ok {
		listenerT.EnterMultiplicativeExpr(s)
	}
}

func (s *MultiplicativeExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(styx_generalListener); ok {
		listenerT.ExitMultiplicativeExpr(s)
	}
}

func (s *MultiplicativeExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case styx_generalVisitor:
		return t.VisitMultiplicativeExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *styx_generalParser) MultiplicativeExpr() (localctx IMultiplicativeExprContext) {
	this := p
	_ = this

	localctx = NewMultiplicativeExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, styx_generalParserRULE_multiplicativeExpr)
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

	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(27)
			p.Match(styx_generalParserINT)
		}
		p.SetState(32)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == styx_generalParserT__1 || _la == styx_generalParserT__2 {
			{
				p.SetState(28)
				_la = p.GetTokenStream().LA(1)

				if !(_la == styx_generalParserT__1 || _la == styx_generalParserT__2) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(29)
				p.Match(styx_generalParserINT)
			}

			p.SetState(34)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(35)
			p.Match(styx_generalParserID)
		}
		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == styx_generalParserT__1 || _la == styx_generalParserT__2 {
			{
				p.SetState(36)
				_la = p.GetTokenStream().LA(1)

				if !(_la == styx_generalParserT__1 || _la == styx_generalParserT__2) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(37)
				p.Match(styx_generalParserID)
			}

			p.SetState(42)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(43)
			p.Match(styx_generalParserINT)
		}
		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == styx_generalParserT__1 || _la == styx_generalParserT__2 {
			{
				p.SetState(44)
				_la = p.GetTokenStream().LA(1)

				if !(_la == styx_generalParserT__1 || _la == styx_generalParserT__2) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(45)
				p.Match(styx_generalParserID)
			}

			p.SetState(50)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(51)
			p.Match(styx_generalParserID)
		}
		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == styx_generalParserT__1 || _la == styx_generalParserT__2 {
			{
				p.SetState(52)
				_la = p.GetTokenStream().LA(1)

				if !(_la == styx_generalParserT__1 || _la == styx_generalParserT__2) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(53)
				p.Match(styx_generalParserINT)
			}

			p.SetState(58)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IAdditiveExprContext is an interface to support dynamic dispatch.
type IAdditiveExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAdditiveExprContext differentiates from other interfaces.
	IsAdditiveExprContext()
}

type AdditiveExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditiveExprContext() *AdditiveExprContext {
	var p = new(AdditiveExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = styx_generalParserRULE_additiveExpr
	return p
}

func (*AdditiveExprContext) IsAdditiveExprContext() {}

func NewAdditiveExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditiveExprContext {
	var p = new(AdditiveExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = styx_generalParserRULE_additiveExpr

	return p
}

func (s *AdditiveExprContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditiveExprContext) AllMultiplicativeExpr() []IMultiplicativeExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMultiplicativeExprContext); ok {
			len++
		}
	}

	tst := make([]IMultiplicativeExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMultiplicativeExprContext); ok {
			tst[i] = t.(IMultiplicativeExprContext)
			i++
		}
	}

	return tst
}

func (s *AdditiveExprContext) MultiplicativeExpr(i int) IMultiplicativeExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiplicativeExprContext); ok {
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

	return t.(IMultiplicativeExprContext)
}

func (s *AdditiveExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdditiveExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(styx_generalListener); ok {
		listenerT.EnterAdditiveExpr(s)
	}
}

func (s *AdditiveExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(styx_generalListener); ok {
		listenerT.ExitAdditiveExpr(s)
	}
}

func (s *AdditiveExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case styx_generalVisitor:
		return t.VisitAdditiveExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *styx_generalParser) AdditiveExpr() (localctx IAdditiveExprContext) {
	this := p
	_ = this

	localctx = NewAdditiveExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, styx_generalParserRULE_additiveExpr)
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
		p.SetState(61)
		p.MultiplicativeExpr()
	}
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == styx_generalParserT__3 || _la == styx_generalParserT__4 {
		{
			p.SetState(62)
			_la = p.GetTokenStream().LA(1)

			if !(_la == styx_generalParserT__3 || _la == styx_generalParserT__4) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(63)
			p.MultiplicativeExpr()
		}

		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = styx_generalParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = styx_generalParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Assignment() IAssignmentContext {
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

func (s *ExpressionContext) AdditiveExpr() IAdditiveExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdditiveExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdditiveExprContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(styx_generalListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(styx_generalListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case styx_generalVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *styx_generalParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, styx_generalParserRULE_expression)

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

	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(69)
			p.Assignment()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(70)
			p.AdditiveExpr()
		}

	}

	return localctx
}

// IExpressionListContext is an interface to support dynamic dispatch.
type IExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionListContext differentiates from other interfaces.
	IsExpressionListContext()
}

type ExpressionListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionListContext() *ExpressionListContext {
	var p = new(ExpressionListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = styx_generalParserRULE_expressionList
	return p
}

func (*ExpressionListContext) IsExpressionListContext() {}

func NewExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListContext {
	var p = new(ExpressionListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = styx_generalParserRULE_expressionList

	return p
}

func (s *ExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionListContext) AllExpression() []IExpressionContext {
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

func (s *ExpressionListContext) Expression(i int) IExpressionContext {
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

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(styx_generalListener); ok {
		listenerT.EnterExpressionList(s)
	}
}

func (s *ExpressionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(styx_generalListener); ok {
		listenerT.ExitExpressionList(s)
	}
}

func (s *ExpressionListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case styx_generalVisitor:
		return t.VisitExpressionList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *styx_generalParser) ExpressionList() (localctx IExpressionListContext) {
	this := p
	_ = this

	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, styx_generalParserRULE_expressionList)
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
		p.SetState(73)
		p.Expression()
	}
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == styx_generalParserT__5 {
		{
			p.SetState(74)
			p.Match(styx_generalParserT__5)
		}
		{
			p.SetState(75)
			p.Expression()
		}

		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExpressionsBlockContext is an interface to support dynamic dispatch.
type IExpressionsBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionsBlockContext differentiates from other interfaces.
	IsExpressionsBlockContext()
}

type ExpressionsBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionsBlockContext() *ExpressionsBlockContext {
	var p = new(ExpressionsBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = styx_generalParserRULE_expressionsBlock
	return p
}

func (*ExpressionsBlockContext) IsExpressionsBlockContext() {}

func NewExpressionsBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionsBlockContext {
	var p = new(ExpressionsBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = styx_generalParserRULE_expressionsBlock

	return p
}

func (s *ExpressionsBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionsBlockContext) AllExpression() []IExpressionContext {
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

func (s *ExpressionsBlockContext) Expression(i int) IExpressionContext {
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

func (s *ExpressionsBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionsBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionsBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(styx_generalListener); ok {
		listenerT.EnterExpressionsBlock(s)
	}
}

func (s *ExpressionsBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(styx_generalListener); ok {
		listenerT.ExitExpressionsBlock(s)
	}
}

func (s *ExpressionsBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case styx_generalVisitor:
		return t.VisitExpressionsBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *styx_generalParser) ExpressionsBlock() (localctx IExpressionsBlockContext) {
	this := p
	_ = this

	localctx = NewExpressionsBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, styx_generalParserRULE_expressionsBlock)
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
		p.SetState(81)
		p.Match(styx_generalParserT__6)
	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == styx_generalParserINT || _la == styx_generalParserID {
		{
			p.SetState(82)
			p.Expression()
		}
		{
			p.SetState(83)
			p.Match(styx_generalParserT__7)
		}

		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(90)
		p.Match(styx_generalParserT__8)
	}

	return localctx
}

// IReturnStatementContext is an interface to support dynamic dispatch.
type IReturnStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnStatementContext differentiates from other interfaces.
	IsReturnStatementContext()
}

type ReturnStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStatementContext() *ReturnStatementContext {
	var p = new(ReturnStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = styx_generalParserRULE_returnStatement
	return p
}

func (*ReturnStatementContext) IsReturnStatementContext() {}

func NewReturnStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = styx_generalParserRULE_returnStatement

	return p
}

func (s *ReturnStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStatementContext) ID() antlr.TerminalNode {
	return s.GetToken(styx_generalParserID, 0)
}

func (s *ReturnStatementContext) INT() antlr.TerminalNode {
	return s.GetToken(styx_generalParserINT, 0)
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(styx_generalListener); ok {
		listenerT.EnterReturnStatement(s)
	}
}

func (s *ReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(styx_generalListener); ok {
		listenerT.ExitReturnStatement(s)
	}
}

func (s *ReturnStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case styx_generalVisitor:
		return t.VisitReturnStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *styx_generalParser) ReturnStatement() (localctx IReturnStatementContext) {
	this := p
	_ = this

	localctx = NewReturnStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, styx_generalParserRULE_returnStatement)

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

	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(92)
			p.Match(styx_generalParserT__9)
		}
		{
			p.SetState(93)
			p.Match(styx_generalParserID)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(94)
			p.Match(styx_generalParserT__9)
		}
		{
			p.SetState(95)
			p.Match(styx_generalParserINT)
		}

	}

	return localctx
}

// IProcedureContext is an interface to support dynamic dispatch.
type IProcedureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProcedureContext differentiates from other interfaces.
	IsProcedureContext()
}

type ProcedureContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProcedureContext() *ProcedureContext {
	var p = new(ProcedureContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = styx_generalParserRULE_procedure
	return p
}

func (*ProcedureContext) IsProcedureContext() {}

func NewProcedureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProcedureContext {
	var p = new(ProcedureContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = styx_generalParserRULE_procedure

	return p
}

func (s *ProcedureContext) GetParser() antlr.Parser { return s.parser }

func (s *ProcedureContext) ID() antlr.TerminalNode {
	return s.GetToken(styx_generalParserID, 0)
}

func (s *ProcedureContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ProcedureContext) ExpressionsBlock() IExpressionsBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionsBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionsBlockContext)
}

func (s *ProcedureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcedureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProcedureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(styx_generalListener); ok {
		listenerT.EnterProcedure(s)
	}
}

func (s *ProcedureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(styx_generalListener); ok {
		listenerT.ExitProcedure(s)
	}
}

func (s *ProcedureContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case styx_generalVisitor:
		return t.VisitProcedure(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *styx_generalParser) Procedure() (localctx IProcedureContext) {
	this := p
	_ = this

	localctx = NewProcedureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, styx_generalParserRULE_procedure)

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
		p.SetState(98)
		p.Match(styx_generalParserT__10)
	}
	{
		p.SetState(99)
		p.Match(styx_generalParserID)
	}
	{
		p.SetState(100)
		p.Match(styx_generalParserT__11)
	}
	{
		p.SetState(101)
		p.ExpressionList()
	}
	{
		p.SetState(102)
		p.Match(styx_generalParserT__12)
	}
	{
		p.SetState(103)
		p.ExpressionsBlock()
	}

	return localctx
}
