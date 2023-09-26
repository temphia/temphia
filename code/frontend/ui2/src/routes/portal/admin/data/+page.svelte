<script lang="ts">
    import { getContext } from "svelte";
    import { LoadingSpinner, PortalService } from "$lib/core";

    const app = getContext("__app__") as PortalService;

    app.nav.admin_data_groups("default");
</script>

<LoadingSpinner />
