<script lang="ts">
  import Codepanel from "./codepanel.svelte";
  import Outputpanel from "./outputpanel.svelte";
  import Layout from "./_layout.svelte";

  export let files = ["server.js", "client.js"];
  export let file = "server.js";
  let editor;
  let code = "";
  let loading = false;

  const modified = {};

  const changeFile = (tofile) => {
    if (file === tofile) {
      return;
    }

    if (editor && file) {
      modified[file] = editor.getValue();
    }

    const old = modified[tofile];
    if (old) {
      code = old;
    } else {
      code = "// data";
    }
    file = tofile;
  };
</script>

<Layout {file} {changeFile} {files}>
  <svelte:fragment slot="code">
    {#key file}
      <Codepanel bind:editor {file} {code} {loading} />
    {/key}
  </svelte:fragment>
  <svelte:fragment slot="output">
    <Outputpanel />
  </svelte:fragment>
</Layout>
