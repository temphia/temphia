<script lang="ts">
  import type { PortalService } from "../admin/core";
  import ActionPicker from "../admin/core/action_picker.svelte";

  export let app: PortalService;

  const resourceTypes = [
    {
      name: "Quick Sheets",
      icon: "table",
      info: "New Data sheets for storing data.",
      action: () => {
        app.utils.small_modal_close();        
        app.nav.data_sheets_new()

      },
    },
    {
      name: "From Store",
      icon: "shopping-bag",
      info: "You can import required data tables from store, if you have proper permissions.",
      action: () => {
        app.utils.small_modal_close();
        app.nav.repo_loader();
      },
    },
    {
      name: "From Builder",
      icon: "document-add",
      info: "You can design schema for data table but it needs knowledge about database system.",
      action: () => {
        app.utils.small_modal_close();
      },
    },
  ];
</script>

<ActionPicker actions={resourceTypes} title="Create new data tables." />
