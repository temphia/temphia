<script lang="ts">
    import NavWrapper from "./nav_wrapper.svelte";
    import type { Launcher } from "$lib/services/portal/launcher/launcher";
    import type { PortalService } from "$lib/services/portal/portal";
    import { setContext } from "svelte";

    export let launcher: Launcher;
    export let app: PortalService;

    const toast_success = (msg: string) => {};
    const toast_error = (msg: string) => {};
    const big_modal_open = (compo: any, options: object) => {};
    const big_modal_close = () => {};
    const small_modal_open = (compo: any, options: object) => {};
    const small_modal_close = () => {};
    const notification_toggle = () => {};
    const pending = app.notifier.is_pending_read;

    app.inject({
        toast_success,
        toast_error,
        big_modal_open,
        big_modal_close,
        small_modal_open,
        small_modal_close,
        notification_toggle,
    });

    setContext("__app__", app);
</script>

<NavWrapper
    {launcher}
    pending_notification={$pending}
    on:logout={(ev) => {}}
    on:notification_toggle={(ev) => {}}
    on:open_executors={(ev) => {}}
    on:self_profile={(ev) => {}}
>
    <svelte:fragment>
        <slot />
    </svelte:fragment>
</NavWrapper>
