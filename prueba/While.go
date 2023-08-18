package prueba

import "main/parser"

func (l *Visitor) VisitPwhile(ctx *parser.PwhileContext, entorno *Ambiente) interface{} {
	exp := l.Visit(ctx.Expr(), entorno).(Valor)
	valor := ""
	for exp.val.(bool) {
		valor += l.Visit(ctx.Stmt(), entorno).(string)
		exp = l.Visit(ctx.Expr(), entorno).(Valor)
	}

	return valor
}
