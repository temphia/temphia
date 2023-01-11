import { authURL } from "../../lib/utils/site";
import { saveOperatorData } from "../operator/service";

export const OpLogin = (baseURL: string, user: string, password: string) => {
  return fetch(`${baseURL}/z/operator/login`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json;charset=utf-8",
    },
    body: JSON.stringify({
      user,
      password,
    }),
  });
};

export const doLogin = async (op_username: string, op_password: string) => {
  const resp = await OpLogin(window.location.origin, op_username, op_password);
  if (!resp.ok) {
    return resp.text();
  }
  const data = await resp.json();
  saveOperatorData(data);
  window.location.href = `${window.location.origin}/z/operator`;
};

export const gotoLoginPage = (tenant: string, group: string) => {
  window.location.href = authURL({ tenant_id: tenant, user_group: group });
};
