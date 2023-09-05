<script lang="ts">
    import { getContext } from "svelte";
    import { AutoForm, PortalService } from "$lib/core";
    import { params } from "$lib/params";
  
    export let pid = $params["pid"];
  
    const app = getContext("__app__") as PortalService;
    const api = app.api_manager.get_admin_plug_api();
  
    let message = "";
  
    const save = async (_data) => {
      console.log("@@data", _data);
  
      const resp = await api.new_plug_state(pid, {
        key: _data["key"],
        value: _data["value"] || "",
        options: {
          tag1: _data["tag1"] || "",
          tag2: _data["tag2"] || "",
          tag3: _data["tag3"] || "",
        },
      });
      if (!resp.ok) {
        message = resp.data;
        return;
      }
      app.nav.admin_plug_states(pid);
    };
  </script>
  
  <AutoForm
    {message}
    schema={{
      fields: [
        {
          name: "key",
          ftype: "TEXT",
          key_name: "key",
        },
        {
          name: "Value",
          ftype: "LONG_TEXT",
          key_name: "value",
        },
  
        {
          name: "Tag1",
          ftype: "TEXT",
          key_name: "tag1",
        },
        {
          name: "Tag2",
          ftype: "TEXT",
          key_name: "tag1",
        },
        {
          name: "Tag3",
          ftype: "TEXT",
          key_name: "tag1",
        },
      ],
      name: "New Plug State",
      required_fields: ["key"],
    }}
    onSave={save}
    data={{}}
  />
  