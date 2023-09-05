<script lang="ts">
  import { getContext } from "svelte";
  import { AutoForm, PortalService } from "$lib/core";

  import { params } from "$lib/params";

  const app = getContext("__app__") as PortalService;
  const api = app.api_manager.get_admin_plug_api();

  let message = "";
  let data = { plug_id: $params["pid"], agent_id: $params["aid"] };

  if (app.nav.options) {
    data = { ...data, ...app.nav.options };
  }

  const save = async (_data) => {
    const resp = await api.new_agent_resource($params["pid"], $params["aid"], {
      ...data,
      ..._data,
    });
    if (!resp.ok) {
      message = resp.data;
      return;
    }
    app.nav.admin_agents($params["pid"]);
  };
</script>

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
        name: "Resource Id",
        ftype: "TEXT",
        key_name: "resource_id",
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
    name: "New Agent Resource",
    required_fields: [],
  }}
  onSave={save}
  {data}
/>
