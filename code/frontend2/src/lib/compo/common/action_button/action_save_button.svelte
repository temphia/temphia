<script>
  import ActionBtn from "./action_button.svelte";
  export let onClick;
  export let name = "Add";
  export let icon_name = "plus";
</script>

<ActionBtn {onClick} {name} {icon_name} />
