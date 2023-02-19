<script lang="ts">
  import { getContext } from "svelte";
  import { AutoForm, PortalService } from "../core";

  const app = getContext("__app__") as PortalService;
  const api = app.api_manager.get_admin_bprint_api();

  let message = "";

  const save = async (_data) => {
    console.log("@@data", _data);

    if (!_data["type"]) {
      message = "needs valid type";
      return;
    }

    if (!_data["slug"]) {
      message = "needs slug";
      return;
    }

    const resp = await api.create(_data);
    if (!resp.ok) {
      message = resp.data;
      return;
    }
    app.nav.admin_bprints()
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
        name: "Slug",
        ftype: "TEXT",
        key_name: "slug",
      },
      {
        name: "Type",
        ftype: "SELECT",
        key_name: "type",
        options: [
          "bundle",
          "data_group",
          "data_table",
          "data_sheet",
          "plug",
          "resource",
        ],
      },
      {
        name: "Description",
        ftype: "LONG_TEXT",
        key_name: "description",
      },
    ],
    name: "New Blueprint",
    required_fields: ["slug", "type"],
  }}
  onSave={save}
  data={{}}
/>
