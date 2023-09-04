<script lang="ts">
    import { getContext } from "svelte";
    import type { PortalService } from "$lib/core";
    import { LoadingSpinner } from "$lib/core";
    import { params } from "$lib/params";
    import Importer from "../panels/importer/index.svelte";

    const app = getContext("__app__") as PortalService;
    let rdata;

    const load = async () => {
        const rapi = app.api_manager.get_repo_api();
        const resp = await rapi.getBprint($params["source"], "", $params["item"]);
        if (!resp.ok) {
            return;
        }
        rdata = resp.data;
    };

    load();
</script>

{#if rdata}
    <Importer data={rdata} source={$params["source"]} />
{:else}
    <LoadingSpinner />
{/if}
