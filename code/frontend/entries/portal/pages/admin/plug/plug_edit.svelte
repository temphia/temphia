<script lang="ts">
  import { getContext } from "svelte";
  import { AutoForm, LoadingSpinner, PortalService } from "../core";
  import { params } from "svelte-hash-router";

  export let pid = $params.pid;

  const app = getContext("__app__") as PortalService;
  const api = app.api_manager.get_admin_plug_api();

  let message = "";
  let data = {};
  let loading = true;

  const load = async () => {
    const resp = await api.get_plug(pid);
    if (!resp.ok) {
      message = resp.data;
      return;
    }

    data = resp.data;
    loading = false;
  };

  const save = async (_data) => {
    console.log("@@data", _data);

    const resp = await api.update_plug(pid, _data);
    if (!resp.ok) {
      message = resp.data;
      return;
    }
    app.nav.admin_plugs();
  };

  load();
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
          disabled: true,
        },
        {
          name: "Name",
          ftype: "TEXT",
          key_name: "name",
        },
        {
          name: "Live",
          ftype: "BOOL",
          key_name: "live",
        },
        {
          name: "Dev",
          ftype: "BOOL",
          key_name: "dev",
        },
        {
          name: "Bprint Id",
          ftype: "TEXT",
          key_name: "bprint_id",
        },
        {
          name: "Invoke Policy",
          ftype: "TEXT_POLICY",
          key_name: "invoke_policy",
        },
        {
          name: "Extra Meta",
          ftype: "KEY_VALUE_TEXT",
          key_name: "extra_meta",
        },
      ],
      name: "New Plug",
      required_fields: ["bprint_id"],
    }}
    onSave={save}
    {data}
  />
{/if}
