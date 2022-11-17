<script lang="ts">
  export let secret_id = "";
  export let name = "";
  export let invoker = undefined;

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
    iframe.contentWindow.postMessage("_temphia_plug_load_", "*", [
      channel.port2,
    ]);
  }}
  title={name}
  class="w-full h-full transition-all"
  src="https://picsum.photos/seed/{secret_id}/1300/700"
/>
