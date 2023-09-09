package prueba

import (
	"fmt"
	"main/parser"
	"reflect"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type Argument struct {
	name   string
	valor  any
	tipo   int
	namet  string
	entorn *Ambiente
}

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
		tempor := l.Visit(ctx.Parguments(), entorno).([]Argument)
		newAmbiente := NewAmbiente(entorno)
		for i := 0; i < len(tempor); i++ {
			if temp.(Funcion).Params.([]Param)[i].extern == "_" && tempor[i].name != "" {
				l.Errores = append(l.Errores, Error{mensaje: "Nombre externo " + tempor[i].name + " no valido"})
				return 0
			}
			if temp.(Funcion).Params.([]Param)[i].extern != "_" && tempor[i].name == "" {
				l.Errores = append(l.Errores, Error{mensaje: "Se esperaba nombre externo " + temp.(Funcion).Params.([]Param)[i].extern})
				return 0
			}

			if temp.(Funcion).Params.([]Param)[i].isArray && reflect.TypeOf(tempor[i].valor).Kind() != reflect.Slice {
				l.Errores = append(l.Errores, Error{mensaje: "El parametro debe ser de tipo Array"})
				return 0
			}
			if temp.(Funcion).Params.([]Param)[i].inout && tempor[i].namet == "" {
				l.Errores = append(l.Errores, Error{mensaje: "Se esperaba un apuntador"})
				return 0
			}
			if tempor[i].namet != "" && !temp.(Funcion).Params.([]Param)[i].inout {
				l.Errores = append(l.Errores, Error{mensaje: "No se esperaba un apuntador"})
				return 0
			}
			if temp.(Funcion).Params.([]Param)[i].tipo != tempor[i].tipo {
				l.Errores = append(l.Errores, Error{mensaje: "No se puede asignar " + Str(tempor[i].tipo) + " con " + Str(temp.(Funcion).Params.([]Param)[i].tipo)})
				return 0
			}
			if !temp.(Funcion).Params.([]Param)[i].isArray && reflect.TypeOf(tempor[i].valor).Kind() == reflect.Slice {
				l.Errores = append(l.Errores, Error{mensaje: "No se esperaba un Array"})
				return 0
			}
			if temp.(Funcion).Params.([]Param)[i].inout {
				newAmbiente.save(tempor[i].tipo, temp.(Funcion).Params.([]Param)[i].intern, tempor[i].entorn, tempor[i].namet)
			} else if temp.(Funcion).Params.([]Param)[i].extern == "_" {
				newAmbiente.save(tempor[i].tipo, temp.(Funcion).Params.([]Param)[i].intern, tempor[i].valor, "")
			} else if temp.(Funcion).Params.([]Param)[i].extern == tempor[i].name {
				newAmbiente.save(tempor[i].tipo, temp.(Funcion).Params.([]Param)[i].intern, tempor[i].valor, "")
			}
		}
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

func (l *Visitor) VisitParguments(ctx *parser.PargumentsContext, entorno *Ambiente) interface{} {
	var listArg []Argument
	if len(ctx.AllPargum()) == 1 && ctx.Pargum(0).GetChild(0).GetChildCount() > 1 {
		//num1:1,Temp
		//Compara si es 3 por ejemplo si es num1:1,Temp num1,':',1
		if ctx.Pargum(0).GetChildCount() == 3 {
			fmt.Println(ctx.Pargum(0).GetChild(1))
			for i := 0; i < ctx.Pargum(0).GetChild(2).GetChildCount(); i += 2 {
				if ctx.Pargum(0).GetChild(i) == ctx.Pargum(0).ID() {
					a := l.Visit(ctx.Pargum(0).GetChild(2).GetChild(i).(antlr.RuleContext), entorno).(Valor)
					listArg = append(listArg, Argument{name: ctx.Pargum(0).ID().GetText(), valor: a.val, tipo: a.tipo})

				} else {
					a := l.Visit(ctx.Pargum(0).GetChild(2).GetChild(i).(antlr.RuleContext), entorno).(Valor)
					listArg = append(listArg, Argument{name: "", valor: a.val, tipo: a.tipo})
				}
			}
			//Entraria aqui si se envian 2 o más parametros que son Accesos, etc
		} else if ctx.Pargum(0).GetChild(0).GetChild(1).(antlr.TerminalNode).GetText() == "," {
			if ctx.Pargum(0).GetChild(0).GetChild(0).GetChild(0).GetChildCount() > 0 {
				a := strings.Split(ctx.Pargum(0).GetText(), ",")
				tromp := l.Visit(ctx.Pargum(0).Expr().GetChild(1).(*antlr.TerminalNodeImpl), entorno)
				fmt.Println(tromp)
				for _, elem := range a {
					if val, err := strconv.ParseInt(elem, 10, 64); err == nil {
						listArg = append(listArg, Argument{name: "", valor: val, tipo: INT})
						continue
					} else if val, err := strconv.ParseFloat(elem, 64); err == nil {
						listArg = append(listArg, Argument{name: "", valor: val, tipo: FLOAT})
						continue
					} else if strings.Contains(elem, "\"") {
						listArg = append(listArg, Argument{name: "", valor: strings.ReplaceAll(elem, "\"", ""), tipo: STRING})
						continue
					} else if strings.Contains(elem, "&") {
						cad := strings.ReplaceAll(elem, "&", "")
						temp := entorno.getVar(cad).(Valor)
						listArg = append(listArg, Argument{name: "", valor: temp.val, tipo: temp.tipo, namet: cad, entorn: entorno})
						continue
					}
					if strings.Contains(elem, "+") {
						arr := strings.Split(elem, "+")
						variab := entorno.getVar(arr[0]).(Valor)
						if val, err := strconv.ParseInt(arr[1], 10, 64); err == nil {
							listArg = append(listArg, Argument{name: "", valor: variab.val.(int64) + val, tipo: variab.tipo})
						} else if val, err := strconv.ParseFloat(arr[0], 64); err == nil {
							listArg = append(listArg, Argument{name: "", valor: variab.val.(float64) + val, tipo: variab.tipo})
						} else if variab.tipo == STRING {
							listArg = append(listArg, Argument{name: "", valor: variab.val.(string) + arr[1], tipo: variab.tipo})
						}
					} else if strings.Contains(elem, "-") {
						arr := strings.Split(elem, "-")
						variab := entorno.getVar(arr[0]).(Valor)
						if val, err := strconv.ParseInt(arr[1], 10, 64); err == nil {
							listArg = append(listArg, Argument{name: "", valor: variab.val.(int64) - val, tipo: variab.tipo})
						} else if val, err := strconv.ParseFloat(arr[0], 64); err == nil {
							listArg = append(listArg, Argument{name: "", valor: variab.val.(float64) - val, tipo: variab.tipo})
						}
					} else if strings.Contains(elem, "*") {
						arr := strings.Split(elem, "*")
						variab := entorno.getVar(arr[0]).(Valor)
						if val, err := strconv.ParseInt(arr[1], 10, 64); err == nil {
							listArg = append(listArg, Argument{name: "", valor: variab.val.(int64) * val, tipo: variab.tipo})
						} else if val, err := strconv.ParseFloat(arr[0], 64); err == nil {
							listArg = append(listArg, Argument{name: "", valor: variab.val.(float64) * val, tipo: variab.tipo})
						}
					} else if strings.Contains(elem, "/") {
						arr := strings.Split(elem, "/")
						variab := entorno.getVar(arr[0]).(Valor)
						if val, err := strconv.ParseInt(arr[1], 10, 64); err == nil {
							listArg = append(listArg, Argument{name: "", valor: variab.val.(int64) / val, tipo: INT})
						} else if val, err := strconv.ParseFloat(arr[0], 64); err == nil {
							listArg = append(listArg, Argument{name: "", valor: variab.val.(float64) / val, tipo: FLOAT})
						}
					} else {
						variab := entorno.getVar(elem).(Valor)
						listArg = append(listArg, Argument{name: "", valor: variab.val, tipo: variab.tipo})
					}

				}
			} else if ctx.Pargum(0).GetChild(0).GetChild(0).GetChild(0).GetChildCount() == 0 {
				for i := 0; i < ctx.Pargum(0).GetChild(0).GetChildCount(); i += 2 {
					a := l.Visit(ctx.Pargum(0).GetChild(0).GetChild(i).(antlr.RuleContext), entorno).(Valor)
					if ctx.Pargum(0).PUNTE() != nil {
						listArg = append(listArg, Argument{name: "", valor: a.val, tipo: a.tipo, namet: ctx.Pargum(0).Expr().GetText(), entorn: entorno})
						continue
					}
					listArg = append(listArg, Argument{name: "", valor: a.val, tipo: a.tipo})
				}
			}
			return listArg
			//Si solo se envia 1 parametro entra a este else
		} else {
			for _, elem := range ctx.AllPargum() {
				a := l.Visit(elem, entorno).(Argument)
				listArg = append(listArg, a)
			}
		}
		return listArg
	}
	//Si el AllArgum es mayor que 1 entonces entra a este for
	for _, elem := range ctx.AllPargum() {
		a := l.Visit(elem, entorno).(Argument)
		listArg = append(listArg, a)
	}
	return listArg
}

func (l *Visitor) VisitPargum(ctx *parser.PargumContext, entorno *Ambiente) interface{} {
	if ctx.ID() != nil {
		exp := l.Visit(ctx.Expr(), entorno).(Valor)
		return Argument{name: ctx.ID().GetText(), valor: exp.val, tipo: exp.tipo}
	}
	exp := l.Visit(ctx.Expr(), entorno).(Valor)
	if ctx.PUNTE() != nil {
		return Argument{name: "", valor: exp.val, tipo: exp.tipo, namet: ctx.Expr().GetText(), entorn: entorno}
	}
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
