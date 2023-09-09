package prueba

import (
	"fmt"
	"main/parser"
	"strconv"
)

func (l *Visitor) VisitEmbebida(ctx *parser.EmbebidaContext, entorno *Ambiente) interface{} {
	if ctx.RINT() != nil {
		temp := l.Visit(ctx.Expr(), entorno).(Valor)
		if temp.tipo == STRING {
			value, _ := strconv.ParseInt(temp.val.(string), 10, 64)
			return Valor{val: value, tipo: INT}
		} else if temp.tipo == FLOAT {
			return Valor{val: int(temp.val.(float64)), tipo: INT}
		}

		l.Errores = append(l.Errores, Error{mensaje: "No se puede convertir el tipo" + Str(temp.tipo)})
		return 0
	} else if ctx.RFLOAT() != nil {
		temp := l.Visit(ctx.Expr(), entorno).(Valor)
		if temp.tipo == STRING {
			value, _ := strconv.ParseFloat(temp.val.(string), 64)
			return Valor{val: value, tipo: FLOAT}
		} else if temp.tipo == INT {
			return Valor{val: float64(temp.val.(int64)), tipo: FLOAT}
		}
	} else if ctx.RSTRING() != nil {
		temp := l.Visit(ctx.Expr(), entorno).(Valor)
		return Valor{val: fmt.Sprint(temp.val), tipo: STRING}
	}

	return 0
}
