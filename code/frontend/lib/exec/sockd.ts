import { Sockd, SockdMessage } from "../sockd";

export const NewSockdRoom = async (url: string, callback: (msg: SockdMessage) => void) => {
  const sockd = new Sockd(url);
  sockd.SetHandler(callback)
  
  await sockd.Init();

  return sockd;
};
