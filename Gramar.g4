grammar Gramar;

WS: [ \n\t]+ -> skip;
COMO: '//'~[\n]* ->skip;
COMM: '/*' .*? '*/'->skip;
/*Reservadas */
PRINT: 'print';
TRUE: 'true';
FALSE: 'false';
RINT: 'Int';
RFLOAT: 'Float';
RSTRING: 'String';
RBOOL: 'Bool';
RChar: 'Character';
RVAR: 'var';
LET: 'let';
RETN: 'return';
CONT: 'continue';
BRK: 'break';
APPE: 'append';
REMOV: 'remove';
REMLST: 'removeLast';
EMPT:'IsEmpty';
COUNT: 'count';
ABR: '[';
CIER: ']';
GUION: '_';
POINTS: '...';
INOUT: 'inout';
PUNTE: '&';
COMMA: ',';
TPOINTS: ':';
/*Expresiones regulares */
CADENA: '"' ( ~["\\] | '\\' [nrt\\"\\] )* '"';
ID: [a-zA-Z_][a-zA-Z0-9]*;
INT: [0-9]+;
DOUBLE: [0-9]+'.'[0-9]+;
CHARAC: '\'' (~["\r\n] | '\'')? '\''; 

s: block EOF;

block: (production)*;

production: prin
    | pasignA
    | declara
    | asign
    | pif
    | pswitch
    | pwhile
    | pfor
    | pguard
    | pdeclarArray
    | pespecial
    | pfuncion
    | pllamada          
    | preturn
    | CONT
    | BRK
;

preturn: RETN expr?
;

pasignA: ID '[' expr ']' '=' expr ';'?
;

pfuncion: 'func' ID '(' ')' ('->' tipo)? '{' block '}'
    | 'func' ID '(' pparams ')' ('->' tipo)? '{' block '}'
;

pllamada: ID '(' ')' ';'?           
    | ID '(' argumento ')' ';'?    
;

pparams: (pparamet (',' pparamet)*)?
;

pparamet: pid=(ID | GUION) ID ':' (INOUT)? tipo
    | pid=(ID | GUION) ID ':' (INOUT)? ABR tipo CIER
;

pdeclarArray: (RVAR | LET) ID ':' '[' tipo ']' pdefinition
    | (RVAR | LET) ID pdefinition
;

pdefinition: '=' '[' (expr (',' expr)*) ']'
;

pguard: 'guard' expr 'else' '{' block opera=(RETN | CONT | BRK)'}'
;

pfor: 'for' ID 'in' pifor '{' block '}';

pifor: expr POINTS expr
    | expr
;

pwhile: 'while' expr '{' block '}';

pswitch: 'switch' expr '{' (ccase)+ pdefault '}';

ccase: 'case' expr ':' block 'break'?
;

 pdefault: 'default' ':' block 'break'?
; 

stmt: '{' block '}'
    | '{' '}'
;


pif: 'if' expr stmt
    | 'if' expr stmt pelse
;

pelse: 'else' pif
    | 'else' stmt
;

prin: PRINT '(' expr ')'';'?
   | PRINT '(' cexpr ')'';'?
;

cexpr: expr ',' cexpr
    | expr
;

declara: (RVAR | LET) ID ':' tipo '=' expr ';'?
    | (RVAR | LET) ID ':' tipo '?'
    | (RVAR | LET) ID '=' expr
;

asign: ID op=('=' | '+=' | '-=') expr 
;

tipo: RINT
    | RFLOAT
    | RSTRING
    | RBOOL
    | RChar
;

pespecial: ID '.' APPE '(' expr ')' ';'?
    | ID '.' REMLST '(' ')' ';'?
    | ID '.' REMOV '(' 'at:' expr ')' ';'?
;


argumento: tipoArg ',' argumento
    | tipoArg
;

tipoArg: PUNTE exp1=expr
    | ID ':' exp2=expr   
    | exp3=expr
;

expr: op=('!'| '-') right=expr                   #Op
    | left=expr op=('/'|'*'|'%') right=expr                #Op
    | left=expr op=('+' | '-') right=expr                  #Op
    | left=expr op=('<' | '>' | '<=' | '>=') right=expr    #Op
    | left=expr op=('==' | '!=') right=expr                #Op
    | left=expr op='&&' right=expr                         #Op
    | left=expr op='||' right=expr                         #Op
    | pllamada                                              #llamada
    | '(' expr ')'                                         #Paren
    | INT                                                  #Literal
    | logico=(TRUE | FALSE)                                #Literal
    | CADENA                                               #Literal
    | DOUBLE                                               #Literal
    | ID                                                   #Acceso
    | ID '.' EMPT ';'?                                     #Especial
    | ID '.' COUNT ';'?                                    #Especial
    | ID '[' expr ']' ';'?                                 #AccesoA
    | RINT '(' expr ')' ';'?                               #Embebida
    | RFLOAT '(' expr ')'                                  #Embebida
    | RSTRING '(' expr ')'';'?                             #Embebida
                    
;
