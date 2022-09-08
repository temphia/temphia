<script lang="ts">
  import EditGroup from "./_edit_group.svelte";
  import Layout from "../layout.svelte";
  import type { PortalApp } from "../../../../../lib/app/portal";
  import { getContext } from "svelte";

  const app: PortalApp = getContext("__app__");

  const create = async (data: any) => {
    const uapi = await app.get_apm().get_user_api();
    return uapi.add_user_group(data);
  };

</script>

<Layout current_item={"user_groups"}>
  <EditGroup {app} new_mode={true} onSave={(slug, data) => create(data)} />
</Layout>
