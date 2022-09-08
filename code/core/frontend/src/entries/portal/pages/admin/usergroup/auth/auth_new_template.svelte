<script lang="ts">
  import type { PortalApp } from "../../../../../../lib/app/portal";

  export let app: PortalApp;
  export let gid: string;

  const navigator = app.navigator;

  const providers = [
    {
      name: "Login with Google",
      type: "oauth",
      provider: "google",
      scopes:
        "https://www.googleapis.com/auth/userinfo.email,https://www.googleapis.com/auth/userinfo.profile,openid",
      auth_url: "https://accounts.google.com/o/oauth2/auth",
      token_url: "https://oauth2.googleapis.com/token",
      icon: "https://icons.duckduckgo.com/ip3/www.google.com.ico",
    },
  ];
</script>

<div class="flex flex-col items-center">
  <div class="w-full">
    <div class="my-2 p-1 bg-white flex border border-gray-200 rounded">
      <div class="flex flex-auto flex-wrap" />
      <input
        placeholder="Search Provider"
        class="p-1 px-2 appearance-none outline-none w-full text-gray-800"
      />
      <div
        class="text-gray-300 w-8 py-1 pl-2 pr-1 border-l flex items-center border-gray-200"
      >
        <button
          class="cursor-pointer w-6 h-6 text-gray-600 outline-none focus:outline-none"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="100%"
            height="100%"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
            class="feather feather-chevron-up w-4 h-4"
          >
            <polyline points="18 15 12 9 6 15" />
          </svg>
        </button>
      </div>
    </div>
  </div>
  <div
    class="shadow bg-white top-100 z-40 w-full lef-0 rounded max-h-select overflow-y-auto svelte-5uyqqj"
  >
    <div class="flex flex-col w-full">
      {#each providers as p}
        <div
          on:click={() => {
            navigator.goto_admin_user_auth_new(gid, p);
            app.simple_modal_close();
          }}
          class="cursor-pointer w-full border-gray-100 rounded-t border-b hover:bg-teal-100"
        >
          <div
            class="flex w-full items-center p-2 pl-2 border-transparent border-l-2 relative hover:border-teal-100"
          >
            <div class="w-6 flex flex-col items-center">
              <div
                class="flex relative w-8 h-8 bg-orange-500 justify-center items-center m-1 mr-2 w-4 h-4 mt-1 rounded-full "
              >
                <img class="rounded-full" alt="A" src={p.icon} />
              </div>
            </div>
            <div class="w-full items-center flex">
              <div class="mx-2 -mt-1  ">
                {p.name}
                <div
                  class="text-xs truncate w-full normal-case font-normal -mt-1 text-gray-500"
                >
                  {p.type}
                </div>

                <span class="bg-gray-800 rounded text-white px-1">
                  {p.provider}
                </span>
              </div>
            </div>
          </div>
        </div>
      {/each}
    </div>
  </div>
</div>
