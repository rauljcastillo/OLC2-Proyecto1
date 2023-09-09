// Code generated from Gramar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Gramar

import "github.com/antlr4-go/antlr/v4"

// BaseGramarListener is a complete listener for a parse tree produced by GramarParser.
type BaseGramarListener struct{}

var _ GramarListener = &BaseGramarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGramarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGramarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGramarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGramarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterS is called when production s is entered.
func (s *BaseGramarListener) EnterS(ctx *SContext) {}

// ExitS is called when production s is exited.
func (s *BaseGramarListener) ExitS(ctx *SContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseGramarListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseGramarListener) ExitBlock(ctx *BlockContext) {}

// EnterProduction is called when production production is entered.
func (s *BaseGramarListener) EnterProduction(ctx *ProductionContext) {}

// ExitProduction is called when production production is exited.
func (s *BaseGramarListener) ExitProduction(ctx *ProductionContext) {}

// EnterPreturn is called when production preturn is entered.
func (s *BaseGramarListener) EnterPreturn(ctx *PreturnContext) {}

// ExitPreturn is called when production preturn is exited.
func (s *BaseGramarListener) ExitPreturn(ctx *PreturnContext) {}

// EnterPasignA is called when production pasignA is entered.
func (s *BaseGramarListener) EnterPasignA(ctx *PasignAContext) {}

// ExitPasignA is called when production pasignA is exited.
func (s *BaseGramarListener) ExitPasignA(ctx *PasignAContext) {}

// EnterPfuncion is called when production pfuncion is entered.
func (s *BaseGramarListener) EnterPfuncion(ctx *PfuncionContext) {}

// ExitPfuncion is called when production pfuncion is exited.
func (s *BaseGramarListener) ExitPfuncion(ctx *PfuncionContext) {}

// EnterPllamada is called when production pllamada is entered.
func (s *BaseGramarListener) EnterPllamada(ctx *PllamadaContext) {}

// ExitPllamada is called when production pllamada is exited.
func (s *BaseGramarListener) ExitPllamada(ctx *PllamadaContext) {}

// EnterParguments is called when production parguments is entered.
func (s *BaseGramarListener) EnterParguments(ctx *PargumentsContext) {}

// ExitParguments is called when production parguments is exited.
func (s *BaseGramarListener) ExitParguments(ctx *PargumentsContext) {}

// EnterPargum is called when production pargum is entered.
func (s *BaseGramarListener) EnterPargum(ctx *PargumContext) {}

// ExitPargum is called when production pargum is exited.
func (s *BaseGramarListener) ExitPargum(ctx *PargumContext) {}

// EnterPparams is called when production pparams is entered.
func (s *BaseGramarListener) EnterPparams(ctx *PparamsContext) {}

// ExitPparams is called when production pparams is exited.
func (s *BaseGramarListener) ExitPparams(ctx *PparamsContext) {}

// EnterPparamet is called when production pparamet is entered.
func (s *BaseGramarListener) EnterPparamet(ctx *PparametContext) {}

// ExitPparamet is called when production pparamet is exited.
func (s *BaseGramarListener) ExitPparamet(ctx *PparametContext) {}

// EnterPdeclarArray is called when production pdeclarArray is entered.
func (s *BaseGramarListener) EnterPdeclarArray(ctx *PdeclarArrayContext) {}

// ExitPdeclarArray is called when production pdeclarArray is exited.
func (s *BaseGramarListener) ExitPdeclarArray(ctx *PdeclarArrayContext) {}

// EnterPdefinition is called when production pdefinition is entered.
func (s *BaseGramarListener) EnterPdefinition(ctx *PdefinitionContext) {}

// ExitPdefinition is called when production pdefinition is exited.
func (s *BaseGramarListener) ExitPdefinition(ctx *PdefinitionContext) {}

// EnterPguard is called when production pguard is entered.
func (s *BaseGramarListener) EnterPguard(ctx *PguardContext) {}

// ExitPguard is called when production pguard is exited.
func (s *BaseGramarListener) ExitPguard(ctx *PguardContext) {}

// EnterPfor is called when production pfor is entered.
func (s *BaseGramarListener) EnterPfor(ctx *PforContext) {}

// ExitPfor is called when production pfor is exited.
func (s *BaseGramarListener) ExitPfor(ctx *PforContext) {}

// EnterPifor is called when production pifor is entered.
func (s *BaseGramarListener) EnterPifor(ctx *PiforContext) {}

// ExitPifor is called when production pifor is exited.
func (s *BaseGramarListener) ExitPifor(ctx *PiforContext) {}

// EnterPwhile is called when production pwhile is entered.
func (s *BaseGramarListener) EnterPwhile(ctx *PwhileContext) {}

