<script lang="ts">
  import { getContext } from "svelte";
  import { AutoForm, PortalService } from "../../core";
  import { params } from "svelte-hash-router";

  export let source = $params.source;
  export let group = $params.group;
  export let table = $params.table;

  const app = getContext("__app__") as PortalService;
  const api = app.api_manager.get_admin_data_api();

  let message = "";
  let data = {};
  
  const save = async (_data) => {
    const resp = await api.add_view(source, group, table, _data);
    if (!resp.ok) {
      message = resp.data;
      return;
    }
    app.nav.admin_data_views(source, group, table);
  };
</script>

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
    name: "New View",
    required_fields: [],
  }}
  onSave={save}
  {data}
/>
