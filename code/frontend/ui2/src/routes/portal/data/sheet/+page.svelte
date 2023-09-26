<script lang="ts">
    import SheetInner from "./_sheet.svelte";
    import { params } from "$lib/params";
</script>

{#key $params["sheet"]}
    <SheetInner
        group={$params["dgroup"]}
        sheetid={$params["dsheet"]}
        source={$params["dsource"]}
    />
{/key}
