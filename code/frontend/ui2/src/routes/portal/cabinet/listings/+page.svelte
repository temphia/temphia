<script lang="ts">
    import FolderView from "../_panels/FolderView.svelte";

    import { getContext } from "svelte";
    import type { PortalService } from "$lib/core";
    import { LoadingSpinner } from "$lib/compo";
    import { current, set } from "./select";

    export let source = "default";

    const app: PortalService = getContext("__app__");
    const cservice = app.get_cabinet_service();

    let data = [];
    let loading = true;
    const load = async () => {
        const capi = cservice.get_source_api(source);
        const resp = await capi.listRoot();
        if (!resp.ok) {
            return;
        }

        data = resp.data.map((element) => ({
            name: element,
            is_dir: true,
            size: "",
            last_modified: "",
        }));

        loading = false;
    };

    load();
</script>

{#if loading}
    <LoadingSpinner />
{:else}
    <FolderView
        selected={$current.item}
        files={data}
        on:open_item={(ev) => {
            app.nav.cab_folder(source, ev.detail["name"]);
        }}
        on:select_item={(ev) => {
            set("", ev.detail["name"]);
        }}
    />
{/if}
