<script lang="ts">
  import { routes, Router, params } from "svelte-hash-router";
  import page_routes from "./pages";
  import MainLayout from "./layout/main.svelte";
  import BigModal from "./layout/big_modal.svelte";
  import NotificationModal from "./notification/modal.svelte";
  import type { App } from "./services";
  routes.set(page_routes);
  let app: App;
  
</script>

<BigModal />
<NotificationModal />

<MainLayout pending_notification={true} on:ntoggle={(ev) => console.log(ev)}>
  <Router />
</MainLayout>
