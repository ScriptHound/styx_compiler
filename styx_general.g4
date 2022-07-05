grammar styx_general;

INT : [0-9]+;
ID: [a-zA-Z]+;

assignment
     :    ID '=' INT
     |    ID '=' ID
     |    ID '=' additiveExpr
     ;

multiplicativeExpr 
     :    INT (('*'|'/') INT)*
     |    ID (('*'|'/') ID)*
     |    INT (('*'|'/') ID)*
     |    ID (('*'|'/') INT)*
     ;

additiveExpr :
     multiplicativeExpr (('+'| '-') multiplicativeExpr)*
     ;

declarator
     : ID
     ;

declaration
     :    
     ;

declarationList
     : declaration+
     ;

functionDefinition
    :     declarator declarationList? compoundStatement
    ;

WS
    :   [ \t\r\n]+ -> skip
    ;