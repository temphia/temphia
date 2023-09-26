<script lang="ts">
  import {
    commaArrayDecode,
    commaArryEncode,
  } from "../../datatable/core/fields/field";

  export let value: any;
  export let onChange: (_value: any) => void;
  export let options: string[];

  $: _selected = commaArrayDecode(value);

  const onOptChange = (opt: string) => (ev) => {
    if (ev.target.checked) {
      _selected = [...new Set([..._selected, opt])];
    } else {
      _selected = [...new Set(_selected.filter((val) => val !== opt))];
    }
    onChange(commaArryEncode(_selected));
  };
</script>

<div class="flex flex-col w-full p-1 overflow-auto">
  <div
    class="flex flex-col pl-4 p-1 space-y-1 border border-dashed rounded-lg bg-gray-50 text-gray-800"
    style="min-height: 2rem;"
  >
    {#each options as opt}
      <label>
        <input
          type="checkbox"
          checked={_selected.includes(opt)}
          on:change={onOptChange(opt)}
          class="form-checkbox h-5 w-5 text-gray-600"
        />
        {opt}
      </label>
    {/each}
  </div>
</div>
