export interface HttpResponse {
  ok: boolean;
  status: number;
  data: any;
}

export class Http {
  baseURL: string;
  headers: any;

  constructor(baseURL: string, headers: any) {
    this.baseURL = baseURL;
    this.headers = headers
  }

  replace_headers(headers: any) {
    this.headers = headers;
  }

  async get(path: string): Promise<HttpResponse> {
    const resp = await fetch(`${this.baseURL}${path}`, {
      method: "GET",
      headers: this.headers,
    });

    if (resp.ok) {
      return Promise.resolve({
        ok: true,
        data: await resp.json(),
        status: resp.status,
      });
    }

    return Promise.resolve({
      ok: false,
      data: await resp.text(),
      status: resp.status,
    });
  }

  async post(path: string, data: any): Promise<HttpResponse> {
    const resp = await fetch(`${this.baseURL}${path}`, {
      method: "POST",
      headers: this.headers,
      body: JSON.stringify(data),
    });

    if (resp.ok) {
      return Promise.resolve({
        ok: true,
        data: await resp.json(),
        status: resp.status,
      });
    }

    return Promise.resolve({
      ok: false,
      data: await resp.text(),
      status: resp.status,
    });
  }

  async postForm(path: string, auth: boolean, data: any) {
    return await fetch(`${this.baseURL}${path}`, {
      method: "POST",
      headers: auth ? { Authorization: this.headers["Authorization"] } : {},
      body: data,
    });
  }

  async patch(path: string, data: any): Promise<HttpResponse> {
    const resp = await fetch(`${this.baseURL}${path}`, {
      method: "PATCH",
      headers: this.headers,
      body: JSON.stringify(data),
    });

    if (resp.ok) {
      return Promise.resolve({
        ok: true,
        data: await resp.json(),
        status: resp.status,
      });
    }

    return Promise.resolve({
      ok: false,
      data: await resp.text(),
      status: resp.status,
    });
  }

  async delete(path: string, data?: any): Promise<HttpResponse> {
    const resp = await fetch(`${this.baseURL}${path}`, {
      method: "DELETE",
      headers: this.headers,
      body: data ? JSON.stringify(data) : data,
    });

    if (resp.ok) {
      return Promise.resolve({
        ok: true,
        data: await resp.json(),
        status: resp.status,
      });
    }

    return Promise.resolve({
      ok: false,
      data: await resp.text(),
      status: resp.status,
    });
  }
}
