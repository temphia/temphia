<script lang="ts">
  import { getContext } from "svelte";
  import { params } from "svelte-hash-router";
  import { LoadingSpinner, PortalService } from "../../core";

  export let pid = $params.pid;

  let datas = [];
  let loading = true;
  let page = 0;
  const app = getContext("__app__") as PortalService;
  const api = app.api_manager.get_admin_plug_api();

  const load = async () => {
    const resp = await api.list_plug_state(pid, page);
    if (!resp.ok) {
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
  <div class="p-2 flex flex-col">
    {#each datas as data}
      <div>
        <pre>
            {JSON.stringify(data)}
        </pre>
      </div>
    {/each}
  </div>
{/if}
