<script lang="ts">
  import { getContext } from "svelte";
  import { AutoForm, LoadingSpinner, PortalService } from "../core";

  const app = getContext("__app__") as PortalService;
  const api = app.api_manager.get_admin_tenant_api();

  let message = "";
  let data = {};
  let loading = true;

  const load = async () => {
    const resp = await api.get();
    if (!resp.ok) {
      return;
    }

    data = resp.data;
    loading = false;
  };

  load();

  const save = async (_data) => {
    const resp = await api.edit(_data);
    if (!resp.ok) {
      message = resp.data;
      return;
    }
    app.nav.admin_resources();
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
          ftype: "TEXT",
          key_name: "slug",
          disabled: true,
        },

        {
          name: "Bio",
          ftype: "LONG_TEXT",
          key_name: "slug",
        },

        {
          name: "Default DataSource",
          ftype: "TEXT",
          key_name: "default_dyn",
        },

        {
          name: "Default User Group",
          ftype: "TEXT",
          key_name: "default_ugroup",
        },

        {
          name: "Disable p2p",
          ftype: "BOOL",
          key_name: "disable_p2p",
        },

        // {
        //   name: "SMTP User",
        //   ftype: "TEXT",
        //   key_name: "smtp_user",
        // },
        // {
        //   name: "SMTP Password",
        //   ftype: "TEXT",
        //   key_name: "smtp_pass",
        // },

        {
          name: "Extra Meta",
          ftype: "KEY_VALUE_TEXT",
          key_name: "extra_meta",
        },
      ],
      name: "Edit Tenant",
      required_fields: [],
    }}
    onSave={save}
    {data}
  />
{/if}
