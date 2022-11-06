<script lang="ts">
  import { getContext } from "svelte";
  import { AutoForm, LoadingSpinner, PortalService } from "../../core";
  import { params } from "svelte-hash-router";

  export let ttype = $params.ttype;
  export let id = $params.id;

  const app = getContext("__app__") as PortalService;
  const api = app.api_manager.get_admin_target_api();

  let message = "";
  let data = {};
  let loading = true;

  const load = async () => {
    const resp = await api.getHook(ttype, Number(id));
    if (!resp.ok) {
      return;
    }

    data = resp.data;
    loading = false;
  };

  load();

  const save = async (_data) => {
    const resp = await api.updateHook(ttype, Number(id), _data);
    if (!resp.ok) {
      message = resp.data;
      return;
    }
    app.nav.admin_target_hooks();
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
          name: "Target Type",
          ftype: "TEXT",
          key_name: "target_type",
          disabled: true,
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
          name: "Agent Id",
          ftype: "TEXT",
          key_name: "agent_id",
        },

        {
          name: "Handler",
          ftype: "TEXT",
          key_name: "handler",
        },

        {
          name: "Exec Meta",
          ftype: "KEY_VALUE_TEXT",
          key_name: "exec_meta",
        },

        {
          name: "Extra Meta",
          ftype: "KEY_VALUE_TEXT",
          key_name: "extra_meta",
        },
      ],
      name: "New Target Hook",
      required_fields: [],
    }}
    onSave={save}
    data={{}}
  />
{/if}
