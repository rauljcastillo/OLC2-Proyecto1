package prueba

import (
	"main/parser"

	"github.com/antlr4-go/antlr/v4"
)

type CASE struct {
	val  interface{}
	bloc antlr.ParseTree
}

func (l *Visitor) VisitPswitch(ctx *parser.PswitchContext, entorno *Ambiente) interface{} {
	expres := l.Visit(ctx.Expr(), entorno).(Valor)
	for i := 0; ctx.Ccase(i) != nil; i++ {
		if a := l.Visit(ctx.Ccase(i), entorno).(CASE); a.val == expres.val {
			newEntorno := NewAmbiente(entorno)
			return l.Visit(a.bloc, newEntorno)
		}
	}
	return l.Visit(ctx.Pdefault(), entorno)

}

func (l *Visitor) VisitCcase(ctx *parser.CcaseContext, entorno *Ambiente) interface{} {
	a := l.Visit(ctx.Expr(), entorno).(Valor)
	return CASE{val: a.val, bloc: ctx.Block()}
}

func (l *Visitor) VisitPdefault(ctx *parser.PdefaultContext, entorno *Ambiente) interface{} {
	amb := NewAmbiente(entorno)
	return l.Visit(ctx.Block(), amb)
}
