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



/*Expresiones regulares */
CADENA: '"' (~["\r\n] | '""')* '"';
ID: [a-zA-Z_][a-zA-Z0-9]*;
INT: [0-9]+;
DOUBLE: [0-9]+'.'[0-9]+;

s: block EOF;

block: (production)*;

production: prin
    | declara
    | asign
    | pif
    | pswitch
    | pwhile
    | pfor
;

pfor: 'for' ID 'in' pifor '{' block '}';

pifor: INT '...' INT
    | expr
;

pwhile: 'while' expr stmt;

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

prin: PRINT '(' expr ')';


declara: RVAR ID ':' tipo '=' expr
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



expr: op=('!'| '-') right=expr                             #Op
    | left=expr op=('*' | '/' | '%') right=expr            #Op
    | left=expr op=('+' | '-') right=expr                  #Op
    | left=expr op=('<' | '>' | '<=' | '>=') right=expr    #Op
    | left=expr op=('==' | '!=') right=expr                #Op
    | '(' expr ')'                                         #Paren
    | INT                                                  #Literal
    | logico=(TRUE | FALSE)                                #Literal
    | ID                                                   #Acceso
    | CADENA                                               #Literal
    | DOUBLE                                               #Literal
;