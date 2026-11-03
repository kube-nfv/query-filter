// Code generated from grammar/Filter.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
  	"sync"
	"unicode"
	"github.com/antlr4-go/antlr/v4"
)
// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter


type FilterLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

var FilterLexerLexerStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  ChannelNames           []string
  ModeNames              []string
  LiteralNames           []string
  SymbolicNames          []string
  RuleNames              []string
  PredictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func filterlexerLexerInit() {
  staticData := &FilterLexerLexerStaticData
  staticData.ChannelNames = []string{
    "DEFAULT_TOKEN_CHANNEL", "HIDDEN",
  }
  staticData.ModeNames = []string{
    "DEFAULT_MODE",
  }
  staticData.LiteralNames = []string{
    "", "'filter'", "'='", "';'", "'('", "')'", "','", "'/'", "'eq'", "'neq'", 
    "'gt'", "'lt'", "'gte'", "'lte'", "'in'", "'nin'", "'cont'", "'ncont'",
  }
  staticData.SymbolicNames = []string{
    "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
    "", "QUOTED", "STRING", "WS",
  }
  staticData.RuleNames = []string{
    "T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8", 
    "T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16", 
    "QUOTED", "STRING", "WS",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 0, 20, 122, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 
	4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 
	10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 
	7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 1, 0, 1, 
	0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 
	4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 
	8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 
	12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 
	1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 
	16, 1, 17, 1, 17, 1, 17, 1, 17, 5, 17, 104, 8, 17, 10, 17, 12, 17, 107, 
	9, 17, 1, 17, 1, 17, 1, 18, 4, 18, 112, 8, 18, 11, 18, 12, 18, 113, 1, 
	19, 4, 19, 117, 8, 19, 11, 19, 12, 19, 118, 1, 19, 1, 19, 0, 0, 20, 1, 
	1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 
	23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 
	1, 0, 3, 1, 0, 39, 39, 9, 0, 33, 33, 35, 38, 42, 43, 45, 46, 48, 58, 63, 
	90, 95, 95, 97, 122, 126, 126, 3, 0, 9, 10, 13, 13, 32, 32, 125, 0, 1, 
	1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 
	1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 
	17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 
	0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 
	0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 
	0, 0, 1, 41, 1, 0, 0, 0, 3, 48, 1, 0, 0, 0, 5, 50, 1, 0, 0, 0, 7, 52, 1, 
	0, 0, 0, 9, 54, 1, 0, 0, 0, 11, 56, 1, 0, 0, 0, 13, 58, 1, 0, 0, 0, 15, 
	60, 1, 0, 0, 0, 17, 63, 1, 0, 0, 0, 19, 67, 1, 0, 0, 0, 21, 70, 1, 0, 0, 
	0, 23, 73, 1, 0, 0, 0, 25, 77, 1, 0, 0, 0, 27, 81, 1, 0, 0, 0, 29, 84, 
	1, 0, 0, 0, 31, 88, 1, 0, 0, 0, 33, 93, 1, 0, 0, 0, 35, 99, 1, 0, 0, 0, 
	37, 111, 1, 0, 0, 0, 39, 116, 1, 0, 0, 0, 41, 42, 5, 102, 0, 0, 42, 43, 
	5, 105, 0, 0, 43, 44, 5, 108, 0, 0, 44, 45, 5, 116, 0, 0, 45, 46, 5, 101, 
	0, 0, 46, 47, 5, 114, 0, 0, 47, 2, 1, 0, 0, 0, 48, 49, 5, 61, 0, 0, 49, 
	4, 1, 0, 0, 0, 50, 51, 5, 59, 0, 0, 51, 6, 1, 0, 0, 0, 52, 53, 5, 40, 0, 
	0, 53, 8, 1, 0, 0, 0, 54, 55, 5, 41, 0, 0, 55, 10, 1, 0, 0, 0, 56, 57, 
	5, 44, 0, 0, 57, 12, 1, 0, 0, 0, 58, 59, 5, 47, 0, 0, 59, 14, 1, 0, 0, 
	0, 60, 61, 5, 101, 0, 0, 61, 62, 5, 113, 0, 0, 62, 16, 1, 0, 0, 0, 63, 
	64, 5, 110, 0, 0, 64, 65, 5, 101, 0, 0, 65, 66, 5, 113, 0, 0, 66, 18, 1, 
	0, 0, 0, 67, 68, 5, 103, 0, 0, 68, 69, 5, 116, 0, 0, 69, 20, 1, 0, 0, 0, 
	70, 71, 5, 108, 0, 0, 71, 72, 5, 116, 0, 0, 72, 22, 1, 0, 0, 0, 73, 74, 
	5, 103, 0, 0, 74, 75, 5, 116, 0, 0, 75, 76, 5, 101, 0, 0, 76, 24, 1, 0, 
	0, 0, 77, 78, 5, 108, 0, 0, 78, 79, 5, 116, 0, 0, 79, 80, 5, 101, 0, 0, 
	80, 26, 1, 0, 0, 0, 81, 82, 5, 105, 0, 0, 82, 83, 5, 110, 0, 0, 83, 28, 
	1, 0, 0, 0, 84, 85, 5, 110, 0, 0, 85, 86, 5, 105, 0, 0, 86, 87, 5, 110, 
	0, 0, 87, 30, 1, 0, 0, 0, 88, 89, 5, 99, 0, 0, 89, 90, 5, 111, 0, 0, 90, 
	91, 5, 110, 0, 0, 91, 92, 5, 116, 0, 0, 92, 32, 1, 0, 0, 0, 93, 94, 5, 
	110, 0, 0, 94, 95, 5, 99, 0, 0, 95, 96, 5, 111, 0, 0, 96, 97, 5, 110, 0, 
	0, 97, 98, 5, 116, 0, 0, 98, 34, 1, 0, 0, 0, 99, 105, 5, 39, 0, 0, 100, 
	104, 8, 0, 0, 0, 101, 102, 5, 39, 0, 0, 102, 104, 5, 39, 0, 0, 103, 100, 
	1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 104, 107, 1, 0, 0, 0, 105, 103, 1, 0, 
	0, 0, 105, 106, 1, 0, 0, 0, 106, 108, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 
	108, 109, 5, 39, 0, 0, 109, 36, 1, 0, 0, 0, 110, 112, 7, 1, 0, 0, 111, 
	110, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 113, 114, 
	1, 0, 0, 0, 114, 38, 1, 0, 0, 0, 115, 117, 7, 2, 0, 0, 116, 115, 1, 0, 
	0, 0, 117, 118, 1, 0, 0, 0, 118, 116, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 
	119, 120, 1, 0, 0, 0, 120, 121, 6, 19, 0, 0, 121, 40, 1, 0, 0, 0, 5, 0, 
	103, 105, 113, 118, 1, 6, 0, 0,
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

// FilterLexerInit initializes any static state used to implement FilterLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewFilterLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func FilterLexerInit() {
  staticData := &FilterLexerLexerStaticData
  staticData.once.Do(filterlexerLexerInit)
}

// NewFilterLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewFilterLexer(input antlr.CharStream) *FilterLexer {
  FilterLexerInit()
	l := new(FilterLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
  staticData := &FilterLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Filter.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// FilterLexer tokens.
const (
	FilterLexerT__0 = 1
	FilterLexerT__1 = 2
	FilterLexerT__2 = 3
	FilterLexerT__3 = 4
	FilterLexerT__4 = 5
	FilterLexerT__5 = 6
	FilterLexerT__6 = 7
	FilterLexerT__7 = 8
	FilterLexerT__8 = 9
	FilterLexerT__9 = 10
	FilterLexerT__10 = 11
	FilterLexerT__11 = 12
	FilterLexerT__12 = 13
	FilterLexerT__13 = 14
	FilterLexerT__14 = 15
	FilterLexerT__15 = 16
	FilterLexerT__16 = 17
	FilterLexerQUOTED = 18
	FilterLexerSTRING = 19
	FilterLexerWS = 20
)

