<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalService } from "../../../services";
  import { AutoForm } from "../core";

  const app = getContext("__app__") as PortalService;
  const rapi = app.api_manager.get_admin_repo_api();

  let message = "";

  const saveHandle = async (_data) => {
    const resp = await rapi.new(_data);
    if (resp.status !== 200) {
      message = resp.data;
      return;
    }

    app.nav.admin_repos();
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
        name: "Provider",
        ftype: "TEXT",
        key_name: "provider",
      },

      {
        name: "URL",
        ftype: "TEXT",
        key_name: "url",
      },

      {
        name: "Extra Meta",
        ftype: "KEY_VALUE_TEXT",
        key_name: "extra_meta",
      },
    ],
    name: "New Repo",
    required_fields: [],
  }}
  onSave={saveHandle}
  data={{}}
/>
