<script lang="ts">
  import { routes, Router, params } from "svelte-hash-router";
  import page_routes from "./pages/page_routes";
  import MainLayout from "./layout/main.svelte";
  import Modal from "./layout/modal.svelte";
  import NotificationModal from "./notification/modal.svelte";
  import type { PortalService } from "./services";
  import { onMount, setContext } from "svelte";

  routes.set(page_routes);

  export let app: PortalService;

  const notifier = app.notifier;
  const nstate = app.notifier.state;

  $: console.log("@PARAMS |>", $params, "@ROUTES |>", $routes);

  // binds
  let big_modal_close;
  let big_modal_open;
  let small_modal_close;
  let small_modal_open;
  let notification_toggle;

  onMount(() =>
    app.inject({
      big_modal_close,
      big_modal_open,
      small_modal_close,
      small_modal_open,
      notification_toggle,
      toast_error: null,
      toast_success: null,
    })
  );

  setContext("__app__", app);
</script>

<Modal
  bind:close_big={big_modal_close}
  bind:show_big={big_modal_open}
  bind:show_small={small_modal_open}
  bind:close_small={small_modal_close}
/>

<NotificationModal
  bind:toggle={notification_toggle}
  on:ndelete={(ev) => notifier.delete_message(ev.detail)}
  on:nread={(ev) => notifier.read_message(ev.detail)}
  on:refresh={() => {}}
  loading={$nstate.loading}
  messages={$nstate.messages}
/>

<MainLayout
  launcher={app.launcher}
  pending_notification={!!$nstate.messages}
  on:notification_toggle={notification_toggle}
  on:self_profile={app.nav.self_profile}
>
  <Router />
</MainLayout>
