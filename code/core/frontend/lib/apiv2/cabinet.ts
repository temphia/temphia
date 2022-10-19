import { Http } from "./http";

export class CabinetAPI {
  http: Http;
  constructor(baseUrl: string, token: string) {
    this.http = new Http(baseUrl, {
      Authorization: token,
    });
  }

  listRoot() {
    return this.http.get("/cabinet/");
  }

  listFolder(folder: string) {
    return this.http.get(`/cabinet/${folder}`);
  }

  newFolder(folder: string, data: any) {
    return this.http.post(`/cabinet/${folder}`, data);
  }

  getFile(folder: string, fname: string) {
    return this.http.get(`/cabinet/${folder}/file/${fname}`);
  }

  uploadFile(folder: string, fname: string, data: any) {
    return this.http.postForm(`/cabinet/${folder}/file/${fname}`, true, data);
  }

  deleteFile(folder: string, fname: string) {
    return this.http.delete(`/cabinet/${folder}/file/${fname}`);
  }

  getFilePreview(folder: string, fname: string) {
    return this.http.get(`/cabinet/${folder}/preview/${fname}`);
  }
}

export class FolderTktAPI {
  http: Http;
  ticket: string;
  constructor(baseUrl: string, token: string) {
    this.http = new Http(baseUrl, {});
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