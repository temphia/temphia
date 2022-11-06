<script lang="ts">
  import { getContext } from "svelte";
  import { AutoForm, PortalService } from "../core";

  const app = getContext("__app__") as PortalService;
  const api = app.api_manager.get_admin_ugroup_api()

  let message = "";

  const save = async (_data) => {
    const resp = await api.new(_data);
    if (!resp.ok) {
      message = resp.data;
      return;
    }
    app.nav.admin_resources()
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
        ftype: "TEXT_SLUG",
        key_name: "slug",
        slug_gen: null, // fixme 
      },

      {
        name: "Icon",
        ftype: "TEXT",
        key_name: "icon",
      },

      {
        name: "Enable Password",
        ftype: "BOOL",
        key_name: "enable_pass_auth",
      },

      {
        name: "Scopes",
        ftype: "MULTI_TEXT",
        key_name: "scopes",
      },

      {
        name: "Open Sign Up",
        ftype: "BOOL",
        key_name: "open_sign_up",
      },

      {
        name: "Extra Meta",
        ftype: "KEY_VALUE_TEXT",
        key_name: "extra_meta",
      },
    ],
    name: "New User Group",
    required_fields: [],
  }}
  onSave={save}
  data={{}}
/>
