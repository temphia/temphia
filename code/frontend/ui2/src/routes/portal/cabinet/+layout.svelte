<script>
    import { params } from "$lib/params";
    $: _paths = ($params["folder"] || "").split("/");
    
</script>

<div class="h-full p-0 md:p-2">
    <div class="card p-2 variant-glass-primary h-full">
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

                <div>
                    <button type="button" class="btn variant-filled-primary">
                        <span>(icon)</span>
                        <span>Button</span>
                    </button>
                </div>
            </div>
        </header>
        <section class="p-2">
            <slot />
        </section>
    </div>
</div>
