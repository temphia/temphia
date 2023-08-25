<script>
  import ActionBtn from "./action_button.svelte";
  export let onClick;
  export let name = "Delete";
</script>

<ActionBtn {onClick} {name} icon_name="trash" bclass="bg-red-400" />
