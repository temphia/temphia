<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import Tab from "./_tab.svelte";

  export let files = [];
  export let changeFile;
  export let file;

  export let onRun = () => {};
  export let onSave = () => {};
  export let onReset = () => {};
</script>

<div class="p-1 flex justify-between">
  <div class="grow h-10 flex flex-row flex-nowrap">
    <Tab name={file} onClick={changeFile} {files} />
  </div>

  <div class="flex gap-1 items-center">
    <button
      on:click={onRun}
      class="p-1 rounded bg-green-500 shadow hover:bg-green-900 flex text-white text-sm sm:text-base uppercase"
    >
      <Icon name="play" class="h-4 w-4 md:h-6 md:w-6" solid />
      run
    </button>

    <button
    on:click={onReset}
    class="p-1 rounded bg-red-400 shadow hover:bg-red-600 flex text-white uppercase"
  >
    <Icon name="refresh" class="h-4 w-4 md:h-6 md:w-6" solid />
    reset
  </button>

    <button
      on:click={onSave}
      class="p-1 rounded bg-blue-400 shadow hover:bg-blue-600 flex text-white uppercase"
    >
      <Icon name="save" class="h-4 w-4 md:h-6 md:w-6" solid />
      save
    </button>

    <button
      class="p-1 rounded bg-blue-200 shadow hover:bg-blue-600 flex text-white"
    >
      <Icon name="information-circle" class="h-4 w-4 md:h-6 md:w-6" solid />
    </button>
  </div>
</div>

<div class="flex flex-grow max-h-screen flex-col md:flex-row bg-white h-full">
  <div class="md:w-1/2 md:h-full h-1/2">
    <slot name="code" />
  </div>

  <div class="md:w-1/2 md:h-full h-1/2 flex flex-col gap-2 p-2 overflow-auto">
    <slot name="output" />
  </div>
</div>
