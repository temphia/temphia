<script lang="ts">
  import { getContext } from "svelte";
  import { AutoForm, PortalService } from "../../../core";
  import { params } from "svelte-hash-router";

  export let pid = $params.pid;
  export let aid = $params.pid;

  const app = getContext("__app__") as PortalService;
  const api = app.api_manager.get_admin_plug_api();

  let message = "";

  const save = async (_data) => {
    const resp = await api.new_agent_link(pid, aid, _data);
    if (!resp.ok) {
      message = resp.data;
      return;
    }
    app.nav.admin_agents(pid);
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
        name: "From Plug",
        ftype: "TEXT",
        key_name: "from_plug_id",
        disabled: true,
      },
      {
        name: "From Agent",
        ftype: "TEXT",
        key_name: "from_agent_id",
        disabled: true,
      },

      {
        name: "To Plug",
        ftype: "TEXT",
        key_name: "to_plug_id",
      },
      {
        name: "To Agent",
        ftype: "TEXT",
        key_name: "to_agent_id",
      },

      {
        name: "To Handler",
        ftype: "TEXT",
        key_name: "to_handler",
      },
      {
        name: "Extra Meta",
        ftype: "KEY_VALUE_TEXT",
        key_name: "extra_meta",
      },
    ],
    name: "New Link",
    required_fields: [],
  }}
  onSave={save}
  data={{ from_agent_id: aid, from_plug_id: pid }}
/>
