<script lang="ts">
  export let data: object;
  export let index: number;
  export let is_open: boolean;
  export let toggleFunc: () => void;
</script>

<button on:click|stopPropagation|preventDefault={toggleFunc}>
  <div class="transition-all {is_open ? 'arrow-down' : 'arrow-right'}" />
</button>

<span class="bg-blue-100 p-1 rounded">
  {index}
</span>

<span class="bg-pink-100 p-1 rounded">
  {data["log_event_id"] || ""}
</span>

{#if is_open}
  <pre class="bg-slate-100 rounded">
        {JSON.stringify(data, null, "\t")}
    </pre>
{:else}
  {#each Object.entries(data) as [k, v]}
    <div class="flex flex-nowrap">
      <span class="px-1 bg-slate-100 rounded">{k}</span>:
      {#if typeof v === "object"}
        {#if Array.isArray(v)}
          <span class="border border-slate-400 rounded">
            {#each v as vv}
              <span class="whitespace-nowrap">{vv}</span>
            {/each}
          </span>
        {:else}
          <span class="border border-slate-400 rounded">
            {#each Object.entries(v) as [kv, vv]}
              <span class="whitespace-nowrap">{kv} => {vv} </span>
            {/each}
          </span>
        {/if}
      {:else}
        <span class="whitespace-nowrap">{v}</span>
      {/if}
    </div>
  {/each}
{/if}

<style>
  .arrow-right {
    width: 0;
    height: 0;
    border-top: 0.5rem solid transparent;
    border-bottom: 0.5rem solid transparent;
    border-left: 0.5rem solid rgb(78, 78, 78);
  }

  .arrow-down {
    width: 0;
    height: 0;
    border-left: 0.5rem solid transparent;
    border-right: 0.5rem solid transparent;

    border-top: 0.5rem solid rgb(78, 78, 78);
  }
</style>
