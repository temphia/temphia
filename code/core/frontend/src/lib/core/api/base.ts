import axios, { AxiosResponse, AxiosInstance, AxiosRequestConfig } from "axios";

interface BaseOptions {
  url: string;
  user_token: string;
  path: string[];
  service_opts?: any;
}

export class ApiBase {
  _user_token: string;
  _session_token: string;
  _http: AxiosInstance;
  _api_base_url: string;
  _raw_http: AxiosInstance;

  _service_options: object;
  _service_path: string[];
  _service_resp_payload: any;

  constructor(opts: BaseOptions) {
    this._user_token = opts.user_token;
    this._api_base_url = opts.url;
    this._service_options = opts.service_opts;
    this._session_token = "";
    this._http = null;
    this._service_path = opts.path;

    this.intercept_request = this.intercept_request.bind(this);
    this.intercept_request_err = this.intercept_request_err.bind(this);
    this._raw_http = axios.create({
      baseURL: opts.url,
    });
  }

  async init() {
    let resp = await this.refresh_token();
    this._service_resp_payload = resp.data["service_payload"] || null;
    this._session_token = resp.data.token;

    this._http = axios.create({
      headers: { Authorization: this._session_token },
      baseURL: this._api_base_url,
    });

    this._http.interceptors.request.use(
      this.intercept_request,
      this.intercept_request_err
    );
  }

  async refresh_token() {
    return this._raw_http.post(`/auth/refresh`, {
      user_token: this._user_token,
      options: this._service_options,
      path: this._service_path,
    });
  }

  intercept_request(config: AxiosRequestConfig) {
    return config;
  }

  intercept_request_err(error: any) {
    // fixme => if error is 401, refresh the token
    return Promise.reject(error);
  }

  get<T = any, R = AxiosResponse<T>>(
    url: string,
    config?: AxiosRequestConfig
  ): Promise<R> {
    return this._http.get(url, config);
  }

  post<T = any, R = AxiosResponse<T>>(
    url: string,
    data?: any,
    config?: AxiosRequestConfig
  ): Promise<R> {
    return this._http.post(url, data, config);
  }

  put<T = any, R = AxiosResponse<T>>(
    url: string,
    data?: any,
    config?: AxiosRequestConfig
  ): Promise<R> {
    return this._http.put(url, data, config);
  }

  patch<T = any, R = AxiosResponse<T>>(
    url: string,
    data?: any,
    config?: AxiosRequestConfig
  ): Promise<R> {
    return this._http.patch(url, data, config);
  }

  delete<T = any, R = AxiosResponse<T>>(
    url: string,
    config?: AxiosRequestConfig
  ): Promise<R> {
    return this._http.delete(url, config);
  }
}
