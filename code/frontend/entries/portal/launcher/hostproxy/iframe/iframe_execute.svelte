<script lang="ts">
  import { onMount } from "svelte";
  import { iframeTemplateBuild } from "../../../../../lib/engine/template";
  import type { ExecInstanceOptions } from "../../../services/engine/exec_type";

  export let name: string;
  export let exec_data: ExecInstanceOptions;
  export let secret_id: string;

  let iframe: HTMLIFrameElement | null;
  const channel = new MessageChannel();

  const src = iframeTemplateBuild({
    plug: exec_data["plug"],
    agent: exec_data["agent"],
    base_url: exec_data["base_url"],
    entry_name: exec_data["entry"],
    exec_loader: exec_data["exec_loader"],
    js_plug_script: exec_data["js_plug_script"],
    style_file: exec_data["style"],
    token: exec_data["token"] || "",
    ext_scripts: exec_data["ext_scripts"] || {},
    parent_secret: secret_id,
    startup_payload: {},
  });

  const onmessage = (ev) => {
    console.log("@message from iframe", ev);
    // fixme => pass to invoker
  };

  channel.port1.onmessage = onmessage;
</script>

<iframe
  bind:this={iframe}
  on:load={(ev) => {
    console.log("@port_transfer_from_HOST");
    iframe.contentWindow.postMessage("port_transfer", "*", [channel.port2]);
  }}
  title={name}
  class="w-full h-full transition-all"
  srcdoc={src}
/>
