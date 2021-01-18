// Code generated from Kl.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Kl

import "github.com/antlr/antlr4/runtime/Go/antlr"

// KlListener is a complete listener for a parse tree produced by KlParser.
type KlListener interface {
	antlr.ParseTreeListener

	// EnterStat is called when entering the stat production.
	EnterStat(c *StatContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// EnterFun is called when entering the fun production.
	EnterFun(c *FunContext)

	// EnterArgs is called when entering the args production.
	EnterArgs(c *ArgsContext)

	// EnterCallFun is called when entering the callFun production.
	EnterCallFun(c *CallFunContext)

	// EnterInit is called when entering the init production.
	EnterInit(c *InitContext)

	// ExitStat is called when exiting the stat production.
	ExitStat(c *StatContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)

	// ExitFun is called when exiting the fun production.
	ExitFun(c *FunContext)

	// ExitArgs is called when exiting the args production.
	ExitArgs(c *ArgsContext)

	// ExitCallFun is called when exiting the callFun production.
	ExitCallFun(c *CallFunContext)

	// ExitInit is called when exiting the init production.
	ExitInit(c *InitContext)
}
