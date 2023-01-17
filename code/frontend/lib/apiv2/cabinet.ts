import type { ApiBase } from "./base";
import { Http } from "./http";

export class CabinetAPI {
  base: ApiBase;
  source: string;
  constructor(source: string, base: ApiBase) {
    this.source = source;
    this.base = base;
  }

  listRoot() {
    return this.base.get(`/cabinet/${this.source}/`);
  }

  listFolder(folder: string) {
    return this.base.get(`/cabinet/${this.source}/${folder}`);
  }

  newFolder(folder: string) {
    return this.base.post(`/cabinet/${this.source}/${folder}`, {});
  }

  getFile(folder: string, fname: string) {
    return this.base.get(`/cabinet/${this.source}/${folder}/file/${fname}`);
  }

  uploadFile(folder: string, fname: string, data: any) {
    return this.base.postForm(
      `/cabinet/${this.source}/${folder}/file/${fname}`,
      true,
      data
    );
  }

  deleteFile(folder: string, fname: string) {
    return this.base.delete(`/cabinet/${this.source}/${folder}/file/${fname}`);
  }

  getFilePreview(folder: string, fname: string) {
    return `${this.base.base_url}/cabinet/${this.source}/${folder}/preview/${fname}?token=${this.base.user_token}`;
  }
}

export class FolderTktAPI {
  http: Http;
  ticket: string;
  base_url: string;
  constructor(baseUrl: string, token: string) {
    this.http = new Http(baseUrl, {});
    this.ticket = token;
    this.base_url = baseUrl;
  }

  //  /folder/:ticket/

  list() {
    return this.http.get(`/folder/${this.ticket}`);
  }

  getFile(file: string) {
    return this.http.get(`/folder/${this.ticket}/${file}`);
  }

  getFileUrl(file: string) {
    return `${this.base_url}/folder/${this.ticket}/${file}`;
  }

  getFilePreviewUrl(file: string) {
    return `${this.base_url}/folder/${this.ticket}/${file}/preview`;
  }

  uploadFile(file: string, data: any) {
    return this.http.post(`/folder/${this.ticket}/${file}`, data);
  }

  // downgraded_ticket() {}

  deleteFile(file: string) {
    return this.http.delete(`/folder/${this.ticket}/${file}`);
  }
}
