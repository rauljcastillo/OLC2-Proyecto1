package prueba

import "main/parser"

func (l *Visitor) VisitPif(ctx *parser.PifContext, entorno *Ambiente) interface{} {
	condition := l.Visit(ctx.Expr(), entorno).(Valor)
	if condition.val == nil {
		return Valor{}
	}

	if condition.val.(bool) {
		return l.Visit(ctx.Stmt(), entorno)
	} else if ctx.Pelse() != nil {
		return l.Visit(ctx.Pelse(), entorno)
	}

	return 0
}

func (l *Visitor) VisitPelse(ctx *parser.PelseContext, entorno *Ambiente) interface{} {
	if ctx.Pif() != nil {
		return l.Visit(ctx.Pif(), entorno)
	} else {
		return l.Visit(ctx.Stmt(), entorno)
	}
}
