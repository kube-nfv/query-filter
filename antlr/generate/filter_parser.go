// Code generated from grammar/Filter.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Filter

import (
	"fmt"
	"strconv"
  	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}


type FilterParser struct {
	*antlr.BaseParser
}

var FilterParserStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  LiteralNames           []string
  SymbolicNames          []string
  RuleNames              []string
  PredictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func filterParserInit() {
  staticData := &FilterParserStaticData
  staticData.LiteralNames = []string{
    "", "'filter'", "'='", "';'", "'('", "')'", "','", "'/'", "'eq'", "'neq'", 
    "'gt'", "'lt'", "'gte'", "'lte'", "'in'", "'nin'", "'cont'", "'ncont'",
  }
  staticData.SymbolicNames = []string{
    "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", 
    "", "QUOTED", "STRING", "WS",
  }
  staticData.RuleNames = []string{
    "filter", "filterExpr", "simpleFilterExpr", "simpleFilterExprOne", "simpleFilterExprMulti", 
    "opOne", "opMulti", "attrName", "value", "name", "keyword",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 1, 20, 99, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 
	4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7, 
	10, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 5, 1, 30, 8, 1, 10, 1, 12, 
	1, 33, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 43, 
	8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 50, 8, 3, 10, 3, 12, 3, 53, 9, 
	3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 63, 8, 4, 10, 
	4, 12, 4, 66, 9, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 72, 8, 4, 10, 4, 12, 
	4, 75, 9, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 
	4, 8, 87, 8, 8, 11, 8, 12, 8, 88, 3, 8, 91, 8, 8, 1, 9, 1, 9, 3, 9, 95, 
	8, 9, 1, 10, 1, 10, 1, 10, 0, 0, 11, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 
	20, 0, 3, 1, 0, 8, 13, 1, 0, 14, 17, 2, 0, 1, 1, 8, 17, 97, 0, 22, 1, 0, 
	0, 0, 2, 26, 1, 0, 0, 0, 4, 42, 1, 0, 0, 0, 6, 44, 1, 0, 0, 0, 8, 57, 1, 
	0, 0, 0, 10, 76, 1, 0, 0, 0, 12, 78, 1, 0, 0, 0, 14, 80, 1, 0, 0, 0, 16, 
	90, 1, 0, 0, 0, 18, 94, 1, 0, 0, 0, 20, 96, 1, 0, 0, 0, 22, 23, 5, 1, 0, 
	0, 23, 24, 5, 2, 0, 0, 24, 25, 3, 2, 1, 0, 25, 1, 1, 0, 0, 0, 26, 31, 3, 
	4, 2, 0, 27, 28, 5, 3, 0, 0, 28, 30, 3, 4, 2, 0, 29, 27, 1, 0, 0, 0, 30, 
	33, 1, 0, 0, 0, 31, 29, 1, 0, 0, 0, 31, 32, 1, 0, 0, 0, 32, 3, 1, 0, 0, 
	0, 33, 31, 1, 0, 0, 0, 34, 35, 5, 4, 0, 0, 35, 36, 3, 6, 3, 0, 36, 37, 
	5, 5, 0, 0, 37, 43, 1, 0, 0, 0, 38, 39, 5, 4, 0, 0, 39, 40, 3, 8, 4, 0, 
	40, 41, 5, 5, 0, 0, 41, 43, 1, 0, 0, 0, 42, 34, 1, 0, 0, 0, 42, 38, 1, 
	0, 0, 0, 43, 5, 1, 0, 0, 0, 44, 45, 3, 10, 5, 0, 45, 46, 5, 6, 0, 0, 46, 
	51, 3, 14, 7, 0, 47, 48, 5, 7, 0, 0, 48, 50, 3, 14, 7, 0, 49, 47, 1, 0, 
	0, 0, 50, 53, 1, 0, 0, 0, 51, 49, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 54, 
	1, 0, 0, 0, 53, 51, 1, 0, 0, 0, 54, 55, 5, 6, 0, 0, 55, 56, 3, 16, 8, 0, 
	56, 7, 1, 0, 0, 0, 57, 58, 3, 12, 6, 0, 58, 59, 5, 6, 0, 0, 59, 64, 3, 
	14, 7, 0, 60, 61, 5, 7, 0, 0, 61, 63, 3, 14, 7, 0, 62, 60, 1, 0, 0, 0, 
	63, 66, 1, 0, 0, 0, 64, 62, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 67, 1, 
	0, 0, 0, 66, 64, 1, 0, 0, 0, 67, 68, 5, 6, 0, 0, 68, 73, 3, 16, 8, 0, 69, 
	70, 5, 6, 0, 0, 70, 72, 3, 16, 8, 0, 71, 69, 1, 0, 0, 0, 72, 75, 1, 0, 
	0, 0, 73, 71, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0, 74, 9, 1, 0, 0, 0, 75, 73, 
	1, 0, 0, 0, 76, 77, 7, 0, 0, 0, 77, 11, 1, 0, 0, 0, 78, 79, 7, 1, 0, 0, 
	79, 13, 1, 0, 0, 0, 80, 81, 3, 18, 9, 0, 81, 15, 1, 0, 0, 0, 82, 91, 5, 
	18, 0, 0, 83, 87, 3, 18, 9, 0, 84, 87, 5, 7, 0, 0, 85, 87, 5, 2, 0, 0, 
	86, 83, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0, 86, 85, 1, 0, 0, 0, 87, 88, 1, 
	0, 0, 0, 88, 86, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 91, 1, 0, 0, 0, 90, 
	82, 1, 0, 0, 0, 90, 86, 1, 0, 0, 0, 91, 17, 1, 0, 0, 0, 92, 95, 5, 19, 
	0, 0, 93, 95, 3, 20, 10, 0, 94, 92, 1, 0, 0, 0, 94, 93, 1, 0, 0, 0, 95, 
	19, 1, 0, 0, 0, 96, 97, 7, 2, 0, 0, 97, 21, 1, 0, 0, 0, 9, 31, 42, 51, 
	64, 73, 86, 88, 90, 94,
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

// FilterParserInit initializes any static state used to implement FilterParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewFilterParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func FilterParserInit() {
  staticData := &FilterParserStaticData
  staticData.once.Do(filterParserInit)
}

// NewFilterParser produces a new parser instance for the optional input antlr.TokenStream.
func NewFilterParser(input antlr.TokenStream) *FilterParser {
	FilterParserInit()
	this := new(FilterParser)
	this.BaseParser = antlr.NewBaseParser(input)
  staticData := &FilterParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Filter.g4"

	return this
}


// FilterParser tokens.
const (
	FilterParserEOF = antlr.TokenEOF
	FilterParserT__0 = 1
	FilterParserT__1 = 2
	FilterParserT__2 = 3
	FilterParserT__3 = 4
	FilterParserT__4 = 5
	FilterParserT__5 = 6
	FilterParserT__6 = 7
	FilterParserT__7 = 8
	FilterParserT__8 = 9
	FilterParserT__9 = 10
	FilterParserT__10 = 11
	FilterParserT__11 = 12
	FilterParserT__12 = 13
	FilterParserT__13 = 14
	FilterParserT__14 = 15
	FilterParserT__15 = 16
	FilterParserT__16 = 17
	FilterParserQUOTED = 18
	FilterParserSTRING = 19
	FilterParserWS = 20
)

// FilterParser rules.
const (
	FilterParserRULE_filter = 0
	FilterParserRULE_filterExpr = 1
	FilterParserRULE_simpleFilterExpr = 2
	FilterParserRULE_simpleFilterExprOne = 3
	FilterParserRULE_simpleFilterExprMulti = 4
	FilterParserRULE_opOne = 5
	FilterParserRULE_opMulti = 6
	FilterParserRULE_attrName = 7
	FilterParserRULE_value = 8
	FilterParserRULE_name = 9
	FilterParserRULE_keyword = 10
)

// IFilterContext is an interface to support dynamic dispatch.
type IFilterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FilterExpr() IFilterExprContext

	// IsFilterContext differentiates from other interfaces.
	IsFilterContext()
}

type FilterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilterContext() *FilterContext {
	var p = new(FilterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FilterParserRULE_filter
	return p
}

func InitEmptyFilterContext(p *FilterContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FilterParserRULE_filter
}

func (*FilterContext) IsFilterContext() {}

func NewFilterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterContext {
	var p = new(FilterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterParserRULE_filter

	return p
}

func (s *FilterContext) GetParser() antlr.Parser { return s.parser }

func (s *FilterContext) FilterExpr() IFilterExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFilterExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFilterExprContext)
}

func (s *FilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FilterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.EnterFilter(s)
	}
}

func (s *FilterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.ExitFilter(s)
	}
}




func (p *FilterParser) Filter() (localctx IFilterContext) {
	localctx = NewFilterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FilterParserRULE_filter)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(22)
		p.Match(FilterParserT__0)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(23)
		p.Match(FilterParserT__1)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(24)
		p.FilterExpr()
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IFilterExprContext is an interface to support dynamic dispatch.
type IFilterExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSimpleFilterExpr() []ISimpleFilterExprContext
	SimpleFilterExpr(i int) ISimpleFilterExprContext

	// IsFilterExprContext differentiates from other interfaces.
	IsFilterExprContext()
}

type FilterExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilterExprContext() *FilterExprContext {
	var p = new(FilterExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FilterParserRULE_filterExpr
	return p
}

func InitEmptyFilterExprContext(p *FilterExprContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FilterParserRULE_filterExpr
}

func (*FilterExprContext) IsFilterExprContext() {}

func NewFilterExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterExprContext {
	var p = new(FilterExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterParserRULE_filterExpr

	return p
}

func (s *FilterExprContext) GetParser() antlr.Parser { return s.parser }

func (s *FilterExprContext) AllSimpleFilterExpr() []ISimpleFilterExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISimpleFilterExprContext); ok {
			len++
		}
	}

	tst := make([]ISimpleFilterExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISimpleFilterExprContext); ok {
			tst[i] = t.(ISimpleFilterExprContext)
			i++
		}
	}

	return tst
}

