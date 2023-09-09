package prueba

import (
	"main/parser"
	"strconv"
	"strings"
)

func (l *Visitor) VisitPdeclarArray(ctx *parser.PdeclarArrayContext, entorno *Ambiente) interface{} {
	a := l.Visit(ctx.Pdefinition(), entorno).([]string)
	if ctx.Tipo() != nil {
		tipe := l.Visit(ctx.Tipo(), entorno)
		if tipe == INT {
			var nuevo []int64
			for _, elem := range a {
				el, e := strconv.ParseInt(elem, 10, 64)
				if e == nil {
					nuevo = append(nuevo, el)
				} else {
					l.Errores = append(l.Errores, Error{mensaje: "El arreglo es de tipo INT"})
					return 0
				}

			}
			entorno.save(INT, ctx.ID().GetText(), nuevo, "")
		} else if tipe == FLOAT {
			var nuevo []float64
			for _, elem := range a {
				el, e := strconv.ParseFloat(elem, 64)
				if e == nil {
					nuevo = append(nuevo, el)
				} else {
					l.Errores = append(l.Errores, Error{mensaje: "El arreglo es de tipo FLOAT"})
					return 0
				}

			}
			entorno.save(FLOAT, ctx.ID().GetText(), nuevo, "")
		} else if tipe == STRING {
			entorno.save(STRING, ctx.ID().GetText(), a, "")
		}

		return 0
	}
	if _, err := strconv.ParseInt(a[0], 10, 64); err == nil {
		var nuevo []int64
		for _, elem := range a {
			el, e := strconv.ParseInt(elem, 10, 64)
			if e == nil {
				nuevo = append(nuevo, el)
			} else {
				l.Errores = append(l.Errores, Error{mensaje: "El arreglo es de tipo INT"})
				return 0
			}

		}
		entorno.save(INT, ctx.ID().GetText(), nuevo, "")
	} else if _, err := strconv.ParseFloat(a[0], 64); err == nil {
		var nuevo []float64
		for _, elem := range a {
			el, e := strconv.ParseFloat(elem, 64)
			if e == nil {
				nuevo = append(nuevo, el)
			} else {
				l.Errores = append(l.Errores, Error{mensaje: "El arreglo es de tipo FLOAT"})
				return 0
			}

		}
		entorno.save(INT, ctx.ID().GetText(), nuevo, "")
	}

	return 0
}

func (l *Visitor) VisitPdefinition(ctx *parser.PdefinitionContext, entorno *Ambiente) interface{} {

	if ctx.GetChildCount() > 1 {
		var array []string

		temp := strings.Split(ctx.Expr(0).GetText(), ",")
		for _, elem := range temp {
			array = append(array, strings.ReplaceAll(elem, "\"", ""))
		}

		return array
	}

	return 0
}
