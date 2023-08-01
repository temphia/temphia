import { Http } from "./http";

export class AuthAPI {
  http: Http;
  constructor(baseUrl: string, site_token: string) {
    this.http = new Http(baseUrl, {
      "Content-Type": "application/json",
      Authorization: site_token,
    });
  }

  list_methods = async (ugroup?: string) => {
    return this.http.get(`/auth?ugroup=${ugroup}`);
  };

  login_next = async (data: any) => {
    return this.http.post("/auth/login/next", data);
  };

  login_submit = async (data: any) => {
    return this.http.post("/auth/login/submit", data);
  };

  altauth_generate = async (id: number, data: any) => {
    return this.http.post(`/auth/alt/${id}/generate`, data);
  };

  altauth_next = async (id: number, stage: string, data: any) => {
    return this.http.post(`/auth/alt/${id}/next/${stage}`, data);
  };

  altauth_submit = async (id: number, data: any) => {
    return this.http.post(`/auth/alt/${id}/submit`, data);
  };

  finish = async (data: any) => {
    return this.http.post("/auth/finish", data);
  };

  signup_next = async (data: any) => {
    return this.http.post("/auth/signup/next", data);
  };

  signup_submit = async (data: any) => {
    return this.http.post("/auth/signup/submit", data);
  };

  reset_submit = async (data: any) => {
    return this.http.post("/reset/submit", data);
  };

  reset_finish = async (data: any) => {
    return this.http.post("/reset/finish", data);
  };

  about = async (user_token: string) => {
    const http = new Http(this.http.baseURL, {
      "Content-Type": "application/json",
      Authorization: user_token,
    });

    return http.get("/auth/about");
  };
}
