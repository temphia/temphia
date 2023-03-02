<script lang="ts">
  import { getContext, setContext } from "svelte";
  import { LoadingSpinner, PortalService } from "../../core";
  import { params } from "svelte-hash-router";
  import ValueInner from "./_value_inner.svelte";

  export let source = $params.source;
  export let group = $params.group;
  export let table = $params.table;
  export let id = $params.id;

  const app = getContext("__app__") as PortalService;
  const api = app.api_manager.get_admin_data_api();

  let message = "";
  let data = {};
  let loading = true;
  let columnsLoadings = true;
  let columns = [];

  const load = async () => {
    const resp = await api.get_view(source, group, table, id);
    if (!resp.ok) return;

    data = resp.data;
    loading = false;
  };

  const loadColumns = async () => {
    const dapi = app.api_manager.get_admin_data_api();
    const resp = await dapi.list_column(source, group, table);
    if (!resp.ok) {
      message = resp.data;
      columnsLoadings = false;
      return;
    }
    columns = resp.data;
    columnsLoadings = false;
  };

  load();
  loadColumns();

  const save = async (_data) => {
    const resp = await api.edit_view(source, group, table, id, _data);
    if (!resp.ok) {
      message = resp.data;
      return;
    }
    app.nav.admin_data_views(source, group, table);
  };

  setContext("__data_context__", {
    get_modal: () => ({
      open: app.utils.small_modal_open,
      close: app.utils.small_modal_close,
    }),
    table_service: null,
  });
</script>

{#if loading || columnsLoadings}
  <LoadingSpinner />
{:else}
  <div class="h-full w-full bg-indigo-100 p-10 overflow-auto">
    <div class="p-5 bg-white w-full ">
      <div class="text-2xl text-indigo-900">Edit View</div>
      <p class="text-red-500">{message || ""}</p>

      <ValueInner {columns} {data} onSave={(_data) => save(_data)} />
    </div>
  </div>
{/if}
