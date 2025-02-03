// Code generated from grammar/Filter.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Filter

import "github.com/antlr4-go/antlr/v4"


// FilterListener is a complete listener for a parse tree produced by FilterParser.
type FilterListener interface {
	antlr.ParseTreeListener

	// EnterFilter is called when entering the filter production.
	EnterFilter(c *FilterContext)

	// EnterFilterExpr is called when entering the filterExpr production.
	EnterFilterExpr(c *FilterExprContext)

	// EnterSimpleFilterExpr is called when entering the simpleFilterExpr production.
	EnterSimpleFilterExpr(c *SimpleFilterExprContext)

	// EnterSimpleFilterExprOne is called when entering the simpleFilterExprOne production.
	EnterSimpleFilterExprOne(c *SimpleFilterExprOneContext)

	// EnterSimpleFilterExprMulti is called when entering the simpleFilterExprMulti production.
	EnterSimpleFilterExprMulti(c *SimpleFilterExprMultiContext)

	// EnterOpOne is called when entering the opOne production.
	EnterOpOne(c *OpOneContext)

	// EnterOpMulti is called when entering the opMulti production.
	EnterOpMulti(c *OpMultiContext)

	// EnterAttrName is called when entering the attrName production.
	EnterAttrName(c *AttrNameContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// ExitFilter is called when exiting the filter production.
	ExitFilter(c *FilterContext)

	// ExitFilterExpr is called when exiting the filterExpr production.
	ExitFilterExpr(c *FilterExprContext)

	// ExitSimpleFilterExpr is called when exiting the simpleFilterExpr production.
	ExitSimpleFilterExpr(c *SimpleFilterExprContext)

	// ExitSimpleFilterExprOne is called when exiting the simpleFilterExprOne production.
	ExitSimpleFilterExprOne(c *SimpleFilterExprOneContext)

	// ExitSimpleFilterExprMulti is called when exiting the simpleFilterExprMulti production.
	ExitSimpleFilterExprMulti(c *SimpleFilterExprMultiContext)

	// ExitOpOne is called when exiting the opOne production.
	ExitOpOne(c *OpOneContext)

	// ExitOpMulti is called when exiting the opMulti production.
	ExitOpMulti(c *OpMultiContext)

	// ExitAttrName is called when exiting the attrName production.
	ExitAttrName(c *AttrNameContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)
}
