<script lang="ts">
  import Kveditor from "$lib/compo/common/kveditor.svelte";
  import Layout from "./_layout.svelte";

  export let onAdd = async (name: string, opts: object) => {};

  let name = "";

  let options = {};
</script>

<Layout title="New Sheet" onClick={() => onAdd(name, options)}>
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
    <label class="block mb-2 text-sm font-bold text-gray-700" for="opts"
      >Options</label
    >
    <Kveditor />
  </div>
</Layout>
