<script lang="ts">
    import { getContext } from "svelte";
    import { LoadingSpinner, PortalService } from "$lib/core";
    import { params } from "$lib/params";
    import { isImage } from "$lib/utils";

    export let source = $params["source"] || "default";
    export let file = $params["file"];
    export let folder = $params["folder"];

    const app = getContext("__app__") as PortalService;

    if (isImage(file || "")) {
        app.nav.cab_image_file(source, folder, file);
    } else {
        app.nav.cab_text_file(source, folder, file);
    }
</script>

<LoadingSpinner />
