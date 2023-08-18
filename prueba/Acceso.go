package prueba

import "main/parser"

func (l *Visitor) VisitAcceso(ctx *parser.AccesoContext, entorno *Ambiente) interface{} {
	temp := entorno.getVar(ctx.ID().GetText()).(Valor)
	if temp.val != nil {
		return temp
	}
	l.Errores = append(l.Errores, Error{mensaje: "No se encontro la varialbe " + ctx.ID().GetText()})
	return Valor{}
}
