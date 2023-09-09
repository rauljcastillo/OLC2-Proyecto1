package prueba

import (
	"main/parser"
	"reflect"
)

type Temporal struct {
	a int64
	b int64
}

func (l *Visitor) VisitPfor(ctx *parser.PforContext, entorno *Ambiente) interface{} {
	ambiente := NewAmbiente(entorno)
	a := l.Visit(ctx.Pifor(), entorno)
	if val, ok := a.(Temporal); ok {
		ambiente.save(INT, ctx.ID().GetText(), nil, "")
		for i := val.a; i <= val.b; i++ {
			ambiente.updateVar(ctx.ID().GetText(), int64(i))
			a := l.Visit(ctx.Block(), ambiente).(Sentences)
			if a.val == BRK {
				break
			}
		}
	} else if value, ok := a.(Valor); ok {
		if reflect.TypeOf(value.val).Kind() == reflect.Slice {
			if value.tipo == INT {
				ambiente.save(INT, ctx.ID().GetText(), nil, "")
				for _, a := range value.val.([]int64) {
					ambiente.updateVar(ctx.ID().GetText(), a)
					if a, ok := l.Visit(ctx.Block(), ambiente).(Sentences); ok {
						if a.val == BRK {
							break
						}

					}
				}
			} else if value.tipo == STRING {
				ambiente.save(STRING, ctx.ID().GetText(), nil, "")
				for _, a := range value.val.([]string) {
					ambiente.updateVar(ctx.ID().GetText(), a)
					if a, ok := l.Visit(ctx.Block(), ambiente).(Sentences); ok {
						if a.val == BRK {
							break
						}

					}
				}
			} else if value.tipo == FLOAT {
				ambiente.save(FLOAT, ctx.ID().GetText(), nil, "")
				for _, a := range value.val.([]float64) {
					ambiente.updateVar(ctx.ID().GetText(), a)
					if a, ok := l.Visit(ctx.Block(), ambiente).(Sentences); ok {
						if a.val == BRK {
							break
						}

					}
				}
			}
		} else if value.tipo == STRING {
			ambiente.save(CHAR, ctx.ID().GetText(), nil, "")
			for _, a := range value.val.(string) {
				ambiente.updateVar(ctx.ID().GetText(), string(a))
				if a, ok := l.Visit(ctx.Block(), ambiente).(Sentences); ok {
					if a.val == BRK {
						break
					}

				}
			}
		}

	}

	return 0
}

func (l *Visitor) VisitPifor(ctx *parser.PiforContext, entorno *Ambiente) interface{} {
	if ctx.POINTS() != nil {
		val1 := l.Visit(ctx.Expr(0), entorno).(Valor)
		val2 := l.Visit(ctx.Expr(1), entorno).(Valor)
		return Temporal{a: val1.val.(int64), b: val2.val.(int64)}
	} else {
		return l.Visit(ctx.Expr(0), entorno)
	}
	l.Errores = append(l.Errores, Error{mensaje: "Error en la expresion for"})
	return 0

}
