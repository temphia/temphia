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
  import { DemoCSS, DemoHTML } from "./demo";

  let rootElem;
  let editor: grapejs.Editor;

  onMount(() => {
    editor = grapejs.init({
      container: rootElem,
      components: DemoHTML,
      style: DemoCSS,
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
      ],
      pluginsOpts: {},
      storageManager: false,
    });
  });
</script>

<svelte:head>
  <link
    rel="stylesheet"
    href="https://unpkg.com/grapesjs/dist/css/grapes.min.css"
  />
  <script src="https://unpkg.com/grapesjs-lory-slider"></script>
</svelte:head>

<div bind:this={rootElem}>Site Builder</div>
