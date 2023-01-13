<script lang="ts">
  import Field from "./field.svelte";
  import type * as f from "./field";
  export let data: typeof f.data;
</script>

{#each data.items as item}
  <Field
    html_attr={item.html_attr}
    info=""
    name={item.name}
    options={item.options}
    type={item.type}
    value={null}
  />
{/each}

<div class="flex justify-end items-center p-1">
  <button
    class="p-1 rounded bg-green-500 shadow hover:bg-green-900 flex text-white"
  >
    <svg
      xmlns="http://www.w3.org/2000/svg"
      class="h-5 w-5"
      viewBox="0 0 20 20"
      fill="currentColor"
    >
      <path
        fill-rule="evenodd"
        d="M10 18a8 8 0 100-16 8 8 0 000 16zM9.555 7.168A1 1 0 008 8v4a1 1 0 001.555.832l3-2a1 1 0 000-1.664l-3-2z"
        clip-rule="evenodd"
      />
    </svg>
    Submit</button
  >
</div>
