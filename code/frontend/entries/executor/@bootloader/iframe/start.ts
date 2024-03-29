import { sleep } from "yootils";
import type { LoaderOptions } from "../../../../lib/engine/plug";
import { initRegistry, plugStart } from "../../../../lib/engine/putils";
import { generateId } from "../../../../lib/utils";
import { Env } from "../../../portal/launcher/env";
import { IFramePipe } from "./iframe_pipe";
import { fakeStorage } from "./fakestorage";

export default () => {
  console.log("Iframe Exec start here..");

  try {
    if (!window["localStorage"]) {
      window["localStorage"] = fakeStorage();
    } else {
      // if we do not have sandbox and browser donot support
      //   credentialless/anonymous iframes #hacky_as_hell
      if (localStorage.getItem("_temphia_authed_key_")) {
        delete window["localStorage"];
        window["localStorage"] = fakeStorage();
      }
    }
  } catch (error) {
    console.log("@execption", error);
  }

  initRegistry();

  let transfered_port: MessagePort;

  const handle_port_transfer = (ev) => {
    if (ev.data !== "port_transfer") {
      console.log("wrong event listener", ev);
      return;
    }

    transfered_port = ev.ports[0];
    console.log("@received_port_@guest", transfered_port);
    window.removeEventListener("message", handle_port_transfer);
    env_init(null);
  };

  const env_init = async (ev) => {
    const opts = window["__loader_options__"] as LoaderOptions;
    if (!opts) {
      console.log("Loader Options not found");
      return;
    }

    console.log("iframe portal opts @=>", opts);

    const pipe = new IFramePipe(opts.parent_secret, transfered_port);

    const target = document.getElementById("plugroot");

    const env = new Env({
      agent: opts.agent,
      plug: opts.plug,
      token: opts.token,
      api_base_url: opts.api_base_url,
      parent_secret: opts.parent_secret,
      pipe: pipe,
      registry: window["__registry__"],
      tenant_id: opts.tenant_id,
      target: target,
      startup_payload: opts.startup_payload,
    });

    await env.init();

    pipe.send(generateId(), "env_loaded", {});

    plugStart({
      plug: opts.plug,
      agent: opts.agent,
      entry: opts.entry,
      env: env,
      target: target,
      exec_loader: opts.exec_loader,
    });
  };

  window.addEventListener("message", handle_port_transfer, false);
};
