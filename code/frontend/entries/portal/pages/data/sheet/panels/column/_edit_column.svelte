<script lang="ts">
  import type { SheetColumn } from "../../sheets";
  import Layout from "../row/_layout.svelte";

  export let column: SheetColumn;
</script>

<Layout title="Edit Column" onClick={() => {}}>
  <div class="mb-4">
    <label class="block mb-2 text-sm font-bold text-gray-700" for="name"
      >Name</label
    >
    <input
      class="w-full px-3 py-2 text-sm leading-tight text-gray-700 border rounded shadow appearance-none focus:outline-none focus:shadow-outline"
      id="name"
      type="text"
      value={column.name}
    />
  </div>
</Layout>
