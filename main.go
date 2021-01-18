package main

import (
	"awesomeProject/parser"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main()  {
	input,err := antlr.NewFileStream("test.kl")
	if err != nil{
		panic(err)
	}

	lexer := parser.NewKlLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewKlParser(tokens)

	antlr.ParseTreeWalkerDefault.Walk(&KlListener{}, p.Init())
}

type KlListener struct {
	*parser.BaseKlListener
}

func (k *KlListener)EnterFun(ctx *parser.FunContext)  {
	ctx.AddChild(&parser.CallFunContext{})
	//fmt.Println(ctx.GetText())
}

func (k *KlListener)EnterCallFun(ctx *parser.CallFunContext)  {
	fmt.Println(ctx.GetText())
}

func (k *KlListener)EnterInit(ctx *parser.InitContext)  {
	//fmt.Println(ctx.GetText())
}