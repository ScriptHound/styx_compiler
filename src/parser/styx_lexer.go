// Code generated from styx.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type styxLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var styxlexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func styxlexerLexerInit() {
	staticData := &styxlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'hello'",
	}
	staticData.symbolicNames = []string{
		"", "", "ID", "WS",
	}
	staticData.ruleNames = []string{
		"T__0", "ID", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 3, 25, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 1, 4, 1, 15, 8, 1, 11, 1, 12, 1, 16, 1, 2, 4, 2,
		20, 8, 2, 11, 2, 12, 2, 21, 1, 2, 1, 2, 0, 0, 3, 1, 1, 3, 2, 5, 3, 1, 0,
		2, 1, 0, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 26, 0, 1, 1, 0, 0, 0, 0,
		3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 1, 7, 1, 0, 0, 0, 3, 14, 1, 0, 0, 0, 5,
		19, 1, 0, 0, 0, 7, 8, 5, 104, 0, 0, 8, 9, 5, 101, 0, 0, 9, 10, 5, 108,
		0, 0, 10, 11, 5, 108, 0, 0, 11, 12, 5, 111, 0, 0, 12, 2, 1, 0, 0, 0, 13,
		15, 7, 0, 0, 0, 14, 13, 1, 0, 0, 0, 15, 16, 1, 0, 0, 0, 16, 14, 1, 0, 0,
		0, 16, 17, 1, 0, 0, 0, 17, 4, 1, 0, 0, 0, 18, 20, 7, 1, 0, 0, 19, 18, 1,
		0, 0, 0, 20, 21, 1, 0, 0, 0, 21, 19, 1, 0, 0, 0, 21, 22, 1, 0, 0, 0, 22,
		23, 1, 0, 0, 0, 23, 24, 6, 2, 0, 0, 24, 6, 1, 0, 0, 0, 3, 0, 16, 21, 1,
		6, 0, 0,
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

// styxLexerInit initializes any static state used to implement styxLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewstyxLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func StyxLexerInit() {
	staticData := &styxlexerLexerStaticData
	staticData.once.Do(styxlexerLexerInit)
}

// NewstyxLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewstyxLexer(input antlr.CharStream) *styxLexer {
	StyxLexerInit()
	l := new(styxLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &styxlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "styx.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// styxLexer tokens.
const (
	styxLexerT__0 = 1
	styxLexerID   = 2
	styxLexerWS   = 3
)
