export class OperatorAPI {
  token: string;
  baseURL: string;
  tenantURL: string;

  constructor(token: string, baseURL: string) {
    this.token = token;
    this.baseURL = baseURL;
    this.tenantURL = `${this.baseURL}/z/operator/tenant`;
  }

  create_tenant = async (data: object) => {
    let response = await fetch(this.tenantURL, {
      method: "POST",
      headers: this.header(),
      body: JSON.stringify(data),
    });
    if (response.ok) {
      return response.json();
    }
    return response.text();
  };

  list_tenant = async () => {
    let response = await fetch(this.tenantURL, {
      method: "GET",
      headers: this.header(),
    });
    if (response.ok) {
      return response.json();
    }
    return response.text();
  };

  update_tenant = async (data: object) => {
    let response = await fetch(this.tenantURL, {
      method: "PATCH",
      headers: this.header(),
      body: JSON.stringify(data),
    });
    if (response.ok) {
      return response.json();
    }
    return response.text();
  };

  delete_tenant = async (id: string) => {
    let response = await fetch(this.tenantURL, {
      method: "DELETE",
      headers: this.header(),
      body: JSON.stringify({ slug: id }),
    });
    if (response.ok) {
      return response.json();
    }
    return response.text();
  };

  // private

  header = () => ({
    "Content-Type": "application/json;charset=utf-8",
    Authorization: this.token,
  });
}

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
