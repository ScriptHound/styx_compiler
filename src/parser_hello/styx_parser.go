// Code generated from styx.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // styx

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

type styxParser struct {
	*antlr.BaseParser
}

var styxParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func styxParserInit() {
	staticData := &styxParserStaticData
	staticData.literalNames = []string{
		"", "'hello'",
	}
	staticData.symbolicNames = []string{
		"", "", "ID", "WS",
	}
	staticData.ruleNames = []string{
		"r",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 3, 6, 2, 0, 7, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 4, 0,
		2, 1, 0, 0, 0, 2, 3, 5, 1, 0, 0, 3, 4, 5, 2, 0, 0, 4, 1, 1, 0, 0, 0, 0,
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

// styxParserInit initializes any static state used to implement styxParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewstyxParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func StyxParserInit() {
	staticData := &styxParserStaticData
	staticData.once.Do(styxParserInit)
}

// NewstyxParser produces a new parser instance for the optional input antlr.TokenStream.
func NewstyxParser(input antlr.TokenStream) *styxParser {
	StyxParserInit()
	this := new(styxParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &styxParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "styx.g4"

	return this
}

// styxParser tokens.
const (
	styxParserEOF  = antlr.TokenEOF
	styxParserT__0 = 1
	styxParserID   = 2
	styxParserWS   = 3
)

// styxParserRULE_r is the styxParser rule.
const styxParserRULE_r = 0

// IRContext is an interface to support dynamic dispatch.
type IRContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRContext differentiates from other interfaces.
	IsRContext()
}

type RContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRContext() *RContext {
	var p = new(RContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = styxParserRULE_r
	return p
}

func (*RContext) IsRContext() {}

func NewRContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RContext {
	var p = new(RContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = styxParserRULE_r

	return p
}

func (s *RContext) GetParser() antlr.Parser { return s.parser }

func (s *RContext) ID() antlr.TerminalNode {
	return s.GetToken(styxParserID, 0)
}

func (s *RContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(styxListener); ok {
		listenerT.EnterR(s)
	}
}

func (s *RContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(styxListener); ok {
		listenerT.ExitR(s)
	}
}

func (p *styxParser) R() (localctx IRContext) {
	this := p
	_ = this

	localctx = NewRContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, styxParserRULE_r)

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
		p.SetState(2)
		p.Match(styxParserT__0)
	}
	{
		p.SetState(3)
		p.Match(styxParserID)
	}

	return localctx
}
