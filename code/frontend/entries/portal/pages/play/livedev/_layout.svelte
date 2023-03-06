<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import Tab from "./_tab.svelte";

  import { ClientJS, ServerJS } from "./conts";

  export let changeFile;
  export let file;
</script>

<div class="p-1 flex justify-between">
  <div class="grow h-10 flex flex-row flex-nowrap overflow-hidden">
    <Tab
      name="server.js"
      active={file === ServerJS}
      onClick={() => changeFile(ServerJS)}
    />
    <Tab
      name="client.js"
      active={file === ClientJS}
      onClick={() => changeFile(ClientJS)}
    />
  </div>

  <div class="flex gap-1 items-center">
    <button
      class="p-1 rounded bg-green-500 shadow hover:bg-green-900 flex text-white text-sm sm:text-base"
    >
      <Icon name="play" class="h-4 w-4 md:h-6 md:w-6" solid />
      SUBMIT
    </button>

    <button
      class="p-1 rounded bg-blue-300 shadow hover:bg-blue-600 flex text-white"
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
