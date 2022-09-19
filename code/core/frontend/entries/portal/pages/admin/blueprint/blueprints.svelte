<script lang="ts">
  import { AutoTable, FloatingAdd } from "../../../../../shared";
  import Layout from "../layout.svelte";
  import BprintEditor from "./editor/editor.svelte";
  import { getContext } from "svelte";
  import type { PortalApp } from "../../../app";
  import Issuer from "./issuer/issuer.svelte";
  import BuildPicker from "./_build_picker.svelte";
  import { instance } from "./bprint_util";
  import InstanceBundlePicker from "./_instance_bundle_picker.svelte";

  const app: PortalApp = getContext("__app__");

  let bprints = [];

  const load = async () => {
    const api = await app.get_apm().get_bprint_api();
    const resp = await api.bprint_list();
    bprints = resp.data;
  };

  load();

  const do_instance = async (id: string) => {
    const bprint = bprints.filter((v) => v.id === id)[0];
    const file = bprint["files"].filter(
      (v) => v !== "schema.json" || v !== "schema.yaml"
    )[0];

    instance(app, bprint["type"], bprint, file, InstanceBundlePicker);
  };
</script>

<Layout>
  <AutoTable
    show_drop={true}
    color={["type"]}
    action_key="id"
    actions={[
      {
        Name: "Instance",
        Class: "bg-blue-400",
        icon: "document-download",
        Action: do_instance,
      },
      {
        Name: "Edit",
        Action: app.navigator.goto_admin_bprint_page,
        icon: "pencil-alt",
        drop: true,
      },
      {
        Name: "Issue",
        Action: (bid) => {
          app.simple_modal_open(Issuer, { app, bid });
        },
        drop: true,
        icon: "terminal",
      },

      {
        Name: "Open File Editor",
        Action: (bid) => {
          app.simple_modal_open(BuildPicker, { app, bid });
        },
        drop: true,
        icon: "beaker",
      },

      {
        Name: "Delete",
        drop: true,
        icon: "trash",
        Action: async (id) => {
          const api = await app.get_apm().get_bprint_api();
          const resp = await api.bprint_remove(id);
        },
      },
    ]}
    key_names={[
      ["id", "ID"],
      ["name", "Name"],
      ["slug", "Slug"],
      ["type", "Type"],
      ["sub_type", "Sub Type"],
    ]}
    datas={bprints}
  />
</Layout>

<FloatingAdd onClick={() => app.big_modal_open(BprintEditor, {})} />
