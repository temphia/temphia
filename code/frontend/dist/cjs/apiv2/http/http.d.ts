export interface HttpResponse {
    ok: boolean;
    status: number;
    data: any;
}
export declare class Http {
    baseURL: string;
    headers: any;
    constructor(baseURL: string, headers: any);
    replace_headers(headers: any): void;
    get(path: string): Promise<HttpResponse>;
    post(path: string, data: any): Promise<HttpResponse>;
    patch(path: string, data: any): Promise<HttpResponse>;
    put(path: string, data: any): Promise<HttpResponse>;
    jsonMethod(path: string, method: string, data: any): Promise<HttpResponse>;
    postForm(path: string, auth: boolean, data: any): Promise<Response>;
    patchForm(path: string, auth: boolean, data: any): Promise<Response>;
    delete(path: string, data?: any): Promise<HttpResponse>;
}
