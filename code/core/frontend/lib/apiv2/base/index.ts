import { Http } from "../http";

export class ApiBase {
  http: Http;
  token: string;

  constructor(baseUrl: string) {
    this.http = new Http(baseUrl, {
      "Content-Type": "application/json",
      Authorization: this.token,
    });
  }

  async get(path: string) {
    return this.http.get(path);
  }

  async post(path: string, data: any) {
    return this.http.post(path, data);
  }

  async patch(path: string, data: any) {
    return this.http.patch(path, data);
  }

  async delete(path: string, data?: any) {
    return this.http.delete(path, data);
  }
}
