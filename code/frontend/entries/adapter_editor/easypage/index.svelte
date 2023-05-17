<script lang="ts">
  import { routes } from "svelte-hash-router";
  import Router from "svelte-hash-router";
  import Tailwind from "../../xcompo/common/_tailwind.svelte";
  import { EasypageService } from "./service/easypage";
  import { onMount, setContext } from "svelte";
  import { Modal } from "./core";

  import Start from "./page/start.svelte";
  import Page from "./page/page.svelte";
  import Post from "./page/post.svelte";
  import RawEdit from "./page/raw_edit.svelte";
  import Visual from "./page/visual.svelte";

  export let env: any;

  routes.set({
    "/": Start,
    "/page/:pid": Page,
    "/post/:pid": Post,
    "/raw/:pid": RawEdit,
    "/visual/:pid": Visual,
  });

  let big_open;
  let big_close;
  let small_open;
  let small_close;

  let service = new EasypageService(env);

  setContext("__easypage_service__", service);

  onMount(() => {
    service.modal = {
      big_open,
      big_close,
      small_open,
      small_close,
    };
  });
</script>

<Modal
  bind:show_big={big_open}
  bind:close_big={big_close}
  bind:show_small={small_open}
  bind:close_small={small_close}
/>

<Tailwind />

<Router />
