<script lang="ts">
  import { getContext } from "svelte";
  import { AutoForm, LoadingSpinner, PortalService } from "../../core";
  import { params } from "svelte-hash-router";

  export let ugroup = $params.ugroup;
  export let id = $params.id;

  const app = getContext("__app__") as PortalService;
  const api = app.api_manager.get_admin_ugroup_api();

  let message = "";
  let data = {};
  let loading = true;

  const load = async () => {
    const resp = await api.getAuth(ugroup, id);
    if (!resp.ok) {
      return;
    }

    data = resp.data;
  };

  load();

  const save = async (_data) => {
    const resp = await api.updateAuth(ugroup, id, _data);
    if (!resp.ok) {
      message = resp.data;
      return;
    }
    app.nav.admin_ugroup_auths(ugroup);
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
          name: "Id",
          ftype: "INT",
          key_name: "id",
          disabled: true,
        },

        {
          name: "Name",
          ftype: "TEXT",
          key_name: "name",
        },

        {
          name: "Type",
          ftype: "TEXT",
          key_name: "type",
        },

        {
          name: "Provider",
          ftype: "TEXT",
          key_name: "provider",
        },

        {
          name: "User Group",
          ftype: "TEXT",
          key_name: "user_group",
          disabled: true,
        },

        {
          name: "Scopes",
          ftype: "MULTI_TEXT",
          key_name: "scopes",
        },

        {
          name: "New User If Not Exist",
          ftype: "BOOL",
          key_name: "newuser_ifnot_exists",
        },

        {
          name: "Policy",
          ftype: "TEXT_POLICY",
          key_name: "policy",
        },

        {
          name: "Extra Meta",
          ftype: "KEY_VALUE_TEXT",
          key_name: "extra_meta",
        },
      ],
      name: "Edit User Group Auth",
      required_fields: [],
    }}
    onSave={save}
    {data}
  />
{/if}
