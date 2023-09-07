<script>
  import Avatar from "../../_shared/user/avatar.svelte";
  export let name = undefined;
  export let url = undefined;
</script>

<Avatar user={name} {url} />
