<script>
  import App from "./app.svelte";

  export const run_app = (app) => {};
  export const close_app = (name) => {};

  let live_apps = [{ name: "test1" }, { name: "test2" }];
  let active_app = "test1";
</script>

<div class="p-4 shadow rounded border m-2">
  {#each live_apps as lapp}
    <div style="display: {lapp['name'] === active_app ? 'block' : 'none'};">
      <App name={lapp["name"]} />
    </div>
  {/each}
</div>

<h2>APP BAR</h2>

<div class="flex p-2 border shadow">
  {#each live_apps as lapp}
    <button
      class="p-2 border shadow rounded"
      on:click={() => {
        active_app = lapp["name"];
      }}
    >
      {lapp["name"]}
    </button>
  {/each}
</div>

<button
  class="rounded bg-slate-400 border "
  on:click={() => {
    live_apps = [...live_apps, { name: "TEST3" }];
  }}
>
  Add
</button>
