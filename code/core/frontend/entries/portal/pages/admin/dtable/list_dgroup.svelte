<script lang="ts">
  import { AutoTable, FloatingAdd } from "../../../../../shared";
  import Layout from "../layout.svelte";

  import { getContext } from "svelte";
  import { DynAdminAPI } from "./dtable2";
  import type { PortalApp } from "../../../app";

  let groups = [];
  let sources = [];
  let current_source = "";
  let loaded = false;

  const app: PortalApp = getContext("__app__");
  const dynapi = new DynAdminAPI(app);

  dynapi.load_sources().then((resp) => {
    sources = resp.data;
    if (!sources.length) {
      return;
    }

    current_source = sources[0];
  });

  $: {
    if (current_source !== "") {
      dynapi.load_groups(current_source).then((resp) => {
        groups = resp.data;
        loaded = true;
      });
    }
  }
</script>

<Layout current_item="dtable">
  <div class="h-full w-full p-2">
    <div class="bg-white rounded p-2 flex flex-col">
      <div class="flex justify-end">
        <label class="p-1">
          <select
            class="p-1 border bg-gray-50 rounded"
            bind:value={current_source}
          >
            {#each sources as source}
              <option>{source}</option>
            {/each}
          </select>
        </label>
      </div>

      {#if loaded}
        <AutoTable
          action_key="slug"
          actions={[
            {
              Class: "bg-green-400",
              Name: "explore",
              icon: "book-open",
              Action: (id) => {
                dynapi.goto_dgroup(current_source, id);
              },
            },
            {
              Name: "Edit",
              Action: (grp) => {
                app.navigator.goto_dgroup_edit(current_source, grp);
              },
              icon: "pencil-alt",
            },
            {
              Name: "Delete",
              Class: "bg-red-400",
              Action: async (grp) => {
                await dynapi.delete_dgroup(current_source, grp);
                dynapi.load_groups(current_source).then((resp) => {
                  groups = resp.data;
                  loaded = true;
                });
              },
              icon: "trash",
            },
          ]}
          key_names={[
            ["name", "Name"],
            ["description", "Description"],
            ["source_db", "source"],
          ]}
          datas={groups}
        />
      {/if}
    </div>
  </div>
</Layout>

<FloatingAdd onClick={() => {}} />
