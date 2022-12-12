<script lang="ts">
  import { getContext } from "svelte";
  import {
    AutoTable,
    LoadingSpinner,
    FloatingAdd,
    PortalService,
  } from "../../core";
  import { params } from "svelte-hash-router";
  import NewPick from "./_new_pick.svelte";

  const app = getContext("__app__") as PortalService;

  let sources = [];
  let datas = [];
  let loading = true;
  const api = app.api_manager.get_admin_data_api();

  const load = async (src) => {
    if (!src) {
      return;
    }

    const sreq = app.api_manager.self_data.get_data_sources();

    const resp = await api.list_group(src);
    if (!resp.ok) {
      return;
    }

    datas = resp.data;
    loading = false;

    sources = await sreq;
  };

  $: load($params.source);

  // actions

  const action_edit = (id: string) =>
    app.nav.admin_data_group($params.source, id);
  const action_explore = (id: string) =>
    app.nav.admin_data_tables($params.source, id);
  const action_delete = async (id: string) => {};
  const action_new = () => {
    app.utils.small_modal_open(NewPick, { app });
  };
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <div class="flex justify-end pt-2 pr-2">
    <select
      class="px-2 py-1 rounded-full bg-white hover:text-white hover:bg-slate-500 border border-slate-600"
      value={$params.source}
      on:change={(ev) => app.nav.admin_data_groups(ev.target["value"])}
    >
      {#each sources || [] as source}
        <option value={source}>{source}</option>
      {/each}
    </select>
  </div>

  <AutoTable
    action_key="slug"
    show_drop={true}
    actions={[
      {
        Name: "Edit",
        Action: action_edit,
        icon: "pencil-alt",
      },
      {
        Class: "bg-green-400",
        Name: "explore",
        icon: "book-open",
        Action: action_explore,
      },

      {
        Class: "bg-green-400",
        Name: "Query",
        icon: "database",
        Action: (id) => app.nav.admin_data_query($params.source, id),
        drop: true,
      },

      {
        Name: "Delete",
        Class: "bg-red-400",
        Action: action_delete,
        icon: "trash",
        drop: true,
      },
    ]}
    key_names={[
      ["name", "Name"],
      ["slug", "Slug"],
      ["description", "Description"],
    ]}
    color={[]}
    {datas}
  />
{/if}

<FloatingAdd onClick={action_new} />
