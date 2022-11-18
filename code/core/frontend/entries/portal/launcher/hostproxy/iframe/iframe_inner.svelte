<script lang="ts">
  import type { Instance } from "../../../services/engine/launcher";

  export let options: Instance;
  export let source: string;

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
  srcdoc={source}
/>
