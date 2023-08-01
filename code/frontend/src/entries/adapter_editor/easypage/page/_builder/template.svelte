<script lang="ts">
  import type grapesjs from "grapesjs";
  import { onMount, tick } from "svelte";
  import type { EasypageService } from "../../service/easypage";
  import { editorFactory } from "./core";

  import { template_ext } from "./template_ext";

  export let page_slug: string;
  export let service: EasypageService;

  let rootElem;
  let editor: grapesjs.Editor;

  onMount(async () => {
    if (!rootElem) {
      await tick();
    }

    editor = editorFactory(
      rootElem,
      page_slug,
      template_ext(service, "<h1>FIXME: Load from api </h1>")
    );
  });
</script>

<svelte:head>
  <link
    rel="stylesheet"
    href="https://unpkg.com/grapesjs/dist/css/grapes.min.css"
  />
  <!-- <script src="https://unpkg.com/grapesjs-lory-slider"></script> -->
</svelte:head>

<div bind:this={rootElem}>Site Builder</div>
