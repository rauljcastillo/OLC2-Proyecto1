package prueba

import "main/parser"

func (l *Visitor) VisitAcceso(ctx *parser.AccesoContext, entorno *Ambiente) interface{} {
	temp := entorno.getVar(ctx.ID().GetText()).(Valor)
	if temp.val == nil {
		l.Errores = append(l.Errores, Error{mensaje: "No se encontro la varialbe " + ctx.ID().GetText()})
		return Valor{}
	}
	if temp.temporal != "" {
		return temp.val.(*Ambiente).getVar(temp.temporal)
	}
	return temp

}
