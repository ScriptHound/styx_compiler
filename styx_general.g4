grammar styx_general;

INT : [0-9]+;
ID: [a-zA-Z_\-]+;

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

expression
     : assignment
     | additiveExpr
     ;

expressionList
     : expression (',' expression)*
     ;

expressionsBlock
     : '{' (expression ';')* '}'
     ;

returnStatement
     : 'return' ID
     | 'return' INT
     ;

procedure
     : 'proc' ID '(' expressionList ')' expressionsBlock
     ;


WS
    :   [ \t\r\n]+ -> skip
    ;