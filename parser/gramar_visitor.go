// Code generated from Gramar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Gramar

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by GramarParser.
type GramarVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by GramarParser#s.
	VisitS(ctx *SContext) interface{}

	// Visit a parse tree produced by GramarParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by GramarParser#production.
	VisitProduction(ctx *ProductionContext) interface{}

	// Visit a parse tree produced by GramarParser#preturn.
	VisitPreturn(ctx *PreturnContext) interface{}

	// Visit a parse tree produced by GramarParser#pasignA.
	VisitPasignA(ctx *PasignAContext) interface{}

	// Visit a parse tree produced by GramarParser#pfuncion.
	VisitPfuncion(ctx *PfuncionContext) interface{}

	// Visit a parse tree produced by GramarParser#pllamada.
	VisitPllamada(ctx *PllamadaContext) interface{}

	// Visit a parse tree produced by GramarParser#parguments.
	VisitParguments(ctx *PargumentsContext) interface{}

	// Visit a parse tree produced by GramarParser#pargum.
	VisitPargum(ctx *PargumContext) interface{}

	// Visit a parse tree produced by GramarParser#pparams.
	VisitPparams(ctx *PparamsContext) interface{}

	// Visit a parse tree produced by GramarParser#pparamet.
	VisitPparamet(ctx *PparametContext) interface{}

	// Visit a parse tree produced by GramarParser#pdeclarArray.
	VisitPdeclarArray(ctx *PdeclarArrayContext) interface{}

	// Visit a parse tree produced by GramarParser#pdefinition.
	VisitPdefinition(ctx *PdefinitionContext) interface{}

	// Visit a parse tree produced by GramarParser#pguard.
	VisitPguard(ctx *PguardContext) interface{}

	// Visit a parse tree produced by GramarParser#pfor.
	VisitPfor(ctx *PforContext) interface{}

	// Visit a parse tree produced by GramarParser#pifor.
	VisitPifor(ctx *PiforContext) interface{}

	// Visit a parse tree produced by GramarParser#pwhile.
	VisitPwhile(ctx *PwhileContext) interface{}

	// Visit a parse tree produced by GramarParser#pswitch.
	VisitPswitch(ctx *PswitchContext) interface{}

	// Visit a parse tree produced by GramarParser#ccase.
	VisitCcase(ctx *CcaseContext) interface{}

	// Visit a parse tree produced by GramarParser#pdefault.
	VisitPdefault(ctx *PdefaultContext) interface{}

	// Visit a parse tree produced by GramarParser#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by GramarParser#pif.
	VisitPif(ctx *PifContext) interface{}

	// Visit a parse tree produced by GramarParser#pelse.
	VisitPelse(ctx *PelseContext) interface{}

	// Visit a parse tree produced by GramarParser#prin.
	VisitPrin(ctx *PrinContext) interface{}

	// Visit a parse tree produced by GramarParser#declara.
	VisitDeclara(ctx *DeclaraContext) interface{}

	// Visit a parse tree produced by GramarParser#asign.
	VisitAsign(ctx *AsignContext) interface{}

	// Visit a parse tree produced by GramarParser#tipo.
	VisitTipo(ctx *TipoContext) interface{}

	// Visit a parse tree produced by GramarParser#pespecial.
	VisitPespecial(ctx *PespecialContext) interface{}

	// Visit a parse tree produced by GramarParser#Op.
	VisitOp(ctx *OpContext) interface{}

	// Visit a parse tree produced by GramarParser#llamada.
	VisitLlamada(ctx *LlamadaContext) interface{}

	// Visit a parse tree produced by GramarParser#AccesoA.
	VisitAccesoA(ctx *AccesoAContext) interface{}

	// Visit a parse tree produced by GramarParser#Literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by GramarParser#Acceso.
	VisitAcceso(ctx *AccesoContext) interface{}

	// Visit a parse tree produced by GramarParser#Especial.
	VisitEspecial(ctx *EspecialContext) interface{}

	// Visit a parse tree produced by GramarParser#Paren.
	VisitParen(ctx *ParenContext) interface{}
}
