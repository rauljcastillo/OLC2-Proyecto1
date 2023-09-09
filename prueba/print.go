package prueba

import (
	"fmt"
	"main/parser"
)

var cadena string = ""

func (l *Visitor) VisitPrint(ctx *parser.PrinContext, entorno *Ambiente) interface{} {

	if ctx.Cexpr() != nil {
		l.Visit(ctx.Cexpr(), entorno)
		l.Cadena += cadena + "\n"
		cadena = ""
		return 0
	}

	a := l.Visit(ctx.Expr(), entorno).(Valor)
	if a.val != nil {
		l.Cadena += fmt.Sprint(a.val) + "\n"
	}
	return 0
}

func (l *Visitor) VisitCexpr(ctx *parser.CexprContext, entorno *Ambiente) interface{} {
	valor := l.Visit(ctx.Expr(), entorno).(Valor)
	if valor.val != nil {
		cadena += fmt.Sprint(valor.val)
	}
	l.Visit(ctx.Cexpr(), entorno)
	return 0
}
