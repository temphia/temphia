<script>
  import ActionBtn from "./action_button.svelte";
  export let onClick;
  export let name = "Edit";
</script>

<ActionBtn {onClick} {name} icon_name={"pencil"} bclass="bg-green-400" />
