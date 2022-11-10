<script lang="ts">
  import { getContext } from "svelte";
  import { AutoForm, LoadingSpinner, PortalService } from "../../core";
  import { params } from "svelte-hash-router";

  export let source = $params.source;
  export let group = $params.group;
  export let table = $params.table;
  export let id = $params.id;

  const app = getContext("__app__") as PortalService;
  const api = app.api_manager.get_admin_data_api();

  let message = "";
  let data = {};
  let loading = true;

  const load = async () => {
    const resp = await api.get_view(source, group, table, id);
    if (!resp.ok) return;

    data = resp.data;
    loading = false;
  };

  load();

  const save = async (_data) => {
    const resp = await api.edit_view(source, group, table, id, _data);
    if (!resp.ok) {
      message = resp.data;
      return;
    }
    app.nav.admin_data_views(source, group, table);
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
          name: "Extra Meta",
          ftype: "KEY_VALUE_TEXT",
          key_name: "extra_meta",
        },
      ],
      name: "Edit View",
      required_fields: [],
    }}
    onSave={save}
    {data}
  />
{/if}
