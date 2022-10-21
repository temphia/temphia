<script lang="ts">
  import { getContext } from "svelte";
  import { apiURL, portalURL } from "../../../../lib/utils/site";
  import type { AuthApp } from "../../app";

  export let tenant_name;
  export let tenant_id;
  export let user_id;
  export let full_name;
  export let group_name;
  export let bio;
  export let return_url = `${window.location.origin}/z/portal`;

  let show_timeout = !!return_url;

  let seconds = 5;
  const it = setInterval(() => {
    if (seconds <= 0) {
      clearInterval(it);
      window.location.href = return_url;
      return;
    }

    seconds -= 1;
  }, 1000);

  const cancel = () => {
    clearInterval(it);
    show_timeout = false;
  };

  const app: AuthApp = getContext("_auth_app_");
</script>

<div class="my-5 mx-auto border" style="max-width: 500px;">
  <div class="bg-white p-1 border rounded">
    <div class="rounded bg-green-500 text-white w-full p-1 mb-2">
      You are Logged in

      {#if show_timeout}
        <span class="text-md"
          >, returning <a href={return_url}
            >back to previous page in {seconds} seconds</a
          ></span
        >
        <button class="text-blue-600" on:click={cancel}>Cancel</button>
      {/if}
    </div>

    <div class="image overflow-hidden">
      <img
        class="h-auto w-12 mx-auto rounded-full border"
        src={`${apiURL(tenant_id)}/user_profile_image/${user_id}`}
        alt=""
      />
    </div>
    <h1 class="text-gray-900 font-bold text-xl leading-8 my-1">
      {full_name}
    </h1>
    <h3 class="text-gray-600 font-lg text-semibold leading-6">
      {group_name}
    </h3>
    <p class="text-sm text-gray-500 hover:text-gray-600 leading-6">
      {bio}
    </p>
    <ul
      class="bg-gray-100 text-gray-600 hover:text-gray-700 hover:shadow py-2 px-3 mt-3 divide-y rounded shadow-sm"
    >
      <li class="flex items-center py-3">
        <span>User Id</span>
        <span class="ml-auto bg-gray-300 rounded p-1">{user_id}</span>
      </li>
      <li class="flex items-center py-3">
        <span>Organization</span>
        <span class="ml-auto bg-gray-300 rounded p-1"
          >{tenant_name + " [" + tenant_id + "]"}</span
        >
      </li>
    </ul>

    <div class="flex flex-col gap-1 mt-1">
      {#if return_url}
        <button
          on:click={() => {
            cancel();
            window.location.href = return_url;
          }}
          class="p-2 bg-blue-400 hover:bg-blue-600 text-white font-semibold rounded"
          >Go Back</button
        >
      {/if}

      <button
        on:click={() => {
          cancel();
          window.location.href = `${window.location.origin}/z/portal`;
        }}
        class="p-2 bg-blue-400 hover:bg-blue-600 text-white font-semibold rounded"
        >Go to Portal</button
      >
      <button
        on:click={() => {
          cancel();
          window.location.href = window.location.origin;
        }}
        class="p-2 bg-blue-400 hover:bg-blue-600 text-white font-semibold rounded"
        >Home</button
      >

      <button
        on:click={() => {
          cancel();
          app.clear_authed_data();
          app.nav.goto_login_page();
        }}
        class="p-2 bg-red-400 hover:bg-red-600 text-white font-semibold rounded"
        >Logout</button
      >
    </div>
  </div>
</div>
