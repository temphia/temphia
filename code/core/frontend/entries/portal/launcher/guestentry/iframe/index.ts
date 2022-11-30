import type { LoaderOptions } from "../../../../../lib/engine/plug";
import { initRegistry, plugStart } from "../../../../../lib/engine/putils";
import { generateId } from "../../../../../lib/utils";
import { Env } from "../../env";
import { IFramePipe } from "./iframe_pipe";

console.log("init registry");
initRegistry();

let transfered_port: MessagePort;

const handle_port_transfer = (ev: MessageEvent<any>) => {
  transfered_port = ev.ports[0];
  window.removeEventListener("message", handle_port_transfer);
};

const env_init = async (ev: MessageEvent<any>) => {
  const opts = window["__loader_options__"] as LoaderOptions;
  if (!opts) {
    console.log("Loader Options not found");
    return;
  }

  console.log("iframe portal opts @=>", opts);

  const pipe = new IFramePipe(opts.parent_secret, transfered_port);

  const env = new Env({
    agent: opts.agent,
    plug: opts.plug,
    token: opts.token,
    base_url: opts.base_url,
    parent_secret: opts.parent_secret,
    pipe: pipe,
    registry: window["__registry__"],
  });

  await env.init();

  pipe.send(generateId(), "env_loaded", {});

  plugStart({
    plug: opts.plug,
    agent: opts.agent,
    entry: opts.entry,
    env: env,
    target: document.getElementById("plugroot"),
    exec_loader: opts.exec_loader,
    payload: null,
  });
};

window.addEventListener("message", handle_port_transfer, false);
window.addEventListener("load", env_init, false);
