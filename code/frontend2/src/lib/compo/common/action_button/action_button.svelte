<script>
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  export let onClick;
  export let name = "";
  export let icon_name = ""
  export let bclass = "bg-blue-400";
  let loading = false;

  const handler = async () => {
    loading = true;
    const r = onClick();
    if (r instanceof Promise) {
      await r;
    }
    loading = false;
  };
</script>

<button
  on:click={handler}
  class="p-1 text-white text-sm font-semibold flex self-center shadow rounded hover:scale-110 {bclass}"
>
  <Icon name={icon_name} class="h-5 w-5 {loading ? 'animate-spin' : ''}" />

  {name}
</button>
