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
CADENA: '"' (~["\r\n] | '""')* '"';
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
    | ID '(' parguments ')' ';'?    
;

parguments: pargum (',' pargum)*
;

pargum:  
     PUNTE? expr
    |ID TPOINTS expr
;

pparams: (pparamet (',' pparamet)*)?
;

pparamet: pid=(ID | GUION) ID ':' (INOUT)? tipo
    | pid=(ID | GUION) ID ':' (INOUT)? ABR tipo CIER
;

pdeclarArray: RVAR ID ':' '[' tipo ']' pdefinition
    | RVAR ID pdefinition
;

pdefinition: '=' '[' (expr (',' expr)*)?  ']'
    | '=' '[' expr ']'
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

prin: PRINT '(' expr ')'';'?;


declara: RVAR ID ':' tipo '=' expr ';'?
    | RVAR ID ':' tipo '?'
    | RVAR ID '=' expr
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


expr: op=('!'| '-') right=expr                             #Op
    | left=expr op=('*' | '/' | '%') right=expr            #Op
    | left=expr op=('+' | '-') right=expr                  #Op
    | left=expr op=('<' | '>' | '<=' | '>=') right=expr    #Op
    | left=expr op=('==' | '!=') right=expr                #Op
    | left=expr op='&&' right=expr                         #Op
    | left=expr op='||' right=expr                         #Op
    | left=expr op=',' right=expr                          #Op
    | '(' expr ')'                                         #Paren
    | INT                                                  #Literal
    | logico=(TRUE | FALSE)                                #Literal
    | CADENA                                               #Literal
    | DOUBLE                                               #Literal
    | ID                                                   #Acceso
    | ID '.' EMPT ';'?                                     #Especial
    | ID '.' COUNT ';'?                                    #Especial
    | ID '[' expr ']' ';'?                                 #AccesoA
    | pllamada                                             #llamada                                           
;