func (s *FilterExprContext) SimpleFilterExpr(i int) ISimpleFilterExprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleFilterExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleFilterExprContext)
}

func (s *FilterExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FilterExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.EnterFilterExpr(s)
	}
}

func (s *FilterExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.ExitFilterExpr(s)
	}
}




func (p *FilterParser) FilterExpr() (localctx IFilterExprContext) {
	localctx = NewFilterExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FilterParserRULE_filterExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(26)
		p.SimpleFilterExpr()
	}
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == FilterParserT__2 {
		{
			p.SetState(27)
			p.Match(FilterParserT__2)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(28)
			p.SimpleFilterExpr()
		}


		p.SetState(33)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// ISimpleFilterExprContext is an interface to support dynamic dispatch.
type ISimpleFilterExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SimpleFilterExprOne() ISimpleFilterExprOneContext
	SimpleFilterExprMulti() ISimpleFilterExprMultiContext

	// IsSimpleFilterExprContext differentiates from other interfaces.
	IsSimpleFilterExprContext()
}

type SimpleFilterExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleFilterExprContext() *SimpleFilterExprContext {
	var p = new(SimpleFilterExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FilterParserRULE_simpleFilterExpr
	return p
}

func InitEmptySimpleFilterExprContext(p *SimpleFilterExprContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FilterParserRULE_simpleFilterExpr
}

func (*SimpleFilterExprContext) IsSimpleFilterExprContext() {}

func NewSimpleFilterExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleFilterExprContext {
	var p = new(SimpleFilterExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterParserRULE_simpleFilterExpr

	return p
}

func (s *SimpleFilterExprContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleFilterExprContext) SimpleFilterExprOne() ISimpleFilterExprOneContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleFilterExprOneContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleFilterExprOneContext)
}

func (s *SimpleFilterExprContext) SimpleFilterExprMulti() ISimpleFilterExprMultiContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleFilterExprMultiContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleFilterExprMultiContext)
}

func (s *SimpleFilterExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleFilterExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SimpleFilterExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.EnterSimpleFilterExpr(s)
	}
}

func (s *SimpleFilterExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.ExitSimpleFilterExpr(s)
	}
}




