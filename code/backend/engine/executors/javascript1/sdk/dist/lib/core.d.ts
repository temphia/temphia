export declare const core: {
    log: (message: string) => void;
    log_lazy: (message: string) => number;
    lazy_log_send: () => void;
    sleep: (t: number) => void;
    self_file: (file: string) => [string, string];
};
