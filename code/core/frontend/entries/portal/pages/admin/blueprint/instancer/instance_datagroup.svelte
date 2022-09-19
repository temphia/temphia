<script lang="ts">
  import DataGroup from "./_data_group.svelte";
  import { getContext } from "svelte";
  import type { PortalApp } from "../../../../app";
  import Layout from "../../layout.svelte";

  export let bid;
  export let file;

  const app: PortalApp = getContext("__app__");

  let loading = true;
  let data: any;

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
  <DataGroup {data} {bid} {file} />
</Layout>
