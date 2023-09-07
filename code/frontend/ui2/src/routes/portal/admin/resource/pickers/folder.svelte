<script lang="ts">
  import { ResourceFolder } from "./resources";
  import { LoadingSpinner } from "$lib/core";
  import type { CabinetService } from "$lib/services/cabinet/cabinet";
  import type { PortalService } from "$lib/core";

  import ActionPicker from "$lib/core/action_picker.svelte";

  export let service: PortalService;

  let loading = true;

  let source = "";
  let folder = "";

  let sources = [];
  let folders = [];

  let mode: "pick_source" | "pick_folder" = "pick_source";
  let cs: CabinetService;

  const load = async () => {
    cs = service.get_cabinet_service();
    sources = await cs.get_cab_sources();
    loading = false;
  };

  load();

  const pick_source = async (sdata) => {
    source = sdata["data"];
    loading = true;

    const dapi = service.api_manager.get_cabinet(source);
    const resp = await dapi.listRoot();
    if (!resp.ok) {
      return;
    }
    folders = resp.data;
    mode = "pick_folder";
    loading = false;
  };

  const pick_group = async (sdata) => {
    folder = sdata["data"];

    service.nav.admin_resource_new({
      type: ResourceFolder,
      target: `${source}/${folder}`,
    });

    service.utils.small_modal_close();
  };
</script>

{#if loading}
  <LoadingSpinner classes="" />
{:else if mode === "pick_source"}
  <ActionPicker
    actions={sources.map((src) => ({
      action: pick_source,
      icon: "hashtag",
      info: "",
      name: src,
      data: src,
    }))}
    title="Pick Cabinet Source"
  />
{:else if mode === "pick_folder"}
  <ActionPicker
    actions={folders.map((f) => ({
      action: pick_group,
      icon: "folder",
      info: "",
      name: f,
      data: f,
    }))}
    title="Pick Cabinet Folder"
  />
{:else}
  <div>end</div>
{/if}
