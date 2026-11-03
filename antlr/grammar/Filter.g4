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
    : name
    ;

// Quoted when it holds ',', a quote or ')' (SOL 013 5.2.2); unquoted, it may span several tokens.
value
    : QUOTED
    | (name | '/' | '=')+
    ;

name
    : STRING
    | keyword
    ;

keyword
    : 'filter' | 'eq' | 'neq' | 'gt' | 'lt' | 'gte' | 'lte' | 'in' | 'nin' | 'cont' | 'ncont'
    ;

QUOTED
    : '\'' ( ~'\'' | '\'\'' )* '\''
    ;

STRING
    : [-a-zA-Z0-9@_.:?&%+~#!*$]+
    ;

WS
    : [ \t\r\n]+ -> skip
    ;
