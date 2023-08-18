package prueba

import (
	"main/parser"
	"strconv"
)

type Temporal struct {
	a int
	b int
}

func (l *Visitor) VisitPfor(ctx *parser.PforContext, entorno *Ambiente) interface{} {
	ambiente := NewAmbiente(entorno)
	a := l.Visit(ctx.Pifor(), entorno)
	cadena := ""
	if val, ok := a.(Temporal); ok {
		ambiente.save(INT, ctx.ID().GetText(), nil)
		for i := val.a; i <= val.b; i++ {
			ambiente.updateVar(ctx.ID().GetText(), int64(i))
			cadena += l.Visit(ctx.Block(), ambiente).(string)
		}
	}

	return cadena

}

func (l *Visitor) VisitPifor(ctx *parser.PiforContext, entorno *Ambiente) interface{} {
	if ctx.INT(0) != nil {
		val1, _ := strconv.Atoi(ctx.INT(0).GetText())
		val2, _ := strconv.Atoi(ctx.INT(1).GetText())
		return Temporal{a: val1, b: val2}
	}

	return 0
}
