<script>
    import { params } from "$lib/params";
    import Icon from "@krowten/svelte-heroicons/Icon.svelte";
    $: _paths = ($params["folder"] || "").split("/");
</script>

<div class="h-full p-0 md:p-2">
    <div class="card p-2 variant-glass-secondary h-full">
        <header class="card-header">
            <div class="flex justify-between">
                <ol class="breadcrumb">
                    <li class="crumb">
                        <a class="anchor" href="/z/pages/portal/cabinet">Home</a
                        >
                    </li>

                    {#each _paths as path}
                        <li class="crumb-separator" aria-hidden>/</li>

                        <li class="crumb">
                            <a class="anchor" href="#">{path}</a>
                        </li>
                    {/each}
                </ol>

                {#if !$params["file"]}
                    <div class="flex flex-row gap-1">
                        <button class="btn btn-sm variant-filled-primary">
                            <Icon name="cloud-upload" class="h-4 w4" />
                            <span class="hidden md:inline">Upload</span>
                        </button>
                        <button class="btn btn-sm variant-filled-primary">
                            <Icon name="folder-add" class="h-4 w4" />
                            <span class="hidden md:inline">New Folder</span>
                        </button>
                    </div>
                {/if}
            </div>
        </header>
        <section class="p-2">
            <slot />
        </section>
    </div>
</div>
