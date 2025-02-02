grammar Filter;

filter: 'filter=' simpleFilterExpr (';' simpleFilterExpr)*;
simpleFilterExpr: '(' (opOne ',' attrName ',' value) | (opMulti ',' attrName ',' value (',' value)*) ')';
opOne: 'eq' | 'neq' | 'gt' | 'lt' | 'gte' | 'lte';
opMulti: 'in' | 'nin' | 'cont' | 'ncont';
attrName: ID ( '/' ID )*;
value: STRING | NUMBER;
ID: [a-zA-Z_][a-zA-Z0-9_]*;
STRING: '"' ~["]* '"';
NUMBER: [0-9]+;
WS: [ \t\r\n]+ -> skip
