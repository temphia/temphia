<script lang="ts">
  import { getContext } from "svelte";
  import { Autotable, LoadingSpinner } from "../core";
  import type { EasypageService } from "../service/easypage";
  import Layout from "./_layout.svelte";

  const service = getContext("__easypage_service__") as EasypageService;

  let loading = true;
  let message = "";

  let datas = [];

  const load = async () => {
    loading = true;

    await service.load();

    datas = await service.loadPages();

    console.log("@@@datas", datas);

    loading = false;
  };

  const action_edit = (id: string, data: object) => {
    switch (data["type"]) {
      case "post":
        location.hash = `/post/${id}`;
        break;
      case "page":
        location.hash = `/page/${id}`;
        break;
      default:
        location.hash = `/raw/${id}`;
        break;
    }
  };

  const action_delete = async (id: string) => {
    loading = true;

    await service.deletePage(id);

    load();
  };

  console.log("@SERVICE", service);

  load();
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <Layout
    onRefresh={() => {
      load();
    }}
    onSave={async (data) => {
      const resp = await service.addPage(data["slug"], data);
      if (!resp.ok) {
        console.log(resp);
        return;
      }

      service.modal.small_close();

      load();
    }}
  >
    <Autotable
      action_key={"slug"}
      show_drop={true}
      actions={[
        {
          Name: "Edit",
          Action: action_edit,
          icon: "pencil-alt",
        },
        {
          Name: "Visual Editor",
          Action: (id) => {
            location.hash = `/visual/${id}`;
          },
          icon: "pencil-alt",
          drop: true,
        },
        {
          Name: "Raw Editor",
          Action: (id) => {
            location.hash = `/raw/${id}`;
          },
          icon: "pencil-alt",
          drop: true,
        },
        {
          Name: "Delete",
          Action: action_delete,
          drop: true,
          icon: "trash",
        },
      ]}
      {datas}
      color={["type"]}
      key_names={[
        ["slug", "Slug"],
        ["type", "Type"],
        ["name", "Name"],
      ]}
    />
  </Layout>
{/if}
