<script lang="ts">
  import Field from "../../field/field.svelte";
  import type { PageFormService, Response } from "../../service";
  import PlayIcon from "@krowten/svelte-heroicons/icons/PlayIcon.svelte";
  import { createEventDispatcher } from "svelte";

  export let data: Response;

  const field_data = { ...(data.data || {}) };

  const dispatcher = createEventDispatcher();

  $: console.log("@field_data", field_data);
</script>

{#if data.ok}
  {#if data.message}
    <div class="p-4 rounded bg-gray-50">
      <p>{data.message}</p>
    </div>
  {/if}

  {#each (data.items || []) as item}
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

  {#if !data.final}
    <div class="flex justify-end items-center">
      <button
        on:click={() => dispatcher("submit", field_data)}
        class="p-1 rounded bg-green-500 shadow hover:bg-green-900 flex text-white"
      >
        <PlayIcon class="h-6 w-6 mr-1" />
        Submit</button
      >
    </div>
  {/if}
{:else}
  <div class="bg-red-50 rounded p-4">
    <p class="text-red-500 ">{data.message}</p>
  </div>
{/if}
