<script lang="ts">
  import { getContext } from "svelte";
  import Issuer from "./issuer/issuer.svelte";
  import InstanceBundlePicker from "./instancer/bundle_picker.svelte";

  import {
    AutoTable,
    LoadingSpinner,
    FloatingAdd,
    PortalService,
  } from "../core";
  import { instance_helper } from "./instancer";

  let datas = [];
  let loading = true;
  const app = getContext("__app__") as PortalService;

  const load = async () => {
    const api = app.api_manager.get_admin_bprint_api();
    const resp = await api.list();
    if (!resp.ok) {
      return;
    }

    datas = resp.data;
    loading = false;
  };

  load();

  // actions

  const action_instance = (id: string) => {
    const bprint = datas.filter((v) => v.id === id)[0];
    const file = bprint["files"].filter(
      (v) => v !== "schema.json" || v !== "schema.yaml"
    )[0];

    instance_helper(app, bprint["type"], bprint, file, InstanceBundlePicker);
  };

  const action_edit = (id: string) => app.nav.admin_bprint(id);
  const action_issue = (id: string) =>
    app.utils.small_modal_open(Issuer, { app, bid: id });

  const action_file_edit = (id: string) => {};
  

  const action_delete = async (id: string) => {
    const api = app.api_manager.get_admin_bprint_api();
    await api.delete(id);
    load();
  };
  const action_new = async () => {};
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <AutoTable
    show_drop={true}
    color={["type"]}
    action_key="id"
    actions={[
      {
        Name: "Instance",
        Class: "bg-blue-400",
        icon: "document-download",
        Action: action_instance,
      },
      {
        Name: "Edit",
        Action: action_edit,
        icon: "pencil-alt",
        drop: true,
      },
      {
        Name: "Issue",
        Action: action_issue,
        drop: true,
        icon: "terminal",
      },

      {
        Name: "Open File Editor",
        Action: action_file_edit,
        drop: true,
        icon: "beaker",
      },

      {
        Name: "Delete",
        drop: true,
        icon: "trash",
        Action: action_delete,
      },
    ]}
    key_names={[
      ["id", "ID"],
      ["name", "Name"],
      ["slug", "Slug"],
      ["type", "Type"],
      ["sub_type", "Sub Type"],
    ]}
    {datas}
  />
{/if}

<FloatingAdd onClick={action_new} />
