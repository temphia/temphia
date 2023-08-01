<script lang="ts">
  import type { PortalService } from "../../core";
  import ActionPicker from "../../core/action_picker.svelte";
  import {
    TargetHookTypeDataTableHook,
    TargetHookTypeDomainHook,
    TargetHookTypeStartUp,
  } from "../target";
  export let service: PortalService;

  const actions = [
    {
      name: TargetHookTypeDataTableHook,
      icon: "table",
      info: "Datatable hook",
      action: () => {
        service.nav.admin_target_hook_new();
        service.utils.small_modal_close();
      },
    },

    {
      name: TargetHookTypeDomainHook,
      icon: "globe-alt",
      info: "Domain hook",
      action: () => {
        service.nav.admin_target_hook_new();
        service.utils.small_modal_close();
      },
    },

    {
      name: TargetHookTypeStartUp,
      icon: "play",
      info: "Startup hook when node starts",
      action: () => {
        service.nav.admin_target_hook_new();
        service.utils.small_modal_close();
      },
    },
  ];
</script>

<ActionPicker {actions} title="New Target hook type" />
