package prueba

import (
	"main/parser"
	"reflect"
)

func (l *Visitor) VisitAccesoA(ctx *parser.AccesoAContext, entorno *Ambiente) interface{} {
	variab := entorno.getVar(ctx.ID().GetText()).(Valor)
	if reflect.TypeOf(variab.val).Kind() != reflect.Slice {
		l.Errores = append(l.Errores, Error{mensaje: "Se esperaba un array"})
		return 0
	}
	exp := l.Visit(ctx.Expr(), entorno).(Valor)
	if variab.tipo == STRING {
		return Valor{val: variab.val.([]string)[exp.val.(int64)], tipo: STRING}
	} else if variab.tipo == INT {
		return Valor{val: variab.val.([]int64)[exp.val.(int64)], tipo: INT}
	} else if variab.tipo == FLOAT {
		return Valor{val: variab.val.([]float64)[exp.val.(int64)], tipo: FLOAT}
	}
	return 0
}
