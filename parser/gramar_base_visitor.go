// Code generated from Gramar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Gramar

import "github.com/antlr4-go/antlr/v4"

type BaseGramarVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseGramarVisitor) VisitS(ctx *SContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramarVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramarVisitor) VisitProduction(ctx *ProductionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramarVisitor) VisitPfor(ctx *PforContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramarVisitor) VisitPifor(ctx *PiforContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramarVisitor) VisitPwhile(ctx *PwhileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramarVisitor) VisitPswitch(ctx *PswitchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramarVisitor) VisitCcase(ctx *CcaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramarVisitor) VisitPdefault(ctx *PdefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramarVisitor) VisitStmt(ctx *StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramarVisitor) VisitPif(ctx *PifContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramarVisitor) VisitPelse(ctx *PelseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramarVisitor) VisitPrin(ctx *PrinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramarVisitor) VisitDeclara(ctx *DeclaraContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramarVisitor) VisitAsign(ctx *AsignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramarVisitor) VisitTipo(ctx *TipoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramarVisitor) VisitOp(ctx *OpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramarVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramarVisitor) VisitAcceso(ctx *AccesoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramarVisitor) VisitParen(ctx *ParenContext) interface{} {
	return v.VisitChildren(ctx)
}
