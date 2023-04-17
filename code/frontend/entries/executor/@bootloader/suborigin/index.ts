import { initRegistry, plugStart } from "../../../../lib/engine/putils";
import { generateId } from "../../../../lib/utils";
import { Env } from "../../../portal/launcher/env";

(() => {
  console.log("Iframe Exec start..");

  initRegistry();

  let transfered_port: MessagePort;
  let launch_options: {
    token: string;
    parent_secret: string;
  };

  const handle_port_transfer = (ev) => {
    if (ev.data !== "port_transfer") {
      console.log("wrong event listener", ev);
      return;
    }

    transfered_port = ev.ports[0];

    console.log("@received_port_@guest", transfered_port);
    window.removeEventListener("message", handle_port_transfer);
    env_init();
  };

  const env_init = async () => {
    const opts = window["__loader_options__"] as {
      api_base_url: string;
      plug_id: string;
      agent_id: string;
      entry_name: string;
      exec_loader: string;
      tenant_id: string;
    };
    if (!opts) {
      console.log("Loader Options not found");
      return;
    }

    console.log("iframe portal opts @=>", opts);

    const target = document.getElementById("plugroot")

    const env = new Env({
      agent: opts.agent_id,
      plug: opts.plug_id,
      token: "",
      api_base_url: opts.api_base_url,
      parent_secret: "",
      pipe: null,
      registry: window["__registry__"],
      tenant_id: opts.tenant_id,
      target,
      startup_payload: opts["startup_payload"] || {}
    });

    await env.init();

    //   pipe.send(generateId(), "env_loaded", {});

    plugStart({
      plug: opts.plug_id,
      agent: opts.agent_id,
      entry: opts.entry_name,
      env: env,
      target,
      exec_loader: opts.exec_loader,
    });
  };

  window.addEventListener("message", handle_port_transfer, false);
})();
