// Code generated from Kl.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Kl

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 15, 93, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 7, 4, 30, 10, 4, 12, 4, 14, 4, 33, 11, 4, 3, 5, 3, 5, 7,
	5, 37, 10, 5, 12, 5, 14, 5, 40, 11, 5, 3, 5, 3, 5, 3, 5, 7, 5, 45, 10,
	5, 12, 5, 14, 5, 48, 11, 5, 3, 5, 3, 5, 7, 5, 52, 10, 5, 12, 5, 14, 5,
	55, 11, 5, 3, 5, 3, 5, 3, 5, 7, 5, 60, 10, 5, 12, 5, 14, 5, 63, 11, 5,
	3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 7, 6, 70, 10, 6, 12, 6, 14, 6, 73, 11, 6,
	3, 7, 3, 7, 3, 7, 7, 7, 78, 10, 7, 12, 7, 14, 7, 81, 11, 7, 3, 7, 3, 7,
	3, 8, 3, 8, 3, 8, 7, 8, 88, 10, 8, 12, 8, 14, 8, 91, 11, 8, 3, 8, 2, 2,
	9, 2, 4, 6, 8, 10, 12, 14, 2, 3, 3, 2, 12, 14, 2, 96, 2, 16, 3, 2, 2, 2,
	4, 21, 3, 2, 2, 2, 6, 23, 3, 2, 2, 2, 8, 34, 3, 2, 2, 2, 10, 66, 3, 2,
	2, 2, 12, 74, 3, 2, 2, 2, 14, 89, 3, 2, 2, 2, 16, 17, 7, 12, 2, 2, 17,
	18, 7, 3, 2, 2, 18, 19, 5, 4, 3, 2, 19, 20, 7, 4, 2, 2, 20, 3, 3, 2, 2,
	2, 21, 22, 9, 2, 2, 2, 22, 5, 3, 2, 2, 2, 23, 24, 7, 12, 2, 2, 24, 25,
	7, 5, 2, 2, 25, 31, 7, 12, 2, 2, 26, 27, 7, 6, 2, 2, 27, 28, 7, 12, 2,
	2, 28, 30, 7, 12, 2, 2, 29, 26, 3, 2, 2, 2, 30, 33, 3, 2, 2, 2, 31, 29,
	3, 2, 2, 2, 31, 32, 3, 2, 2, 2, 32, 7, 3, 2, 2, 2, 33, 31, 3, 2, 2, 2,
	34, 38, 7, 7, 2, 2, 35, 37, 7, 5, 2, 2, 36, 35, 3, 2, 2, 2, 37, 40, 3,
	2, 2, 2, 38, 36, 3, 2, 2, 2, 38, 39, 3, 2, 2, 2, 39, 41, 3, 2, 2, 2, 40,
	38, 3, 2, 2, 2, 41, 42, 7, 12, 2, 2, 42, 46, 7, 8, 2, 2, 43, 45, 5, 6,
	4, 2, 44, 43, 3, 2, 2, 2, 45, 48, 3, 2, 2, 2, 46, 44, 3, 2, 2, 2, 46, 47,
	3, 2, 2, 2, 47, 49, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 49, 53, 7, 9, 2, 2,
	50, 52, 7, 5, 2, 2, 51, 50, 3, 2, 2, 2, 52, 55, 3, 2, 2, 2, 53, 51, 3,
	2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 56, 3, 2, 2, 2, 55, 53, 3, 2, 2, 2, 56,
	61, 7, 10, 2, 2, 57, 60, 5, 2, 2, 2, 58, 60, 5, 12, 7, 2, 59, 57, 3, 2,
	2, 2, 59, 58, 3, 2, 2, 2, 60, 63, 3, 2, 2, 2, 61, 59, 3, 2, 2, 2, 61, 62,
	3, 2, 2, 2, 62, 64, 3, 2, 2, 2, 63, 61, 3, 2, 2, 2, 64, 65, 7, 11, 2, 2,
	65, 9, 3, 2, 2, 2, 66, 71, 5, 4, 3, 2, 67, 68, 7, 6, 2, 2, 68, 70, 5, 4,
	3, 2, 69, 67, 3, 2, 2, 2, 70, 73, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 71, 72,
	3, 2, 2, 2, 72, 11, 3, 2, 2, 2, 73, 71, 3, 2, 2, 2, 74, 75, 7, 12, 2, 2,
	75, 79, 7, 8, 2, 2, 76, 78, 5, 10, 6, 2, 77, 76, 3, 2, 2, 2, 78, 81, 3,
	2, 2, 2, 79, 77, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80, 82, 3, 2, 2, 2, 81,
	79, 3, 2, 2, 2, 82, 83, 7, 9, 2, 2, 83, 13, 3, 2, 2, 2, 84, 88, 5, 2, 2,
	2, 85, 88, 5, 8, 5, 2, 86, 88, 5, 12, 7, 2, 87, 84, 3, 2, 2, 2, 87, 85,
	3, 2, 2, 2, 87, 86, 3, 2, 2, 2, 88, 91, 3, 2, 2, 2, 89, 87, 3, 2, 2, 2,
	89, 90, 3, 2, 2, 2, 90, 15, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 12, 31, 38,
	46, 53, 59, 61, 71, 79, 87, 89,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'='", "';'", "' '", "','", "'func'", "'('", "')'", "'{'", "'}'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "Id", "STRING", "Num", "SW",
}

