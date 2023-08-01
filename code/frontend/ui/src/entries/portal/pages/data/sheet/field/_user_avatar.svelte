<script>
  import Avatar from "../../shared/user/avatar.svelte";
  export let name = undefined;
  export let url = undefined;
</script>

<Avatar user={name} {url} />
