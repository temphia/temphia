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
    const resp = await api.new_agent_ext(pid, aid, _data);
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
        name: "Plug",
        ftype: "TEXT",
        key_name: "plug_id",
        disabled: true,
      },
      {
        name: "Agent",
        ftype: "TEXT",
        key_name: "agent_id",
        disabled: true,
      },
      {
        name: "File",
        ftype: "TEXT",
        key_name: "ref_file",
      },

      {
        name: "Blueprint",
        ftype: "TEXT",
        key_name: "bprint_id",
      },

      {
        name: "Extra Meta",
        ftype: "KEY_VALUE_TEXT",
        key_name: "extra_meta",
      },
    ],
    name: "New Extension",
    required_fields: [],
  }}
  onSave={save}
  data={{ agent_id: aid, plug_id: pid }}
/>
