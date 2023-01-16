<script lang="ts">
  import { getContext } from "svelte";
  import {
    AutoTable,
    LoadingSpinner,
    FloatingAdd,
    PortalService,
  } from "../../core";
  import ChangeEmail from "../../user/actions/change_email.svelte";
  import ResetPassword from "../../user/actions/reset_password.svelte";
  import { params } from "svelte-hash-router";

  export let ugroup = $params.ugroup;

  const app = getContext("__app__") as PortalService;

  let datas = [];
  let loading = true;

  const load = async () => {
    const api = await app.api_manager.get_ugroup_tkt_api(ugroup);
    if (!api) {
      return
    }

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

  // syncme => ../../user/users.svelte

  const action_reset_password = (id: string) => {
    app.utils.small_modal_open(ResetPassword, {
      uid: id,
      onComplete: (opts) => {
        console.log("RESET PASSWORD", opts);
        app.utils.small_modal_close();
      },
    });
  };
  const action_email_change = (id: string) => {
    app.utils.small_modal_open(ChangeEmail, {
      uid: id,
      onComplete: (opts) => {
        console.log("CHANGE EMAIL", opts);
        app.utils.small_modal_close();
      },
    });
  };
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <AutoTable
    action_key="user_id"
    show_drop={true}
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
        drop: true,
      },
      {
        Name: "Reset Password",
        Action: action_reset_password,
        icon: "lock-open",
        drop: true,
      },

      {
        Name: "Change Email",
        Action: action_email_change,
        icon: "at-symbol",
        drop: true,
      },
      {
        Name: "Disable",
        Action: (id) => {},
        icon: "user-remove",
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
