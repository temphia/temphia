import { Sockd } from "../../../lib/sockd";

export const sockd_load = async (url: string, handler: any): Promise<Sockd> => {
  const sockd = new Sockd({
    OnHandler: handler,
    URL: url,
  });

  await sockd.init();
  return sockd;
};
