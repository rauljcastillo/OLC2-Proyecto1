// Code generated from Gramar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Gramar

import "github.com/antlr4-go/antlr/v4"

// GramarListener is a complete listener for a parse tree produced by GramarParser.
type GramarListener interface {
	antlr.ParseTreeListener

	// EnterS is called when entering the s production.
	EnterS(c *SContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterProduction is called when entering the production production.
	EnterProduction(c *ProductionContext)

	// EnterPreturn is called when entering the preturn production.
	EnterPreturn(c *PreturnContext)

	// EnterPasignA is called when entering the pasignA production.
	EnterPasignA(c *PasignAContext)

	// EnterPfuncion is called when entering the pfuncion production.
	EnterPfuncion(c *PfuncionContext)

	// EnterPllamada is called when entering the pllamada production.
	EnterPllamada(c *PllamadaContext)

	// EnterParguments is called when entering the parguments production.
	EnterParguments(c *PargumentsContext)

	// EnterPargum is called when entering the pargum production.
	EnterPargum(c *PargumContext)

	// EnterPparams is called when entering the pparams production.
	EnterPparams(c *PparamsContext)

	// EnterPparamet is called when entering the pparamet production.
	EnterPparamet(c *PparametContext)

	// EnterPdeclarArray is called when entering the pdeclarArray production.
	EnterPdeclarArray(c *PdeclarArrayContext)

	// EnterPdefinition is called when entering the pdefinition production.
	EnterPdefinition(c *PdefinitionContext)

	// EnterPguard is called when entering the pguard production.
	EnterPguard(c *PguardContext)

	// EnterPfor is called when entering the pfor production.
	EnterPfor(c *PforContext)

	// EnterPifor is called when entering the pifor production.
	EnterPifor(c *PiforContext)

	// EnterPwhile is called when entering the pwhile production.
	EnterPwhile(c *PwhileContext)

	// EnterPswitch is called when entering the pswitch production.
	EnterPswitch(c *PswitchContext)

	// EnterCcase is called when entering the ccase production.
	EnterCcase(c *CcaseContext)

	// EnterPdefault is called when entering the pdefault production.
	EnterPdefault(c *PdefaultContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterPif is called when entering the pif production.
	EnterPif(c *PifContext)

	// EnterPelse is called when entering the pelse production.
	EnterPelse(c *PelseContext)

	// EnterPrin is called when entering the prin production.
	EnterPrin(c *PrinContext)

	// EnterDeclara is called when entering the declara production.
	EnterDeclara(c *DeclaraContext)

	// EnterAsign is called when entering the asign production.
	EnterAsign(c *AsignContext)

	// EnterTipo is called when entering the tipo production.
	EnterTipo(c *TipoContext)

	// EnterPespecial is called when entering the pespecial production.
	EnterPespecial(c *PespecialContext)

	// EnterOp is called when entering the Op production.
	EnterOp(c *OpContext)

	// EnterLlamada is called when entering the llamada production.
	EnterLlamada(c *LlamadaContext)

	// EnterAccesoA is called when entering the AccesoA production.
	EnterAccesoA(c *AccesoAContext)

	// EnterLiteral is called when entering the Literal production.
	EnterLiteral(c *LiteralContext)

	// EnterAcceso is called when entering the Acceso production.
	EnterAcceso(c *AccesoContext)

	// EnterEspecial is called when entering the Especial production.
	EnterEspecial(c *EspecialContext)

	// EnterParen is called when entering the Paren production.
	EnterParen(c *ParenContext)

	// ExitS is called when exiting the s production.
	ExitS(c *SContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitProduction is called when exiting the production production.
	ExitProduction(c *ProductionContext)

	// ExitPreturn is called when exiting the preturn production.
	ExitPreturn(c *PreturnContext)

	// ExitPasignA is called when exiting the pasignA production.
	ExitPasignA(c *PasignAContext)

	// ExitPfuncion is called when exiting the pfuncion production.
	ExitPfuncion(c *PfuncionContext)

	// ExitPllamada is called when exiting the pllamada production.
	ExitPllamada(c *PllamadaContext)

	// ExitParguments is called when exiting the parguments production.
	ExitParguments(c *PargumentsContext)

	// ExitPargum is called when exiting the pargum production.
	ExitPargum(c *PargumContext)

	// ExitPparams is called when exiting the pparams production.
	ExitPparams(c *PparamsContext)

	// ExitPparamet is called when exiting the pparamet production.
	ExitPparamet(c *PparametContext)

	// ExitPdeclarArray is called when exiting the pdeclarArray production.
	ExitPdeclarArray(c *PdeclarArrayContext)

	// ExitPdefinition is called when exiting the pdefinition production.
	ExitPdefinition(c *PdefinitionContext)

	// ExitPguard is called when exiting the pguard production.
	ExitPguard(c *PguardContext)

	// ExitPfor is called when exiting the pfor production.
	ExitPfor(c *PforContext)

	// ExitPifor is called when exiting the pifor production.
	ExitPifor(c *PiforContext)

	// ExitPwhile is called when exiting the pwhile production.
	ExitPwhile(c *PwhileContext)

	// ExitPswitch is called when exiting the pswitch production.
	ExitPswitch(c *PswitchContext)

	// ExitCcase is called when exiting the ccase production.
	ExitCcase(c *CcaseContext)

	// ExitPdefault is called when exiting the pdefault production.
	ExitPdefault(c *PdefaultContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitPif is called when exiting the pif production.
	ExitPif(c *PifContext)

	// ExitPelse is called when exiting the pelse production.
	ExitPelse(c *PelseContext)

	// ExitPrin is called when exiting the prin production.
	ExitPrin(c *PrinContext)

	// ExitDeclara is called when exiting the declara production.
	ExitDeclara(c *DeclaraContext)

	// ExitAsign is called when exiting the asign production.
	ExitAsign(c *AsignContext)

	// ExitTipo is called when exiting the tipo production.
	ExitTipo(c *TipoContext)

	// ExitPespecial is called when exiting the pespecial production.
	ExitPespecial(c *PespecialContext)

	// ExitOp is called when exiting the Op production.
	ExitOp(c *OpContext)

	// ExitLlamada is called when exiting the llamada production.
	ExitLlamada(c *LlamadaContext)

	// ExitAccesoA is called when exiting the AccesoA production.
	ExitAccesoA(c *AccesoAContext)

	// ExitLiteral is called when exiting the Literal production.
	ExitLiteral(c *LiteralContext)

	// ExitAcceso is called when exiting the Acceso production.
	ExitAcceso(c *AccesoContext)

	// ExitEspecial is called when exiting the Especial production.
	ExitEspecial(c *EspecialContext)

	// ExitParen is called when exiting the Paren production.
	ExitParen(c *ParenContext)
}
