<script lang="ts">
  import StartPage from "./pages/start/start.svelte";

  import { data } from "./field/field";
  import Tailwind from "../xcompo/common/_tailwind.svelte";
  export let env: any;
</script>

<div class="h-full w-full p-4 bg-blue-50">
  <div class="p-4 bg-white rounded">
    <h1
      class="text-gray-800 font-lg font-bold tracking-normal leading-tight mb-4"
    >
      Form 1
    </h1>
    <StartPage {data} />
  </div>
</div>

<Tailwind />
