<script lang="ts">
  import type { PortalService } from "../../core";
  import ActionPicker from "../../core/action_picker.svelte";
  export let service: PortalService;

  const actions = [
    {
      name: "datatable_hook",
      icon: "hashtag",
      info: "Datatable hook",
      action: () => {
        service.nav.admin_target_hook_new();
        service.utils.small_modal_close();
      },
    },

    {
      name: "domain_hook",
      icon: "hashtag",
      info: "Domain hook",
      action: () => {
        service.nav.admin_target_hook_new();
        service.utils.small_modal_close();
      },
    },

    {
      name: "startup",
      icon: "hashtag",
      info: "Startup hook when node starts",
      action: () => {
        service.nav.admin_target_hook_new();
        service.utils.small_modal_close();
      },
    },
  ];
</script>

<ActionPicker {actions} title="New Target hook type" />
