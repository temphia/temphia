<script lang="ts">
  import NewResource from "./_new_resource.svelte";
  import Layout from "../layout.svelte";
  import { getContext } from "svelte";
  import type { PortalApp } from "../../../app";

  const app: PortalApp = getContext("__app__");

  const save = async (data: object) => {
    const rapi = await app.get_apm().get_resource_api();
    const resp = await rapi.resource_create(data);
    // fixme => show message
    if (resp.status !== 200) {
      console.log("Err", resp);
      return;
    }

    app.navigator.goto_admin_resources_page();
  };
</script>

<Layout current_item="resources">
  <NewResource saveFn={save} {app} />
</Layout>
