<script>
  import ActionBtn from "./action_button.svelte";
  export let onClick;
  export let name = "Add";
</script>

<ActionBtn {onClick} {name} icon_name="plus" />
