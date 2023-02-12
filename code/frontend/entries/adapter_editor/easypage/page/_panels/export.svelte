<script lang="ts">
  import type { EasypageService } from "../../service/easypage";

  export let service: EasypageService;

  let export_as = "download";
  let name = `site_${new Date().toLocaleDateString()}.zip`;
</script>

<div class="flex flex-col">
  <div class="mb-4">
    <label class="block mb-2 text-sm font-bold text-gray-700" for="type">
      Export As
    </label>

    <select
      bind:value={export_as}
      class="w-full px-3 py-2 text-sm leading-tight text-gray-700 border rounded shadow appearance-none focus:outline-none focus:shadow-outline"
      id="type"
    >
      <option value="download">Zip Download</option>
      <option value="folder_zip">Zip to Cabinet Folder</option>
      <option value="folder">Cabinet Folder</option>
    </select>
  </div>

  {#if export_as !== "folder"}
    <div class="mb-4">
      <label class="block mb-2 text-sm font-bold text-gray-700" for="name">
        Name
      </label>
      <input
        class="w-full px-3 py-2 text-sm leading-tight text-gray-700 border rounded shadow appearance-none focus:outline-none focus:shadow-outline"
        id="name"
        type="text"
        bind:value={name}
        placeholder="name"
      />
    </div>
  {/if}

  <button
    class="w-full px-4 py-2 font-bold text-white bg-blue-500 rounded-full hover:bg-blue-700 focus:outline-none focus:shadow-outline"
    type="button"
  >
    Start
  </button>
</div>
