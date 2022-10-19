<script lang="ts">
  import { routes, Router, params } from "svelte-hash-router";
  import page_routes from "./pages";
  import Tailwind from "../shared/_compo1/common/_tailwind.svelte";
  import MainLayout from "./layout/main.svelte";
  import BigModal from "./layout/big_modal.svelte";
  import build from "./services/builder";
  import NotificationModal from "./notification/modal.svelte";
  routes.set(page_routes);

  const app = build();
</script>

<BigModal />
<NotificationModal />

<MainLayout pending_notification={true} on:ntoggle={(ev) => console.log(ev)}>
  <Router />
</MainLayout>

<Tailwind />
