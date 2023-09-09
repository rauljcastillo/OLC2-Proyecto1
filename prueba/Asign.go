package prueba

import (
	"main/parser"
)

func (l *Visitor) VisitAsign(ctx *parser.AsignContext, entorno *Ambiente) interface{} {
	temp := l.Visit(ctx.Expr(), entorno).(Valor)         //Retorna el valor a asignar a la variable
	variab := entorno.getVar(ctx.ID().GetText()).(Valor) //Retorna la variable guardada

	if temp.val == nil {
		return Valor{}
	}
	if variab.val == nil {
		l.Errores = append(l.Errores, Error{mensaje: "No existe la variable " + ctx.ID().GetText()})
		return Valor{}
	}
	operador := ctx.GetOp().GetText()
	if operador == "=" {
		if variab.temporal != "" {
			if temp.tipo == variab.tipo {
				variab.val.(*Ambiente).updateVar(variab.temporal, temp.val)
			} else if temp.tipo == INT && variab.tipo == FLOAT {
				variab.val.(*Ambiente).updateVar(variab.temporal, float64(temp.val.(int64)))
			}
			return 0
		}
		if temp.tipo == variab.tipo {
			entorno.updateVar(ctx.ID().GetText(), temp.val)
			return 0
		} else if temp.tipo == INT && variab.tipo == FLOAT {
			entorno.updateVar(ctx.ID().GetText(), float64(temp.val.(int64)))
			return 0
		}
		l.Errores = append(l.Errores, Error{mensaje: "No se puede operar " + Str(variab.tipo) + " con " + Str(temp.tipo)})
	} else if operador == "+=" {
		if variab.temporal != "" {
			if temp.tipo == STRING && variab.tipo == STRING {
				a := variab.val.(*Ambiente).getVar(variab.temporal).(Valor)
				variab.val.(*Ambiente).updateVar(variab.temporal, a.val.(string)+temp.val.(string))
				return 0
			} else if temp.tipo == variab.tipo && temp.tipo == INT {
				a := variab.val.(*Ambiente).getVar(variab.temporal).(Valor)
				variab.val.(*Ambiente).updateVar(variab.temporal, a.val.(int64)+temp.val.(int64))
				return 0
			} else if temp.tipo == variab.tipo && temp.tipo == FLOAT {
				a := variab.val.(*Ambiente).getVar(variab.temporal).(Valor)
				variab.val.(*Ambiente).updateVar(variab.temporal, a.val.(float64)+temp.val.(float64))
				return 0
			} else if temp.tipo == INT && variab.tipo == FLOAT {
				a := variab.val.(*Ambiente).getVar(variab.temporal).(Valor)
				variab.val.(*Ambiente).updateVar(variab.temporal, a.val.(float64)+float64(temp.val.(int64)))
				return 0
			}
		}
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
		if variab.temporal != "" {
			if temp.tipo == variab.tipo && temp.tipo == INT {
				a := variab.val.(*Ambiente).getVar(variab.temporal).(Valor)
				variab.val.(*Ambiente).updateVar(variab.temporal, a.val.(int64)+temp.val.(int64))
				return 0
			} else if temp.tipo == variab.tipo && temp.tipo == FLOAT {
				a := variab.val.(*Ambiente).getVar(variab.temporal).(Valor)
				variab.val.(*Ambiente).updateVar(variab.temporal, a.val.(float64)+temp.val.(float64))
				return 0
			} else if temp.tipo == INT && variab.tipo == FLOAT {
				a := variab.val.(*Ambiente).getVar(variab.temporal).(Valor)
				variab.val.(*Ambiente).updateVar(variab.temporal, a.val.(float64)+float64(temp.val.(int64)))
				return 0
			}
		}
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
