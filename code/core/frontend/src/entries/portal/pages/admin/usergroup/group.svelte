<script lang="ts">
  import EditGroup from "./_edit_group.svelte";
  import Layout from "../layout.svelte";
  import type { PortalApp } from "../../../../../lib/app/portal";
  import { getContext } from "svelte";

  export let id;

  const app: PortalApp = getContext("__app__");

  let group;

  const load = async () => {
    const uapi = await app.get_apm().get_user_api();
    const resp = await uapi.get_user_group(id);
    group = resp.data;
  };

  const update = async (id: string, data: any) => {
    const uapi = await app.get_apm().get_user_api();
    return uapi.update_user_group(id, data);
  };

  load();
</script>

<Layout loading={!group} current_item={"user_groups"}>
  {#if group}
    <EditGroup {app} {...group} onSave={update} />
  {/if}
</Layout>
