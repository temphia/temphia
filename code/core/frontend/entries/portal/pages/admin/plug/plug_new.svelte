<script lang="ts">
  import { getContext } from "svelte";
  import { AutoForm, PortalService } from "../core";

  const app = getContext("__app__") as PortalService;
  const api = app.api_manager.get_admin_plug_api();

  let message = "";

  const save = async (_data) => {
    console.log("@@data", _data);

    const resp = await api.new_plug(_data);
    if (!resp.ok) {
      message = resp.data;
      return;
    }
    app.nav.admin_plugs();
  };
</script>

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
  data={{}}
/>
