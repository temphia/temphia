export class Http {
  baseURL: string;
  headers: any;

  constructor(baseURL: string, headers: any) {
    this.baseURL = baseURL;
  }

  replace_headers(headers: any) {
    this.headers = headers;
  }

  async get(path: string) {
    return fetch(`${this.baseURL}/${path}`, {
      method: "GET",
      headers: this.headers,
    });
  }

  async post(path: string, data: any) {
    return fetch(`${this.baseURL}/${path}`, {
      method: "POST",
      headers: this.headers,
      body: JSON.stringify(data),
    });
  }

  async patch(path: string, data: any) {
    return fetch(`${this.baseURL}/${path}`, {
      method: "PATCH",
      headers: this.headers,
      body: JSON.stringify(data),
    });
  }

  async delete(path: string, data?: any) {
    return fetch(`${this.baseURL}/${path}`, {
      method: "DELETE",
      headers: this.headers,
      body: data ? JSON.stringify(data) : data,
    });
  }
}
