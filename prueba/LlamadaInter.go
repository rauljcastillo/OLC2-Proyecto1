package prueba

import "main/parser"

func (l *Visitor) VisitLlamada(ctx *parser.LlamadaContext, entorno *Ambiente) interface{} {
	return l.Visit(ctx.Pllamada(), entorno)
}
