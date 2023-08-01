<script lang="ts">
  import { getContext } from "svelte";
  import { AutoForm, LoadingSpinner, PortalService } from "../../core";
  import { params } from "svelte-hash-router";

  export let source = $params.source;
  export let group = $params.group;

  const app = getContext("__app__") as PortalService;
  const api = app.api_manager.get_admin_data_api();

  let message = "";
  let data = {};
  let loading = true;

  const load = async () => {
    const resp = await api.get_group(source, group);
    if (!resp.ok) {
      return;
    }

    data = resp.data;
    loading = false;
  };

  load();

  const save = async (_data) => {
    const resp = await api.edit_group(source, group, _data);
    if (!resp.ok) {
      message = resp.data;
      return;
    }
    app.nav.admin_data_groups(source);
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
          disabled: true,
        },

        {
          name: "Renderer",
          ftype: "SELECT",
          key_name: "renderer",
          disabled: true,
          options: ["", "default", "sheet"],
        },

        {
          name: "Description",
          ftype: "LONG_TEXT",
          key_name: "description",
        },

        {
          name: "Cabinet Source",
          ftype: "TEXT",
          key_name: "cabinet_source",
        },

        {
          name: "Cabinet Folder",
          ftype: "TEXT",
          key_name: "cabinet_folder",
        },

        {
          name: "Active",
          ftype: "BOOL",
          key_name: "active",
        },

        {
          name: "Extra Meta",
          ftype: "KEY_VALUE_TEXT",
          key_name: "extra_meta",
        },
      ],
      name: "Edit Group",
      required_fields: [],
    }}
    onSave={save}
    {data}
  />
{/if}
