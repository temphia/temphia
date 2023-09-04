<script>
    import User from "./user.svelte";
    import { params } from "$lib/params";
</script>

<User id={$params["id"]} />
