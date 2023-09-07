<script lang="ts">
  import type { SheetService } from "$lib/services/data";
  import Avatar from "../../_shared/user/avatar.svelte";
  import type { SheetColumn } from "../sheets";

  export let service: SheetService;
  export let column: SheetColumn;
  export let onUserAdd = (user) => {};

  let users = [];
  let loading = true;

  const load = async () => {
    loading = true;
    const resp = await service.api.list_users({
      target_type: "sheet",
      target: `${service.sheetid}/${column.__id}`,
    });
    if (!resp.ok) {
      return;
    }

    users = resp.data;
    loading = false;
  };

  load();
</script>

<div class="flex flex-col p-1">
  {#each users as usr}
    {@const userid = usr["user_id"]}
    <div
      class="flex hover:bg-blue-200 justify-between p-1 rounded cursor-pointer"
    >
      <Avatar user={userid} url={service.profile_genrator(userid)} />

      <span class="text-gray-500">{usr["full_name"] || ""}</span>
      <strong
        class="bg-gray-100 text-xs text-gray-500 flex-inline self-center rounded-lg p-0.5"
        >{usr["group"] || ""}</strong
      >
      <button
        on:click={() => onUserAdd(userid)}
        class="hover:bg-blue-500 w-10 rounded-full text-gray-600 hover:text-white  font-semibold text-xs bg-blue-100"
        >+</button
      >
    </div>
  {/each}
</div>
