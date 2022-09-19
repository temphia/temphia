<script lang="ts">
  import type { Column } from "../../manager/dtypes";
  import { commaArrayDecode, commaArryEncode } from "../field";
  export let value: any;
  export let column: Column;
  export let onChange: (_value: any) => void;

  let options = column.options || [];

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

<div class="flex flex-col w-full h-full p-1 overflow-auto">
  <div
    class="flex flex-col pl-4 p-1 space-y-1 border border-dashed rounded-lg bg-gray-50 text-gray-800"
    style="min-height: 2rem;"
  >
    {#each options as opt}
      <label>
        <input
          type="checkbox"
          on:change={onOptChange(opt)}
          class="form-checkbox h-5 w-5 text-gray-600"
        />
        {opt}
      </label>
    {/each}
  </div>
</div>
