<script lang="ts">
    import { getContext } from "svelte";
    import type { PortalService } from "$lib/core";
    import { params } from "$lib/params";
    import Icon from "@krowten/svelte-heroicons/Icon.svelte";
    import NewFolder from "../_panels/new_folder.svelte";
    import { current } from "./select";

    $: _paths = ($params["folder"] || "").split("/");

    const app: PortalService = getContext("__app__");

    const cservice = app.get_cabinet_service();
    const capi = cservice.get_source_api($params["source"]);

    $: __epoch = 1;

    const complete_new_folder = async (name) => {
        await capi.newFolder(($params["folder"] || "") + "/" + name);
        __epoch = __epoch + 1;
        app.utils.small_modal_close();
    };
</script>

<div class="h-full p-0 md:p-2">
    <div class="card p-2 variant-glass-secondary h-full">
        <header class="card-header">
            <div class="flex justify-between">
                <ol class="breadcrumb">
                    <li class="crumb">
                        <a
                            class="anchor"
                            href="/z/pages/portal/cabinet/listings">Home</a
                        >
                    </li>

                    {#each _paths as path, i}
                        <li class="crumb-separator" aria-hidden>/</li>

                        <li class="crumb">
                            <a
                                class="anchor"
                                href="/z/pages/portal/cabinet/listings/folder?source={$params[
                                    'source'
                                ]}&folder={_paths.slice(0, i + 1).join('/')}"
                                >{path}</a
                            >
                        </li>
                    {/each}
                </ol>

                {#if !$params["file"]}
                    <div class="flex flex-row gap-1">
                        {#if $params["folder"]}
                            <button
                                class="btn btn-sm variant-filled-primary"
                                on:click={() =>
                                    app.nav.cab_uploader(
                                        $params["source"],
                                        $params["folder"]
                                    )}
                            >
                                <Icon name="cloud-upload" class="h-4 w4" />
                                <span class="hidden md:inline">Upload</span>
                            </button>
                        {/if}

                        <button
                            class="btn btn-sm variant-filled-primary"
                            on:click={() => {
                                app.utils.small_modal_open(NewFolder, {
                                    onNewName: complete_new_folder,
                                });
                            }}
                        >
                            <Icon name="folder-add" class="h-4 w4" />
                            <span class="hidden md:inline">New Folder</span>
                        </button>

                        {#if $current.item && $current.folder}
                            <button
                                class="btn btn-sm variant-filled-primary"
                                on:click={() => {}}
                            >
                                <Icon name="trash" class="h-4 w4" />
                                <span class="hidden md:inline">Delete</span>
                            </button>
                        {/if}
                    </div>
                {/if}
            </div>
        </header>
        <section class="p-2">
            {#key __epoch}
                <slot />
            {/key}
        </section>
    </div>
</div>
