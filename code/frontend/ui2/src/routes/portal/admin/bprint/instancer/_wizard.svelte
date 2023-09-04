<script lang="ts">
    import { getContext } from "svelte";
    import { LoadingSpinner, PortalService } from "$lib/core";
    import { Stepper, Step } from "$lib/compo/stepper";
    import Preview from "./_preview.svelte";
    import FinalTable from "./_final.svelte";

    export let bid: string;
    export let appjson: object;
    export let bprintjson: object;

    const app = getContext("__app__") as PortalService;
    const bapi = app.api_manager.get_admin_bprint_api();

    let instanceing = false;
    let message = "";
    let finished = false;
    let fdata = {};

    const onNextHandler = async (e: any) => {
        if (e.detail["step"] === 0) {
            instanceing = true;

            const resp = await bapi.instance(bid, {
                auto: true,
                instancer_type: "app",
                file: "app.json",
            });

            if (!resp.ok) {
                finished = true;
                message = resp.data;
                instanceing = false
                return;
            }

            fdata = resp.data;
            finished = true;
            instanceing = false
        }
    };
</script>

<div class="card p-2 bg-white border shadow mx-auto" style="max-width: 750px;">
    <Stepper on:next={onNextHandler} buttonCompleteLabel="">
        <Step back_locked locked={instanceing}>
            <svelte:fragment slot="header">Blueprint items</svelte:fragment>
            <Preview {bid} items={appjson} />

            <p class="text-sm italic my-1">
                Click next to instance items automatically or click manual to
                instance individual items.
            </p>
        </Step>

        <Step back_locked={finished}>
            <svelte:fragment slot="header"
                >{instanceing ? "Instancing" : "Finished"}</svelte:fragment
            >

            {#if instanceing}
                <LoadingSpinner classes="" />
            {:else if message}
                <p class="bg-red-500">
                    {message}
                </p>
            {:else}
                <FinalTable installed_items={fdata} />
            {/if}
        </Step>
    </Stepper>
</div>
