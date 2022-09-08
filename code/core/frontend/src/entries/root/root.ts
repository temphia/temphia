
import { OpLogin } from "../../lib/core/api/operator";
import { saveOperatorData } from "../operator/operator";


export const doLogin = async (op_username: string, op_password: string) => {
    const resp = await OpLogin(
        window.location.origin,
        op_username,
        op_password
    );
    if (!resp.ok) {
        return resp.text();
    }
    const data = await resp.json();
    saveOperatorData(data);
    window.location.href = `${window.location.origin}/z/operator`;
}


export const gotoLoginPage = (tenant: string, group:string) => {
    window.location.href = `${window.location.origin}/z/auth?tenant_id=${tenant}&ugroup=${group}`;
}