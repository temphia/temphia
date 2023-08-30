<script>
    import ActionBtn from "./action_button.svelte";
    export let onClick;
  </script>
  
  <ActionBtn {onClick} name="Select" icon_name={"check"} bclass="bg-green-400" />
  