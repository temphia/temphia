<script>
    import { LoadingSpinner } from "$lib/core";

    let loading = false;
    export let before = true;

    export let run;
</script>

<div class="grid place-items-center">
    {#if before}
        <div class="card mt-10">
            {#if loading}
                <LoadingSpinner />
            {:else}
                <header class="card-header">
                    Do you want to execute app ?
                </header>
                <section class="p-4">
                    <button
                        on:click={() => {
                            loading = true;
                            run();
                        }}
                        class="btn btn-sm variant-filled w-20">Execute</button
                    >
                </section>
            {/if}
        </div>
    {:else}
        <slot />
    {/if}
</div>
