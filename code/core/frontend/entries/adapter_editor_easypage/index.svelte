<script lang="ts">
  import { routes } from "svelte-hash-router";
  import Router from "svelte-hash-router";

  import Start from "./page/start.svelte";
  import Page from "./page/page.svelte";
  import Tailwind from "../xcompo/common/_tailwind.svelte";
  import { EasypageService } from "./service/easypage";
  import { setContext } from "svelte";

  export let env: any;

  routes.set({
    "/": Start,
    "/page/:pid": Page,
  });

  let service = new EasypageService(env);

  setContext("__easypage_service__", service);

  console.log("SERVICE |>", service);
</script>

<Tailwind />
<Router />
