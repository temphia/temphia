<script lang="ts">
  import { getContext } from "svelte";
  import { AutoForm, LoadingSpinner, PortalService } from "../../core";
  import { params } from "svelte-hash-router";

  export let source = $params.source;
  export let group = $params.group;
  export let table = $params.table;
  export let column = $params.column;

  const app = getContext("__app__") as PortalService;
  const api = app.api_manager.get_admin_data_api();

  let message = "";
  let data = {};
  let loading = true;

  const load = async () => {
    const resp = await api.get_column(source, group, table, column);
    if (!resp.ok) {
      return;
    }

    data = resp.data;
    loading = false;
  };

  load();

  const save = async (_data) => {
    const resp = await api.edit_column(source, group, table, column, _data);
    if (!resp.ok) {
      message = resp.data;
      return;
    }
    app.nav.admin_data_columns(source, group, table);
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
          name: "Column Type",
          ftype: "TEXT",
          key_name: "ctype",
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
          name: "Order Id",
          ftype: "INT",
          key_name: "order_id",
        },

        {
          name: "Pattern",
          ftype: "TEXT",
          key_name: "pattern",
        },
        {
          name: "Strict pattern",
          ftype: "BOOL",
          key_name: "strict_pattern",
        },

        {
          name: "Ref Id",
          ftype: "TEXT",
          key_name: "ref_id",
        },

        {
          name: "Ref Type",
          ftype: "TEXT",
          key_name: "ref_type",
        },

        {
          name: "Ref Target",
          ftype: "TEXT",
          key_name: "ref_target",
        },

        {
          name: "Ref Object",
          ftype: "TEXT",
          key_name: "ref_object",
        },
        {
          name: "Ref Copy",
          ftype: "TEXT",
          key_name: "ref_copy",
        },

        {
          name: "Extra Meta",
          ftype: "KEY_VALUE_TEXT",
          key_name: "extra_meta",
        },
      ],
      name: "Edit Column",
      required_fields: [],
    }}
    onSave={save}
    {data}
  />
{/if}
