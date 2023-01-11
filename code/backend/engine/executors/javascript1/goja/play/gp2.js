"use strict";
// ncc

var resp1 = _multi_return_err("err example");
_log(typeof resp1[0] + "__err___" + typeof resp1[1]);

_log((resp1[1] instanceof GoError) + "_____");

_log(JSON.stringify(resp1));

var resp2 = _multi_return_not_err("no err example");
_log(typeof resp2[0] + "__no_err___" + typeof resp2[1]);

_log(JSON.stringify(resp2));

var resp3 = _byte_return("_byte_return");
_log(typeof resp3[0] + "__byte_return___" + typeof resp3[1]);
_log(JSON.stringify(resp3));


var resp4 = _byte_return2("_byte_return2");
_log(typeof resp4[0] + "__byte_return2___" + typeof resp4[1]);
_log(Object.keys(resp4[0]));


_log(JSON.stringify(resp4) + "@"+typeof resp4[0]);

