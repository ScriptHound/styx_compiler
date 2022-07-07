grammar Styx_refactoring;

INT: [0-9]+;
ID: [a-zA-Z_]+;
SEMICOLON: ';';
COMMA: ',';
WS
    :   [ \t\r\n]+ -> skip
    ;

prog
    :   (statement SEMICOLON|expression SEMICOLON)+ EOF
    ;


statement
    : functionDefinition
    | functionCall
    | assignment
    ;

expression
    : multiplyExpr ('+'|'-') multiplyExpr
    ;

multiplyExpr
    : (INT|ID) ('*'|'/') (INT|ID)
    ;

statementsBlock
    : '{' ((statement|expression) SEMICOLON)* '}'
    ;

functionArgs:
    '('(ID COMMA|expression COMMA)*')';

functionDefinition
    : 'proc' ID functionArgs statementsBlock;

functionCall
    : ID functionArgs;

assignment
    : ID '=' expression
    | ID '=' ID
    | ID '=' INT
    ;
