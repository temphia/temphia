<script lang="ts">
  import { getContext } from "svelte";
  import { AutoForm, LoadingSpinner, PortalService } from "../../core";
  import { params } from "svelte-hash-router";

  export let did = $params.did;

  const app = getContext("__app__") as PortalService;
  const api = app.api_manager.get_admin_tenant_api();

  let message = "";
  let data = {};
  let loading = true;

  const load = async () => {
    const resp = await api.get_domain(did);
    if (!resp.ok) {
      return;
    }

    data = resp.data;
    loading = false;
  };

  load();

  const save = async (_data) => {
    const resp = await api.edit_domain(did, _data);
    if (!resp.ok) {
      message = resp.data;
      return;
    }
    app.nav.admin_tenant();
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
          name: "About",
          ftype: "LONG_TEXT",
          key_name: "about",
        },

        {
          name: "Default User Group",
          ftype: "LONG_TEXT",
          key_name: "default_ugroup",
        },

        {
          name: "CORS Policy",
          ftype: "TEXT_POLICY",
          key_name: "cors_policy",
        },

        {
          name: "Adapter Policy",
          ftype: "TEXT_POLICY",
          key_name: "adapter_policy",
        },

        {
          name: "Adapter Type",
          ftype: "TEXT",
          key_name: "adapter_type",
        },

        {
          name: "Adapter Options",
          ftype: "KEY_VALUE_TEXT",
          key_name: "adapter_opts",
        },

        {
          name: "Adapter Cabinet Source",
          ftype: "TEXT",
          key_name: "adapter_cab_source",
        },

        {
          name: "Adapter Cabinet Folder",
          ftype: "TEXT",
          key_name: "adapter_cab_folder",
        },

        {
          name: "Adapter Template Blueprints",
          ftype: "MULTI_TEXT",
          key_name: "adapter_template_bprints",
        },

        {
          name: "Extra Meta",
          ftype: "KEY_VALUE_TEXT",
          key_name: "extra_meta",
        },
      ],
      name: "Update Domain",
      required_fields: [],
    }}
    onSave={save}
    {data}
  />
{/if}
