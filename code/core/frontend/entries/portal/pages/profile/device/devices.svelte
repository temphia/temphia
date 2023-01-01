<script lang="ts">
  import { AutoTable, FloatingAdd, PortalService } from "../../admin/core";
  import { getContext } from "svelte";
  import PairDevice from "./_pair_device.svelte";
  import LoadingSpinner from "../../../../xcompo/common/loading_spinner.svelte";

  const app = getContext("__app__") as PortalService;

  const sapi = app.api_manager.get_self_api();

  let datas = [];
  let loading = true;

  const load = async () => {
    const resp = await sapi.list_devices();
    if (resp.status !== 200) {
      console.log("Err", resp);
      return;
    }

    datas = resp.data;
    loading = false;
  };

  load();
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <div class="p-4 w-full h-full bg-indigo-100">
    <AutoTable
      action_key="id"
      actions={[
        {
          Name: "Delete",
          Class: "bg-red-400",
          Action: async (id) => {
            await sapi.delete_device(id);
            load();
          },
        },
      ]}
      key_names={[
        ["id", "ID"],
        ["name", "Name"],
        ["device_type", "Device Type"],
        ["last_addr", "Last Address"],
      ]}
      color={["device_type"]}
      {datas}
    />
  </div>
{/if}

<FloatingAdd onClick={() => app.utils.small_modal_open(PairDevice, {})} />
