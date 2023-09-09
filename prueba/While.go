package prueba

import "main/parser"

func (l *Visitor) VisitPwhile(ctx *parser.PwhileContext, entorno *Ambiente) interface{} {
	nuevoAmb := NewAmbiente(entorno)
	exp := l.Visit(ctx.Expr(), nuevoAmb).(Valor)
	for exp.val.(bool) {
		nuevoAmb1 := NewAmbiente(nuevoAmb)
		if value, ok := l.Visit(ctx.Block(), nuevoAmb1).(Sentences); ok {
			if value.val == BRK {
				break
			}
		}
		exp = l.Visit(ctx.Expr(), nuevoAmb1).(Valor)
	}

	return 0
}
