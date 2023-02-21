"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.AuthAPI = void 0;
const http_1 = require("./http");
class AuthAPI {
    constructor(baseUrl, site_token) {
        this.list_methods = async (ugroup) => {
            return this.http.get(`/auth?ugroup=${ugroup}`);
        };
        this.login_next = async (data) => {
            return this.http.post("/auth/login/next", data);
        };
        this.login_submit = async (data) => {
            return this.http.post("/auth/login/submit", data);
        };
        this.altauth_generate = async (id, data) => {
            return this.http.post(`/auth/alt/${id}/generate`, data);
        };
        this.altauth_next = async (id, stage, data) => {
            return this.http.post(`/auth/alt/${id}/next/${stage}`, data);
        };
        this.altauth_submit = async (id, data) => {
            return this.http.post(`/auth/alt/${id}/submit`, data);
        };
        this.finish = async (data) => {
            return this.http.post("/auth/finish", data);
        };
        this.signup_next = async (data) => {
            return this.http.post("/auth/signup/next", data);
        };
        this.signup_submit = async (data) => {
            return this.http.post("/auth/signup/submit", data);
        };
        this.reset_submit = async (data) => {
            return this.http.post("/reset/submit", data);
        };
        this.reset_finish = async (data) => {
            return this.http.post("/reset/finish", data);
        };
        this.about = async (user_token) => {
            const http = new http_1.Http(this.http.baseURL, {
                "Content-Type": "application/json",
                Authorization: user_token,
            });
            return http.get("/auth/about");
        };
        this.http = new http_1.Http(baseUrl, {
            "Content-Type": "application/json",
            Authorization: site_token,
        });
    }
}
exports.AuthAPI = AuthAPI;
