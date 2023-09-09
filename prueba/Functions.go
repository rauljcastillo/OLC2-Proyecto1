package prueba

import (
	"main/parser"
)

type Param struct {
	extern  string
	intern  string
	tipo    int
	isArray bool
	inout   bool
}

func (l *Visitor) VisitPfuncion(ctx *parser.PfuncionContext, entorno *Ambiente) interface{} {
	if ctx.Tipo() != nil {
		tipo := Num(ctx.Tipo().GetText())
		if ctx.Pparams() != nil {
			a := l.Visit(ctx.Pparams(), entorno).([]Param)
			if !entorno.saveFunctions(tipo, ctx.ID().GetText(), a, ctx.Block()) {
				l.Errores = append(l.Errores, Error{mensaje: "Ya existe una funcion con el mismo nombre"})
			}

		} else {
			if !entorno.saveFunctions(tipo, ctx.ID().GetText(), nil, ctx.Block()) {
				l.Errores = append(l.Errores, Error{mensaje: "Ya existe una funcion con el mismo nombre"})
			}
		}

	} else {
		if ctx.Pparams() != nil {
			a := l.Visit(ctx.Pparams(), entorno).([]Param)
			if !entorno.saveFunctions(VOID, ctx.ID().GetText(), a, ctx.Block()) {
				l.Errores = append(l.Errores, Error{mensaje: "Ya existe una funcion con el mismo nombre"})
			}
		} else {
			if !entorno.saveFunctions(VOID, ctx.ID().GetText(), nil, ctx.Block()) {
				l.Errores = append(l.Errores, Error{mensaje: "Ya existe una funcion con el mismo nombre"})
			}
		}
	}

	return 0
}

func (l *Visitor) VisitPparams(ctx *parser.PparamsContext, entorno *Ambiente) interface{} {
	var temp []Param
	for _, elem := range ctx.AllPparamet() {
		a := l.Visit(elem, entorno).(Param)
		temp = append(temp, a)
	}
	return temp
	/*
		var temp []Param
		ctx.Pparamet(0).ABR()
		tempo := ctx.AllID()
		tipos := ctx.AllTipo()
		cont := 0
		for i := 0; i < len(tempo); i += 2 {
			temp = append(temp, Param{extern: tempo[i].GetText(), intern: tempo[i+1].GetText(), tipo: Num(tipos[cont].GetStart().GetText())})
			cont++
		}
		return temp
	*/
}

func (l *Visitor) VisitPparamet(ctx *parser.PparametContext, entorno *Ambiente) interface{} {
	if ctx.ABR() != nil {
		if ctx.GetPid().GetText() != "_" {
			if ctx.INOUT() != nil {
				return Param{extern: ctx.GetPid().GetText(), intern: ctx.ID(0).GetText(), tipo: Num(ctx.Tipo().GetText()), isArray: true, inout: true}
			} else {
				return Param{extern: ctx.GetPid().GetText(), intern: ctx.ID(0).GetText(), tipo: Num(ctx.Tipo().GetText()), isArray: true, inout: false}
			}

		}
		if ctx.INOUT() != nil {
			return Param{extern: ctx.GetPid().GetText(), intern: ctx.ID(0).GetText(), tipo: Num(ctx.Tipo().GetText()), isArray: true, inout: true}
		} else {
			return Param{extern: ctx.GetPid().GetText(), intern: ctx.ID(0).GetText(), tipo: Num(ctx.Tipo().GetText()), isArray: true, inout: false}
		}

	}
	if ctx.GetPid().GetText() != "_" {
		if ctx.INOUT() != nil {
			return Param{extern: ctx.GetPid().GetText(), intern: ctx.ID(1).GetText(), tipo: Num(ctx.Tipo().GetText()), isArray: false, inout: true}
		} else {
			return Param{extern: ctx.GetPid().GetText(), intern: ctx.ID(1).GetText(), tipo: Num(ctx.Tipo().GetText()), isArray: false, inout: false}
		}

	}
	if ctx.INOUT() != nil {
		return Param{extern: "_", intern: ctx.ID(0).GetText(), tipo: Num(ctx.Tipo().GetText()), isArray: false, inout: true}
	} else {
		return Param{extern: "_", intern: ctx.ID(0).GetText(), tipo: Num(ctx.Tipo().GetText()), isArray: false, inout: false}
	}

}
