<script lang="ts">
  import grapejs from "grapesjs";
  import webpagePlugin from "grapesjs-preset-webpage";
  import basicPlugin from "grapesjs-blocks-basic";
  import gjsForms from "grapesjs-plugin-forms";
  import navPlugin from "grapesjs-navbar";
  import customCodePlugin from "grapesjs-custom-code";
  import blkFlexboxPlugin from "grapesjs-blocks-flexbox";
  import stgrPlugin from "grapesjs-style-gradient";
  import styleFilter from "grapesjs-style-filter";
  import tabPlugin from "grapesjs-tabs";
  import toolTipPlugin from "grapesjs-tooltip";
  import twPlugin from "grapesjs-tailwind";

  import { onMount } from "svelte";
  import { easyPageStore } from "./grape_ext";
  import type { EasypageService } from "../../service/easypage";

  export let page_slug: string;
  export let service: EasypageService;

  let rootElem;
  let editor: grapejs.Editor;

  onMount(() => {
    editor = grapejs.init({
      container: rootElem,
      plugins: [
        webpagePlugin,
        basicPlugin,
        gjsForms,
        navPlugin,
        stgrPlugin,
        styleFilter,
        customCodePlugin,
        blkFlexboxPlugin,
        "grapesjs-lory-slider",
        tabPlugin,
        toolTipPlugin,
        twPlugin,
        easyPageStore(service),
      ],

      pluginsOpts: {},
      storageManager: {
        type: "easypage-store",
        stepsBeforeSave: 3,
        options: {
          "easypage-store": {
            page_slug,
          },
        },
      },
    });
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
