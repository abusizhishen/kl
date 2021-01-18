grammar Kl;

Id: [a-zA-Z_]+;
STRING:'"' [a-zA-Z_]+ '"';
Num: [0-9]+;
stat: Id '=' value ';';
value: (Id|Num|STRING);

SW: [\n\r ]+ ->skip;

param
    :Id' 'Id(',' Id Id)*;
fun: 'func' ' '* Id '(' param*  ')' ' '*'{'
    (stat|callFun)*
'}';

args: value (','value)*;
callFun: Id'('args*')';

init : (stat|fun|callFun)*;