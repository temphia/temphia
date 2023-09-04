<script lang="ts">
    import { getContext } from "svelte";
    import type { PortalService } from "$lib/core";
    import { params } from "$lib/params";
    import Layout from "./_layout.svelte";
    import { EFrame } from "$lib/compo/eframe";

    let data = $params;

    const pid = data["pid"];
    const aid = data["aid"];
    const app = getContext("__app__") as PortalService;
    const api = app.api_manager.get_engine_api();

    let before = true;
    let url = "";
    let token = "";

    const run = async () => {
        const resp = await api.launch_agent({
            plug_id: pid,
            agent_id: aid,
        });
        if (!resp.ok) {
            return;
        }

        url = `http://${resp.data["domain"]}:${
            location.port || 80
        }/z/pages/agent/inject`;
        token = resp.data["token"];
        before = false;
    };
</script>

<Layout {run} bind:before>
    <EFrame
        exec_data={{
            agent_id: aid,
            plug_id: pid,
            exec_token: token,
            tenant_id: "",
            startup_payload: null,
        }}
        {url}
        name="App"
    />
</Layout>
