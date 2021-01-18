// Code generated from Kl.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Kl

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseKlListener is a complete listener for a parse tree produced by KlParser.
type BaseKlListener struct{}

var _ KlListener = &BaseKlListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseKlListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseKlListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseKlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseKlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStat is called when production stat is entered.
func (s *BaseKlListener) EnterStat(ctx *StatContext) {}

// ExitStat is called when production stat is exited.
func (s *BaseKlListener) ExitStat(ctx *StatContext) {}

// EnterValue is called when production value is entered.
func (s *BaseKlListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseKlListener) ExitValue(ctx *ValueContext) {}

// EnterParam is called when production param is entered.
func (s *BaseKlListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BaseKlListener) ExitParam(ctx *ParamContext) {}

// EnterFun is called when production fun is entered.
func (s *BaseKlListener) EnterFun(ctx *FunContext) {}

// ExitFun is called when production fun is exited.
func (s *BaseKlListener) ExitFun(ctx *FunContext) {}

// EnterArgs is called when production args is entered.
func (s *BaseKlListener) EnterArgs(ctx *ArgsContext) {}

// ExitArgs is called when production args is exited.
func (s *BaseKlListener) ExitArgs(ctx *ArgsContext) {}

// EnterCallFun is called when production callFun is entered.
func (s *BaseKlListener) EnterCallFun(ctx *CallFunContext) {}

// ExitCallFun is called when production callFun is exited.
func (s *BaseKlListener) ExitCallFun(ctx *CallFunContext) {}

// EnterInit is called when production init is entered.
func (s *BaseKlListener) EnterInit(ctx *InitContext) {}

// ExitInit is called when production init is exited.
func (s *BaseKlListener) ExitInit(ctx *InitContext) {}
