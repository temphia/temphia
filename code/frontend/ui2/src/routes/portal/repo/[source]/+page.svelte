<script lang="ts">
    import { getContext } from "svelte";
    import type { PortalService } from "$lib/core";
    import Skeleton from "../panels/_skeleton.svelte";
    import Listings from "../panels/listings.svelte";

    export let data;

    const app = getContext("__app__") as PortalService;
    const rapi = app.api_manager.get_repo_api();

    let sources;
    let loading = true;
    let items = [];

    const load = async (__source: string) => {
        loading = true;
        sources = await app.api_manager.self_data.get_repo_sources();
        const resp = await rapi.load(__source);
        if (!resp.ok) {
            return;
        }
        items = resp.data;
        loading = false;
    };

    $: load(data["source"]);
</script>

{#if loading}
    <Skeleton />
{:else if data["source"]}
    <Listings
        onChangeSource={(next) => app.nav.repo_source(next)}
        onItemSelect={(item) => {
            app.nav.repo_item(
                data["source"],
                item["group"] || item["type"],
                item["slug"]
            );
        }}
        currentSource={data["source"]}
        {items}
        {sources}
    />
{/if}
