<script lang="ts">
  import type grapesjs from "grapesjs";
  import { onMount, tick } from "svelte";
  import { page_ext } from "./page_ext";
  import type { EasypageService } from "../../service/easypage";
  import { editorFactory } from "./core";

  export let page_slug: string;
  export let service: EasypageService;

  let rootElem;
  let editor: grapesjs.Editor;

  onMount(async () => {
    if (!rootElem) {
      await tick();
    }

    editor = editorFactory(rootElem, page_slug, page_ext(service));
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
