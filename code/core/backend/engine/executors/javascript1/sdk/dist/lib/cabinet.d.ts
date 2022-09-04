export declare class CabFolder {
    _folder: string;
    constructor(folder: string);
    add_file: (file: string) => string;
    list_folder: (file: string) => [string[], string];
    get_file: (file: string) => [ArrayBuffer, string];
    get_file_str: (file: string) => [string, string];
    del_file: (file: string) => string;
    gen_ticket: (opts: any) => [string, string];
}
