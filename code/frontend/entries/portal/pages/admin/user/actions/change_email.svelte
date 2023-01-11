<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import type { AutoForm, PortalService } from "../../core";

  export let uid: string;
  export let app: PortalService;
  const api = app.api_manager.get_admin_user_api();

  let message = "";
  let loading = true;

  const save = async (_data) => {};
</script>

<div class="p-4">
  <div class="text-2xl text-indigo-900">Change Email</div>
  <p class="text-red-500 ">{message}</p>
  <div class="flex-col flex py-3">
    <label for="field-0" class="pb-2 text-gray-700 font-semibold"
      >New Email</label
    >
    <input
      id="field-0"
      type="email"
      class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
    />
  </div>

  <div class="flex-col flex py-3">
    <label for="field-1" class="pb-2 text-gray-700 font-semibold"
      >Pending Verify Email</label
    >
    <input
      type="checkbox"
      id="field-1"
      class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
    />
  </div>
  <div class="flex justify-end py-3">
    <button
      class="p-1 text-white text-sm font-semibold flex self-center shadow rounded hover:scale-110 bg-blue-400"
    >
      <Icon name="save" class="h-5 w-5" />

      Save</button
    >
  </div>
</div>
