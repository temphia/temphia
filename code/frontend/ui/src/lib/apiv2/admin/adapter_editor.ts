import { ApiBase } from "../base";

export class AdapterEditorAPI {
  base: ApiBase;
  constructor(base_url: string, tenant_id: string, token: string) {
    this.base = new ApiBase(base_url, tenant_id, token);
  }

  perform_action(name: string, data: any) {
    return this.base.post(`/admin/adapter_editor/action/${name}`, data);
  }

  self_update(data: any) {
    return this.base.post("/admin/adapter_editor/", data);
  }

  self_reset() {
    return this.base.post("/admin/adapter_editor/reset", {});
  }

  // app
  list_apps() {
    return this.base.get("/admin/adapter_editor/app");
  }

  new_app(data: any) {
    return this.base.post("/admin/adapter_editor/app", data);
  }

  get_app(id: number) {
    return this.base.get(`/admin/adapter_editor/app/${id}`);
  }

  update_app(id: number, data: any) {
    return this.base.post(`/admin/adapter_editor/app/${id}`, data);
  }

  delete_app(id: number) {
    return this.base.delete(`/admin/adapter_editor/app/${id}`);
  }

  // hook

  list_hooks() {
    return this.base.get("/admin/adapter_editor/hook");
  }

  new_hook(data: any) {
    return this.base.post("/admin/adapter_editor/hook", data);
  }

  get_hook(id: number) {
    return this.base.get(`/admin/adapter_editor/hook/${id}`);
  }

  update_hook(id: number, data: any) {
    return this.base.post(`/admin/adapter_editor/hook/${id}`, data);
  }

  delete_hook(id: number) {
    return this.base.delete(`/admin/adapter_editor/hook/${id}`);
  }
}
