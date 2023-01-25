<script lang="ts">
  import Kveditor from "../../../../../xcompo/common/kveditor.svelte";
  import { SheetColTypes, SheetColTypeText } from "../sheets";
  import Layout from "./_layout.svelte";

  export let onAdd = (name: string, ctype: string, opts: object) => {};

  let name = "";
  let ctype = SheetColTypeText;
  let options = {};
</script>

<Layout title="New Column" onClick={() => onAdd(name, ctype, options)}>
  <div class="mb-4">
    <label class="block mb-2 text-sm font-bold text-gray-700" for="name"
      >Name</label
    >
    <input
      class="w-full px-3 py-2 text-sm leading-tight text-gray-700 border rounded shadow appearance-none focus:outline-none focus:shadow-outline"
      id="name"
      type="text"
      bind:value={name}
      placeholder="name"
    />
  </div>

  <div class="mb-4">
    <label class="block mb-2 text-sm font-bold text-gray-700" for="type"
      >Type</label
    >

    <select
      id="type"
      bind:value={ctype}
      class="w-full px-3 py-2 text-sm leading-tight text-gray-700 border rounded shadow appearance-none focus:outline-none focus:shadow-outline"
    >
      {#each SheetColTypes as st}
        <option value={st}>{st}</option>
      {/each}
    </select>
  </div>

  <div class="mb-4">
    <label class="block mb-2 text-sm font-bold text-gray-700" for="opts"
      >Options</label
    >
    <Kveditor />
  </div>
</Layout>
