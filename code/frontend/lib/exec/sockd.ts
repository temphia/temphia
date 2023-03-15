import { Sockd } from "../sockd";

export const NewSockdRoom = async (url: string) => {
  const sockd = new Sockd(url);
  await sockd.Init();
  return sockd;
};
