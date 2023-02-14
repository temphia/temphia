<script lang="ts">
  import type { PortalService } from "../../core";
  import ActionPicker from "../../core/action_picker.svelte";
  import {
    TargetAppTypeDataSheetWidget,
    TargetAppTypeDataTableWidget,
    TargetAppTypeDomainWidget,
    TargetAppTypeUserGroupApp,
  } from "../target";

  export let service: PortalService;

  const actions = [
    {
      name: TargetAppTypeUserGroupApp,
      icon: "user-group",
      info: "App for group of people",
      action: () => {
        service.nav.admin_target_app_new();
        service.utils.small_modal_close();
      },
    },

    {
      name: TargetAppTypeDataTableWidget,
      icon: "table",
      info: "Datatable widget",
      action: () => {},
    },

    {
      name: TargetAppTypeDataSheetWidget,
      icon: "table",
      info: "DataSheet widget",
      action: () => {},
    },

    {
      name: TargetAppTypeDomainWidget,
      icon: "globe-alt",
      info: "Domain widget",
      action: () => {},
    },
  ];
</script>

<ActionPicker {actions} title="New Target app type" />
