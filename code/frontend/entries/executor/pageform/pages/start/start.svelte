<script lang="ts">
  import Field from "../../field/field.svelte";
  import type { LoadResponse } from "../../service";
  import PlayIcon from "@krowten/svelte-heroicons/icons/PlayIcon.svelte";
  export let data: LoadResponse;

  const field_data = { ...(data.data || {}) };

  $: console.log("@field_data", field_data);
</script>

{#each data.items as item}
  <Field
    html_attr={item.html_attr}
    info={item.info}
    name={item.name}
    options={item.options}
    type={item.type}
    value={field_data[item.name]}
    onChange={(newval) => {
      field_data[item.name] = newval;
    }}
  />
{/each}

<div class="grow" />

<div class="flex justify-end items-center">
  <button
    on:click={() => {
      console.log("@submit");
    }}
    class="p-1 rounded bg-green-500 shadow hover:bg-green-900 flex text-white"
  >
    <PlayIcon class="h-6 w-6 mr-1" />
    Submit</button
  >
</div>
