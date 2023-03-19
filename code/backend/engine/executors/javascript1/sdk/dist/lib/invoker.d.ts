export declare const invoker: {
    name: () => string;
    exec_method: (method: string, path: string, data: any) => string;
    context_user: () => any;
    context_user_info: () => [any, string];
    context_user_message: () => any;
};
