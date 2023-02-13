<script lang="ts">
  import { getContext } from "svelte";
  import {
    AutoTable,
    LoadingSpinner,
    FloatingAdd,
    PortalService,
  } from "../../core";
  import TopActions from "../../core/top_actions.svelte";
  import NewPicker from "./_new_picker.svelte";
  export let ttype = undefined;
  export let target = undefined;

  const app = getContext("__app__") as PortalService;

  export let action_new = () => {
    app.utils.small_modal_open(NewPicker, { service: app });
  };

  let datas = [];
  let loading = true;
  const api = app.api_manager.get_admin_target_api();

  const load = async () => {
    let resp;

    if (!ttype) {
      resp = await api.listApp();
    } else {
      resp = await api.listAppByType(ttype, target);
    }

    if (!resp.ok) {
      return;
    }

    datas = resp.data;
    loading = false;
  };

  load();

  // actions

  const action_edit = (id: string, data: object) =>
    app.nav.admin_target_app_edit(data["target_type"], Number(id));
  const action_delete = async (id: string, data: object) => {
    const resp = await api.deleteApp(data["target_type"], Number(id));
    load();
  };

  const actions = {
    Hooks: () => app.nav.admin_target_hooks(),
  };

  if (ttype) {
    actions["All Apps"] = () => app.nav.admin_target_apps();
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
