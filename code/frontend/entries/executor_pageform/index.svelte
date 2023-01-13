<script lang="ts">
  import PageForm from "./page_form.svelte";
  import { data } from "./field";
</script>

<div class="h-full w-full p-4 bg-blue-50">
  <div class="p-4 bg-white rounded">
    <h1
      class="text-gray-800 font-lg font-bold tracking-normal leading-tight mb-4"
    >
      Form 1
    </h1>
    <PageForm {data} />
  </div>
</div>
