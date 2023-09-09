// Este especial es para las funciones de los arreglos
package prueba

import (
	"main/parser"
)

func (l *Visitor) VisitPespecial(ctx *parser.PespecialContext, entorno *Ambiente) interface{} {
	if ctx.REMOV() != nil {
		valor := l.Visit(ctx.Expr(), entorno).(Valor)
		if valor.tipo != INT {
			l.Errores = append(l.Errores, Error{mensaje: "El indice debe ser un entero"})
			return 0
		}

		temp := entorno.getVar(ctx.ID().GetText()).(Valor)
		if temp.temporal != "" {
			variabTemp := temp.val.(*Ambiente).getVar(temp.temporal).(Valor)
			if variabTemp.tipo == FLOAT {
				variabTemp.val = append(variabTemp.val.([]float64)[:valor.val.(int64)], variabTemp.val.([]float64)[valor.val.(int64)+1:]...)
				temp.val.(*Ambiente).updateVar(temp.temporal, variabTemp.val)
			} else if variabTemp.tipo == INT {
				variabTemp.val = append(variabTemp.val.([]int64)[:valor.val.(int64)], variabTemp.val.([]int64)[valor.val.(int64)+1:]...)
				temp.val.(*Ambiente).updateVar(temp.temporal, variabTemp.val)
			} else if variabTemp.tipo == STRING {
				variabTemp.val = append(variabTemp.val.([]string)[:valor.val.(int64)], variabTemp.val.([]string)[valor.val.(int64)+1:]...)
				temp.val.(*Ambiente).updateVar(temp.temporal, variabTemp.val)
			}
			return 0
		}
		if temp.tipo == FLOAT {
			temp.val = append(temp.val.([]float64)[:valor.val.(int64)], temp.val.([]float64)[valor.val.(int64)+1:]...)
			entorno.updateVar(ctx.ID().GetText(), temp.val)
		} else if temp.tipo == INT {
			temp.val = append(temp.val.([]int64)[:valor.val.(int64)], temp.val.([]int64)[valor.val.(int64)+1:]...)
			entorno.updateVar(ctx.ID().GetText(), temp.val)
		} else if temp.tipo == STRING {
			temp.val = append(temp.val.([]string)[:valor.val.(int64)], temp.val.([]string)[valor.val.(int64)+1:]...)
			entorno.updateVar(ctx.ID().GetText(), temp.val)
		}
	} else if ctx.APPE() != nil {
		valor := l.Visit(ctx.Expr(), entorno).(Valor)
		temp := entorno.getVar(ctx.ID().GetText()).(Valor)
		if temp.temporal != "" {
			varibTemp := temp.val.(*Ambiente).getVar(temp.temporal).(Valor)
			if varibTemp.tipo == FLOAT {
				if valor.tipo == FLOAT {
					varibTemp.val = append(varibTemp.val.([]float64), valor.val.(float64))
					temp.val.(*Ambiente).updateVar(temp.temporal, varibTemp.val)
					//entorno.updateVar(ctx.ID().GetText(), temp.val)
				} else {
					l.Errores = append(l.Errores, Error{mensaje: "No se puede asignar" + Str(valor.tipo) + " con " + Str(temp.tipo)})
				}
			} else if varibTemp.tipo == INT {
				if valor.tipo == INT {
					varibTemp.val = append(varibTemp.val.([]int64), valor.val.(int64))
					temp.val.(*Ambiente).updateVar(temp.temporal, varibTemp.val)
					//entorno.updateVar(ctx.ID().GetText(), temp.val)
				} else {
					l.Errores = append(l.Errores, Error{mensaje: "No se puede asignar" + Str(valor.tipo) + " con " + Str(temp.tipo)})
				}
			} else if varibTemp.tipo == STRING {
				if valor.tipo == STRING {
					varibTemp.val = append(varibTemp.val.([]string), valor.val.(string))
					temp.val.(*Ambiente).updateVar(temp.temporal, varibTemp.val)
					//entorno.updateVar(ctx.ID().GetText(), temp.val)
				} else {
					l.Errores = append(l.Errores, Error{mensaje: "No se puede asignar" + Str(valor.tipo) + " con " + Str(temp.tipo)})
				}
			}
			return 0
		}
		if temp.tipo == FLOAT {
			if valor.tipo == FLOAT {
				temp.val = append(temp.val.([]float64), valor.val.(float64))
				entorno.updateVar(ctx.ID().GetText(), temp.val)
			} else {
				l.Errores = append(l.Errores, Error{mensaje: "No se puede asignar" + Str(valor.tipo) + " con " + Str(temp.tipo)})
			}
		} else if temp.tipo == INT {
			if valor.tipo == INT {
				temp.val = append(temp.val.([]int64), valor.val.(int64))
				entorno.updateVar(ctx.ID().GetText(), temp.val)
			} else {
				l.Errores = append(l.Errores, Error{mensaje: "No se puede asignar" + Str(valor.tipo) + " con " + Str(temp.tipo)})
			}
		} else if temp.tipo == STRING {
			if valor.tipo == STRING {
				temp.val = append(temp.val.([]string), valor.val.(string))
				entorno.updateVar(ctx.ID().GetText(), temp.val)
			} else {
				l.Errores = append(l.Errores, Error{mensaje: "No se puede asignar" + Str(valor.tipo) + " con " + Str(temp.tipo)})
			}
		}

	} else if ctx.REMLST() != nil {
		temp := entorno.getVar(ctx.ID().GetText()).(Valor)
		if temp.temporal != "" {
			variabTemp := temp.val.(*Ambiente).getVar(temp.temporal).(Valor)
			if temp.tipo == INT {
				variabTemp.val = variabTemp.val.([]int64)[:len(variabTemp.val.([]int64))-1]
				temp.val.(*Ambiente).updateVar(temp.temporal, variabTemp.val)
				//entorno.updateVar(ctx.ID().GetText(), temp.val)
			} else if temp.tipo == FLOAT {
				variabTemp.val = variabTemp.val.([]float64)[:len(variabTemp.val.([]float64))-1]
				temp.val.(*Ambiente).updateVar(temp.temporal, variabTemp.val)
				//entorno.updateVar(ctx.ID().GetText(), temp.val)
			} else if temp.tipo == STRING {
				variabTemp.val = variabTemp.val.([]string)[:len(variabTemp.val.([]string))-1]
				temp.val.(*Ambiente).updateVar(temp.temporal, variabTemp.val)
				//entorno.updateVar(ctx.ID().GetText(), temp.val)
			}
			return 0
		}
		if temp.tipo == INT {
			temp.val = temp.val.([]int64)[:len(temp.val.([]int64))-1]
			entorno.updateVar(ctx.ID().GetText(), temp.val)
		} else if temp.tipo == FLOAT {
			temp.val = temp.val.([]float64)[:len(temp.val.([]float64))-1]
			entorno.updateVar(ctx.ID().GetText(), temp.val)
		} else if temp.tipo == STRING {
			temp.val = temp.val.([]string)[:len(temp.val.([]string))-1]
			entorno.updateVar(ctx.ID().GetText(), temp.val)
		}

	}
	return 0
}

