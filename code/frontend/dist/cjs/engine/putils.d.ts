export declare const initRegistry: () => void;
export declare const plugStart: (opts: {
    exec_loader?: string;
    plug: string;
    agent: string;
    entry: string;
    env: any;
    target: HTMLElement;
    payload?: any;
}) => Promise<void>;
