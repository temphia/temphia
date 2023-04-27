<script lang="ts">
  import type { Element } from "../../service";
  export let element: Element;
  export let data;
</script>

<pre class="p-1 rounded bg-gray-100 overflow-auto">
    <code>{JSON.stringify(data, null, 4)}</code>
</pre>
