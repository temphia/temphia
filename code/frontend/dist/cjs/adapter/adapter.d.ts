import type { AdapterEditorAPI } from "../apiv2/admin/adapter_editor";
export declare class AdapterEditorEnv {
    api: AdapterEditorAPI;
    domain_name: string;
    constructor(api: AdapterEditorAPI, domain_name: string);
}
