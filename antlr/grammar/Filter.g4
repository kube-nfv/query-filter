grammar Filter;

filter
    : 'filter' '=' filterExpr
    ;

filterExpr
    : simpleFilterExpr (';' simpleFilterExpr)*
    ;

simpleFilterExpr
    : '(' simpleFilterExprOne ')'
    | '(' simpleFilterExprMulti ')'
    ;

simpleFilterExprOne
    : opOne ',' attrName ( '/' attrName )* ',' value
    ;

simpleFilterExprMulti
    : opMulti ',' attrName ( '/' attrName )* ',' value ( ',' value )*
    ;

opOne
    : 'eq'
    | 'neq'
    | 'gt'
    | 'lt'
    | 'gte'
    | 'lte'
    ;

opMulti
    : 'in'
    | 'nin'
    | 'cont'
    | 'ncont'
    ;

attrName
    : STRING
    ;

value
    : STRING
    ;

STRING
    : [-a-zA-Z0-9@_]+    // Basic string support, can be extended
    ;

WS
    : [ \t\r\n]+ -> skip
    ;
