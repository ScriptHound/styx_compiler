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
		"", "'='", "'*'", "'/'", "'+'", "'-'", "','", "'{'", "';'", "'}'", "'return'",
		"'proc'", "'('", "')'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "INT", "ID",
		"WS",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "INT", "ID", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 16, 84, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1,
		12, 1, 13, 4, 13, 69, 8, 13, 11, 13, 12, 13, 70, 1, 14, 4, 14, 74, 8, 14,
		11, 14, 12, 14, 75, 1, 15, 4, 15, 79, 8, 15, 11, 15, 12, 15, 80, 1, 15,
		1, 15, 0, 0, 16, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17,
		9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 1, 0, 3, 1,
		0, 48, 57, 4, 0, 45, 45, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 13, 13,
		32, 32, 86, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7,
		1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0,
		15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0,
		0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0,
		0, 0, 31, 1, 0, 0, 0, 1, 33, 1, 0, 0, 0, 3, 35, 1, 0, 0, 0, 5, 37, 1, 0,
		0, 0, 7, 39, 1, 0, 0, 0, 9, 41, 1, 0, 0, 0, 11, 43, 1, 0, 0, 0, 13, 45,
		1, 0, 0, 0, 15, 47, 1, 0, 0, 0, 17, 49, 1, 0, 0, 0, 19, 51, 1, 0, 0, 0,
		21, 58, 1, 0, 0, 0, 23, 63, 1, 0, 0, 0, 25, 65, 1, 0, 0, 0, 27, 68, 1,
		0, 0, 0, 29, 73, 1, 0, 0, 0, 31, 78, 1, 0, 0, 0, 33, 34, 5, 61, 0, 0, 34,
		2, 1, 0, 0, 0, 35, 36, 5, 42, 0, 0, 36, 4, 1, 0, 0, 0, 37, 38, 5, 47, 0,
		0, 38, 6, 1, 0, 0, 0, 39, 40, 5, 43, 0, 0, 40, 8, 1, 0, 0, 0, 41, 42, 5,
		45, 0, 0, 42, 10, 1, 0, 0, 0, 43, 44, 5, 44, 0, 0, 44, 12, 1, 0, 0, 0,
		45, 46, 5, 123, 0, 0, 46, 14, 1, 0, 0, 0, 47, 48, 5, 59, 0, 0, 48, 16,
		1, 0, 0, 0, 49, 50, 5, 125, 0, 0, 50, 18, 1, 0, 0, 0, 51, 52, 5, 114, 0,
		0, 52, 53, 5, 101, 0, 0, 53, 54, 5, 116, 0, 0, 54, 55, 5, 117, 0, 0, 55,
		56, 5, 114, 0, 0, 56, 57, 5, 110, 0, 0, 57, 20, 1, 0, 0, 0, 58, 59, 5,
		112, 0, 0, 59, 60, 5, 114, 0, 0, 60, 61, 5, 111, 0, 0, 61, 62, 5, 99, 0,
		0, 62, 22, 1, 0, 0, 0, 63, 64, 5, 40, 0, 0, 64, 24, 1, 0, 0, 0, 65, 66,
		5, 41, 0, 0, 66, 26, 1, 0, 0, 0, 67, 69, 7, 0, 0, 0, 68, 67, 1, 0, 0, 0,
		69, 70, 1, 0, 0, 0, 70, 68, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 28, 1,
		0, 0, 0, 72, 74, 7, 1, 0, 0, 73, 72, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75,
		73, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76, 30, 1, 0, 0, 0, 77, 79, 7, 2, 0,
		0, 78, 77, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 78, 1, 0, 0, 0, 80, 81,
		1, 0, 0, 0, 81, 82, 1, 0, 0, 0, 82, 83, 6, 15, 0, 0, 83, 32, 1, 0, 0, 0,
		4, 0, 70, 75, 80, 1, 6, 0, 0,
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
	styx_generalLexerT__0  = 1
	styx_generalLexerT__1  = 2
	styx_generalLexerT__2  = 3
	styx_generalLexerT__3  = 4
	styx_generalLexerT__4  = 5
	styx_generalLexerT__5  = 6
	styx_generalLexerT__6  = 7
	styx_generalLexerT__7  = 8
	styx_generalLexerT__8  = 9
	styx_generalLexerT__9  = 10
	styx_generalLexerT__10 = 11
	styx_generalLexerT__11 = 12
	styx_generalLexerT__12 = 13
	styx_generalLexerINT   = 14
	styx_generalLexerID    = 15
	styx_generalLexerWS    = 16
)