func (p *FilterParser) SimpleFilterExpr() (localctx ISimpleFilterExprContext) {
	localctx = NewSimpleFilterExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, FilterParserRULE_simpleFilterExpr)
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(34)
			p.Match(FilterParserT__3)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(35)
			p.SimpleFilterExprOne()
		}
		{
			p.SetState(36)
			p.Match(FilterParserT__4)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(38)
			p.Match(FilterParserT__3)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(39)
			p.SimpleFilterExprMulti()
		}
		{
			p.SetState(40)
			p.Match(FilterParserT__4)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}


errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// ISimpleFilterExprOneContext is an interface to support dynamic dispatch.
type ISimpleFilterExprOneContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OpOne() IOpOneContext
	AllAttrName() []IAttrNameContext
	AttrName(i int) IAttrNameContext
	Value() IValueContext

	// IsSimpleFilterExprOneContext differentiates from other interfaces.
	IsSimpleFilterExprOneContext()
}

type SimpleFilterExprOneContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleFilterExprOneContext() *SimpleFilterExprOneContext {
	var p = new(SimpleFilterExprOneContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FilterParserRULE_simpleFilterExprOne
	return p
}

func InitEmptySimpleFilterExprOneContext(p *SimpleFilterExprOneContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FilterParserRULE_simpleFilterExprOne
}

func (*SimpleFilterExprOneContext) IsSimpleFilterExprOneContext() {}

func NewSimpleFilterExprOneContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleFilterExprOneContext {
	var p = new(SimpleFilterExprOneContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterParserRULE_simpleFilterExprOne

	return p
}

func (s *SimpleFilterExprOneContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleFilterExprOneContext) OpOne() IOpOneContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpOneContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpOneContext)
}

func (s *SimpleFilterExprOneContext) AllAttrName() []IAttrNameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAttrNameContext); ok {
			len++
		}
	}

	tst := make([]IAttrNameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAttrNameContext); ok {
			tst[i] = t.(IAttrNameContext)
			i++
		}
	}

	return tst
}

func (s *SimpleFilterExprOneContext) AttrName(i int) IAttrNameContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttrNameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttrNameContext)
}

func (s *SimpleFilterExprOneContext) Value() IValueContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *SimpleFilterExprOneContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleFilterExprOneContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SimpleFilterExprOneContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.EnterSimpleFilterExprOne(s)
	}
}

func (s *SimpleFilterExprOneContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.ExitSimpleFilterExprOne(s)
	}
}




func (p *FilterParser) SimpleFilterExprOne() (localctx ISimpleFilterExprOneContext) {
	localctx = NewSimpleFilterExprOneContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, FilterParserRULE_simpleFilterExprOne)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		p.OpOne()
	}
	{
		p.SetState(45)
		p.Match(FilterParserT__5)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(46)
		p.AttrName()
	}
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == FilterParserT__6 {
		{
			p.SetState(47)
			p.Match(FilterParserT__6)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(48)
			p.AttrName()
		}


		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(54)
		p.Match(FilterParserT__5)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(55)
		p.Value()
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// ISimpleFilterExprMultiContext is an interface to support dynamic dispatch.
type ISimpleFilterExprMultiContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OpMulti() IOpMultiContext
	AllAttrName() []IAttrNameContext
	AttrName(i int) IAttrNameContext
	AllValue() []IValueContext
	Value(i int) IValueContext

	// IsSimpleFilterExprMultiContext differentiates from other interfaces.
	IsSimpleFilterExprMultiContext()
}

type SimpleFilterExprMultiContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleFilterExprMultiContext() *SimpleFilterExprMultiContext {
	var p = new(SimpleFilterExprMultiContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FilterParserRULE_simpleFilterExprMulti
	return p
}

func InitEmptySimpleFilterExprMultiContext(p *SimpleFilterExprMultiContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FilterParserRULE_simpleFilterExprMulti
}

func (*SimpleFilterExprMultiContext) IsSimpleFilterExprMultiContext() {}

func NewSimpleFilterExprMultiContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleFilterExprMultiContext {
	var p = new(SimpleFilterExprMultiContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterParserRULE_simpleFilterExprMulti

	return p
}

func (s *SimpleFilterExprMultiContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleFilterExprMultiContext) OpMulti() IOpMultiContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpMultiContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpMultiContext)
}

