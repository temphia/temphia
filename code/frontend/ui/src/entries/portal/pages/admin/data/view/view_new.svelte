<script lang="ts">
  import { getContext, setContext } from "svelte";
  import { LoadingSpinner, PortalService } from "../../core";
  import { params } from "svelte-hash-router";
  import ValueInner from "./_value_inner.svelte";

  export let source = $params.source;
  export let group = $params.group;
  export let table = $params.table;

  const app = getContext("__app__") as PortalService;
  const api = app.api_manager.get_admin_data_api();

  let message = "";
  let loading = true;
  let columns = [];
  const save = async (_data) => {

    if (!_data["selects"]) {
      _data["selects"] = columns.map((v) => v["slug"]);
    }

    const resp = await api.add_view(source, group, table, _data);
    if (!resp.ok) {
      message = resp.data;
      return;
    }
    app.nav.admin_data_views(source, group, table);
  };

  const load = async () => {
    const dapi = app.api_manager.get_admin_data_api();
    const resp = await dapi.list_column(source, group, table);
    if (!resp.ok) {
      message = resp.data;
      loading = false;
      return;
    }
    columns = resp.data;
    loading = false;
  };

  load();

  setContext("__data_context__", {
    get_modal: () => ({
      open: app.utils.small_modal_open,
      close: app.utils.small_modal_close,
    }),
    table_service: null,
  });
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <div class="h-full w-full bg-indigo-100 p-10 overflow-auto">
    <div class="p-5 bg-white w-full ">
      <div class="text-2xl text-indigo-900">Add View</div>
      <p class="text-red-500">{message || ""}</p>

      <ValueInner {columns} data={{}} onSave={save} />
    </div>
  </div>
{/if}
