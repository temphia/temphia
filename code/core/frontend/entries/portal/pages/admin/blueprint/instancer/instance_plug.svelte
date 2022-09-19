<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalApp } from "../../../../app";
  import Layout from "../../layout.svelte";
  import Plug from "./_plug.svelte";
  import type { PlugRawSchema } from "./instance";

  export let bid: string;
  export let file: string;

  const app: PortalApp = getContext("__app__");

  let loading = true;
  let data: PlugRawSchema;

  const load = async () => {
    const bapi = await app.get_apm().get_bprint_api();
    const resp = await bapi.bprint_get_file(bid, file);
    if (resp.status !== 200) {
      console.log(resp);
      return;
    }
    console.log("@file", resp.data);
    data = resp.data;
    loading = false;
  };

  load();
</script>

<Layout {loading}>
  <Plug {data} {app} {bid} {file} />
</Layout>
