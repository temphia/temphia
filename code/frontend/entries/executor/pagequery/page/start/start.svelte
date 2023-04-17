<script lang="ts">
  import Ceditor from "../../../../xcompo/ceditor/ceditor.svelte";
  import type { LoadResponse } from "../../service";

  export let onSubmit = async (data): Promise<any> => {};
  export let onNext = (data) => {};
  export let data: LoadResponse;
  export let startup_payload;

  let editor;
  let getParamData;

  let show = false;
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
  <div
    class="flex flex-wrap justify-end text-sm items-center gap-1 text-gray-700"
  >
    <label>
      Script

      <select class="p-1 rounded w-32" value={data["first_stage"]}>
        {#each Object.keys(data["stages"]) as skey}
          <option>{skey}</option>
        {/each}
      </select>
    </label>
    <label>
      <input bind:checked={show} type="checkbox" class="" />
    </label>
  </div>

  <p class="text-red-500">{message}</p>

  <div class="p-1 flex-grow">
    <Ceditor bind:editor {code} />
  </div>

  <div />

  <div class="flex flex-wrap justify-end text-sm items-center gap-2">
    <div>
      <details>
        <summary> Exec Data </summary>

        <code class="p-2 rounded bg-gray-100">
          <pre>
            {JSON.stringify(startup_payload)}
          </pre>
        </code>
      </details>
    </div>

    <button
      on:click={submit}
      class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded text-sm px-4 py-2 flex"
    >
      Submit</button
    >
  </div>
</div>
