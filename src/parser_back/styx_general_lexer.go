// Code generated from styx_general.g4 by ANTLR 4.10.1. DO NOT EDIT.

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

type styx_generalLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var styx_generallexerLexerStaticData struct {
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

func styx_generallexerLexerInit() {
	staticData := &styx_generallexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'*'", "'+'", "'-'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "INT",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "INT",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 4, 17, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 0, 0, 4, 1, 1, 3, 2, 5, 3,
		7, 4, 1, 0, 1, 1, 0, 48, 57, 16, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0,
		5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 1, 9, 1, 0, 0, 0, 3, 11, 1, 0, 0, 0, 5,
		13, 1, 0, 0, 0, 7, 15, 1, 0, 0, 0, 9, 10, 5, 42, 0, 0, 10, 2, 1, 0, 0,
		0, 11, 12, 5, 43, 0, 0, 12, 4, 1, 0, 0, 0, 13, 14, 5, 45, 0, 0, 14, 6,
		1, 0, 0, 0, 15, 16, 7, 0, 0, 0, 16, 8, 1, 0, 0, 0, 1, 0, 0,
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

// styx_generalLexerInit initializes any static state used to implement styx_generalLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// Newstyx_generalLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func Styx_generalLexerInit() {
	staticData := &styx_generallexerLexerStaticData
	staticData.once.Do(styx_generallexerLexerInit)
}

// Newstyx_generalLexer produces a new lexer instance for the optional input antlr.CharStream.
func Newstyx_generalLexer(input antlr.CharStream) *styx_generalLexer {
	Styx_generalLexerInit()
	l := new(styx_generalLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &styx_generallexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "styx_general.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// styx_generalLexer tokens.
const (
	styx_generalLexerT__0 = 1
	styx_generalLexerT__1 = 2
	styx_generalLexerT__2 = 3
	styx_generalLexerINT  = 4
)
