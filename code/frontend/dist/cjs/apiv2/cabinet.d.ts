import type { ApiBase } from "./base";
import { Http } from "./http";
export declare class CabinetAPI {
    base: ApiBase;
    source: string;
    constructor(source: string, base: ApiBase);
    listRoot(): Promise<import("./http").HttpResponse>;
    listFolder(folder: string): Promise<import("./http").HttpResponse>;
    newFolder(folder: string): Promise<import("./http").HttpResponse>;
    getFile(folder: string, fname: string): Promise<import("./http").HttpResponse>;
    uploadFile(folder: string, fname: string, data: any): Promise<Response>;
    deleteFile(folder: string, fname: string): Promise<import("./http").HttpResponse>;
    getFilePreview(folder: string, fname: string): string;
}
export declare class FolderTktAPI {
    http: Http;
    ticket: string;
    base_url: string;
    constructor(baseUrl: string, token: string);
    list(): Promise<import("./http").HttpResponse>;
    getFile(file: string): Promise<import("./http").HttpResponse>;
    getFileUrl(file: string): string;
    getFilePreviewUrl(file: string): string;
    uploadFile(file: string, data: any): Promise<import("./http").HttpResponse>;
    deleteFile(file: string): Promise<import("./http").HttpResponse>;
}
