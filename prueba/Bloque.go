package prueba

import "main/parser"

func (l *Visitor) VisitStmt(ctx *parser.StmtContext, entorno *Ambiente) interface{} {
	nuevo := NewAmbiente(entorno)
	return l.Visit(ctx.Block(), nuevo)
}
