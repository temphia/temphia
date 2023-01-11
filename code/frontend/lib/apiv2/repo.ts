import type { ApiBase } from "./base";

// fixme => change this to StoreAPI

export class RepoAPI {
  base: ApiBase;
  constructor(base: ApiBase) {
    this.base = base;
  }

  list() {
    return this.base.get("/repo/");
  }

  load(id: string) {
    return this.base.get(`/repo/${id}`);
  }

  getBprint(id: string, group: string, slug: string) {
    return this.base.get(`/repo/${id}/${group}/${slug}`);
  }

  getBprintFile(id: string, group: string, slug: string, file: string) {
    return this.base.get(`/repo/${id}/${group}/${slug}/${file}`);
  }
}