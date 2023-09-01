<script>
    import User from "./user.svelte";

    export let data;
</script>

<User id={data["user"]} />
