<script lang="ts">
    import NavWrapper from "./nav_wrapper.svelte";
    import type { Launcher } from "$lib/services/portal/launcher/launcher";
    import type { PortalService } from "$lib/services/portal/portal";
    import { setContext } from "svelte";

    export let launcher: Launcher;
    export let app: PortalService;

    const pending = app.notifier.is_pending_read;




    setContext("__app__", app);
</script>

<NavWrapper {launcher} pending_notification={$pending}>
    <svelte:fragment>
        <slot />
    </svelte:fragment>
</NavWrapper>
