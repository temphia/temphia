<script lang="ts">
  import type { PortalService } from "../../core";
  import ActionPicker from "../../core/action_picker.svelte";
  import {
    TargetAppTypeDataSheetWidget,
    TargetAppTypeDataTableWidget,
    TargetAppTypeDomainWidget,
    TargetAppTypeUserGroupApp,
  } from "../target";
  import NewUserGroupApp from "./_new_user_group_app.svelte";

  export let service: PortalService;

  const actions = [
    {
      name: TargetAppTypeUserGroupApp,
      icon: "user-group",
      info: "App for group of people",
      action: () => {
        service.utils.small_modal_open(NewUserGroupApp, { service});
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
