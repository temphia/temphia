import { Sockd, SockdMessage } from "../sockd";

export class SockdAPI {
  base_url: string;
  constructor(base_url: string) {
    this.base_url = base_url;
  }

  async user(token: string, on_handler: (message: SockdMessage) => void) {
    const sockd = new Sockd({
      OnHandler: on_handler,
      URL: `${this.base_url}/sockd/user/ws?token=${token}`,
    });

    await sockd.init();
    return sockd;
  }

  async data(token: string, on_handler: (message: SockdMessage) => void) {
    const sockd = new Sockd({
      OnHandler: on_handler,
      URL: `${this.base_url}/sockd/data/ws?token=${token}`,
    });

    await sockd.init();
    return sockd;
  }

  async room(token: string, on_handler: (message: SockdMessage) => void) {
    const sockd = new Sockd({
      OnHandler: on_handler,
      URL: `${this.base_url}/sockd/room/ws?token=${token}`,
    });

    await sockd.init();
    return sockd;
  }

  async dev(token: string, on_handler: (message: SockdMessage) => void) {
    const sockd = new Sockd({
      OnHandler: on_handler,
      URL: `${this.base_url}/sockd/dev/ws?token=${token}`,
    });

    await sockd.init();
    return sockd;
  }
}
