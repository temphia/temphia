import { Sockd } from "../sockd";

export class SockdAPI {
  base_url: string;
  constructor(base_url: string) {
    this.base_url = base_url;
  }

  async user() {
    const sockd = new Sockd({
      OnHandler: null,
      URL: `${this.base_url}/sockd/user/ws`,
    });

    await sockd.init();
    return sockd;
  }

  async data() {
    const sockd = new Sockd({
      OnHandler: null,
      URL: `${this.base_url}/sockd/data/ws`,
    });

    await sockd.init();
    return sockd;
  }

  async room() {
    const sockd = new Sockd({
      OnHandler: null,
      URL: `${this.base_url}/sockd/room/ws`,
    });

    await sockd.init();
    return sockd;
  }
  async dev() {
    const sockd = new Sockd({
      OnHandler: null,
      URL: `${this.base_url}/sockd/dev/ws`,
    });

    await sockd.init();
    return sockd;
  }
}
