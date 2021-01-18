// Code generated from Kl.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 15, 75, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 4, 14, 9, 14, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10,
	3, 10, 3, 11, 6, 11, 52, 10, 11, 13, 11, 14, 11, 53, 3, 12, 3, 12, 6, 12,
	58, 10, 12, 13, 12, 14, 12, 59, 3, 12, 3, 12, 3, 13, 6, 13, 65, 10, 13,
	13, 13, 14, 13, 66, 3, 14, 6, 14, 70, 10, 14, 13, 14, 14, 14, 71, 3, 14,
	3, 14, 2, 2, 15, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19,
	11, 21, 12, 23, 13, 25, 14, 27, 15, 3, 2, 5, 5, 2, 67, 92, 97, 97, 99,
	124, 3, 2, 50, 59, 5, 2, 12, 12, 15, 15, 34, 34, 2, 78, 2, 3, 3, 2, 2,
	2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2,
	2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2,
	2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3,
	2, 2, 2, 3, 29, 3, 2, 2, 2, 5, 31, 3, 2, 2, 2, 7, 33, 3, 2, 2, 2, 9, 35,
	3, 2, 2, 2, 11, 37, 3, 2, 2, 2, 13, 42, 3, 2, 2, 2, 15, 44, 3, 2, 2, 2,
	17, 46, 3, 2, 2, 2, 19, 48, 3, 2, 2, 2, 21, 51, 3, 2, 2, 2, 23, 55, 3,
	2, 2, 2, 25, 64, 3, 2, 2, 2, 27, 69, 3, 2, 2, 2, 29, 30, 7, 63, 2, 2, 30,
	4, 3, 2, 2, 2, 31, 32, 7, 61, 2, 2, 32, 6, 3, 2, 2, 2, 33, 34, 7, 34, 2,
	2, 34, 8, 3, 2, 2, 2, 35, 36, 7, 46, 2, 2, 36, 10, 3, 2, 2, 2, 37, 38,
	7, 104, 2, 2, 38, 39, 7, 119, 2, 2, 39, 40, 7, 112, 2, 2, 40, 41, 7, 101,
	2, 2, 41, 12, 3, 2, 2, 2, 42, 43, 7, 42, 2, 2, 43, 14, 3, 2, 2, 2, 44,
	45, 7, 43, 2, 2, 45, 16, 3, 2, 2, 2, 46, 47, 7, 125, 2, 2, 47, 18, 3, 2,
	2, 2, 48, 49, 7, 127, 2, 2, 49, 20, 3, 2, 2, 2, 50, 52, 9, 2, 2, 2, 51,
	50, 3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2, 53, 54, 3, 2, 2,
	2, 54, 22, 3, 2, 2, 2, 55, 57, 7, 36, 2, 2, 56, 58, 9, 2, 2, 2, 57, 56,
	3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 57, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2,
	60, 61, 3, 2, 2, 2, 61, 62, 7, 36, 2, 2, 62, 24, 3, 2, 2, 2, 63, 65, 9,
	3, 2, 2, 64, 63, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 64, 3, 2, 2, 2, 66,
	67, 3, 2, 2, 2, 67, 26, 3, 2, 2, 2, 68, 70, 9, 4, 2, 2, 69, 68, 3, 2, 2,
	2, 70, 71, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 73,
	3, 2, 2, 2, 73, 74, 8, 14, 2, 2, 74, 28, 3, 2, 2, 2, 7, 2, 53, 59, 66,
	71, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'='", "';'", "' '", "','", "'func'", "'('", "')'", "'{'", "'}'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "Id", "STRING", "Num", "SW",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"Id", "STRING", "Num", "SW",
}

type KlLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewKlLexer(input antlr.CharStream) *KlLexer {

	l := new(KlLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Kl.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// KlLexer tokens.
const (
	KlLexerT__0   = 1
	KlLexerT__1   = 2
	KlLexerT__2   = 3
	KlLexerT__3   = 4
	KlLexerT__4   = 5
	KlLexerT__5   = 6
	KlLexerT__6   = 7
	KlLexerT__7   = 8
	KlLexerT__8   = 9
	KlLexerId     = 10
	KlLexerSTRING = 11
	KlLexerNum    = 12
	KlLexerSW     = 13
)
