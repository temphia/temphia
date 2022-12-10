<script lang="ts">
  import { getContext } from "svelte";
  import { AutoForm, PortalService } from "../../core";
  import { params } from "svelte-hash-router";

  export let ugroup = $params.ugroup;

  const app = getContext("__app__") as PortalService;
  const api = app.api_manager.get_admin_ugroup_api();

  let message = "";

  const save = async (_data) => {
    const resp = await api.new(_data);
    if (!resp.ok) {
      message = resp.data;
      return;
    }
    app.nav.admin_ugroup_auths(ugroup)
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
        name: "Type",
        ftype: "TEXT",
        key_name: "type",
      },

      {
        name: "Provider",
        ftype: "TEXT",
        key_name: "provider",
      },

      {
        name: "Scopes",
        ftype: "MULTI_TEXT",
        key_name: "scopes",
      },

      {
        name: "New User If Not Exist",
        ftype: "BOOL",
        key_name: "newuser_ifnot_exists",
      },

      {
        name: "Policy",
        ftype: "TEXT_POLICY",
        key_name: "policy",
      },

      {
        name: "Extra Meta",
        ftype: "KEY_VALUE_TEXT",
        key_name: "extra_meta",
      },
    ],
    name: "New User Group Auth",
    required_fields: [],
  }}
  onSave={save}
  data={{ user_group: ugroup }}
/>
