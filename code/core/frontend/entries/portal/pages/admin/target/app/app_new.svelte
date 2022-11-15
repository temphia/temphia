<script lang="ts">
  import { getContext } from "svelte";
  import { AutoForm, PortalService } from "../../core";
  import { params } from "svelte-hash-router";
  export let ttype = $params.ttype;
  export let target = undefined;

  const app = getContext("__app__") as PortalService;
  const api = app.api_manager.get_admin_target_api();

  let message = "";
  let data = {};
  if (ttype) {
    data["target_type"] = ttype;
  }

  if (target) {
    data["target"] = target;
  }
  const save = async (_data) => {
    const resp = await api.newApp(_data["target_type"], _data);
    if (!resp.ok) {
      message = resp.data;
      return;
    }
    app.nav.admin_target_apps();
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
        name: "Target Type",
        ftype: "TEXT",
        key_name: "target_type",
        disabled: !!ttype,
        options: [
          "user_app",
          "auth_app",
          "domain_widget_app",
          "domain_editor_app",
        ],
      },
      {
        name: "Target",
        ftype: "TEXT",
        key_name: "target",
        disabled: !!target,
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
        name: "Exec Domain",
        ftype: "INT",
        key_name: "exec_domain",
      },
      {
        name: "Exec Meta",
        ftype: "KEY_VALUE_TEXT",
        key_name: "exec_meta",
      },

      {
        name: "Extra Meta",
        ftype: "KEY_VALUE_TEXT",
        key_name: "extra_meta",
      },
    ],
    name: "New Target App",
    required_fields: [],
  }}
  onSave={save}
  {data}
/>
