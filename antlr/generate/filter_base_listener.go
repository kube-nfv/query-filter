// Code generated from grammar/Filter.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Filter

import "github.com/antlr4-go/antlr/v4"

// BaseFilterListener is a complete listener for a parse tree produced by FilterParser.
type BaseFilterListener struct{}

var _ FilterListener = &BaseFilterListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFilterListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFilterListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFilterListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFilterListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFilter is called when production filter is entered.
func (s *BaseFilterListener) EnterFilter(ctx *FilterContext) {}

// ExitFilter is called when production filter is exited.
func (s *BaseFilterListener) ExitFilter(ctx *FilterContext) {}

// EnterFilterExpr is called when production filterExpr is entered.
func (s *BaseFilterListener) EnterFilterExpr(ctx *FilterExprContext) {}

// ExitFilterExpr is called when production filterExpr is exited.
func (s *BaseFilterListener) ExitFilterExpr(ctx *FilterExprContext) {}

// EnterSimpleFilterExpr is called when production simpleFilterExpr is entered.
func (s *BaseFilterListener) EnterSimpleFilterExpr(ctx *SimpleFilterExprContext) {}

// ExitSimpleFilterExpr is called when production simpleFilterExpr is exited.
func (s *BaseFilterListener) ExitSimpleFilterExpr(ctx *SimpleFilterExprContext) {}

// EnterSimpleFilterExprOne is called when production simpleFilterExprOne is entered.
func (s *BaseFilterListener) EnterSimpleFilterExprOne(ctx *SimpleFilterExprOneContext) {}

// ExitSimpleFilterExprOne is called when production simpleFilterExprOne is exited.
func (s *BaseFilterListener) ExitSimpleFilterExprOne(ctx *SimpleFilterExprOneContext) {}

// EnterSimpleFilterExprMulti is called when production simpleFilterExprMulti is entered.
func (s *BaseFilterListener) EnterSimpleFilterExprMulti(ctx *SimpleFilterExprMultiContext) {}

// ExitSimpleFilterExprMulti is called when production simpleFilterExprMulti is exited.
func (s *BaseFilterListener) ExitSimpleFilterExprMulti(ctx *SimpleFilterExprMultiContext) {}

// EnterOpOne is called when production opOne is entered.
func (s *BaseFilterListener) EnterOpOne(ctx *OpOneContext) {}

// ExitOpOne is called when production opOne is exited.
func (s *BaseFilterListener) ExitOpOne(ctx *OpOneContext) {}

// EnterOpMulti is called when production opMulti is entered.
func (s *BaseFilterListener) EnterOpMulti(ctx *OpMultiContext) {}

// ExitOpMulti is called when production opMulti is exited.
func (s *BaseFilterListener) ExitOpMulti(ctx *OpMultiContext) {}

// EnterAttrName is called when production attrName is entered.
func (s *BaseFilterListener) EnterAttrName(ctx *AttrNameContext) {}

// ExitAttrName is called when production attrName is exited.
func (s *BaseFilterListener) ExitAttrName(ctx *AttrNameContext) {}

// EnterValue is called when production value is entered.
func (s *BaseFilterListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseFilterListener) ExitValue(ctx *ValueContext) {}
