<script lang="ts">
  import { getContext } from "svelte";
  import { Autotable, LoadingSpinner } from "../core";
  import type { EasypageService } from "../service/easypage";
  import Layout from "./_layout.svelte";
  import Link from "./_panels/link.svelte";

  const service = getContext("__easypage_service__") as EasypageService;

  let loading = true;
  let message = "";

  let datas = [];

  const load = async () => {
    loading = true;

    await service.load()
    

    loading = false;
  };

  const action_visit = (id: string) => {
    const url = "http://localhost"

    const u = new URL(url || "");

    let domain_name = service.env.domain_name;
    if (!domain_name || domain_name === "*") {
      domain_name = location.hostname;
    }

    if (!domain_name) {
      domain_name = u.hostname;
    }

    console.log("@domain_name", domain_name);

    service.modal.small_open(Link, {
      domain: `http://${domain_name}:${u.port || "80"}`,
      slug: id,
      service,
    });
  };

  const action_edit = (id: string, data: object) => {
    if (data["type"] === "post") {
      location.hash = `/post/${id}`;
    } else {
      location.hash = `/page/${id}`;
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
      actions={[
        {
          Name: "Visit",
          Action: action_visit,
          Class: "bg-green-400",
          icon: "link",
        },

        {
          Name: "Edit",
          Action: action_edit,
          icon: "pencil-alt",
        },
        {
          Name: "Delete",
          Class: "bg-red-400",
          Action: action_delete,
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
