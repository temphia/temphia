<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import type { LoadResponse, ExecData } from "../../service";
  import Editor from "./_editor.svelte";
  import ExecRows from "./_exec_rows.svelte";
  import Paramform from "./_paramform.svelte";
  import Tabbed from "./_tabbed.svelte";

  export let onSubmit = async (data): Promise<void> => {};
  export let data: LoadResponse;
  export let startup_payload: ExecData;
  export let message = "";

  let getParamData;
  let getCodeValue;
  let tabmode;

  let loading = false;

  const submit = async () => {
    if (loading) return;

    let code = getCodeValue();
    let param_data = {};
    if (getParamData) {
      param_data = getParamData();
    }

    loading = true;
    await onSubmit({
      code,
      param_data,
    });

    loading = false;
  };
</script>

<div class="rounded bg-white p-2 h-full">
  <p class="text-red-500">{message}</p>

  <Editor {data} bind:getCodeValue />

  <Tabbed modes={["Params", "Context"]} bind:mode={tabmode} />

  {#if tabmode === "Params"}
    <Paramform title={data.title} bind:getParamData />
  {:else}
    <ExecRows
      cells={startup_payload.cells}
      columns={startup_payload.columns}
      rows={startup_payload.rows}
    />
  {/if}

  <div class="flex flex-wrap justify-end text-sm items-center gap-2">
    <button
      on:click={submit}
      class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded text-sm px-4 py-2 flex"
    >
      {#if loading}
        <Icon name="globe" class="h-4 w-4 animate-bounce" solid />
      {/if}

      Submit</button
    >
  </div>
</div>
