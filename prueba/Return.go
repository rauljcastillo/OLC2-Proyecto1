package prueba

import (
	"main/parser"
)

func (l *Visitor) VisitPreturn(ctx *parser.PreturnContext, entorno *Ambiente) interface{} {
	if ctx.Expr() != nil {
		exp := l.Visit(ctx.Expr(), entorno).(Valor)
		return Sentences{val: RTN, expres: exp}
	}
	return 0
}
