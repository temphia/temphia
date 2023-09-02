<script lang="ts">
    import { getContext } from "svelte";
    import type { PortalService } from "$lib/core";
    import Skeleton from "./panels/_skeleton.svelte";

    const app = getContext("__app__") as PortalService;
    let sources;

    (async () => {
        sources = await app.api_manager.self_data.get_repo_sources();
        app.nav.repo_source(Object.entries(sources)[0][0]);
    })();
</script>

<Skeleton />
