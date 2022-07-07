// Code generated from Styx_refactoring.g4 by ANTLR 4.10.1. DO NOT EDIT.

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

type Styx_refactoringLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var styx_refactoringlexerLexerStaticData struct {
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

func styx_refactoringlexerLexerInit() {
	staticData := &styx_refactoringlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'+'", "'-'", "'*'", "'/'", "'{'", "'}'", "'('", "')'", "'proc'",
		"'='", "", "", "';'", "','",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "INT", "ID", "SEMICOLON",
		"COMMA", "WS",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "INT", "ID", "SEMICOLON", "COMMA", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 15, 75, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 1, 0,
		1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6,
		1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 4, 10,
		56, 8, 10, 11, 10, 12, 10, 57, 1, 11, 4, 11, 61, 8, 11, 11, 11, 12, 11,
		62, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 4, 14, 70, 8, 14, 11, 14, 12, 14,
		71, 1, 14, 1, 14, 0, 0, 15, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7,
		15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 1, 0, 3,
		1, 0, 48, 57, 3, 0, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32,
		77, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0,
		0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0,
		0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1,
		0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 1, 31,
		1, 0, 0, 0, 3, 33, 1, 0, 0, 0, 5, 35, 1, 0, 0, 0, 7, 37, 1, 0, 0, 0, 9,
		39, 1, 0, 0, 0, 11, 41, 1, 0, 0, 0, 13, 43, 1, 0, 0, 0, 15, 45, 1, 0, 0,
		0, 17, 47, 1, 0, 0, 0, 19, 52, 1, 0, 0, 0, 21, 55, 1, 0, 0, 0, 23, 60,
		1, 0, 0, 0, 25, 64, 1, 0, 0, 0, 27, 66, 1, 0, 0, 0, 29, 69, 1, 0, 0, 0,
		31, 32, 5, 43, 0, 0, 32, 2, 1, 0, 0, 0, 33, 34, 5, 45, 0, 0, 34, 4, 1,
		0, 0, 0, 35, 36, 5, 42, 0, 0, 36, 6, 1, 0, 0, 0, 37, 38, 5, 47, 0, 0, 38,
		8, 1, 0, 0, 0, 39, 40, 5, 123, 0, 0, 40, 10, 1, 0, 0, 0, 41, 42, 5, 125,
		0, 0, 42, 12, 1, 0, 0, 0, 43, 44, 5, 40, 0, 0, 44, 14, 1, 0, 0, 0, 45,
		46, 5, 41, 0, 0, 46, 16, 1, 0, 0, 0, 47, 48, 5, 112, 0, 0, 48, 49, 5, 114,
		0, 0, 49, 50, 5, 111, 0, 0, 50, 51, 5, 99, 0, 0, 51, 18, 1, 0, 0, 0, 52,
		53, 5, 61, 0, 0, 53, 20, 1, 0, 0, 0, 54, 56, 7, 0, 0, 0, 55, 54, 1, 0,
		0, 0, 56, 57, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 22,
		1, 0, 0, 0, 59, 61, 7, 1, 0, 0, 60, 59, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0,
		62, 60, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 24, 1, 0, 0, 0, 64, 65, 5,
		59, 0, 0, 65, 26, 1, 0, 0, 0, 66, 67, 5, 44, 0, 0, 67, 28, 1, 0, 0, 0,
		68, 70, 7, 2, 0, 0, 69, 68, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 69, 1,
		0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 73, 74, 6, 14, 0, 0, 74,
		30, 1, 0, 0, 0, 4, 0, 57, 62, 71, 1, 6, 0, 0,
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

// Styx_refactoringLexerInit initializes any static state used to implement Styx_refactoringLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewStyx_refactoringLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func Styx_refactoringLexerInit() {
	staticData := &styx_refactoringlexerLexerStaticData
	staticData.once.Do(styx_refactoringlexerLexerInit)
}

// NewStyx_refactoringLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewStyx_refactoringLexer(input antlr.CharStream) *Styx_refactoringLexer {
	Styx_refactoringLexerInit()
	l := new(Styx_refactoringLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &styx_refactoringlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Styx_refactoring.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// Styx_refactoringLexer tokens.
const (
	Styx_refactoringLexerT__0      = 1
	Styx_refactoringLexerT__1      = 2
	Styx_refactoringLexerT__2      = 3
	Styx_refactoringLexerT__3      = 4
	Styx_refactoringLexerT__4      = 5
	Styx_refactoringLexerT__5      = 6
	Styx_refactoringLexerT__6      = 7
	Styx_refactoringLexerT__7      = 8
	Styx_refactoringLexerT__8      = 9
	Styx_refactoringLexerT__9      = 10
	Styx_refactoringLexerINT       = 11
	Styx_refactoringLexerID        = 12
	Styx_refactoringLexerSEMICOLON = 13
	Styx_refactoringLexerCOMMA     = 14
	Styx_refactoringLexerWS        = 15
)
