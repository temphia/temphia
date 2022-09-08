<script lang="ts">
  export let source = "";
  export let path = "";
  export let sources = [];
</script>

<div class="w-full h-screen p-2">
  <div class="h-full w-full flex flex-col rounded border">
    <header class="flex-shrink h-12 border-b pb-1">
      <div class="flex justify-between border-b p-1">
        <div class="flex">
          <slot name="actions_left" />
        </div>

        <div class="flex">
          <slot name="actions_right" />
          <select class="rounded p-2 bg-white border">
            {#each sources as src}
              <option on:click={() => {}}> {src}</option>
            {/each}
          </select>
        </div>
      </div>
    </header>
    <div class="flex-grow h-32 p-2 overflow-auto">
      <div class="bg-gray-50 h-full w-full rounded">
        <slot name="body" />
      </div>
    </div>
  </div>
</div>
