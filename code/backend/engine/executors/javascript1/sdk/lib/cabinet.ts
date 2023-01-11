declare function _cab_add_file(folder: string, file: string): string
declare function _cab_list_folder(folder: string, file: string): [string[], string]
declare function _cab_get_file(folder: string, file: string): [ArrayBuffer, string]
declare function _cab_get_file_str(folder: string, file: string): [string, string]
declare function _cab_del_file(folder: string, file: string): string
declare function _cab_generate_ticket(folder: string, opts: any): [string, string]

export class CabFolder {
    _folder: string
    constructor(folder: string) {
        this._folder = folder
    }

    add_file = (file: string): string => {
        return _cab_add_file(this._folder, file)
    }

    list_folder = (file: string): [string[], string] => {
        return _cab_list_folder(this._folder, file)
    }

    get_file = (file: string): [ArrayBuffer, string] => {
        return _cab_get_file(this._folder, file)
    }

    get_file_str = (file: string): [string, string] => {
        return _cab_get_file_str(this._folder, file)
    }

    del_file = (file: string): string => {
        return _cab_del_file(this._folder, file)
    }

    gen_ticket = (opts: any): [string, string] => {
        return _cab_generate_ticket(this._folder, opts)
    }
}

