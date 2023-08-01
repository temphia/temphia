<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";

  export let uid: string;
  export let onComplete: (opts: any) => void;

  let message = "";
  let loading = true;

  let new_email = "";
  let pending_verify = false;

  const save = (_data) => {
    onComplete({
      new_email,
      pending_verify,
    });
  };
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
      bind:value={new_email}
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
      bind:checked={pending_verify}
      class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
    />
  </div>
  <div class="flex justify-end py-3">
    <button
      on:click={save}
      class="p-1 text-white text-sm font-semibold flex self-center shadow rounded hover:scale-110 bg-blue-400"
    >
      <Icon name="save" class="h-5 w-5" />

      Save</button
    >
  </div>
</div>
