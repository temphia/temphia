<script lang="ts">
    import { createEventDispatcher } from "svelte";
    import type { ExecData } from "./eframe";

    export let name: string;
    export let exec_data: ExecData;
    export let iframe: HTMLIFrameElement = null;
    export let url: string;
    export let chan = new MessageChannel();

    const dispatcher = createEventDispatcher();

    const onFrameMessage = (ev) => {
        console.log("@onFramwData", ev)

        const data = JSON.parse(ev.data);
        if (data["mtype"] !== "get_exec_data") {
            dispatcher("eframe_message", data);
            return;
        }

        console.log("sending_exec_data")

        chan.port1.postMessage(
            JSON.stringify({
                mtype: "exec_data",
                data: exec_data,
            }),
            []
        );
    };

    chan.port1.onmessage = onFrameMessage;
</script>

<iframe
    bind:this={iframe}
    on:load={(ev) => {
        iframe.contentWindow.postMessage("transfer_port", "*", [chan.port2]);
    }}
    src={url}
    title={name}
    class="border-green-200 w-full h-full transition-all"
    allow="accelerometer; ambient-light-sensor; autoplay; battery; camera; clipboard-write; document-domain; encrypted-media; fullscreen; geolocation; gyroscope; layout-animations; legacy-image-formats; magnetometer; microphone; midi; oversized-images; payment; picture-in-picture; publickey-credentials-get; sync-xhr; usb; vr ; wake-lock; xr-spatial-tracking"
    sandbox="allow-forms allow-modals allow-popups allow-popups-to-escape-sandbox allow-same-origin allow-scripts allow-downloads allow-storage-access-by-user-activation"
/>
