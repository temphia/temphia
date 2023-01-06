<script lang="ts">
  import { getContext } from "svelte";
  import { Autotable, LoadingSpinner } from "../core";
  import type { EasypageService } from "../service/easypage";
  import Layout from "./_layout.svelte";

  const service = getContext("__easypage_service__") as EasypageService;

  let loading = true;
  let message = "";

  let datas = [
    { slug: "index", name: "Main Page" },
    { slug: "test", name: "Test page" },
  ];

  const load = async () => {
    const resp = await service.listPage();
    if (!resp.ok) {
      message = resp.data;
      console.log("Err", resp.data);
      return;
    }

    datas = resp.data;
    loading = false;
  };

  load();
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <Layout>
    <Autotable
      action_key={"slug"}
      actions={[
        {
          Name: "Visit",
          Action: (id) => {},
          Class: "bg-green-400",
          icon: "link",
        },

        {
          Name: "Edit",
          Action: (id) => {
            location.hash = `/page/${id}`;
          },
          icon: "pencil-alt",
        },
        {
          Name: "Delete",
          Class: "bg-red-400",
          Action: () => {},
          icon: "trash",
        },
      ]}
      {datas}
      key_names={[
        ["slug", "Slug"],
        ["name", "Name"],
      ]}
    />
  </Layout>
{/if}
