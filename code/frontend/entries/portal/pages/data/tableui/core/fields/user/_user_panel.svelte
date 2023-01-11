<script lang="ts">
  export let profile;

  const users = [
    { name: "John Doe", id: "user1", bio: "this a example user", active: true },
  ];
</script>

<div class="h-full flex flex-col">
  <div class="grow">
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
            >Bio</th
          >
          <th
            class="px-5 py-3 border-b-2 border-gray-200 bg-gray-100 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider"
            >Active</th
          >
        </tr></thead
      >
      <tbody>
        <tr>
          {#each users as user}
            <td class="px-5 py-5 border-b border-gray-200 bg-white text-sm"
              ><div class="flex items-center">
                <input type="checkbox" class="h-6 w-6 m-1" />

                <div class="flex-shrink-0 w-10 h-10">
                  <img
                    class="w-24 h-auto rounded-full"
                    src={profile(user.id)}
                    alt=""
                  />
                </div>
                <div class="ml-3">
                  <p class="text-gray-900 whitespace-no-wrap">{user.name}</p>
                </div>
              </div></td
            >
            <td class="px-5 py-5 border-b border-gray-200 bg-white text-sm"
              ><p class="text-gray-900 whitespace-no-wrap">{user.id}</p></td
            >
            <td class="px-5 py-5 border-b border-gray-200 bg-white text-sm">
              {user.bio}
            </td>
            <td class="px-5 py-5 border-b border-gray-200 bg-white text-sm"
              ><p class="text-gray-900 whitespace-no-wrap">{user.active}</p></td
            >
          {/each}
        </tr>
      </tbody>
    </table>
  </div>

  <div class="grow-0 flex justify-end">
    <button class="p-1 bg-slate-400 rounded">Select</button>
  </div>
</div>
