<script lang="ts">
  import type { PortalService } from "../../core";
  import ActionPicker from "../../core/action_picker.svelte";
  export let service: PortalService;

  const actions = [
    {
      name: "user_group_app",
      icon: "hashtag",
      info: "App for group of people",
      action: () => {
        service.nav.admin_target_app_new();
        service.utils.small_modal_close()
      },
    },

    {
      name: "datatable_widget",
      icon: "hashtag",
      info: "Datatable widget",
      action: () => {},
    },

    {
      name: "datasheet_widget",
      icon: "hashtag",
      info: "DataSheet widget",
      action: () => {},
    },

    {
      name: "domain_widget",
      icon: "hashtag",
      info: "Domain widget",
      action: () => {},
    },
  ];
</script>

<ActionPicker {actions} title="New Target app type" />
