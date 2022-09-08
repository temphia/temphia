export class DevTktAPI {
  url: string;
  tkt: string;
  constructor(base_url: string, ticket: string) {
    this.url = base_url;
    this.tkt = ticket;
  }

  async exec_reset(pid: string, data: any) {
    const resp = await fetch(`${this.url}/dev/exec/reset/plug/${pid}`, {
      body: JSON.stringify(data),
      method: "POST",
      headers: {"Authorization": this.tkt}
    });
    return resp.json();
  }

  async exec_run(pid: string, aid: string, method: string, data: string) {
    const resp = await fetch(
      `${this.url}/dev/exec/run/plug/${pid}/agent/${aid}/${method}`,
      {
        body: data,
        method: "POST",
        headers: {"Authorization": this.tkt}
      }
    );

    if (resp.status !== 200) {
      return {
        status: resp.status,
        data: await resp.text(),
      };
    }

    return {
      status: resp.status,
      data: await resp.json(),
    };
  }
}
