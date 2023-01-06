<script lang="ts">
  import { getContext } from "svelte";
  import {
    AutoTable,
    LoadingSpinner,
    FloatingAdd,
    PortalService,
  } from "../../core";
  import TopActions from "../../core/top_actions.svelte";

  export let ttype = undefined;
  export let target = undefined;
  export let action_new = () => app.nav.admin_target_hook_new();

  const app = getContext("__app__") as PortalService;

  let datas = [];
  let loading = true;
  const api = app.api_manager.get_admin_target_api();

  const load = async () => {
    let resp;

    if (!ttype) {
      resp = await api.listHook();
    } else {
      resp = await api.listHookByType(ttype, target);
    }

    if (!resp.ok) {
      return;
    }

    datas = resp.data;
    loading = false;
  };

  load();

  // actions

  const action_edit = (id: string) =>
    app.nav.admin_target_hook_edit(ttype, Number(id));
  const action_delete = async (id: string, data: object) => {
    const resp = await api.deleteHook(data["target_type"], Number(id));
    load();
  };

  const actions = {
    Apps: () => app.nav.admin_target_apps(),
  };

  if (ttype) {
    actions["All Hooks"] = () => app.nav.admin_target_hooks();
  }
</script>

<TopActions {actions} />

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
      ["target_type", "Type"],
      ["target", "Target"],
    ]}
    color={["target_type"]}
    {datas}
  />
{/if}

<FloatingAdd onClick={action_new} />