// ExitPwhile is called when production pwhile is exited.
func (s *BaseGramarListener) ExitPwhile(ctx *PwhileContext) {}

// EnterPswitch is called when production pswitch is entered.
func (s *BaseGramarListener) EnterPswitch(ctx *PswitchContext) {}

// ExitPswitch is called when production pswitch is exited.
func (s *BaseGramarListener) ExitPswitch(ctx *PswitchContext) {}

// EnterCcase is called when production ccase is entered.
func (s *BaseGramarListener) EnterCcase(ctx *CcaseContext) {}

// ExitCcase is called when production ccase is exited.
func (s *BaseGramarListener) ExitCcase(ctx *CcaseContext) {}

// EnterPdefault is called when production pdefault is entered.
func (s *BaseGramarListener) EnterPdefault(ctx *PdefaultContext) {}

// ExitPdefault is called when production pdefault is exited.
func (s *BaseGramarListener) ExitPdefault(ctx *PdefaultContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BaseGramarListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseGramarListener) ExitStmt(ctx *StmtContext) {}

// EnterPif is called when production pif is entered.
func (s *BaseGramarListener) EnterPif(ctx *PifContext) {}

// ExitPif is called when production pif is exited.
func (s *BaseGramarListener) ExitPif(ctx *PifContext) {}

// EnterPelse is called when production pelse is entered.
func (s *BaseGramarListener) EnterPelse(ctx *PelseContext) {}

// ExitPelse is called when production pelse is exited.
func (s *BaseGramarListener) ExitPelse(ctx *PelseContext) {}

// EnterPrin is called when production prin is entered.
func (s *BaseGramarListener) EnterPrin(ctx *PrinContext) {}

// ExitPrin is called when production prin is exited.
func (s *BaseGramarListener) ExitPrin(ctx *PrinContext) {}

// EnterDeclara is called when production declara is entered.
func (s *BaseGramarListener) EnterDeclara(ctx *DeclaraContext) {}

// ExitDeclara is called when production declara is exited.
func (s *BaseGramarListener) ExitDeclara(ctx *DeclaraContext) {}

// EnterAsign is called when production asign is entered.
func (s *BaseGramarListener) EnterAsign(ctx *AsignContext) {}

// ExitAsign is called when production asign is exited.
func (s *BaseGramarListener) ExitAsign(ctx *AsignContext) {}

// EnterTipo is called when production tipo is entered.
func (s *BaseGramarListener) EnterTipo(ctx *TipoContext) {}

// ExitTipo is called when production tipo is exited.
func (s *BaseGramarListener) ExitTipo(ctx *TipoContext) {}

// EnterPespecial is called when production pespecial is entered.
func (s *BaseGramarListener) EnterPespecial(ctx *PespecialContext) {}

// ExitPespecial is called when production pespecial is exited.
func (s *BaseGramarListener) ExitPespecial(ctx *PespecialContext) {}

// EnterOp is called when production Op is entered.
func (s *BaseGramarListener) EnterOp(ctx *OpContext) {}

// ExitOp is called when production Op is exited.
func (s *BaseGramarListener) ExitOp(ctx *OpContext) {}

// EnterLlamada is called when production llamada is entered.
func (s *BaseGramarListener) EnterLlamada(ctx *LlamadaContext) {}

// ExitLlamada is called when production llamada is exited.
func (s *BaseGramarListener) ExitLlamada(ctx *LlamadaContext) {}

// EnterAccesoA is called when production AccesoA is entered.
func (s *BaseGramarListener) EnterAccesoA(ctx *AccesoAContext) {}

// ExitAccesoA is called when production AccesoA is exited.
func (s *BaseGramarListener) ExitAccesoA(ctx *AccesoAContext) {}

// EnterLiteral is called when production Literal is entered.
func (s *BaseGramarListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production Literal is exited.
func (s *BaseGramarListener) ExitLiteral(ctx *LiteralContext) {}

// EnterAcceso is called when production Acceso is entered.
func (s *BaseGramarListener) EnterAcceso(ctx *AccesoContext) {}

// ExitAcceso is called when production Acceso is exited.
func (s *BaseGramarListener) ExitAcceso(ctx *AccesoContext) {}

// EnterEspecial is called when production Especial is entered.
func (s *BaseGramarListener) EnterEspecial(ctx *EspecialContext) {}

// ExitEspecial is called when production Especial is exited.
func (s *BaseGramarListener) ExitEspecial(ctx *EspecialContext) {}

// EnterParen is called when production Paren is entered.
func (s *BaseGramarListener) EnterParen(ctx *ParenContext) {}

// ExitParen is called when production Paren is exited.
func (s *BaseGramarListener) ExitParen(ctx *ParenContext) {}
