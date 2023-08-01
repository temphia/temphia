<script lang="ts">
  import { getContext } from "svelte";
  import {
    AutoTable,
    LoadingSpinner,
    FloatingAdd,
    PortalService,
  } from "../core";
  import AssignAgent from "./_assign_agent.svelte";
  import PickResourceType from "./_pick_resource_type.svelte";

  const app = getContext("__app__") as PortalService;

  let datas = [];
  let loading = true;
  const api = app.api_manager.get_admin_resource_api();

  const load = async () => {
    const resp = await api.list();
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
    const resp = await api.delete(id);
    if (!resp.ok) {
      return;
    }
    load();
  };

  const action_new = () =>
    app.utils.small_modal_open(PickResourceType, { app });
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <AutoTable
    action_key="id"
    show_drop={true}
    actions={[
      {
        Name: "Edit",
        Action: action_edit,
        icon: "pencil-alt",
      },

      {
        Name: "Assign agent",
        Action: (id, data) => {
          app.utils.small_modal_open(AssignAgent, {
            service: app,
            rid: id,
          });
        },
        icon: "link",
        drop: true,
      },

      {
        Name: "Delete",
        Action: action_delete,
        icon: "trash",
        drop: true,
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