func (l *Visitor) VisitEspecial(ctx *parser.EspecialContext, entorno *Ambiente) interface{} {
	if ctx.EMPT() != nil {
		temp := entorno.getVar(ctx.ID().GetText()).(Valor)
		if temp.temporal != "" {
			varibTemp := temp.val.(*Ambiente).getVar(temp.temporal).(Valor)
			if varibTemp.tipo == INT {
				return Valor{val: len(varibTemp.val.([]int64)) == 0, tipo: BOOL}
			} else if varibTemp.tipo == FLOAT {
				return Valor{val: len(varibTemp.val.([]float64)) == 0, tipo: BOOL}
			} else if varibTemp.tipo == STRING {
				return Valor{val: len(varibTemp.val.([]string)) == 0, tipo: BOOL}
			}
			return 0
		}
		if temp.tipo == INT {
			return Valor{val: len(temp.val.([]int64)) == 0, tipo: BOOL}
		} else if temp.tipo == FLOAT {
			return Valor{val: len(temp.val.([]float64)) == 0, tipo: BOOL}
		} else if temp.tipo == STRING {
			return Valor{val: len(temp.val.([]string)) == 0, tipo: BOOL}
		}
	} else if ctx.COUNT() != nil {
		temp := entorno.getVar(ctx.ID().GetText()).(Valor)
		if temp.temporal != "" {
			variabTemp := temp.val.(*Ambiente).getVar(temp.temporal).(Valor)
			if variabTemp.tipo == INT {
				return Valor{val: len(variabTemp.val.([]int64)), tipo: INT}
			} else if variabTemp.tipo == FLOAT {
				return Valor{val: len(variabTemp.val.([]float64)), tipo: INT}
			} else if variabTemp.tipo == STRING {
				return Valor{val: len(variabTemp.val.([]string)), tipo: INT}
			}
			return 0
		}
		if temp.tipo == INT {
			return Valor{val: len(temp.val.([]int64)), tipo: INT}
		} else if temp.tipo == FLOAT {
			return Valor{val: len(temp.val.([]float64)), tipo: INT}
		} else if temp.tipo == STRING {
			return Valor{val: len(temp.val.([]string)), tipo: INT}
		}
	}

	return 0
}