var ruleNames = []string{
	"stat", "value", "param", "fun", "args", "callFun", "init",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type KlParser struct {
	*antlr.BaseParser
}

func NewKlParser(input antlr.TokenStream) *KlParser {
	this := new(KlParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Kl.g4"

	return this
}

// KlParser tokens.
const (
	KlParserEOF    = antlr.TokenEOF
	KlParserT__0   = 1
	KlParserT__1   = 2
	KlParserT__2   = 3
	KlParserT__3   = 4
	KlParserT__4   = 5
	KlParserT__5   = 6
	KlParserT__6   = 7
	KlParserT__7   = 8
	KlParserT__8   = 9
	KlParserId     = 10
	KlParserSTRING = 11
	KlParserNum    = 12
	KlParserSW     = 13
)

// KlParser rules.
const (
	KlParserRULE_stat    = 0
	KlParserRULE_value   = 1
	KlParserRULE_param   = 2
	KlParserRULE_fun     = 3
	KlParserRULE_args    = 4
	KlParserRULE_callFun = 5
	KlParserRULE_init    = 6
)

// IStatContext is an interface to support dynamic dispatch.
type IStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatContext differentiates from other interfaces.
	IsStatContext()
}

type StatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatContext() *StatContext {
	var p = new(StatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KlParserRULE_stat
	return p
}

func (*StatContext) IsStatContext() {}

func NewStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatContext {
	var p = new(StatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlParserRULE_stat

	return p
}

func (s *StatContext) GetParser() antlr.Parser { return s.parser }

func (s *StatContext) Id() antlr.TerminalNode {
	return s.GetToken(KlParserId, 0)
}

func (s *StatContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *StatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlListener); ok {
		listenerT.EnterStat(s)
	}
}

func (s *StatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlListener); ok {
		listenerT.ExitStat(s)
	}
}

func (p *KlParser) Stat() (localctx IStatContext) {
	localctx = NewStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, KlParserRULE_stat)

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
		p.SetState(14)
		p.Match(KlParserId)
	}
	{
		p.SetState(15)
		p.Match(KlParserT__0)
	}
	{
		p.SetState(16)
		p.Value()
	}
	{
		p.SetState(17)
		p.Match(KlParserT__1)
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KlParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) Id() antlr.TerminalNode {
	return s.GetToken(KlParserId, 0)
}

func (s *ValueContext) Num() antlr.TerminalNode {
	return s.GetToken(KlParserNum, 0)
}

func (s *ValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(KlParserSTRING, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *KlParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, KlParserRULE_value)
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
		p.SetState(19)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<KlParserId)|(1<<KlParserSTRING)|(1<<KlParserNum))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IParamContext is an interface to support dynamic dispatch.
type IParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamContext() *ParamContext {
	var p = new(ParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KlParserRULE_param
	return p
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlParserRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) AllId() []antlr.TerminalNode {
	return s.GetTokens(KlParserId)
}

func (s *ParamContext) Id(i int) antlr.TerminalNode {
	return s.GetToken(KlParserId, i)
}

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlListener); ok {
		listenerT.EnterParam(s)
	}
}

func (s *ParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlListener); ok {
		listenerT.ExitParam(s)
	}
}

func (p *KlParser) Param() (localctx IParamContext) {
	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, KlParserRULE_param)
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
		p.SetState(21)
		p.Match(KlParserId)
	}
	{
		p.SetState(22)
		p.Match(KlParserT__2)
	}
	{
		p.SetState(23)
		p.Match(KlParserId)
	}
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == KlParserT__3 {
		{
			p.SetState(24)
			p.Match(KlParserT__3)
		}
		{
			p.SetState(25)
			p.Match(KlParserId)
		}
		{
			p.SetState(26)
			p.Match(KlParserId)
		}

		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFunContext is an interface to support dynamic dispatch.
type IFunContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunContext differentiates from other interfaces.
	IsFunContext()
}

type FunContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunContext() *FunContext {
	var p = new(FunContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KlParserRULE_fun
	return p
}

func (*FunContext) IsFunContext() {}

func NewFunContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunContext {
	var p = new(FunContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlParserRULE_fun

	return p
}

func (s *FunContext) GetParser() antlr.Parser { return s.parser }

func (s *FunContext) Id() antlr.TerminalNode {
	return s.GetToken(KlParserId, 0)
}

func (s *FunContext) AllParam() []IParamContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParamContext)(nil)).Elem())
	var tst = make([]IParamContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParamContext)
		}
	}

	return tst
}

func (s *FunContext) Param(i int) IParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParamContext)
}

func (s *FunContext) AllStat() []IStatContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatContext)(nil)).Elem())
	var tst = make([]IStatContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatContext)
		}
	}

	return tst
}

func (s *FunContext) Stat(i int) IStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *FunContext) AllCallFun() []ICallFunContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICallFunContext)(nil)).Elem())
	var tst = make([]ICallFunContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICallFunContext)
		}
	}

	return tst
}

func (s *FunContext) CallFun(i int) ICallFunContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICallFunContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICallFunContext)
}

func (s *FunContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlListener); ok {
		listenerT.EnterFun(s)
	}
}

func (s *FunContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlListener); ok {
		listenerT.ExitFun(s)
	}
}

func (p *KlParser) Fun() (localctx IFunContext) {
	localctx = NewFunContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, KlParserRULE_fun)
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
		p.SetState(32)
		p.Match(KlParserT__4)
	}
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == KlParserT__2 {
		{
			p.SetState(33)
			p.Match(KlParserT__2)
		}

		p.SetState(38)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(39)
		p.Match(KlParserId)
	}
	{
		p.SetState(40)
		p.Match(KlParserT__5)
	}
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == KlParserId {
		{
			p.SetState(41)
			p.Param()
		}

		p.SetState(46)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(47)
		p.Match(KlParserT__6)
	}
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == KlParserT__2 {
		{
			p.SetState(48)
			p.Match(KlParserT__2)
		}

		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(54)
		p.Match(KlParserT__7)
	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == KlParserId {
		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(55)
				p.Stat()
			}

		case 2:
			{
				p.SetState(56)
				p.CallFun()
			}

		}

		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(62)
		p.Match(KlParserT__8)
	}

	return localctx
}

// IArgsContext is an interface to support dynamic dispatch.
type IArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgsContext differentiates from other interfaces.
	IsArgsContext()
}

type ArgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgsContext() *ArgsContext {
	var p = new(ArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KlParserRULE_args
	return p
}

func (*ArgsContext) IsArgsContext() {}

func NewArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgsContext {
	var p = new(ArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlParserRULE_args

	return p
}

func (s *ArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgsContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *ArgsContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlListener); ok {
		listenerT.EnterArgs(s)
	}
}

func (s *ArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlListener); ok {
		listenerT.ExitArgs(s)
	}
}

func (p *KlParser) Args() (localctx IArgsContext) {
	localctx = NewArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, KlParserRULE_args)
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
		p.SetState(64)
		p.Value()
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == KlParserT__3 {
		{
			p.SetState(65)
			p.Match(KlParserT__3)
		}
		{
			p.SetState(66)
			p.Value()
		}

		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ICallFunContext is an interface to support dynamic dispatch.
type ICallFunContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCallFunContext differentiates from other interfaces.
	IsCallFunContext()
}

type CallFunContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCallFunContext() *CallFunContext {
	var p = new(CallFunContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KlParserRULE_callFun
	return p
}

func (*CallFunContext) IsCallFunContext() {}

func NewCallFunContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallFunContext {
	var p = new(CallFunContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlParserRULE_callFun

	return p
}

func (s *CallFunContext) GetParser() antlr.Parser { return s.parser }

func (s *CallFunContext) Id() antlr.TerminalNode {
	return s.GetToken(KlParserId, 0)
}

func (s *CallFunContext) AllArgs() []IArgsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgsContext)(nil)).Elem())
	var tst = make([]IArgsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgsContext)
		}
	}

	return tst
}

func (s *CallFunContext) Args(i int) IArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
}

func (s *CallFunContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallFunContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CallFunContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlListener); ok {
		listenerT.EnterCallFun(s)
	}
}

func (s *CallFunContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlListener); ok {
		listenerT.ExitCallFun(s)
	}
}

func (p *KlParser) CallFun() (localctx ICallFunContext) {
	localctx = NewCallFunContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, KlParserRULE_callFun)
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
		p.SetState(72)
		p.Match(KlParserId)
	}
	{
		p.SetState(73)
		p.Match(KlParserT__5)
	}
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<KlParserId)|(1<<KlParserSTRING)|(1<<KlParserNum))) != 0 {
		{
			p.SetState(74)
			p.Args()
		}

		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(80)
		p.Match(KlParserT__6)
	}

	return localctx
}

// IInitContext is an interface to support dynamic dispatch.
type IInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInitContext differentiates from other interfaces.
	IsInitContext()
}

type InitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInitContext() *InitContext {
	var p = new(InitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KlParserRULE_init
	return p
}

func (*InitContext) IsInitContext() {}

func NewInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitContext {
	var p = new(InitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KlParserRULE_init

	return p
}

func (s *InitContext) GetParser() antlr.Parser { return s.parser }

func (s *InitContext) AllStat() []IStatContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatContext)(nil)).Elem())
	var tst = make([]IStatContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatContext)
		}
	}

	return tst
}

func (s *InitContext) Stat(i int) IStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *InitContext) AllFun() []IFunContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunContext)(nil)).Elem())
	var tst = make([]IFunContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunContext)
		}
	}

	return tst
}

func (s *InitContext) Fun(i int) IFunContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunContext)
}

func (s *InitContext) AllCallFun() []ICallFunContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICallFunContext)(nil)).Elem())
	var tst = make([]ICallFunContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICallFunContext)
		}
	}

	return tst
}

func (s *InitContext) CallFun(i int) ICallFunContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICallFunContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICallFunContext)
}

func (s *InitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlListener); ok {
		listenerT.EnterInit(s)
	}
}

func (s *InitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KlListener); ok {
		listenerT.ExitInit(s)
	}
}

func (p *KlParser) Init() (localctx IInitContext) {
	localctx = NewInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, KlParserRULE_init)
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
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == KlParserT__4 || _la == KlParserId {
		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(82)
				p.Stat()
			}

		case 2:
			{
				p.SetState(83)
				p.Fun()
			}

		case 3:
			{
				p.SetState(84)
				p.CallFun()
			}

		}

		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}
