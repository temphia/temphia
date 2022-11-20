import type { ApiBase } from "./base";
import { Http } from "./http";

export class CabinetAPI {
  base: ApiBase;
  source: string
  constructor(source: string, base: ApiBase) {
    this.source = source
    this.base = base;
  }

  listRoot() {
    return this.base.get(`/cabinet/${this.source}/`);
  }


  listFolder(folder: string) {
    return this.base.get(`/cabinet/${this.source}/${folder}`);
  }

  newFolder(folder: string, data: any) {
    return this.base.post(`/cabinet/${this.source}/${folder}`, data);
  }

  getFile(folder: string, fname: string) {
    return this.base.get(`/cabinet/${this.source}/${folder}/file/${fname}`);
  }

  uploadFile(folder: string, fname: string, data: any) {
    return this.base.postForm(`/cabinet/${this.source}/${folder}/file/${fname}`, true, data);
  }

  deleteFile(folder: string, fname: string) {
    return this.base.delete(`/cabinet/${this.source}/${folder}/file/${fname}`);
  }

  getFilePreview(folder: string, fname: string) {
    return this.base.get(`/cabinet/${this.source}/${folder}/preview/${fname}`);
  }
}

export class FolderTktAPI {
  http: Http;
  ticket: string;
  constructor(baseUrl: string, token: string) {
    this.http = new Http(baseUrl, {});
    this.ticket = token;
  }

  //  /folder/:ticket/

  list() {
    return this.http.get(`/folder/${this.ticket}`);
  }

  getFile(file: string) {
    return this.http.get(`/folder/${this.ticket}/${file}`);
  }

  getFilePreview(file: string) {
    return this.http.get(`/folder/${this.ticket}/${file}/preview`);
  }

  uploadFile(file: string, data: any) {
    return this.http.post(`/folder/${this.ticket}/${file}`, data);
  }

  deleteFile(file: string) {
    return this.http.delete(`/folder/${this.ticket}/${file}`);
  }
}
