<script lang="ts">
  import type { PortalService } from "../core";
  import ActionPicker from "../core/action_picker.svelte";
  import DataPicker from "./pickers/datatable.svelte";
  import RoomPicker from "./pickers/room.svelte";

  export let app: PortalService;

  const resourceTypes = [
    {
      name: "Data Table",
      icon: "collection",
      info: "Create a data table resource to perform operations on data",
      action: () => {
        app.utils.small_modal_open(DataPicker, {
          service: app,
        });
      },
    },
    {
      name: "Room",
      icon: "chat-alt",
      info: "Create a room resource that can publish messages",
      action: () => {
        app.utils.small_modal_open(RoomPicker, {
          service: app,
        });
      },
    },
    {
      name: "Folder",
      icon: "folder",
      info: "Create a Folder resource to store and retrive files",
      action: () => app.nav.admin_resource_folder_new(),
    },
    {
      name: "Module",
      icon: "hashtag",
      action: () => {},
      info: "Create a Module resource that can call native extended functionality",
    },
    {
      name: "Custom",
      icon: "cube",
      info: "Create a Custom resource with own scope and attributes",
      action: () => app.nav.admin_resource_new(),
    },
  ];
</script>

<ActionPicker actions={resourceTypes} title="Pick a resource type" />
