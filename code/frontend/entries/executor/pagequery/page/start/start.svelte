<script lang="ts">
  import Ceditor from "../../../../xcompo/ceditor/ceditor.svelte";
  import Paramform from "./paramform/paramform.svelte";

  export let onSubmit = async (data): Promise<any> => {};
  export let onNext = (data) => {};

  let editor;
  let getParamData;

  let show = false;
  let loading = false;
  let message = "";

  let code = ``;
</script>

<div class="rounded bg-white p-2">
  <div class="flex flex-wrap justify-end text-sm items-center gap-1 text-gray-700">
    <label>
      Script

      <select class="p-1 rounded w-32">
        <option>Do Xyz</option>
        <option>Do Xyz</option>
      </select>
    </label>
    <label>
      <input bind:checked={show} type="checkbox" class="" />
    </label>
  </div>

  <p class="text-red-500">{message}</p>

  {#if show}
    <div class="p-1 flex-grow">
      <Ceditor bind:editor {code} />
    </div>
  {/if}

  <Paramform bind:getParamData />

  <div class="flex flex-wrap justify-end text-sm items-center">
    <button
      on:click={async () => {
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
        if (!resp["ok"]) {
          message = resp.data;
          loading = false;
        } else {
          onNext(resp["data"]);
        }
      }}
      class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded text-sm px-4 py-2 flex"
    >
      Search</button
    >
  </div>
</div>
