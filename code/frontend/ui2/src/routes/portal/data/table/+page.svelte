<script lang="ts">
    import { getContext } from "svelte";
    import type { PortalService } from "$lib/core";
    import { LoadingSpinner } from "$lib/core";
    import { params } from "$lib/params";

    export let source = $params["source"] || $params["dsource"] || "default";
    export let group = $params["dgroup"];

    const app: PortalService = getContext("__app__");

    let loading = true;
    const load = async () => {
        const ds = await app.get_data_service();
        const gs = await ds.group_service(source, group);

        const table = gs.default_table();

        if (!table) {
            loading = false;
            return;
        }

        app.nav.data_table_render_page(source, group, table);
    };
    load();
</script>

{#if loading}
    <LoadingSpinner />
{:else}
    <div>Empty Group</div>
{/if}
