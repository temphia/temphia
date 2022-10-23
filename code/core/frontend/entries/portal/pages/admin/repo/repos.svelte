<script lang="ts">
  import { AutoTable, FloatingAdd } from "../../../../xcompo";

  import { getContext } from "svelte";
  import type { PortalService } from "../../../services";
  const app = getContext("__app__") as PortalService;

  const rapi = app.api_manager.get_admin_repo_api();

  let repos = [];

  let loading = true;

  const load = async () => {
    const resp = await rapi.list();
    if (resp.status !== 200) {
      console.log("Err", resp);
      return;
    }

    repos = resp.data;
    loading = false;
  };

  load();
</script>

<AutoTable
  action_key="id"
  actions={[
    {
      Name: "Edit",
      Action: null,
    },
    {
      Name: "Delete",
      Class: "bg-red-400",
      Action: async (rid) => {
        await rapi.delete(rid);
        load();
      },
    },
  ]}
  key_names={[
    ["id", "ID"],
    ["name", "Name"],
    ["provider", "Provider"],
  ]}
  color={["provider"]}
  datas={repos}
/>

<FloatingAdd onClick={() => {}} />
