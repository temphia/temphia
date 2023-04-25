<script lang="ts">
  import type { LoadResponse, ExecData } from "../../service";
  import Editor from "./_editor.svelte";
  import ExecRows from "./_exec_rows.svelte";
  import Paramform from "./_paramform.svelte";
  import Tabbed from "./_tabbed.svelte";

  export let onSubmit = async (data): Promise<any> => {};
  export let onNext = (data) => {};
  export let data: LoadResponse;
  export let startup_payload: ExecData;

  let getParamData;
  let getCodeValue;
  let tabmode;

  let loading = false;
  let message = "";

  const submit = async () => {
    let code = getCodeValue();
    let param_data = {};
    if (getParamData) {
      param_data = getParamData();
    }

    loading = true;
    const resp = await onSubmit({
      code,
      param_data,
    });

    if (resp["ok"]) {
      onNext(resp["data"]);
      return;
    }

    message = resp.data;
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
      Submit</button
    >
  </div>
</div>
