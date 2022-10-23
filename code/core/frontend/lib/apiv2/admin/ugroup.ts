import type { ApiBase } from "../base";

export class AdminUserGroupAPI {
  base: ApiBase;
  constructor(base: ApiBase) {
    this.base = base;
  }

  list() {
    return this.base.get("/admin/ugroup/");
  }
  new(data: any) {
    return this.base.post("/admin/ugroup/", data);
  }

  get(ugroup: string) {
    return this.base.get(`/admin/ugroup/${ugroup}`);
  }

  update(ugroup: string, data: any) {
    return this.base.post(`/admin/ugroup/${ugroup}`, data);
  }

  delete(ugroup: string) {
    return this.base.delete(`/admin/ugroup/${ugroup}`);
  }

  // data

  listData(ugroup: string) {
    return this.base.get(`/admin/ugroup/${ugroup}/data`);
  }

  newData(ugroup: string, data: any) {
    return this.base.post(`/admin/ugroup/${ugroup}/data`, data);
  }
  getData(ugroup: string, id: string) {
    return this.base.get(`/admin/ugroup/${ugroup}/data/${id}`);
  }

  updateData(ugroup: string, id: string, data: any) {
    return this.base.post(`/admin/ugroup/${ugroup}/data/${id}`, data);
  }

  deleteData(ugroup: string, id: string) {
    return this.base.delete(`/admin/ugroup/${ugroup}/data/${id}`);
  }

  // plug

  listPlug(ugroup: string) {
    return this.base.get(`/admin/ugroup/${ugroup}/plug`);
  }
  newPlug(ugroup: string, data: any) {
    return this.base.post(`/admin/ugroup/${ugroup}/plug`, data);
  }
  getPlug(ugroup: string, id: string) {
    return this.base.get(`/admin/ugroup/${ugroup}/plug/${id}`);
  }
  updatePlug(ugroup: string, id: string,data: any) {
    return this.base.post(`/admin/ugroup/${ugroup}/plug/${id}`, data);
  }
  deletePlug(ugroup: string, id: string) {
    return this.base.delete(`/admin/ugroup/${ugroup}/plug/${id}`);
  }

  // auth
  listAuth(ugroup: string) {
    return this.base.get(`/admin/ugroup/${ugroup}/auth`);
  }
  newAuth(ugroup: string, data: any) {
    return this.base.post(`/admin/ugroup/${ugroup}/auth`, data);
  }
  getAuth(ugroup: string, id: string) {
    return this.base.get(`/admin/ugroup/${ugroup}/auth/${id}`);
  }
  updateAuth(ugroup: string, id: string, data: any) {
    return this.base.post(`/admin/ugroup/${ugroup}/auth/${id}`, data);
  }
  deleteAuth(ugroup: string, id: string) {
    return this.base.delete(`/admin/ugroup/${ugroup}/auth/${id}`);
  }
}