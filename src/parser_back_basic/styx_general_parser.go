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
		"", "'='", "'*'", "'/'", "'+'", "'-'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "INT", "ID",
	}
	staticData.ruleNames = []string{
		"assignment", "multiplicativeExpr", "additiveExpr",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 7, 60, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 16, 8, 0, 1, 1, 1, 1, 1, 1, 5, 1, 21,
		8, 1, 10, 1, 12, 1, 24, 9, 1, 1, 1, 1, 1, 1, 1, 5, 1, 29, 8, 1, 10, 1,
		12, 1, 32, 9, 1, 1, 1, 1, 1, 1, 1, 5, 1, 37, 8, 1, 10, 1, 12, 1, 40, 9,
		1, 1, 1, 1, 1, 1, 1, 5, 1, 45, 8, 1, 10, 1, 12, 1, 48, 9, 1, 3, 1, 50,
		8, 1, 1, 2, 1, 2, 1, 2, 5, 2, 55, 8, 2, 10, 2, 12, 2, 58, 9, 2, 1, 2, 0,
		0, 3, 0, 2, 4, 0, 2, 1, 0, 2, 3, 1, 0, 4, 5, 66, 0, 15, 1, 0, 0, 0, 2,
		49, 1, 0, 0, 0, 4, 51, 1, 0, 0, 0, 6, 7, 5, 7, 0, 0, 7, 8, 5, 1, 0, 0,
		8, 16, 5, 6, 0, 0, 9, 10, 5, 7, 0, 0, 10, 11, 5, 1, 0, 0, 11, 16, 5, 7,
		0, 0, 12, 13, 5, 7, 0, 0, 13, 14, 5, 1, 0, 0, 14, 16, 3, 4, 2, 0, 15, 6,
		1, 0, 0, 0, 15, 9, 1, 0, 0, 0, 15, 12, 1, 0, 0, 0, 16, 1, 1, 0, 0, 0, 17,
		22, 5, 6, 0, 0, 18, 19, 7, 0, 0, 0, 19, 21, 5, 6, 0, 0, 20, 18, 1, 0, 0,
		0, 21, 24, 1, 0, 0, 0, 22, 20, 1, 0, 0, 0, 22, 23, 1, 0, 0, 0, 23, 50,
		1, 0, 0, 0, 24, 22, 1, 0, 0, 0, 25, 30, 5, 7, 0, 0, 26, 27, 7, 0, 0, 0,
		27, 29, 5, 7, 0, 0, 28, 26, 1, 0, 0, 0, 29, 32, 1, 0, 0, 0, 30, 28, 1,
		0, 0, 0, 30, 31, 1, 0, 0, 0, 31, 50, 1, 0, 0, 0, 32, 30, 1, 0, 0, 0, 33,
		38, 5, 6, 0, 0, 34, 35, 7, 0, 0, 0, 35, 37, 5, 7, 0, 0, 36, 34, 1, 0, 0,
		0, 37, 40, 1, 0, 0, 0, 38, 36, 1, 0, 0, 0, 38, 39, 1, 0, 0, 0, 39, 50,
		1, 0, 0, 0, 40, 38, 1, 0, 0, 0, 41, 46, 5, 7, 0, 0, 42, 43, 7, 0, 0, 0,
		43, 45, 5, 6, 0, 0, 44, 42, 1, 0, 0, 0, 45, 48, 1, 0, 0, 0, 46, 44, 1,
		0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 50, 1, 0, 0, 0, 48, 46, 1, 0, 0, 0, 49,
		17, 1, 0, 0, 0, 49, 25, 1, 0, 0, 0, 49, 33, 1, 0, 0, 0, 49, 41, 1, 0, 0,
		0, 50, 3, 1, 0, 0, 0, 51, 56, 3, 2, 1, 0, 52, 53, 7, 1, 0, 0, 53, 55, 3,
		2, 1, 0, 54, 52, 1, 0, 0, 0, 55, 58, 1, 0, 0, 0, 56, 54, 1, 0, 0, 0, 56,
		57, 1, 0, 0, 0, 57, 5, 1, 0, 0, 0, 58, 56, 1, 0, 0, 0, 7, 15, 22, 30, 38,
		46, 49, 56,
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
	styx_generalParserEOF  = antlr.TokenEOF
	styx_generalParserT__0 = 1
	styx_generalParserT__1 = 2
	styx_generalParserT__2 = 3
	styx_generalParserT__3 = 4
	styx_generalParserT__4 = 5
	styx_generalParserINT  = 6
	styx_generalParserID   = 7
)

// styx_generalParser rules.
const (
	styx_generalParserRULE_assignment         = 0
	styx_generalParserRULE_multiplicativeExpr = 1
	styx_generalParserRULE_additiveExpr       = 2
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

	p.SetState(15)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(6)
			p.Match(styx_generalParserID)
		}
		{
			p.SetState(7)
			p.Match(styx_generalParserT__0)
		}
		{
			p.SetState(8)
			p.Match(styx_generalParserINT)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(9)
			p.Match(styx_generalParserID)
		}
		{
			p.SetState(10)
			p.Match(styx_generalParserT__0)
		}
		{
			p.SetState(11)
			p.Match(styx_generalParserID)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(12)
			p.Match(styx_generalParserID)
		}
		{
			p.SetState(13)
			p.Match(styx_generalParserT__0)
		}
		{
			p.SetState(14)
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

	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(17)
			p.Match(styx_generalParserINT)
		}
		p.SetState(22)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == styx_generalParserT__1 || _la == styx_generalParserT__2 {
			{
				p.SetState(18)
				_la = p.GetTokenStream().LA(1)

				if !(_la == styx_generalParserT__1 || _la == styx_generalParserT__2) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(19)
				p.Match(styx_generalParserINT)
			}

			p.SetState(24)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(25)
			p.Match(styx_generalParserID)
		}
		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == styx_generalParserT__1 || _la == styx_generalParserT__2 {
			{
				p.SetState(26)
				_la = p.GetTokenStream().LA(1)

				if !(_la == styx_generalParserT__1 || _la == styx_generalParserT__2) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(27)
				p.Match(styx_generalParserID)
			}

			p.SetState(32)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(33)
			p.Match(styx_generalParserINT)
		}
		p.SetState(38)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == styx_generalParserT__1 || _la == styx_generalParserT__2 {
			{
				p.SetState(34)
				_la = p.GetTokenStream().LA(1)

				if !(_la == styx_generalParserT__1 || _la == styx_generalParserT__2) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(35)
				p.Match(styx_generalParserID)
			}

			p.SetState(40)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(41)
			p.Match(styx_generalParserID)
		}
		p.SetState(46)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == styx_generalParserT__1 || _la == styx_generalParserT__2 {
			{
				p.SetState(42)
				_la = p.GetTokenStream().LA(1)

				if !(_la == styx_generalParserT__1 || _la == styx_generalParserT__2) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(43)
				p.Match(styx_generalParserINT)
			}

			p.SetState(48)
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
		p.SetState(51)
		p.MultiplicativeExpr()
	}
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == styx_generalParserT__3 || _la == styx_generalParserT__4 {
		{
			p.SetState(52)
			_la = p.GetTokenStream().LA(1)

			if !(_la == styx_generalParserT__3 || _la == styx_generalParserT__4) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(53)
			p.MultiplicativeExpr()
		}

		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}
