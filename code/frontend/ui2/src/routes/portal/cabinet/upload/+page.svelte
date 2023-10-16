<script lang="ts">
    import { getContext } from "svelte";
    import type { PortalService } from "$lib/core";
    import { params } from "$lib/params";
    import Fpreview from "./_fpreview.svelte";

    export let source = $params["source"] || "default";
    let folder = $params["folder"];

    const app = getContext("__app__") as PortalService;
    const cservice = app.get_cabinet_service();
    const capi = cservice.get_source_api(source);

    let filename = "";
    let file;

    const fileSelect = (ev) => {
        console.log(ev);
        file = ev.target.files[0];
        filename = file.name;
        console.log(file);
    };

    const upload = async () => {
        const formdata = new FormData();
        formdata.append("file", file);
        capi.uploadFile(folder, filename, formdata);
    };
</script>

<form class="mt-4 space-y-3" on:submit|preventDefault={upload}>
    <div class="grid grid-cols-1 space-y-2">
        <span class="text-sm font-bold text-gray-500 tracking-wide"
            >Attach Document</span
        >
        <div class="flex items-center justify-center w-full">
            <label
                class="flex flex-col rounded-lg border-4 border-dashed w-full h-60 p-5 group text-center"
            >
                <input type="file" class="hidden" on:change={fileSelect} />

                <Fpreview {file} {filename} />
            </label>
        </div>
    </div>
    <p class="text-sm text-gray-300">
        <span>File type: any</span>
    </p>
    <div class="flex justify-end">
        {#if file}
            <button class="btn variant-filled-primary"> Upload </button>
        {/if}
    </div>
</form>
