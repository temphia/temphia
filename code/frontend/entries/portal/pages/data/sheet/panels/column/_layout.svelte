<script lang="ts">
  import Button from "../row/edit/_button.svelte";
  export let title;
  export let onSave;
  export let onDelete = undefined;
</script>

<div class="flex flex-col h-full">
  <div class="mb-4 border-b">
    <h1 class="block mb-2 text-lg font-bold text-gray-700">
      {title}
    </h1>
  </div>

  <div class="flex flex-col flex-grow gap-2">
    <slot />
  </div>

  <div class="flex justify-end gap-2 border-t">
    {#if onDelete}
      <Button name="Delete" onClick={onDelete} color="red" />
    {/if}

    <Button name="Save" onClick={onSave} />
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
