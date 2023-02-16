<script lang="ts">
  import type { PortalService } from "../../core";
  import ActionPicker from "../../core/action_picker.svelte";
  import {
    TargetAppTypeDataSheetWidget,
    TargetAppTypeDataTableWidget,
    TargetAppTypeDomainWidget,
    TargetAppTypeUserGroupApp,
  } from "../target";
  import NewDataTableWidget from "./_new_data_table_widget.svelte";
  import NewUserGroupApp from "./_new_user_group_app.svelte";

  export let service: PortalService;

  const actions = [
    {
      name: TargetAppTypeUserGroupApp,
      icon: "user-group",
      info: "App for group of people",
      action: () =>
        service.utils.small_modal_open(NewUserGroupApp, { service }),
    },

    {
      name: TargetAppTypeDataTableWidget,
      icon: "table",
      info: "Datatable widget",
      action: () =>
        service.utils.small_modal_open(NewDataTableWidget, { service }),
    },

    {
      name: TargetAppTypeDataSheetWidget,
      icon: "table",
      info: "DataSheet widget",
      action: () => {
        service.nav.admin_target_app_new({
          target_type: TargetAppTypeDataSheetWidget,
          context_type: "widget.1",
        });

        service.utils.small_modal_close()

      },
    },

    {
      name: TargetAppTypeDomainWidget,
      icon: "globe-alt",
      info: "Domain widget",
      action: () => {
        service.nav.admin_target_app_new({
          target_type: TargetAppTypeDomainWidget,
          context_type: "widget.1",
        });

        service.utils.small_modal_close()

      },
    },
  ];
</script>

<ActionPicker {actions} title="New Target app type" />
