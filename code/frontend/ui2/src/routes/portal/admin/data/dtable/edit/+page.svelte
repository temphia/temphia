<script lang="ts">
  import { getContext } from "svelte";
  import { AutoForm, LoadingSpinner, PortalService } from "$lib/core";
  import { params } from "$lib/params";

  export let source = $params["source"];
  export let group = $params["group"];
  export let table = $params["table"];

  const app = getContext("__app__") as PortalService;
  const api = app.api_manager.get_admin_data_api();

  let message = "";
  let data = {};
  let loading = true;

  const load = async () => {
    const resp = await api.get_table(source, group, table);
    if (!resp.ok) {
      return;
    }

    data = resp.data;
    loading = false;
  };

  load();

  const save = async (_data) => {
    const resp = await api.edit_table(source, group, table, _data);
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
          name: "Description",
          ftype: "LONG_TEXT",
          key_name: "description",
        },

        {
          name: "Icon",
          ftype: "TEXT",
          key_name: "icon",
        },

        {
          name: "Main Column",
          ftype: "TEXT",
          key_name: "main_column",
        },
        {
          name: "Main View",
          ftype: "TEXT",
          key_name: "main_view",
        },

        {
          name: "Activity Type",
          ftype: "TEXT",
          key_name: "activity_type",
          options: [], // fixme
        },

        {
          name: "Sync Type",
          ftype: "TEXT",
          key_name: "sync_type",
          options: [], // fixme
        },

        {
          name: "Extra Meta",
          ftype: "KEY_VALUE_TEXT",
          key_name: "extra_meta",
        },
      ],
      name: "Edit User Table",
      required_fields: [],
    }}
    onSave={save}
    {data}
  />
{/if}
