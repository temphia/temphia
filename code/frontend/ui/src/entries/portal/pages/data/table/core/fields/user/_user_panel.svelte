<script lang="ts">
  import PanelLayout from "../_panel_layout.svelte";
  import type { Column, RowService } from "../../../../../../services/data";

  export let row_service: RowService;
  export let profile_generator: (user: string) => string;
  export let column: Column;
  export let onSelect: (user) => void;

  let selected;

  let users = [];

  let loading = true;
  const load = async () => {
    const resp = await row_service.list_user(column.slug);
    if (!resp.ok) {
      return;
    }

    users = resp.data;
    loading = false;
  };

  load();
</script>

<PanelLayout
  {loading}
  onSelect={() => selected && onSelect(selected)}
  selected={!!selected}
>
  <table class="min-w-full leading-normal">
    <thead
      ><tr
        ><th
          class="px-5 py-3 border-b-2 border-gray-200 bg-gray-100 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider"
          >Name</th
        >

        <th
          class="px-5 py-3 border-b-2 border-gray-200 bg-gray-100 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider"
          >Id</th
        >

        <th
          class="px-5 py-3 border-b-2 border-gray-200 bg-gray-100 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider"
          >Group</th
        >
      </tr></thead
    >
    <tbody>
      {#each users as user}
        <tr>
          <td class="px-5 py-5 border-b border-gray-200 bg-white text-sm"
            ><div class="flex items-center">
              <input
                type="checkbox"
                class="h-6 w-6 m-1"
                on:click={() => {
                  if (selected === user.user_id) {
                    selected = null;
                  } else {
                    selected = user.user_id;
                  }
                }}
                checked={selected === user.user_id}
              />

              <div class="flex-shrink-0 w-10 h-10">
                <img
                  class="w-24 h-auto rounded-full"
                  src={profile_generator(user.user_id)}
                  alt=""
                />
              </div>
              <div class="ml-3">
                <p class="text-gray-900 whitespace-no-wrap">{user.full_name}</p>
              </div>
            </div></td
          >
          <td class="px-5 py-5 border-b border-gray-200 bg-white text-sm"
            ><p class="text-gray-900 whitespace-no-wrap">{user.user_id}</p></td
          >
          <td class="px-5 py-5 border-b border-gray-200 bg-white text-sm">
            {user.group}
          </td>
        </tr>
      {/each}
    </tbody>
  </table>
</PanelLayout>
