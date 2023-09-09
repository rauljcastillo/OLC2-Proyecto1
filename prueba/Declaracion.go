package prueba

import (
	"main/parser"
)

func (l *Visitor) VisitTipo(ctx *parser.TipoContext) interface{} {
	if ctx.RSTRING() != nil {
		return 0
	} else if ctx.RFLOAT() != nil {
		return 1
	} else if ctx.RINT() != nil {
		return 2
	} else if ctx.RBOOL() != nil {
		return 3
	}
	return 0
}

func (l *Visitor) VisitDeclara(ctx *parser.DeclaraContext, entorno *Ambiente) interface{} {

	if ctx.Tipo() != nil && ctx.Expr() != nil {
		tipe := Num(ctx.Tipo().GetText()) //Tipo de la variable
		if exp, ok := l.Visit(ctx.Expr(), entorno).(Valor); ok {
			if tipe == exp.tipo {
				entorno.save(tipe, ctx.ID().GetText(), exp.val, "")
			} else if tipe == FLOAT && exp.tipo == INT {
				entorno.save(tipe, ctx.ID().GetText(), float64(exp.val.(int64)), "")

			} else {
				l.Errores = append(l.Errores, Error{mensaje: "No se puede asignar " + Str(exp.tipo) + " con " + Str(tipe)})
			}
		}

	} else if ctx.Expr() == nil && ctx.Tipo() != nil {
		tipo := Num(ctx.Tipo().GetText())
		entorno.save(tipo, ctx.ID().GetText(), nil, "")

	} else if ctx.Tipo() == nil {
		if exp, ok := l.Visit(ctx.Expr(), entorno).(Valor); ok {
			entorno.save(exp.tipo, ctx.ID().GetText(), exp.val, "")
		}
	}

	return 0
}
