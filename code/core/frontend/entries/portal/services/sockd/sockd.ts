import { Sockd, SockdMessage, ISockd } from "../../../../lib/sockd";

export class SockdService {
  constructor() {}

  build = async (
    url: string,
    on_handler: (message: SockdMessage) => void
  ): Promise<ISockd> => {
    const sockd = new Sockd({
      OnHandler: on_handler,
      URL: url,
    });

    await sockd.init();
    return sockd;
  };
}
