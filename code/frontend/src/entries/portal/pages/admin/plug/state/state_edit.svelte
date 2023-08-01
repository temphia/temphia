<script lang="ts">
  import { getContext } from "svelte";
  import { AutoForm, LoadingSpinner, PortalService } from "../../core";
  import { params } from "svelte-hash-router";

  export let pid = $params.pid;
  export let skey = $params.skey;

  const app = getContext("__app__") as PortalService;
  const api = app.api_manager.get_admin_plug_api();

  let message = "";
  let data = {};
  let loading = true;

  const load = async () => {
    const resp = await api.get_plug_state(pid, skey);
    if (!resp.ok) {
      message = resp.data;
      return;
    }

    data = resp.data;
    loading = false;
  };

  const save = async (_data) => {
    console.log("@@data", _data);

    const resp = await api.update_plug_state(pid, skey, {
      value: _data["value"] || data["value"],
      options: {
        force_ver: _data["force_ver"] || false,
        with_version: _data["with_version"] || false,
        version: _data["version"] || 0,
        set_tag1: _data["set_tag1"] || false,
        set_tag2: _data["set_tag2"] || false,
        set_tag3: _data["set_tag3"] || false,
        tag1: _data["tag1"] || "",
        tag2: _data["tag3"] || "",
        tag3: _data["tag3"] || "",
      },
    });
    if (!resp.ok) {
      message = resp.data;
      return;
    }
    app.nav.admin_plug_states(pid);
  };

  load();
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <AutoForm
    {message}
    schema={{
      fields: [
        {
          name: "Key",
          ftype: "TEXT_SLUG",
          key_name: "key",
          disabled: true,
        },
        {
          name: "Value",
          ftype: "LONG_TEXT",
          key_name: "value",
        },

        {
          name: "Set Tag 1",
          ftype: "BOOL",
          key_name: "set_tag1",
        },
        {
          name: "Tag 1",
          ftype: "TEXT",
          key_name: "tag1",
        },

        {
          name: "Set Tag 2",
          ftype: "BOOL",
          key_name: "set_tag2",
        },
        {
          name: "Tag 2",
          ftype: "TEXT",
          key_name: "tag2",
        },

        {
          name: "Set Tag 3",
          ftype: "BOOL",
          key_name: "set_tag3",
        },
        {
          name: "Tag 3",
          ftype: "TEXT",
          key_name: "tag3",
        },
        {
          name: "Force Version",
          ftype: "BOOL",
          key_name: "force_ver",
        },
        {
          name: "With Version",
          ftype: "BOOL",
          key_name: "with_version",
        },

        {
          name: "Version",
          ftype: "INT",
          key_name: "version",
        },
      ],
      name: "Edit Plug State",
      required_fields: [],
    }}
    onSave={save}
    {data}
  />
{/if}
