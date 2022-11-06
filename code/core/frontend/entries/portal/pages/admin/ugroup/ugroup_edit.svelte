<script lang="ts">
  import { getContext } from "svelte";
  import { AutoForm, LoadingSpinner, PortalService } from "../core";
  import { params } from "svelte-hash-router";

  export let ugroup = $params.ugroup;

  const app = getContext("__app__") as PortalService;
  const api = app.api_manager.get_admin_ugroup_api();

  let message = "";
  let data = {};
  let loading = true;

  const load = async () => {
    const resp = await api.get(ugroup);
    if (!resp.ok) {
      return;
    }

    data = resp.data;
  };

  load();

  const save = async (_data) => {
    const resp = await api.update(ugroup, _data);
    if (!resp.ok) {
      message = resp.data;
      return;
    }
    app.nav.admin_ugroups();
  };
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <AutoForm
    {message}
    schema={{
      fields: [
        {
          name: "Name",
          ftype: "TEXT",
          key_name: "name",
        },
        {
          name: "Slug",
          ftype: "TEXT_SLUG",
          key_name: "slug",
          slug_gen: null, // fixme
        },

        {
          name: "Icon",
          ftype: "TEXT",
          key_name: "icon",
        },

        {
          name: "Enable Password",
          ftype: "BOOL",
          key_name: "enable_pass_auth",
        },

        {
          name: "Scopes",
          ftype: "MULTI_TEXT",
          key_name: "scopes",
        },

        {
          name: "Open Sign Up",
          ftype: "BOOL",
          key_name: "open_sign_up",
        },

        {
          name: "Extra Meta",
          ftype: "KEY_VALUE_TEXT",
          key_name: "extra_meta",
        },
      ],
      name: "Edit User Group",
      required_fields: [],
    }}
    onSave={save}
    {data}
  />
{/if}
