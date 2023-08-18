package prueba

import (
	"main/parser"
)

func (l *Visitor) VisitAsign(ctx *parser.AsignContext, entorno *Ambiente) interface{} {
	temp := l.Visit(ctx.Expr(), entorno).(Valor)
	variab := entorno.getVar(ctx.ID().GetText()).(Valor)

	if temp.val == nil {
		return Valor{}
	}
	if variab.val == nil {
		l.Errores = append(l.Errores, Error{mensaje: "No existe la variable " + ctx.ID().GetText()})
		return Valor{}
	}
	operador := ctx.GetOp().GetText()
	if operador == "=" {
		if temp.tipo == variab.tipo {
			entorno.updateVar(ctx.ID().GetText(), temp.val)
			return 0
		} else if temp.tipo == INT && variab.tipo == FLOAT {
			entorno.updateVar(ctx.ID().GetText(), float64(temp.val.(int64)))
			return 0
		}
		l.Errores = append(l.Errores, Error{mensaje: "No se puede operar " + Str(variab.tipo) + " con " + Str(temp.tipo)})
	} else if operador == "+=" {
		if temp.tipo == STRING && variab.tipo == STRING {
			entorno.updateVar(ctx.ID().GetText(), variab.val.(string)+temp.val.(string))
			return 0
		} else if temp.tipo == variab.tipo && temp.tipo == INT {
			entorno.updateVar(ctx.ID().GetText(), variab.val.(int64)+temp.val.(int64))
			return 0
		} else if temp.tipo == variab.tipo && temp.tipo == FLOAT {
			entorno.updateVar(ctx.ID().GetText(), variab.val.(float64)+temp.val.(float64))
			return 0
		} else if temp.tipo == INT && variab.tipo == FLOAT {
			entorno.updateVar(ctx.ID().GetText(), variab.val.(float64)+float64(temp.val.(int64)))
			return 0
		}
		l.Errores = append(l.Errores, Error{mensaje: "No se puede operar " + Str(variab.tipo) + " con " + Str(temp.tipo)})
	} else if operador == "-=" {
		if temp.tipo == variab.tipo && temp.tipo == INT {
			entorno.updateVar(ctx.ID().GetText(), variab.val.(int64)+temp.val.(int64))
			return 0
		} else if temp.tipo == variab.tipo && temp.tipo == FLOAT {
			entorno.updateVar(ctx.ID().GetText(), variab.val.(float64)+temp.val.(float64))
			return 0
		} else if temp.tipo == INT && variab.tipo == FLOAT {
			entorno.updateVar(ctx.ID().GetText(), variab.val.(float64)+float64(temp.val.(int64)))
			return 0
		}
		l.Errores = append(l.Errores, Error{mensaje: "No se puede operar " + Str(variab.tipo) + " con " + Str(temp.tipo)})
	}

	return 0
}
