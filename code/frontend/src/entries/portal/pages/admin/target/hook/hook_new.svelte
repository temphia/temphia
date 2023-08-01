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

  const nav_options = app.nav.options;
  if (nav_options) {
    data = { ...nav_options };
  }

  const save = async (_data) => {
    const __traget = _data["target_type"] || data["target_type"];
    const resp = await api.newHook(__traget, {
      ...data,
      ..._data,
    });
    if (!resp.ok) {
      message = resp.data;
      return;
    }
    app.nav.admin_target_hooks();
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
        name: "Handler",
        ftype: "TEXT",
        key_name: "handler",
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
    name: "New Target Hook",
    required_fields: [],
  }}
  onSave={save}
  {data}
/>
