package prueba

import "main/parser"

func (l *Visitor) VisitPwhile(ctx *parser.PwhileContext, entorno *Ambiente) interface{} {
	nuevoAmb := NewAmbiente(entorno)
	exp := l.Visit(ctx.Expr(), nuevoAmb).(Valor)
	for exp.val.(bool) {
		a := l.Visit(ctx.Block(), nuevoAmb).(Sentences)
		if a.val == BRK {
			break
		}
		exp = l.Visit(ctx.Expr(), nuevoAmb).(Valor)
	}

	return 0
}
