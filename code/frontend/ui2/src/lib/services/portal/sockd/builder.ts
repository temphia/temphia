import type { SockdMessage } from "./stypes";
import { Sockd } from "./sockd"


export class SockdBuilder {
  constructor() { }

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
