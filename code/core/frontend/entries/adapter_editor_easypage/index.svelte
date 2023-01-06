<script>
  import { routes } from "svelte-hash-router";
  import Router from "svelte-hash-router";

  import Start from "./page/start.svelte";
  import Page from "./page/page.svelte";
  import Tailwind from "../xcompo/common/_tailwind.svelte";

  routes.set({
    "/": Start,
    "/page/:pid": Page,
  });
</script>

<Tailwind />
<Router />
