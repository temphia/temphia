<script lang="ts">
    import NavWrapper from "./nav_wrapper.svelte";
    import type { Launcher } from "$lib/services/portal/launcher/launcher";
    import type { PortalService } from "$lib/services/portal/portal";
    import { getDrawerStore, Drawer } from "@skeletonlabs/skeleton";

    import { onMount, setContext } from "svelte";
    import Notification from "../notification/notification.svelte";
    import Modal from "./_modal.svelte";

    export let launcher: Launcher;
    export let app: PortalService;

    const drawerStore = getDrawerStore();
    const pending = app.notifier.is_pending_read;
    const nstate = app.notifier.state;

    const toast_success = (msg: string) => {};
    const toast_error = (msg: string) => {};

    let big_modal_open;
    let big_modal_close;
    let small_modal_open;
    let small_modal_close;

    const notification_toggle = () => {
        if ($drawerStore.open) {
            drawerStore.close();
        } else {
            drawerStore.open({
                id: "notification",
                bgDrawer: "bg-purple-200 text-white",
                position: "left",
                width: "w-full md:w-[480px]",
                padding: "p-2",
                rounded: "rounded-xl",
            });
        }
    };

    onMount(() => {
        app.inject({
            toast_success,
            toast_error,
            big_modal_open,
            big_modal_close,
            small_modal_open,
            small_modal_close,
            notification_toggle,
        });
    });

    setContext("__app__", app);
</script>

<Modal
    bind:show_big={big_modal_open}
    bind:show_small={small_modal_open}
    bind:close_big={big_modal_close}
    bind:close_small={small_modal_close}
/>

<Drawer>
    <Notification
        loading={$nstate.loading}
        messages={$nstate.messages}
        on:ndelete={(ev) => app.notifier.delete_message(ev.detail)}
        on:nread={(ev) => app.notifier.read_message(ev.detail)}
        on:refresh={() => app.notifier.init()}
        on:toggle_npanel={notification_toggle}
        on:explore_noti={() => app.nav.notifications()}
    />
</Drawer>

<NavWrapper
    {launcher}
    pending_notification={$pending}
    on:logout={() => app.logout()}
    on:notification_toggle={notification_toggle}
    on:open_executors={() => app.nav.launcher()}
    on:self_profile={() => app.nav.self_profile()}
>
    <svelte:fragment>
        <slot />
    </svelte:fragment>
</NavWrapper>
