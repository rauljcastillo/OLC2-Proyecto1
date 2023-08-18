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
		tipe := l.Visit(ctx.Tipo(), entorno).(int)     //Tipo de la variable
		expres := l.Visit(ctx.Expr(), entorno).(Valor) //Tipo del valor se guarda
		if expres.val != nil {
			if tipe == expres.tipo {
				entorno.save(tipe, ctx.ID().GetText(), expres.val)
			} else if tipe == FLOAT && expres.tipo == INT {
				entorno.save(tipe, ctx.ID().GetText(), float64(expres.val.(int64)))

			} else {
				l.Errores = append(l.Errores, Error{mensaje: "No se puede asignar " + Str(expres.tipo) + " con " + Str(tipe)})
			}
		}

	} else if ctx.Expr() == nil && ctx.Tipo() != nil {
		tipo := l.Visit(ctx.Tipo(), entorno).(int)
		entorno.save(tipo, ctx.ID().GetText(), nil)

	} else if ctx.Tipo() == nil {
		exp := l.Visit(ctx.Expr(), entorno).(Valor)
		if exp.val != nil {
			entorno.save(exp.tipo, ctx.ID().GetText(), exp.val)
		}
	}

	return 0
}
