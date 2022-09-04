// tsc gp2.ts


export declare function _multi_return_err(message: string): [any, any]
export declare function _multi_return_not_err(message: string): [any, any]
export declare function _log(message: string): void


const resp1 = _multi_return_err("err example")

_log(JSON.stringify(Object.keys(resp1[1])))
_log(typeof resp1[1])
_log(JSON.stringify(resp1))

const resp2 = _multi_return_not_err("no err example")
_log(JSON.stringify(resp2))