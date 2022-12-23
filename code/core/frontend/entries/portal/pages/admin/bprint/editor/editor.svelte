<script lang="ts">
  import type { PortalService } from "../../core";
  import DataMigrator from "./data_migrator/index.svelte";
  import DataSchema from "./data_schema/index.svelte";

  export let bid: string;
  export let etype: string;
  export let file: string;

  // app.registry.GetAll("bprint_editor_type");
</script>

<div>Editor</div>
