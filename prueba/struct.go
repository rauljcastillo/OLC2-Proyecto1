package prueba

import (
	"main/parser"

	"github.com/antlr4-go/antlr/v4"
)

type Visitor struct {
	antlr.ParseTreeVisitor
	Errores []Error
	Cadena  string
}

func NewVisitor() *Visitor {
	return &Visitor{ParseTreeVisitor: &parser.BaseGramarVisitor{}, Errores: make([]Error, 0)}
}

func (l *Visitor) Visit(tree antlr.ParseTree, entorno *Ambiente) interface{} {
	switch val := tree.(type) {
	case *parser.SContext:
		return l.VisitS(val, entorno)
	case *parser.BlockContext:
		return l.VisitBlock(val, entorno)
	case *parser.ProductionContext:
		return l.VisitProduction(val, entorno)
	case *parser.PrinContext:
		return l.VisitPrint(val, entorno)
	case *parser.LiteralContext:
		return l.VisitLiteral(val)
	case *parser.OpContext:
		return l.VisitOp(val, entorno)
	case *parser.ParenContext:
		return l.VisitParen(val, entorno)
	case *parser.DeclaraContext:
		return l.VisitDeclara(val, entorno)
	case *parser.TipoContext:
		return l.VisitTipo(val)
	case *parser.AccesoContext:
		return l.VisitAcceso(val, entorno)
	case *parser.PifContext:
		return l.VisitPif(val, entorno)
	case *parser.StmtContext:
		return l.VisitStmt(val, entorno)
	case *parser.AsignContext:
		return l.VisitAsign(val, entorno)
	case *parser.PelseContext:
		return l.VisitPelse(val, entorno)
	case *parser.PswitchContext:
		return l.VisitPswitch(val, entorno)
	case *parser.CcaseContext:
		return l.VisitCcase(val, entorno)
	case *parser.PdefaultContext:
		return l.VisitPdefault(val, entorno)
	case *parser.PwhileContext:
		return l.VisitPwhile(val, entorno)
	case *parser.PforContext:
		return l.VisitPfor(val, entorno)
	case *parser.PiforContext:
		return l.VisitPifor(val, entorno)
	case *parser.PguardContext:
		return l.VisitPguard(val, entorno)
	case *parser.PdeclarArrayContext:
		return l.VisitPdeclarArray(val, entorno)
	case *parser.PdefinitionContext:
		return l.VisitPdefinition(val, entorno)
	case *parser.PespecialContext:
		return l.VisitPespecial(val, entorno)
	case *parser.EspecialContext:
		return l.VisitEspecial(val, entorno)
	case *parser.PfuncionContext:
		return l.VisitPfuncion(val, entorno)
	case *parser.PllamadaContext:
		return l.VisitPllamada(val, entorno)
	case *parser.PparamsContext:
		return l.VisitPparams(val, entorno)
	case *parser.PargumentsContext:
		return l.VisitParguments(val, entorno)
	case *parser.PargumContext:
		return l.VisitPargum(val, entorno)
	case *parser.PparametContext:
		return l.VisitPparamet(val, entorno)
	case *parser.PasignAContext:
		return l.VisitPasignA(val, entorno)
	case *parser.AccesoAContext:
		return l.VisitAccesoA(val, entorno)
	case *parser.LlamadaContext:
		return l.VisitLlamada(val, entorno)
	case *parser.PreturnContext:
		return l.VisitPreturn(val, entorno)
	}

	return Valor{val: int64(0), tipo: INT}
}

func (l *Visitor) VisitS(ctx *parser.SContext, entorno *Ambiente) interface{} {
	return l.Visit(ctx.Block(), entorno)
}

func (l *Visitor) VisitBlock(ctx *parser.BlockContext, entorno *Ambiente) interface{} {
	for i := 0; ctx.Production(i) != nil; i++ {
		if ctx.Production(i).CONT() != nil {
			return Sentences{val: CONTINUE}
		} else if ctx.Production(i).BRK() != nil {
			return Sentences{val: BRK}
		}
		if a, ok := l.Visit(ctx.Production(i), entorno).(Sentences); ok {
			return a
		}

	}
	return 0
}

func (l *Visitor) VisitProduction(ctx *parser.ProductionContext, entorno *Ambiente) interface{} {
	if ctx.Prin() != nil {
		return l.Visit(ctx.Prin(), entorno)
	} else if ctx.Declara() != nil {
		return l.Visit(ctx.Declara(), entorno)
	} else if ctx.Pif() != nil {
		return l.Visit(ctx.Pif(), entorno)
	} else if ctx.Asign() != nil {
		return l.Visit(ctx.Asign(), entorno)
	} else if ctx.Pswitch() != nil {
		return l.Visit(ctx.Pswitch(), entorno)
	} else if ctx.Pwhile() != nil {
		return l.Visit(ctx.Pwhile(), entorno)
	} else if ctx.Pfor() != nil {
		return l.Visit(ctx.Pfor(), entorno)
	} else if ctx.Pguard() != nil {
		return l.Visit(ctx.Pguard(), entorno)
	} else if ctx.PdeclarArray() != nil {
		return l.Visit(ctx.PdeclarArray(), entorno)
	} else if ctx.Pespecial() != nil {
		return l.Visit(ctx.Pespecial(), entorno)
	} else if ctx.Pfuncion() != nil {
		return l.Visit(ctx.Pfuncion(), entorno)
	} else if ctx.Pllamada() != nil {
		return l.Visit(ctx.Pllamada(), entorno)
	} else if ctx.PasignA() != nil {
		return l.Visit(ctx.PasignA(), entorno)
	} else if ctx.Preturn() != nil {
		return l.Visit(ctx.Preturn(), entorno)
	}
	return 0
}
