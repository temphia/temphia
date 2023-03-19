<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalService } from "../../../core";
  import Codepanel from "./codepanel.svelte";
  import Outputpanel from "./outputpanel.svelte";
  import Layout from "./_layout.svelte";

  export let files;
  export let file;
  export let bid;
  
  let editor;
  let code = "";
  const modified = {};
  let loading = true;

  const app = getContext("__app__") as PortalService;
  const bapi = app.api_manager.get_admin_bprint_api();

  const load = async (lfile: string) => {
    loading = true;
    const resp = await bapi.get_file(bid, lfile);
    if (!resp.ok) {
      console.log("@resp", resp);
      return;
    }

    if (typeof resp.data === "object") {
      code = JSON.stringify(resp.data, undefined, 2);
    } else {
      code = resp.data;
    }

    loading = false;
  };
  const changeFile = async (tofile) => {
    if (file === tofile) {
      return;
    }

    if (editor && file) {
      modified[file] = editor.getValue();
    }
    file = tofile;

    const old = modified[tofile];
    if (old) {
      code = old;
    } else {
      await load(tofile);
    }
  };

  load(file);
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