func (s *SimpleFilterExprMultiContext) AllAttrName() []IAttrNameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAttrNameContext); ok {
			len++
		}
	}

	tst := make([]IAttrNameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAttrNameContext); ok {
			tst[i] = t.(IAttrNameContext)
			i++
		}
	}

	return tst
}

func (s *SimpleFilterExprMultiContext) AttrName(i int) IAttrNameContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttrNameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttrNameContext)
}

func (s *SimpleFilterExprMultiContext) AllValue() []IValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueContext); ok {
			len++
		}
	}

	tst := make([]IValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueContext); ok {
			tst[i] = t.(IValueContext)
			i++
		}
	}

	return tst
}

func (s *SimpleFilterExprMultiContext) Value(i int) IValueContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *SimpleFilterExprMultiContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleFilterExprMultiContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SimpleFilterExprMultiContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.EnterSimpleFilterExprMulti(s)
	}
}

func (s *SimpleFilterExprMultiContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.ExitSimpleFilterExprMulti(s)
	}
}




func (p *FilterParser) SimpleFilterExprMulti() (localctx ISimpleFilterExprMultiContext) {
	localctx = NewSimpleFilterExprMultiContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FilterParserRULE_simpleFilterExprMulti)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(57)
		p.OpMulti()
	}
	{
		p.SetState(58)
		p.Match(FilterParserT__5)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(59)
		p.AttrName()
	}
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == FilterParserT__6 {
		{
			p.SetState(60)
			p.Match(FilterParserT__6)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(61)
			p.AttrName()
		}


		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(67)
		p.Match(FilterParserT__5)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(68)
		p.Value()
	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == FilterParserT__5 {
		{
			p.SetState(69)
			p.Match(FilterParserT__5)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(70)
			p.Value()
		}


		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IOpOneContext is an interface to support dynamic dispatch.
type IOpOneContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsOpOneContext differentiates from other interfaces.
	IsOpOneContext()
}

type OpOneContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpOneContext() *OpOneContext {
	var p = new(OpOneContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FilterParserRULE_opOne
	return p
}

func InitEmptyOpOneContext(p *OpOneContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FilterParserRULE_opOne
}

func (*OpOneContext) IsOpOneContext() {}

func NewOpOneContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpOneContext {
	var p = new(OpOneContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterParserRULE_opOne

	return p
}

func (s *OpOneContext) GetParser() antlr.Parser { return s.parser }
func (s *OpOneContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpOneContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *OpOneContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.EnterOpOne(s)
	}
}

func (s *OpOneContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.ExitOpOne(s)
	}
}




func (p *FilterParser) OpOne() (localctx IOpOneContext) {
	localctx = NewOpOneContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, FilterParserRULE_opOne)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(76)
		_la = p.GetTokenStream().LA(1)

		if !(((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 16128) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IOpMultiContext is an interface to support dynamic dispatch.
type IOpMultiContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsOpMultiContext differentiates from other interfaces.
	IsOpMultiContext()
}

type OpMultiContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpMultiContext() *OpMultiContext {
	var p = new(OpMultiContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FilterParserRULE_opMulti
	return p
}

func InitEmptyOpMultiContext(p *OpMultiContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FilterParserRULE_opMulti
}

func (*OpMultiContext) IsOpMultiContext() {}

func NewOpMultiContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpMultiContext {
	var p = new(OpMultiContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterParserRULE_opMulti

	return p
}

func (s *OpMultiContext) GetParser() antlr.Parser { return s.parser }
func (s *OpMultiContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpMultiContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *OpMultiContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.EnterOpMulti(s)
	}
}

func (s *OpMultiContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.ExitOpMulti(s)
	}
}




func (p *FilterParser) OpMulti() (localctx IOpMultiContext) {
	localctx = NewOpMultiContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, FilterParserRULE_opMulti)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		_la = p.GetTokenStream().LA(1)

		if !(((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 245760) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IAttrNameContext is an interface to support dynamic dispatch.
type IAttrNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext

	// IsAttrNameContext differentiates from other interfaces.
	IsAttrNameContext()
}

type AttrNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttrNameContext() *AttrNameContext {
	var p = new(AttrNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FilterParserRULE_attrName
	return p
}

func InitEmptyAttrNameContext(p *AttrNameContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FilterParserRULE_attrName
}

func (*AttrNameContext) IsAttrNameContext() {}

func NewAttrNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttrNameContext {
	var p = new(AttrNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterParserRULE_attrName

	return p
}

func (s *AttrNameContext) GetParser() antlr.Parser { return s.parser }

func (s *AttrNameContext) Name() INameContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *AttrNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttrNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AttrNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.EnterAttrName(s)
	}
}

func (s *AttrNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.ExitAttrName(s)
	}
}




func (p *FilterParser) AttrName() (localctx IAttrNameContext) {
	localctx = NewAttrNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, FilterParserRULE_attrName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		p.Name()
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	QUOTED() antlr.TerminalNode
	AllName() []INameContext
	Name(i int) INameContext

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FilterParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FilterParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) QUOTED() antlr.TerminalNode {
	return s.GetToken(FilterParserQUOTED, 0)
}

func (s *ValueContext) AllName() []INameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INameContext); ok {
			len++
		}
	}

	tst := make([]INameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INameContext); ok {
			tst[i] = t.(INameContext)
			i++
		}
	}

	return tst
}

