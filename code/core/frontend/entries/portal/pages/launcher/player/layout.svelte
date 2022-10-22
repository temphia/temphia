<script>
  export let active_apps = [];
  export let current_app = "";
</script>

<div class="flex justify-between border border-gray-100">
  <div class="grow h-10 flex flex-row flex-nowrap overflow-hidden">
    <button
      class="text-gray-600 py-2 px-2 block hover:text-blue-500 focus:outline-none text-blue-500 border-b-2 font-medium border-blue-500"
    >
      Tab 1
    </button><button
      class="text-gray-600 py-2 px-2 block hover:text-blue-500 focus:outline-none"
    >
      Tab 2
    </button><button
      class="text-gray-600 py-2 px-2 block hover:text-blue-500 focus:outline-none"
    >
      Tab 3
    </button><button
      class="text-gray-600 py-2 px-2 block hover:text-blue-500 focus:outline-none"
    >
      Tab 4
    </button>
  </div>

  <div class="grow-0 h-8 w-8 p-1">
    <svg
      class="w-6 h-6 rounded border"
      fill="none"
      viewBox="0 0 24 24"
      stroke-width="1.5"
      stroke="currentColor"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        d="M3.75 6A2.25 2.25 0 016 3.75h2.25A2.25 2.25 0 0110.5 6v2.25a2.25 2.25 0 01-2.25 2.25H6a2.25 2.25 0 01-2.25-2.25V6zM3.75 15.75A2.25 2.25 0 016 13.5h2.25a2.25 2.25 0 012.25 2.25V18a2.25 2.25 0 01-2.25 2.25H6A2.25 2.25 0 013.75 18v-2.25zM13.5 6a2.25 2.25 0 012.25-2.25H18A2.25 2.25 0 0120.25 6v2.25A2.25 2.25 0 0118 10.5h-2.25a2.25 2.25 0 01-2.25-2.25V6zM13.5 15.75a2.25 2.25 0 012.25-2.25H18a2.25 2.25 0 012.25 2.25V18A2.25 2.25 0 0118 20.25h-2.25A2.25 2.25 0 0113.5 18v-2.25z"
      />
    </svg>
  </div>
</div>

<slot />
