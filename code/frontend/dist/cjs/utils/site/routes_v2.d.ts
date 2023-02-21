export declare const apiURL: (tenant_id: string) => string;
export declare const baseURL: () => string;
export declare const portalURL: () => string;
export declare const authURL: (opts?: {
    tenant_id?: string;
    user_group?: string;
}) => string;
