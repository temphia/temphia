import { IFramePipe } from "../../service/engine/pipe";
import type { LoaderOptions } from "../../core/engine/ecore";
import { initRegistry, plugStart } from "../../core/engine/registry";
import { Env } from "../../core/engine/env";

console.log("init registry");
initRegistry();

window.addEventListener(
  "load",
  async () => {
    const opts = window["__loader_options__"] as LoaderOptions;
    if (!opts) {
      console.log("Loader Options not found");
      return;
    }

    console.log("iframe portal opts @=>", opts);

    const pipe = new IFramePipe(opts.parent_secret);

    const env = new Env({
      agent: opts.agent,
      plug: opts.plug,
      token: opts.token,
      base_url: opts.base_url,
      parent_secret: opts.parent_secret,
      pipe,
    });

    await env.init();

    pipe.send("", "env_loaded", {});

    plugStart({
      plug: opts.plug,
      agent: opts.agent,
      entry: opts.entry,
      env: env,
      target: document.getElementById("plugroot"),
      exec_loader: opts.exec_loader,
      payload: null,
    });
  },
  false
);
