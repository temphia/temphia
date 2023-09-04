<script lang="ts">
    import { LoadingSpinner } from "$lib/compo";

    window.addEventListener(
        "message",
        (ev) => {
            console.log("@got_message");
        },
        false
    );
</script>

<LoadingSpinner />