func (s *ValueContext) Name(i int) INameContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.ExitValue(s)
	}
}




func (p *FilterParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, FilterParserRULE_value)
	var _la int

	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FilterParserQUOTED:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(82)
			p.Match(FilterParserQUOTED)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case FilterParserT__0, FilterParserT__1, FilterParserT__6, FilterParserT__7, FilterParserT__8, FilterParserT__9, FilterParserT__10, FilterParserT__11, FilterParserT__12, FilterParserT__13, FilterParserT__14, FilterParserT__15, FilterParserT__16, FilterParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 786310) != 0) {
			p.SetState(86)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case FilterParserT__0, FilterParserT__7, FilterParserT__8, FilterParserT__9, FilterParserT__10, FilterParserT__11, FilterParserT__12, FilterParserT__13, FilterParserT__14, FilterParserT__15, FilterParserT__16, FilterParserSTRING:
				{
					p.SetState(83)
					p.Name()
				}


			case FilterParserT__6:
				{
					p.SetState(84)
					p.Match(FilterParserT__6)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}


			case FilterParserT__1:
				{
					p.SetState(85)
					p.Match(FilterParserT__1)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}



			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(88)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}



	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}


errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	Keyword() IKeywordContext

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FilterParserRULE_name
	return p
}

func InitEmptyNameContext(p *NameContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FilterParserRULE_name
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) STRING() antlr.TerminalNode {
	return s.GetToken(FilterParserSTRING, 0)
}

func (s *NameContext) Keyword() IKeywordContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeywordContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeywordContext)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.ExitName(s)
	}
}




func (p *FilterParser) Name() (localctx INameContext) {
	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, FilterParserRULE_name)
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FilterParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(92)
			p.Match(FilterParserSTRING)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case FilterParserT__0, FilterParserT__7, FilterParserT__8, FilterParserT__9, FilterParserT__10, FilterParserT__11, FilterParserT__12, FilterParserT__13, FilterParserT__14, FilterParserT__15, FilterParserT__16:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(93)
			p.Keyword()
		}



	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}


errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IKeywordContext is an interface to support dynamic dispatch.
type IKeywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsKeywordContext differentiates from other interfaces.
	IsKeywordContext()
}

type KeywordContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeywordContext() *KeywordContext {
	var p = new(KeywordContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FilterParserRULE_keyword
	return p
}

func InitEmptyKeywordContext(p *KeywordContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FilterParserRULE_keyword
}

func (*KeywordContext) IsKeywordContext() {}

func NewKeywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordContext {
	var p = new(KeywordContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FilterParserRULE_keyword

	return p
}

func (s *KeywordContext) GetParser() antlr.Parser { return s.parser }
func (s *KeywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *KeywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.EnterKeyword(s)
	}
}

func (s *KeywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FilterListener); ok {
		listenerT.ExitKeyword(s)
	}
}




func (p *FilterParser) Keyword() (localctx IKeywordContext) {
	localctx = NewKeywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, FilterParserRULE_keyword)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		_la = p.GetTokenStream().LA(1)

		if !(((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 261890) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


