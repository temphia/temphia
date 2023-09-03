<script lang="ts">
    import { getContext } from "svelte";
    import { LoadingSpinner, PortalService } from "$lib/core";
    import Wizard from "./_wizard.svelte";

    export let data;

    let bid = data["slug"];

    const app = getContext("__app__") as PortalService;
    const bapi = app.api_manager.get_admin_bprint_api();

    let loading = true;
    let appjson = {};
    let bprintjson = {};

    const load = async () => {
        const resp1 = await bapi.get(bid);
        if (!resp1.ok) {
            console.log("err", resp1);
            return;
        }

        bprintjson = resp1.data;

        if (resp1.data["type"] !== "app") {
            console.log("not implemented", resp1.data);
            return;
        }

        const resp = await bapi.get_file(bid, "app.json");
        if (!resp.ok) {
            console.log("Err", resp);
            return;
        }

        appjson = resp.data;
        loading = false;
    };

    load();
</script>

<div class="p-4">
    {#if loading}
        <LoadingSpinner />
    {:else}
        <Wizard {bid} {appjson} {bprintjson} />
    {/if}
</div>
