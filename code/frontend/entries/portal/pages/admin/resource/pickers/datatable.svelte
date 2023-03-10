<script lang="ts">
  import { ResourceTypeDgroup } from "../../../../../../lib/entities";
  import { LoadingSpinner } from "../../../../../adapter_editor/easypage/core";
  import type { DataService } from "../../../../services/data";
  import type { PortalService } from "../../core";

  import ActionPicker from "../../core/action_picker.svelte";

  export let service: PortalService;

  let loading = true;

  let data_source = "";
  let data_group = "";
  let options = {};

  let sources = [];
  let groups = [];

  let mode: "pick_source" | "pick_group" = "pick_source";
  let ds: DataService;

  const load = async () => {
    ds = await service.get_data_service();
    sources = ds.sources;
    loading = false;
  };

  load();

  const pick_source = async (sdata) => {
    data_source = sdata["data"];
    loading = true;

    const dapi = service.api_manager.get_admin_data_api();
    const resp = await dapi.list_group(data_source);
    if (!resp.ok) {
      return;
    }
    groups = resp.data;

    mode = "pick_group";
    loading = false;
  };

  const pick_group = async (sdata) => {
    data_group = sdata["data"]["slug"];

    service.nav.admin_resource_new({
      type: ResourceTypeDgroup,
      target: `${data_source}/${data_group}`,
    });
    
    service.utils.small_modal_close()

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
    title="Pick Data Source"
  />
{:else if mode === "pick_group"}
  <ActionPicker
    actions={groups.map((group) => ({
      action: pick_group,
      icon: "database",
      info: group["description"],
      name: group["name"],
      data: group,
    }))}
    title="Pick Data Group"
  />
{:else}
  <div>end</div>
{/if}
