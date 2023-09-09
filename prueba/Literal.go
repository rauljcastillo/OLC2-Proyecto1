package prueba

import (
	"main/parser"
	"strconv"
	"strings"
)

func (v *Visitor) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	switch {
	case ctx.CADENA() != nil:
		temp := ctx.GetText()[1 : len(ctx.GetText())-1]

		temp = strings.ReplaceAll(temp, `\n`, "\n")
		temp = strings.ReplaceAll(temp, `\"`, "\"")
		temp = strings.ReplaceAll(temp, `\t`, "\t")
		temp = strings.ReplaceAll(temp, `\r`, "\r")
		temp = strings.ReplaceAll(temp, `\\`, "\\")

		return Valor{val: temp, tipo: STRING}
	case ctx.GetLogico() != nil:
		val, _ := strconv.ParseBool(ctx.GetLogico().GetText())
		return Valor{val: val, tipo: BOOL}
	case ctx.INT() != nil:
		val, _ := strconv.ParseInt(ctx.INT().GetText(), 10, 64)
		return Valor{val: val, tipo: INT}
	case ctx.DOUBLE() != nil:
		val, _ := strconv.ParseFloat(ctx.DOUBLE().GetText(), 64)
		return Valor{val: val, tipo: FLOAT}
		/*
			case ctx.NIL() != nil:
				return Valor{val: nil}
		*/
	}
	return Valor{}
}
