export class FolderTktAPI {
  url: string;
  constructor(base_url: string, ticket: string) {
    this.url = `${base_url}/tkt/folder/${ticket}`;
  }

  async list() {
    const resp = await fetch(this.url);
    return resp.json();
  }

  async upload_file(file: string, data?: any) {
    const resp = await fetch(`${this.url}/${file}`, {
      method: "POST",
      body: data,
    });
    return resp.json();
  }

  async del_file(file: string, proof?: string) {
    fetch(`${this.url}/${file}?proof=${proof}`, {
      method: "DELETE",
    });
  }

  get_file_link(file: string): string {
    return `${this.url}/${file}`;
  }

  get_file_preview_link(file: string): string {
    return `${this.url}/${file}/preview`;
  }
}
