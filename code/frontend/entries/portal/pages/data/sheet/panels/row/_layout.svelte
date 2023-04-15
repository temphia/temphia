<script lang="ts">
  import Button from "./_button.svelte";

  export let title;
  export let onSave;
  export let onDelete = undefined

  export let new_record = true;
  export let mode: "EDIT" | "RELATIONS" | "HISTORY" = "EDIT";

  export const setMode = (newmode) => {
    if (mode !== newmode) {
      mode = newmode;
    }
  };

  const modes = ["EDIT", "RELATIONS", "HISTORY"];

  let saving = false;
  let deleting = false;
</script>

<div class="flex flex-col h-full">
  <div class="mb-4 border-b">
    {#if new_record}
      <h1 class="block mb-2 text-lg font-bold text-gray-700">
        {title}
      </h1>
    {:else}
      <div class="flex flex-row border flex-nowrap overflow-auto">
        {#each modes as m}
          {#if mode == m}
            <button
              class="p-2 hover:text-blue-500 focus:outline-none text-blue-500 border-b-2 font-medium border-blue-500 inline-flex"
              >{mode}
            </button>
          {:else}
            <button
              on:click={() => setMode(m)}
              class="text-gray-600 p-2 block hover:text-blue-500 focus:outline-none"
              >{m}</button
            >
          {/if}
        {/each}
      </div>
    {/if}
  </div>

  {#if new_record}
    <div class="flex flex-col flex-grow gap-2">
      <slot />
    </div>
  {:else}
    <div class="flex flex-col flex-grow gap-2">
      {#if mode === "EDIT"}
        <slot name="edit">Empty Edit template</slot>
      {:else if mode === "RELATIONS"}
        <slot name="relation">Empty relations template</slot>
      {:else if mode === "HISTORY"}
        <slot name="history">Empty history template</slot>
      {/if}
    </div>
  {/if}

  <div class="flex justify-end gap-2 border-t">
    {#if new_record}
      <Button name="Save" onClick={onSave} loading={saving} />
    {:else}
      <Button name="Save" onClick={onSave} loading={saving} />
      <Button name="Delete" onClick={onDelete} loading={deleting} color="red" />
    {/if}
  </div>
</div>

<div class="hidden">
  <div class="border-l-red-400" />
  <div class="border-l-blue-400" />
  <div class="border-l-green-400" />
  <div class="border-l-gray-400" />
  <div class="border-l-yellow-400" />
  <div class="border-l-zinc-400" />
  <div class="border-l-emerald-400" />
  <div class="border-l-orange-400" />
  <div class="border-l-lime-400" />
  <div class="border-l-amber-400" />
  <div class="border-l-teal-400" />
  <div class="border-l-cyan-400" />
  <div class="border-l-sky-400" />
  <div class="border-l-indigo-400" />
  <div class="border-l-violet-400" />
  <div class="border-l-purple-400" />
  <div class="border-l-fuchsia-400" />
  <div class="border-l-pink-400" />

  <!--  -->

  <div class="bg-red-400" />
  <div class="bg-blue-400" />
  <div class="bg-green-400" />
  <div class="bg-gray-400" />
  <div class="bg-yellow-400" />
  <div class="bg-zinc-400" />
  <div class="bg-emerald-400" />
  <div class="bg-orange-400" />
  <div class="bg-lime-400" />
  <div class="bg-amber-400" />
  <div class="bg-teal-400" />
  <div class="bg-cyan-400" />
  <div class="bg-sky-400" />
  <div class="bg-indigo-400" />
  <div class="bg-violet-400" />
  <div class="bg-purple-400" />
  <div class="bg-fuchsia-400" />
  <div class="bg-pink-400" />
</div>
