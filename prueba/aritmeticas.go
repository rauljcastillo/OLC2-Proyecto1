package prueba

import (
	"main/parser"
)

func (v *Visitor) VisitOp(ctx *parser.OpContext, entorno *Ambiente) interface{} {
	l := v.Visit(ctx.GetLeft(), entorno).(Valor)
	r := v.Visit(ctx.GetRight(), entorno).(Valor)
	op := ctx.GetOp().GetText()
	dom := TablaT(l.tipo, r.tipo)
	if l.val != nil && r.val != nil {
		switch op {
		case "+":
			if dom == FLOAT {
				if _, ok := l.val.(float64); !ok {
					l.val = float64(l.val.(int64))
				}
				if _, ok := r.val.(float64); !ok {
					r.val = float64(r.val.(int64))
				}
				temp := l.val.(float64) + r.val.(float64)
				return Valor{val: temp, tipo: dom}
			} else if dom == INT {
				temp := l.val.(int64) + r.val.(int64)
				return Valor{val: temp, tipo: dom}
			} else if dom == STRING {
				temp := l.val.(string) + r.val.(string)
				return Valor{val: temp, tipo: dom}
			}
			v.Errores = append(v.Errores, Error{mensaje: "No se puede sumar " + Str(l.tipo) + " con " + Str(r.tipo)})
		case "+=":
			if dom == FLOAT {
				if _, ok := l.val.(float64); !ok {
					l.val = float64(l.val.(int64))
				}
				if _, ok := r.val.(float64); !ok {
					r.val = float64(r.val.(int64))
				}
				temp := l.val.(float64) + r.val.(float64)
				return Valor{val: temp, tipo: dom}
			} else if dom == INT {
				temp := l.val.(int64) + r.val.(int64)
				return Valor{val: temp, tipo: dom}
			} else if dom == STRING {
				temp := l.val.(string) + r.val.(string)
				return Valor{val: temp, tipo: dom}
			}
			v.Errores = append(v.Errores, Error{mensaje: "No se puede sumar " + Str(l.tipo) + " con " + Str(r.tipo)})
		case "-":
			if dom == FLOAT {
				if _, ok := l.val.(float64); !ok {
					l.val = float64(l.val.(int64))
				}
				if _, ok := r.val.(float64); !ok {
					r.val = float64(r.val.(int64))
				}
				temp := l.val.(float64) - r.val.(float64)
				return Valor{val: temp, tipo: dom}
			} else if dom == INT {
				temp := l.val.(int64) - r.val.(int64)
				return Valor{val: temp, tipo: dom}
			}
			v.Errores = append(v.Errores, Error{mensaje: "No se puede restar " + Str(l.tipo) + " con " + Str(r.tipo)})
		case "*":
			if dom == FLOAT {
				if _, ok := l.val.(float64); !ok {
					l.val = float64(l.val.(int64))
				}
				if _, ok := r.val.(float64); !ok {
					r.val = float64(r.val.(int64))
				}
				temp := l.val.(float64) * r.val.(float64)
				return Valor{val: temp, tipo: FLOAT}
			} else if dom == INT {
				temp := l.val.(int64) * r.val.(int64)
				return Valor{val: temp, tipo: dom}
			}
			v.Errores = append(v.Errores, Error{mensaje: "No se puede multiplicar " + Str(l.tipo) + " con " + Str(r.tipo)})
		case "/":
			if dom == FLOAT {
				if _, ok := l.val.(float64); !ok {
					l.val = float64(l.val.(int64))
				}
				if _, ok := r.val.(float64); !ok {
					r.val = float64(r.val.(int64))
				}
				if r.val.(float64) != 0 {
					temp := l.val.(float64) * r.val.(float64)
					return Valor{val: temp, tipo: FLOAT}
				}
			} else if dom == INT {
				temp := l.val.(int64) / r.val.(int64)
				return Valor{val: temp, tipo: dom}
			}
			v.Errores = append(v.Errores, Error{mensaje: "No se puede dividir " + Str(l.tipo) + " con " + Str(r.tipo)})
			//
		case ">":
			if dom != 4 {
				if _, ok := l.val.(float64); !ok {
					l.val = float64(l.val.(int64))
				}
				if _, ok := r.val.(float64); !ok {
					r.val = float64(r.val.(int64))
				}
				return Valor{val: l.val.(float64) > r.val.(float64), tipo: BOOL}
			}
			v.Errores = append(v.Errores, Error{mensaje: "No se puede comparar " + Str(l.tipo) + " con " + Str(r.tipo)})
		case "<":
			if dom != 4 {
				if _, ok := l.val.(float64); !ok {
					l.val = float64(l.val.(int64))
				}
				if _, ok := r.val.(float64); !ok {
					r.val = float64(r.val.(int64))
				}

				return Valor{val: l.val.(float64) < r.val.(float64), tipo: BOOL}
			}

			v.Errores = append(v.Errores, Error{mensaje: "No se puede comparar " + Str(l.tipo) + " con " + Str(r.tipo)})
		case "<=":
			if dom != 4 {
				if _, ok := l.val.(float64); !ok {
					l.val = float64(l.val.(int64))
				}
				if _, ok := r.val.(float64); !ok {
					r.val = float64(r.val.(int64))
				}

				return Valor{val: l.val.(float64) <= r.val.(float64), tipo: BOOL}
			}

			v.Errores = append(v.Errores, Error{mensaje: "No se puede comparar " + Str(l.tipo) + " con " + Str(r.tipo)})

		case ">=":
			if dom != 4 {
				if _, ok := l.val.(float64); !ok {
					l.val = float64(l.val.(int64))
				}
				if _, ok := r.val.(float64); !ok {
					r.val = float64(r.val.(int64))
				}

				return Valor{val: l.val.(float64) >= r.val.(float64), tipo: BOOL}
			}

			v.Errores = append(v.Errores, Error{mensaje: "No se puede comparar " + Str(l.tipo) + " con " + Str(r.tipo)})
		case "==":
			if l.tipo == r.tipo {
				return Valor{val: l.val == r.val, tipo: BOOL}
			}
			v.Errores = append(v.Errores, Error{mensaje: "No se puede comparar" + Str(l.tipo) + " con " + Str(r.tipo)})
		case "!":
			if r.tipo == BOOL {
				return Valor{val: !r.val.(bool), tipo: BOOL}
			}
			v.Errores = append(v.Errores, Error{mensaje: "No se puede negar " + Str(r.tipo)})
		case "%":
			if dom == INT {
				return Valor{val: l.val.(int64) % r.val.(int64), tipo: INT}
			}
			v.Errores = append(v.Errores, Error{mensaje: "No se puede operar " + Str(l.tipo) + " con " + Str(r.tipo)})
		case "!=":
			if l.tipo == r.tipo {
				return Valor{val: l.val != r.val, tipo: BOOL}
			}
			v.Errores = append(v.Errores, Error{mensaje: "No se puede operar " + Str(l.tipo) + " con " + Str(r.tipo)})
		case "&&":
			if l.tipo == BOOL && r.tipo == BOOL {
				return Valor{val: l.val.(bool) && r.val.(bool), tipo: BOOL}
			}
			v.Errores = append(v.Errores, Error{mensaje: "No se puede operar " + Str(l.tipo) + " con " + Str(r.tipo)})
		case "||":
			if l.tipo == BOOL && r.tipo == BOOL {
				return Valor{val: l.val.(bool) || r.val.(bool), tipo: BOOL}
			}
			v.Errores = append(v.Errores, Error{mensaje: "No se puede operar " + Str(l.tipo) + " con " + Str(r.tipo)})
		}
	}

	return Valor{}
}

func (v *Visitor) VisitParen(ctx *parser.ParenContext, entorno *Ambiente) interface{} {
	return v.Visit(ctx.Expr(), entorno)
}
