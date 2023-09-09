package prueba

import (
	"main/parser"
	"reflect"
)

type Argument struct {
	name   string
	valor  any
	tipo   int
	namet  string
	entorn *Ambiente
}

type ListaArgum struct {
	pinche []Argument
}

var Tempo []ListaArgum

func (l *Visitor) VisitPllamada(ctx *parser.PllamadaContext, entorno *Ambiente) interface{} {
	temp := entorno.getFuction(ctx.ID().GetText())
	if temp == "" {
		l.Errores = append(l.Errores, Error{mensaje: "No existe la funcion " + ctx.ID().GetText()})
		return 0
	}
	if temp.(Funcion).Params == nil {
		newAmbiente := NewAmbiente(entorno)
		return l.Visit(temp.(Funcion).Block.(parser.IBlockContext), newAmbiente)
	} else {
		Tempo = append(Tempo, ListaArgum{})
		listArg := l.Visit(ctx.Argumento(), entorno).([]Argument)
		newAmbiente := NewAmbiente(entorno)
		for i := 0; i < len(listArg); i++ {
			if temp.(Funcion).Params.([]Param)[i].extern == "_" && listArg[i].name != "" {
				l.Errores = append(l.Errores, Error{mensaje: "Nombre externo " + listArg[i].name + " no valido"})
				return 0
			}
			if temp.(Funcion).Params.([]Param)[i].extern != "_" && listArg[i].name == "" {
				l.Errores = append(l.Errores, Error{mensaje: "Se esperaba nombre externo " + temp.(Funcion).Params.([]Param)[i].extern})
				return 0
			}
			if temp.(Funcion).Params.([]Param)[i].inout && listArg[i].namet == "" {
				l.Errores = append(l.Errores, Error{mensaje: "Se esperaba un apuntador"})
				return 0
			}
			if listArg[i].namet != "" && !temp.(Funcion).Params.([]Param)[i].inout {
				l.Errores = append(l.Errores, Error{mensaje: "No se esperaba un apuntador"})
				return 0
			}
			if temp.(Funcion).Params.([]Param)[i].tipo != listArg[i].tipo {
				l.Errores = append(l.Errores, Error{mensaje: "No se puede asignar " + Str(listArg[i].tipo) + " con " + Str(temp.(Funcion).Params.([]Param)[i].tipo)})
				return 0
			}
			if !temp.(Funcion).Params.([]Param)[i].isArray && reflect.TypeOf(listArg[i].valor).Kind() == reflect.Slice {
				l.Errores = append(l.Errores, Error{mensaje: "No se esperaba un Array"})
				return 0
			}
			if temp.(Funcion).Params.([]Param)[i].inout {
				newAmbiente.save(listArg[i].tipo, temp.(Funcion).Params.([]Param)[i].intern, listArg[i].entorn, listArg[i].namet)
			} else if temp.(Funcion).Params.([]Param)[i].extern == "_" {
				newAmbiente.save(listArg[i].tipo, temp.(Funcion).Params.([]Param)[i].intern, listArg[i].valor, "")
			} else if temp.(Funcion).Params.([]Param)[i].extern == listArg[i].name {
				newAmbiente.save(listArg[i].tipo, temp.(Funcion).Params.([]Param)[i].intern, listArg[i].valor, "")
			}

		}
		Tempo = Tempo[:len(Tempo)-1]
		if value, ok := l.Visit(temp.(Funcion).Block.(parser.IBlockContext), newAmbiente).(Sentences); ok {
			if value.val == RTN && value.expres != nil {
				if temp.(Funcion).tipo == VOID {
					l.Errores = append(l.Errores, Error{mensaje: "La funcion es de tipo void"})
					return 0
				}
				if temp.(Funcion).tipo == value.expres.(Valor).tipo {
					return value.expres
				}

			} else if value.val == RTN && value.expres == nil {
				return 0
			}
		}

	}

	return 0
}

func (v *Visitor) VisitArgumento(ctx *parser.ArgumentoContext, entorno *Ambiente) interface{} {
	l := v.Visit(ctx.TipoArg(), entorno).(Argument)
	Tempo[len(Tempo)-1].pinche = append(Tempo[len(Tempo)-1].pinche, l)
	v.Visit(ctx.Argumento(), entorno)
	return Tempo[len(Tempo)-1].pinche

}

func (l *Visitor) VisitTipoArg(ctx *parser.TipoArgContext, entorno *Ambiente) interface{} {
	if ctx.ID() != nil {
		exp := l.Visit(ctx.GetExp2(), entorno).(Valor)
		return Argument{name: ctx.ID().GetText(), valor: exp.val, tipo: exp.tipo}
	}

	if ctx.PUNTE() != nil {
		temp := entorno.getVar(ctx.GetExp1().GetText()).(Valor)
		if amb, ok := temp.val.(*Ambiente); ok {
			return Argument{name: "", tipo: temp.tipo, namet: temp.temporal, entorn: amb}
		}
		exp := l.Visit(ctx.GetExp1(), entorno).(Valor)
		return Argument{name: "", valor: exp.val, tipo: exp.tipo, namet: ctx.Expr().GetText(), entorn: entorno}
	}

	exp := l.Visit(ctx.GetExp3(), entorno).(Valor)
	if reflect.TypeOf(exp.val).Kind() == reflect.Slice {
		if exp.tipo == STRING {
			a := make([]string, len(exp.val.([]string)))
			copy(a, exp.val.([]string))
			return Argument{name: "", valor: a, tipo: exp.tipo}
		} else if exp.tipo == INT {
			a := make([]int64, len(exp.val.([]int64)))
			copy(a, exp.val.([]int64))
			return Argument{name: "", valor: a, tipo: exp.tipo}
		} else if exp.tipo == FLOAT {
			a := make([]float64, len(exp.val.([]float64)))
			copy(a, exp.val.([]float64))
			return Argument{name: "", valor: a, tipo: exp.tipo}
		}
	}
	return Argument{name: "", valor: exp.val, tipo: exp.tipo}
}
