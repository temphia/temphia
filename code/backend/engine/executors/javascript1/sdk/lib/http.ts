declare function _http1(method: string, url: string, headers: object, body: string): [number, object, string];

export class Response {
    status: number
    headers: object
    body: string

    constructor(status: number, header: object, body: string) {
        this.status = status
        this.headers = header
        this.body = body
    }

    ok = (): boolean => {
        return this.status >= 200 && this.status < 300
    }

    get_header = (key: string) => {
        return this.headers[key]
    }

    json_body = () => {
        return JSON.parse(this.body)
    }
}

export class Request {
    _url: string
    _headers: object
    _body: string

    constructor(url: string) {
        this._url = url
        this._headers = {}
    }

    set_header = (key: string, value: string) => {
        this._headers[key] = value
    }

    set_body = (body: string) => {
        this._body = body
    }

    set_json_body = (value: object) => {
        this._body = JSON.stringify(value)
    }

    get = () => {
        const [status, header, body] = _http1("GET", this._url, this._headers, "")
        return new Response(status, header, body);
    }

    post = () => {
        const [status, header, body] = _http1("POST", this._url, this._headers, this._body)
        return new Response(status, header, body);
    }

    put = () => {
        const [status, header, body] = _http1("PUT", this._url, this._headers, this._body)
        return new Response(status, header, body);
    }

    patch = () => {
        const [status, header, body] = _http1("PATCH", this._url, this._headers, this._body)
        return new Response(status, header, body);
    }

    delete = () => {
        const [status, header, body] = _http1("DELETE", this._url, this._headers, this._body)
        return new Response(status, header, body);
    }

}


