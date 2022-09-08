<script lang="ts">
  import { AutoTable } from "../../../../../../components";
  import type { PortalApp } from "../../../../../../lib/app/portal";

  export let user_group = "";
  export let app: PortalApp;

  let datas = [];
  const load = async () => {
    const uapi = await app.get_apm().get_user_api();
    const resp = await uapi.user_group_list_plug(user_group);
    datas = resp.data;
  };

  load();
</script>

<AutoTable
  action_key="id"
  actions={[
    {
      Name: "Explore",
      Class: "bg-green-400",
      Action: null,
    },
    {
      Name: "Edit",
      Action: null,
    },
    {
      Name: "Delete",
      Class: "bg-red-400",
      Action: null,
    },
  ]}
  key_names={[
    ["id", "id"],
    ["stage_name", "Stage"],
  ]}
  {datas}
/>
