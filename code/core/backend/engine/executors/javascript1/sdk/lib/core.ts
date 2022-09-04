declare function _log(message: string): void
declare function _log_lazy(message: string[]): void
declare function _sleep(time: number): void
declare function _get_self_file_as_str(file: string): [string, string]


const _buffer: string[] = []
export const core = {
    log: (message: string) => _log(message),
    log_lazy: (message: string) => _buffer.push(message),
    lazy_log_send: () => {
        _log_lazy(_buffer)
        _buffer.length = 0
    },
    sleep: (t: number) => _sleep(t),
    self_file: (file: string) => _get_self_file_as_str(file),
    
}
