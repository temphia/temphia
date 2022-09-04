interface Response {
    payload: {
        ok: boolean;
        data?: any;
        message?: string;
    };
}
export declare const utils: {
    is_db_not_found: (err: string) => boolean;
    is_db_already_exists: (err: string) => boolean;
    ok_response: (data: any) => Response;
    err_response: (message: string) => Response;
    ab2str: (buf: ArrayBuffer) => any;
    str2ab: (str: string) => ArrayBuffer;
    is_arraybuffer(value: any): boolean;
    generate_str_id: () => string;
};
export {};
