<script lang="ts">
  import { getContext } from "svelte";
  import {
    AutoTable,
    LoadingSpinner,
    FloatingAdd,
    PortalService,
  } from "../../../core";

  import { params } from "svelte-hash-router";

  const app = getContext("__app__") as PortalService;

  let datas = [];
  let loading = true;
  const api = app.api_manager.get_admin_plug_api();

  const load = async () => {
    const resp = await api.list_agent_resource($params.pid, $params.aid);
    if (!resp.ok) {
      return;
    }

    datas = resp.data;
    loading = false;
  };

  load();

  // actions

  const action_edit = (id: string) => app.nav.admin_resource_edit(id);
  const action_delete = async (id: string) => {
    const resp = await api.delete_agent_resource($params.pid, $params.aid, id);
    if (!resp.ok) {
      return;
    }
    load();
  };

  const action_new = () =>
    app.nav.admin_agent_res_new($params.pid, $params.aid);
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <AutoTable
    action_key="id"
    actions={[
      {
        Name: "Edit",
        Action: action_edit,
        icon: "pencil-alt",
      },
      {
        Name: "Delete",
        Class: "bg-red-400",
        Action: action_delete,
        icon: "trash",
      },
    ]}
    key_names={[
      ["name", "Name"],
      ["id", "Id"],
      ["type", "Type"],
      ["schema", "Schema"],
    ]}
    color={["type"]}
    {datas}
  />
{/if}

<FloatingAdd onClick={action_new} />
