<script lang="ts">
  import { getContext } from "svelte";
  import {
    AutoTable,
    LoadingSpinner,
    FloatingAdd,
    PortalService,
  } from "../../core";

  const app = getContext("__app__") as PortalService;

  let datas = [];
  let loading = true;

  // fixme => user pi tkt
  const api = app.api_manager.get_admin_user_api()


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
  const action_profile = (id: string) => app.nav.user_profile(id);
  const action_delete = async (id: string) => {};
  const action_new = () => {};
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <AutoTable
    action_key="user_id"
    actions={[
      {
        Name: "Profile",
        Action: action_profile,
        icon: "user-circle",
      },

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
      ["user_id", "User Id"],
      ["full_name", "Full Name"],
      ["group_id", "Group"],
      ["created_at", "Created At"],
      ["active", "Active"],
    ]}
    color={["group_id", "active"]}
    {datas}
  />
{/if}

<FloatingAdd onClick={action_new} />
