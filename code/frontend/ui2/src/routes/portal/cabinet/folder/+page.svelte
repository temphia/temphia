<script lang="ts">
    import FolderView from "../_panels/FolderView.svelte";

    import { getContext } from "svelte";
    import type { PortalService } from "$lib/core";
    import { LoadingSpinner } from "$lib/compo";
    import { params} from "$lib/params"

    export let source = $params["source"] || "default";

    let folder = $params["folder"]

    const app: PortalService = getContext("__app__");
    const cservice = app.get_cabinet_service();

    let data = [];
    let loading = true;
    const load = async () => {
        const capi = cservice.get_source_api(source);
        const resp = await capi.listFolder(folder)
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
        files={data}
        on:open_item={(ev) => {
            app.nav.cab_folder(source, ev.detail["name"]);
        }}
    />
{/if}
