<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import { getContext } from "svelte";
  import { LoadingSpinner, PortalService, ActionButton } from "$lib/core";
  import { params } from "$lib/params";

  export let source = $params["source"];
  export let group = $params["group"];
  export let table = $params["table"];

  const app = getContext("__app__") as PortalService;
  const api = app.api_manager.get_admin_data_api();

  let loading = true;
  let seed_no = 10;
  let datas = [
    {
      name: "Notes",
      slug: "notes",
      description: "Lorem ipsum dolor sit amet consectetur adipisicing elit.",
    },
  ];

  const load = async () => {
    const resp = await api.list_column(source, group, table);
    if (!resp.ok) {
      console.log("Err", resp);
      return;
    }
    loading = false;
    datas = resp.data;
  };

  load();

  let done = false;
  let done_data;
  const do_seed = async () => {
    const resp = await api.seed_table(source, group, table, seed_no);
    done_data = resp.data;
    done = true;
    loading = false;
  };
</script>

{#if loading}
  <LoadingSpinner />
{:else if done}
  <div class="flex justify-center py-5">
    <div class="flex flex-col">
      <p>{done_data || ""}</p>

      <!-- /*app.nav.data_table(source, group, table)*/ -->
      <button
        on:click={() => {}} 
        class="p-2 bg-blue-400 text-white rounded">Table Data</button
      >
    </div>
  </div>
{:else}
  <div class="p-4">
    <div class="p-4 bg-white rounded-md">
      <div class="text-2xl text-indigo-900 mb-6">Auto Seeder</div>

      <div>
        <legend class="text-base text-1.5xl font-medium text-gray-900 mb-2"
          >Columns</legend
        >
        <div class="overflow-auto shadow p-1">
          <table class="text-left w-full border">
            <thead
              ><tr>
                <th
                  class="p-1 bg-grey-lightest font-bold uppercase text-sm text-grey-dark border-b border-grey-light"
                  >#</th
                >
                <th
                  class="p-1 bg-grey-lightest font-bold uppercase text-sm text-grey-dark border-b border-grey-light"
                  >Slug</th
                >
                <th
                  class="p-1 bg-grey-lightest font-bold uppercase text-sm text-grey-dark border-b border-grey-light"
                  >Name</th
                >
                <th
                  class="p-1 bg-grey-lightest font-bold uppercase text-sm text-grey-dark border-b border-grey-light"
                  >Description</th
                >
              </tr>
            </thead>
            <tbody>
              {#each datas as data}
                <tr class="hover:bg-grey-lighter">
                  <td class="p-1 border-b border-grey-light">
                    <!-- <input type="checkbox" /> -->
                  </td>
                  <td class="p-1 border-b border-grey-light">{data.slug}</td>
                  <td class="p-1 border-b border-grey-light">{data.name || ""}</td>
                  <td class="p-1 border-b border-grey-light"
                    >{data.description || ""}</td
                  >
                </tr>
              {/each}
            </tbody>
          </table>
        </div>
      </div>
      <div class="flex ">
        <div class="inline-flex py-4">
          <legend class="text-base text-1.5xl font-medium text-gray-900 p-1"
            >No of records</legend
          >
          <input
            type="number"
            bind:value={seed_no}
            class="p-1 border rounded"
          />
        </div>
      </div>

      <div class="flex justify-end p-2">
        <ActionButton onClick={do_seed} name="Finish" icon_name="sparkles" />
      </div>
    </div>
  </div>
{/if}
