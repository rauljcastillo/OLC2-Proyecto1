package prueba

import "main/parser"

func (l *Visitor) VisitPguard(ctx *parser.PguardContext, entorno *Ambiente) interface{} {
	valor := l.Visit(ctx.Expr(), entorno).(Valor)
	if !valor.val.(bool) {
		l.Visit(ctx.Block(), entorno)
		return Sentences{val: Num(ctx.GetOpera().GetText())}
	}

	return 0
}
