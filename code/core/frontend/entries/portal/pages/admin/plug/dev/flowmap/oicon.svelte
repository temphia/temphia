<script>
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  export let name;
  export let classes = "h-5 w-5";
  const objects = {
    plug: "view-grid-add",
    agent: "users",
    resource: "cube",
    link: "link",
    ext: "puzzle",
  };
</script>

<Icon name={objects[name]} class={classes} />
