<script lang="ts">
  import Ceditor from "../../../../xcompo/ceditor/ceditor.svelte";
  import type { LoadResponse, ExecData } from "../../service";
  import ExecRows from "./_exec_rows.svelte";
  import Paramform from "./_paramform.svelte";
  import Templates from "./_templates.svelte";
  export let onSubmit = async (data): Promise<any> => {};
  export let onNext = (data) => {};
  export let data: LoadResponse;
  export let startup_payload: ExecData;

  let editor;
  let getParamData;

  let loading = false;
  let message = "";

  let code = ``;

  const submit = async () => {
    let code = "";
    if (editor) {
      code = editor.getValue();
    }

    const param_data = getParamData();

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

  <div class="p-1 flex-grow">
    <Ceditor bind:editor {code} />
  </div>

  <Templates {data} />

  <Paramform title={data.title} bind:getParamData />

  <ExecRows
    cells={startup_payload.cells}
    columns={startup_payload.columns}
    rows={startup_payload.rows}
  />

  <div class="flex flex-wrap justify-end text-sm items-center gap-2">
    <button
      on:click={submit}
      class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded text-sm px-4 py-2 flex"
    >
      Submit</button
    >
  </div>
</div>
