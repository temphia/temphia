<script lang="ts">
  import Codepanel from "./codepanel.svelte";
  import Outputpanel from "./outputpanel.svelte";
  import Layout from "./_layout.svelte";

  let file;
  let changeFile;
</script>

<Layout {file} {changeFile}>
  <svelte:fragment slot="code">
    <Codepanel bind:file bind:changeFile />
  </svelte:fragment>
  <svelte:fragment slot="output">
    <Outputpanel />
  </svelte:fragment>
</Layout>
