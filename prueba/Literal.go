package prueba

import (
	"main/parser"
	"strconv"
)

func (v *Visitor) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	switch {
	case ctx.CADENA() != nil:
		return Valor{val: ctx.GetText()[1 : len(ctx.GetText())-1], tipo: STRING}
	case ctx.GetLogico() != nil:
		val, _ := strconv.ParseBool(ctx.GetLogico().GetText())
		return Valor{val: val, tipo: BOOL}
	case ctx.INT() != nil:
		val, _ := strconv.ParseInt(ctx.INT().GetText(), 10, 64)
		return Valor{val: val, tipo: INT}
	case ctx.DOUBLE() != nil:
		val, _ := strconv.ParseFloat(ctx.DOUBLE().GetText(), 64)
		return Valor{val: val, tipo: FLOAT}

	}
	return Valor{}
}
