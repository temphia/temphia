export interface AuthedData {
    user_token: string;
    site_token: string;
    tenant_id: string;
}
export interface SiteData {
    tenant_id: string;
    site_token: string;
    user_group?: string;
}
export declare class SiteUtils {
    _site_token: string;
    constructor(site_token?: string);
    isLogged(): boolean;
    gotoLoginPage(): void;
    setAuthedData(data: AuthedData): void;
    getAuthedData(): AuthedData;
    clearAuthedData(): void;
    get(): string;
    set(data: string): void;
    private key;
}
