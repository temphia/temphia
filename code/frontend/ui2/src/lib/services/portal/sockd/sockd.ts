import { Sockd, SockdMessage } from "../../../../lib/sockd";

export class SockdService {
  constructor() {}

  build = async (
    url: string,
    on_handler: (message: SockdMessage) => void
  ): Promise<Sockd> => {
    const sockd = new Sockd(url);
    sockd.SetHandler(on_handler);
    await sockd.Init();
    return sockd;
  };
}
