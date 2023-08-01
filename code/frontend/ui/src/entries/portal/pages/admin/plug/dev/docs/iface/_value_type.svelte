<script type="ts">
  import type { AgentIface, EventType, Method, ValueType } from "./iface";
  export let data: ValueType;
</script>

{#if data.ref}
  <span class="text-green-800 bg-slate-200 rounded">{data.ref}</span>
  <a href="#schema-{data.ref}">{data.ref}</a>
{:else}
  <span class="text-green-800 bg-slate-200 rounded">{data.type}</span>
  <p class="">{data.property || ""}</p>
{/if}
