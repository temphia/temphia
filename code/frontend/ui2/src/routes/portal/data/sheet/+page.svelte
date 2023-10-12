<script lang="ts">
    import SheetInner from "./_sheet.svelte";
    import { params } from "$lib/params";

    $: __group = $params["dgroup"] || $params["group"];
    $: __sheet = $params["dsheet"] || $params["sheetid"];
    $: __source = $params["dsource"] || $params["source"] || "default";
</script>

{#key __sheet}
    <SheetInner group={__group} sheetid={__sheet} source={__source} />
{/key}
