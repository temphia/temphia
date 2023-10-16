<script lang="ts">
    import FolderView from "../../_panels/FolderView.svelte";

    import { getContext } from "svelte";
    import type { PortalService } from "$lib/core";
    import { LoadingSpinner } from "$lib/compo";
    import { params } from "$lib/params";

    export let source = $params["source"] || "default";

    let folder;

    const app: PortalService = getContext("__app__");
    const cservice = app.get_cabinet_service();

    let data = [];
    let loading = true;
    const load = async (_folder) => {
        folder = _folder;
        const capi = cservice.get_source_api(source);
        const resp = await capi.listFolder(_folder);
        if (!resp.ok) {
            return;
        }

        data = resp.data;
        loading = false;
    };

    $: load($params["folder"]);
</script>

{#if loading}
    <LoadingSpinner />
{:else}
    <FolderView
        files={data}
        on:open_item={(ev) => {
            console.log("@EEEE", ev.detail);
            if (ev.detail["is_dir"]) {
                app.nav.cab_folder(source, `${folder}/${ev.detail["name"]}`);
            } else {
                app.nav.cab_file(source, folder, ev.detail["name"]);
            }
        }}
    />
{/if}
