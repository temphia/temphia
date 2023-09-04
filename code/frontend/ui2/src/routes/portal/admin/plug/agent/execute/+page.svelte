<script lang="ts">
    import { getContext } from "svelte";
    import type { PortalService } from "$lib/core";
    import { params } from "$lib/params";
    import Layout from "./_layout.svelte";

    let data = $params;

    const pid = data["pid"];
    const aid = data["aid"];
    const app = getContext("__app__") as PortalService;
    const api = app.api_manager.get_engine_api();

    let before = true;
    let url = "";
    let token = "";
    let iframwref;

    const run = async () => {
        const resp = await api.launch_agent({
            plug_id: pid,
            agent_id: aid,
        });
        if (!resp.ok) {
            return;
        }

        url = `http://${resp.data["domain"]}:${location.port}/z/pages/agent/inject`;
        token = resp.data["token"];
        before = false;
    };
</script>

<Layout {run} bind:before>
    <iframe
        bind:this={iframwref}
        src={url}
        on:load={() => {
            iframwref.contentWindow.postMessage("port_transfer", "*");
        }}
        title="App"
        class="border-green-200 w-full h-full transition-all"
        allow="accelerometer; ambient-light-sensor; autoplay; battery; camera; clipboard-write; document-domain; encrypted-media; fullscreen; geolocation; gyroscope; layout-animations; legacy-image-formats; magnetometer; microphone; midi; oversized-images; payment; picture-in-picture; publickey-credentials-get; sync-xhr; usb; vr ; wake-lock; xr-spatial-tracking"
        sandbox="allow-forms allow-modals allow-popups allow-popups-to-escape-sandbox allow-same-origin allow-scripts allow-downloads allow-storage-access-by-user-activation"
    />
</Layout>
