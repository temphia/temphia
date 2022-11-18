<script lang="ts">
  import type { Instance } from "../../../services/engine/launcher";
  import { iframeTemplateBuild } from "../../../../../lib/engine/template";

  export let options: Instance;
  export let agent: object;

  let iframe: HTMLIFrameElement | null;
  const channel = new MessageChannel();

  const onmessage = (ev) => {
    console.log("@message from iframe", ev);
    // fixme => pass to invoker
  };

  channel.port1.onmessage = onmessage;
</script>

<iframe
  bind:this={iframe}
  on:load={(ev) => {
    iframe.contentWindow.postMessage("port_transfer", "*", [channel.port2]);
  }}
  title={options.name}
  class="w-full h-full transition-all"
  src="https://picsum.photos/seed/{options.id}/1300/700"
/>

<!-- 


    id
name
type
executor
iface_file
entry_file
web_entry
web_script
web_style
web_loader
web_files
env_vars
extra_meta
mod_version
plug_id
tenant_id



 -->
