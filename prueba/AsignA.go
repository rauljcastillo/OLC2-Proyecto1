package prueba

import (
	"main/parser"
	"reflect"
)

func (l *Visitor) VisitPasignA(ctx *parser.PasignAContext, entorno *Ambiente) interface{} {
	variab := entorno.getVar(ctx.ID().GetText()).(Valor)
	if variab.temporal != "" {
		exp1 := l.Visit(ctx.Expr(0), entorno).(Valor)
		exp2 := l.Visit(ctx.Expr(1), entorno).(Valor)
		//Llama a la variable del entorno
		variabTemp := variab.val.(*Ambiente).getVar(variab.temporal).(Valor)

		if variabTemp.tipo == STRING {
			if exp2.tipo == STRING {
				variabTemp.val.([]string)[exp1.val.(int64)] = exp2.val.(string)
				variab.val.(*Ambiente).updateVar(variab.temporal, variabTemp.val)
				return 0
			}
			l.Errores = append(l.Errores, Error{mensaje: "No se puede asignar " + Str(exp2.tipo) + " a String"})
			return 0

		} else if variabTemp.tipo == INT {
			if exp2.tipo == INT {
				variabTemp.val.([]int64)[exp1.val.(int64)] = exp2.val.(int64)
				variab.val.(*Ambiente).updateVar(variab.temporal, variabTemp.val)
				return 0
			}
			l.Errores = append(l.Errores, Error{mensaje: "No se puede asignar " + Str(exp2.tipo) + " a Int"})
			return 0
		} else if variabTemp.tipo == FLOAT {
			if exp2.tipo == FLOAT {
				variabTemp.val.([]float64)[exp1.val.(int64)] = exp2.val.(float64)
				variab.val.(*Ambiente).updateVar(variab.temporal, variabTemp.val)
				return 0
			}
			l.Errores = append(l.Errores, Error{mensaje: "No se puede asignar " + Str(exp2.tipo) + " a Float"})
			return 0
		}
	}
	if reflect.TypeOf(variab.val).Kind() != reflect.Slice {
		l.Errores = append(l.Errores, Error{mensaje: "Se esperaba un Array"})
		return 0
	}
	exp1 := l.Visit(ctx.Expr(0), entorno).(Valor)
	exp2 := l.Visit(ctx.Expr(1), entorno).(Valor)
	if variab.tipo == STRING {
		if exp2.tipo == STRING {
			variab.val.([]string)[exp1.val.(int64)] = exp2.val.(string)
			return 0
		}
		l.Errores = append(l.Errores, Error{mensaje: "No se puede asignar " + Str(exp2.tipo) + " a String"})

	} else if variab.tipo == INT {
		if exp2.tipo == INT {
			variab.val.([]int64)[exp1.val.(int64)] = exp2.val.(int64)
			return 0
		}
		l.Errores = append(l.Errores, Error{mensaje: "No se puede asignar " + Str(exp2.tipo) + " a Int"})
	} else if variab.tipo == FLOAT {
		if exp2.tipo == FLOAT {
			variab.val.([]float64)[exp1.val.(int64)] = exp2.val.(float64)
			return 0
		}
		l.Errores = append(l.Errores, Error{mensaje: "No se puede asignar " + Str(exp2.tipo) + " a Float"})
	}

	return 0
}
