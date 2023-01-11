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
    const resp = await service.load();
    if (!resp.ok) {
      message = resp.data;
      console.log("Err", resp.data);
      return;
    }

    datas = resp.data["pages"] || [];
    loading = false;
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
      const resp = await service.updatePages([...datas, data]);
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
          Action: (id) => {
            const u = new URL(service.api.base.base_url || "");            

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
          },
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
          Action: async (id) => {
            loading = true;

            const newDatas = datas.filter((v) => v["slug"] !== id);
            await service.updatePages(newDatas);
            await service.deletePageData(id);

            load();
          },
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
