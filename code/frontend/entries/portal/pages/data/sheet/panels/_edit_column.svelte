<script lang="ts">
  import Layout from "./_layout.svelte";
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
      placeholder="name"
    />
  </div>
</Layout>
