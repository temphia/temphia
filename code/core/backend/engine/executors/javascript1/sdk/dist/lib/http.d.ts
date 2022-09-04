export declare class Response {
    status: number;
    headers: object;
    body: string;
    constructor(status: number, header: object, body: string);
    ok: () => boolean;
    get_header: (key: string) => any;
    json_body: () => any;
}
export declare class Request {
    _url: string;
    _headers: object;
    _body: string;
    constructor(url: string);
    set_header: (key: string, value: string) => void;
    set_body: (body: string) => void;
    set_json_body: (value: object) => void;
    get: () => Response;
    post: () => Response;
    put: () => Response;
    patch: () => Response;
    delete: () => Response;
}
