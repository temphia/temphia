<script lang="ts">
  import type { AuthService } from "../../../lib/app/auth/auth";
  import AltMethod from "./_methods/alt_method.svelte";
  import Password from "./_methods/password.svelte";
  import Icons from "./icons";

  export let app: AuthService;
  export let alt_methods = [];
  export let password = false;
  export let opensignup = false;

  let password_mode = password;
  let alt_auth_mode = false;

  let selected_method;
  let data;

  const oauthNext = async (method: object) => {
    const resp = await app.generate_alt_auth(Number(method["id"]));
    if (resp.status !== 200) {
      console.log("Error", resp);
      return;
    }
    selected_method = method;
    data = resp.data;
    alt_auth_mode = true;
  };
</script>

<h2 class="text-2xl font-semibold text-center text-gray-700 mb-5">
  Temphia User Login
</h2>

{#if !alt_auth_mode}
  {#if password_mode}
    <Password {app} />
  {/if}

  {#if alt_methods}
    <div class="w-full flex items-center justify-between py-5">
      <hr class="w-full bg-gray-400" />
      <p class="text-base font-medium leading-4 px-2.5 text-gray-400">OR</p>
      <hr class="w-full bg-gray-400  " />
    </div>
  {/if}

  <div class="p-4 flex flex-col border mt-2">
    {#each alt_methods as method}
      {#if method.type === "oauth"}
        <button
          class="w-full p-2 text-gray-600 border rounded-lg shadow-md hover:bg-gray-200 flex justify-center gap-2"
          on:click={() => oauthNext(method)}
        >
          <img src={Icons[method["provider"]] || ""} alt="" />

          {method.name}
        </button>
      {/if}
    {/each}
  </div>

  {#if opensignup}
    <div class="flex items-center justify-between mt-4">
      <span class="w-1/5 border-b md:w-1/4" />

      <a href="#" class="text-xs text-gray-500 uppercase hover:underline"
        >or sign up</a
      >

      <span class="w-1/5 border-b md:w-1/4" />
    </div>
  {/if}
{:else}
  <AltMethod {app} {data} method={selected_method} />
{/if}
