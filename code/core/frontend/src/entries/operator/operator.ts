import { navigate } from "svelte-routing";
import { OperatorAPI } from "../../lib/core/api/operator";

export const goto = (path) => () => navigate(path, { replace: true });

const OPERATOR_DATA = "__operator_data__"

export interface OperatorData {
    token: string
}

let data: OperatorData
let api: OperatorAPI

export const loadOperatorData = () => {
    const dataStr = localStorage.getItem(OPERATOR_DATA);
    if (!dataStr) {
        window.location.href = window.location.origin;
        return
    }

    try {
        data = JSON.parse(dataStr)
    } catch (error) {
        window.location.href = window.location.origin;
    }
    api = new OperatorAPI(data.token, window.location.origin);
}


export const saveOperatorData = (data: OperatorData) => {
    localStorage.setItem(OPERATOR_DATA, JSON.stringify(data));
}

export const listTenant = async () => (api.list_tenant())
export const createTenant = async (data: object) => (api.create_tenant(data))
export const updateTenant = async (data: object) => (api.update_tenant(data))
export const deleteTenant = async (slug: string) => (api.delete_tenant(slug))
export const goto_tenant = (id: string) => navigate(`/operator/tenants/${id}`, { replace: true });
export const ensureTenant = async (slug: string) => (console.log("fixme => ENSURE TENANT"))