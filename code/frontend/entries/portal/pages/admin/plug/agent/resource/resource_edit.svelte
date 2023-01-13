<script lang="ts">
  import { getContext } from "svelte";
  import { AutoForm, LoadingSpinner, PortalService } from "../../../core";

  import { params } from "svelte-hash-router";

  const app = getContext("__app__") as PortalService;
  const api = app.api_manager.get_admin_plug_api();

  let message = "";
  let data = {};
  let loading = true;

  const load = async () => {
    const resp = await api.get_agent_resource(
      $params.pid,
      $params.aid,
      $params.rid
    );
    if (!resp.ok) {
      return;
    }

    data = resp.data;
    loading = false;
  };

  const save = async (_data) => {
    const resp = await api.update_agent_resource(
      $params.pid,
      $params.aid,
      $params.rid,
      _data
    );
    if (!resp.ok) {
      message = resp.data;
      return;
    }
    app.nav.admin_agents($params.pid);
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
          name: "Slug",
          ftype: "TEXT_SLUG",
          key_name: "slug",
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
          name: "Actions",
          ftype: "MULTI_TEXT",
          key_name: "actions",
        },

        {
          name: "Extra Meta",
          ftype: "KEY_VALUE_TEXT",
          key_name: "extra_meta",
        },
      ],
      name: "Edit Agent Resource link",
      required_fields: [],
    }}
    onSave={save}
    {data}
  />
{/if}
