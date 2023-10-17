<script lang="ts">
    import Icon from "@krowten/svelte-heroicons/Icon.svelte";
    import { getContext } from "svelte";
    import { LoadingSpinner, PortalService } from "$lib/core";
    import { params } from "$lib/params";

    export let source = $params["source"] || "default";
    export let file = $params["file"];
    export let folder = $params["folder"];

    let loading = true;
    let code = "";

    const app = getContext("__app__") as PortalService;
    const cservice = app.get_cabinet_service();
    const capi = cservice.get_source_api(source);


    const load = async () => {
        const resp = await capi.getFile(folder, file);
        if (!resp.ok) {
            console.log("@err", resp);
            return;
        }

        if (typeof resp.data === "object") {
            code = JSON.stringify(resp.data, undefined, 2);
        } else {
            code = resp.data;
        }

        loading = false;
    };

    load();
</script>

{#if loading}
    <LoadingSpinner />
{:else}
<div class="p-2 h-full w-full overflow-auto flex mx-auto justify-center items-center">
    <img
        class="max-w-full p-2 bg-white rounded border"
        src={capi.getFilePreview(folder, file)}
        alt=""
        srcset=""
    />
</div>

{/if}
