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
    "", "STRING", "WS",
  }
  staticData.RuleNames = []string{
    "filter", "filterExpr", "simpleFilterExpr", "simpleFilterExprOne", "simpleFilterExprMulti", 
    "opOne", "opMulti", "attrName", "value",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 1, 19, 81, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 
	4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 1, 0, 1, 0, 1, 
	0, 1, 1, 1, 1, 1, 1, 5, 1, 26, 8, 1, 10, 1, 12, 1, 29, 9, 1, 1, 2, 1, 2, 
	1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 39, 8, 2, 1, 3, 1, 3, 1, 3, 1, 
	3, 1, 3, 5, 3, 46, 8, 3, 10, 3, 12, 3, 49, 9, 3, 1, 3, 1, 3, 1, 3, 1, 4, 
	1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 59, 8, 4, 10, 4, 12, 4, 62, 9, 4, 1, 4, 1, 
	4, 1, 4, 1, 4, 5, 4, 68, 8, 4, 10, 4, 12, 4, 71, 9, 4, 1, 5, 1, 5, 1, 6, 
	1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 0, 0, 9, 0, 2, 4, 6, 8, 10, 12, 14, 
	16, 0, 2, 1, 0, 8, 13, 1, 0, 14, 17, 76, 0, 18, 1, 0, 0, 0, 2, 22, 1, 0, 
	0, 0, 4, 38, 1, 0, 0, 0, 6, 40, 1, 0, 0, 0, 8, 53, 1, 0, 0, 0, 10, 72, 
	1, 0, 0, 0, 12, 74, 1, 0, 0, 0, 14, 76, 1, 0, 0, 0, 16, 78, 1, 0, 0, 0, 
	18, 19, 5, 1, 0, 0, 19, 20, 5, 2, 0, 0, 20, 21, 3, 2, 1, 0, 21, 1, 1, 0, 
	0, 0, 22, 27, 3, 4, 2, 0, 23, 24, 5, 3, 0, 0, 24, 26, 3, 4, 2, 0, 25, 23, 
	1, 0, 0, 0, 26, 29, 1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 27, 28, 1, 0, 0, 0, 
	28, 3, 1, 0, 0, 0, 29, 27, 1, 0, 0, 0, 30, 31, 5, 4, 0, 0, 31, 32, 3, 6, 
	3, 0, 32, 33, 5, 5, 0, 0, 33, 39, 1, 0, 0, 0, 34, 35, 5, 4, 0, 0, 35, 36, 
	3, 8, 4, 0, 36, 37, 5, 5, 0, 0, 37, 39, 1, 0, 0, 0, 38, 30, 1, 0, 0, 0, 
	38, 34, 1, 0, 0, 0, 39, 5, 1, 0, 0, 0, 40, 41, 3, 10, 5, 0, 41, 42, 5, 
	6, 0, 0, 42, 47, 3, 14, 7, 0, 43, 44, 5, 7, 0, 0, 44, 46, 3, 14, 7, 0, 
	45, 43, 1, 0, 0, 0, 46, 49, 1, 0, 0, 0, 47, 45, 1, 0, 0, 0, 47, 48, 1, 
	0, 0, 0, 48, 50, 1, 0, 0, 0, 49, 47, 1, 0, 0, 0, 50, 51, 5, 6, 0, 0, 51, 
	52, 3, 16, 8, 0, 52, 7, 1, 0, 0, 0, 53, 54, 3, 12, 6, 0, 54, 55, 5, 6, 
	0, 0, 55, 60, 3, 14, 7, 0, 56, 57, 5, 7, 0, 0, 57, 59, 3, 14, 7, 0, 58, 
	56, 1, 0, 0, 0, 59, 62, 1, 0, 0, 0, 60, 58, 1, 0, 0, 0, 60, 61, 1, 0, 0, 
	0, 61, 63, 1, 0, 0, 0, 62, 60, 1, 0, 0, 0, 63, 64, 5, 6, 0, 0, 64, 69, 
	3, 16, 8, 0, 65, 66, 5, 6, 0, 0, 66, 68, 3, 16, 8, 0, 67, 65, 1, 0, 0, 
	0, 68, 71, 1, 0, 0, 0, 69, 67, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 9, 1, 
	0, 0, 0, 71, 69, 1, 0, 0, 0, 72, 73, 7, 0, 0, 0, 73, 11, 1, 0, 0, 0, 74, 
	75, 7, 1, 0, 0, 75, 13, 1, 0, 0, 0, 76, 77, 5, 18, 0, 0, 77, 15, 1, 0, 
	0, 0, 78, 79, 5, 18, 0, 0, 79, 17, 1, 0, 0, 0, 5, 27, 38, 47, 60, 69,
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
	FilterParserSTRING = 18
	FilterParserWS = 19
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
		p.SetState(18)
		p.Match(FilterParserT__0)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(19)
		p.Match(FilterParserT__1)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(20)
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
		p.SetState(22)
		p.SimpleFilterExpr()
	}
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == FilterParserT__2 {
		{
			p.SetState(23)
			p.Match(FilterParserT__2)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(24)
			p.SimpleFilterExpr()
		}


		p.SetState(29)
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
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(30)
			p.Match(FilterParserT__3)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(31)
			p.SimpleFilterExprOne()
		}
		{
			p.SetState(32)
			p.Match(FilterParserT__4)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
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
			p.SimpleFilterExprMulti()
		}
		{
			p.SetState(36)
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
		p.SetState(40)
		p.OpOne()
	}
	{
		p.SetState(41)
		p.Match(FilterParserT__5)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(42)
		p.AttrName()
	}
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == FilterParserT__6 {
		{
			p.SetState(43)
			p.Match(FilterParserT__6)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(44)
			p.AttrName()
		}


		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(50)
		p.Match(FilterParserT__5)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(51)
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
		p.SetState(53)
		p.OpMulti()
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
		p.AttrName()
	}
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == FilterParserT__6 {
		{
			p.SetState(56)
			p.Match(FilterParserT__6)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(57)
			p.AttrName()
		}


		p.SetState(62)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(63)
		p.Match(FilterParserT__5)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(64)
		p.Value()
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == FilterParserT__5 {
		{
			p.SetState(65)
			p.Match(FilterParserT__5)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(66)
			p.Value()
		}


		p.SetState(71)
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
		p.SetState(72)
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
		p.SetState(74)
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
	STRING() antlr.TerminalNode

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

func (s *AttrNameContext) STRING() antlr.TerminalNode {
	return s.GetToken(FilterParserSTRING, 0)
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
		p.SetState(76)
		p.Match(FilterParserSTRING)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
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


// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode

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

func (s *ValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(FilterParserSTRING, 0)
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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Match(FilterParserSTRING)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
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


