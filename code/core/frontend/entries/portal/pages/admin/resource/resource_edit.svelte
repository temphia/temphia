<script lang="ts">
  import { getContext } from "svelte";
  import { AutoForm, LoadingSpinner, PortalService } from "../core";
  import { params } from "svelte-hash-router";

  export let rid = $params.rid;

  const app = getContext("__app__") as PortalService;
  const api = app.api_manager.get_admin_resource_api();

  let message = "";
  let data = {};
  let loading = true;

  const load = async () => {
    const resp = await api.get(rid);
    if (resp.ok) {
      return;
    }

    data = resp.data;
    loading = false;
  };

  load();

  const save = async (_data) => {
    const resp = await api.update(rid, _data);
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
          name: "Id",
          ftype: "TEXT_SLUG",
          key_name: "id",
        },
        {
          name: "Name",
          ftype: "TEXT",
          key_name: "name",
        },
        // {
        //   name: "Sub Type",
        //   ftype: "TEXT",
        //   key_name: "sub_type",
        // },
        // {
        //   name: "Payload",
        //   ftype: "LONG_TEXT",
        //   key_name: "payload",
        // },
        {
          name: "Type",
          ftype: "TEXT",
          key_name: "type",
          options: ["datagroup", "folder", "room"],
        },

        {
          name: "Target",
          ftype: "TEXT",
          key_name: "target",
        },
        {
          name: "Policy",
          ftype: "TEXT_POLICY",
          key_name: "policy",
        },

        {
          name: "Plug Id",
          ftype: "TEXT",
          key_name: "plug_id",
        },

        {
          name: "Extra Meta",
          ftype: "KEY_VALUE_TEXT",
          key_name: "extra_meta",
        },
      ],
      name: "New Resource",
      required_fields: [],
    }}
    onSave={save}
    {data}
  />
{/if}
