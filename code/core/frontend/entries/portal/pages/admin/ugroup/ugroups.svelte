<script lang="ts">
  import { getContext } from "svelte";
  import {
    AutoTable,
    LoadingSpinner,
    FloatingAdd,
    PortalService,
  } from "../core";
  import TopActions from "../core/top_actions.svelte";

  const app = getContext("__app__") as PortalService;

  let datas = [];
  let loading = true;
  const api = app.api_manager.get_admin_ugroup_api();

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

  const action_edit = (id: string) => {};
  const action_delete = async (id: string) => {
    const resp = await api.delete(id);
    load();
  };

  const action_new = () => app.nav.admin_ugroup_new()
</script>

<TopActions
  actions={{
    "All Users": () => {},
  }}
/>

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
      ["slug", "Slug"],
      ["enable_pass_auth", "Enabled Password"],
      ["Open Sign Up", "open_sign_up"],
    ]}
    color={[]}
    {datas}
  />
{/if}

<FloatingAdd onClick={action_new} />
