<script>
  export let user_id_hints = [
    "john12",
    "john32",
    "ram13",
    "cena66",
    "philip122",
  ];

  export let user_id = "";
  export let full_name = "";
  export let bio = "";
  export let email;
  export let onNext;

  const applyNext = () => {
    onNext({
      user_id,
      full_name,
      bio,
    });
  };
</script>

<div>
  <h3 class="text-xl font-medium text-gray-900 mb-2">
    Fill in your information.
  </h3>
  <div>
    <label for="userid" class="text-sm font-medium text-gray-900 block mb-2"
      >User Id</label
    >
    <input
      bind:value={user_id}
      type="text"
      name="userid"
      id="userid"
      class="bg-gray-100 focus:bg-white border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 "
      placeholder="john12"
    />

    <div class="flex gap-2">
      {#each user_id_hints as hint}
        <a
          class="text-blue-700"
          href="#"
          target="_blank"
          on:click|preventDefault={() => {
            user_id = hint;
          }}>{hint}</a
        >
      {/each}
    </div>
  </div>

  <div>
    <label
      for="email"
      class="text-sm font-medium text-gray-900 block mb-2 mt-4 ">Email</label
    >
    <input
      bind:value={email}
      type="text"
      name="email"
      id="email"
      disabled
      class="bg-gray-100 focus:bg-white border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 "
    />
  </div>

  <div>
    <label
      for="userid"
      class="text-sm font-medium text-gray-900 block mb-2 mt-4 "
      >Full Name</label
    >
    <input
      bind:value={full_name}
      type="text"
      name="name"
      id="name"
      class="bg-gray-100 focus:bg-white border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 "
      placeholder="John Doe"
    />
  </div>

  <div>
    <label for="bio" class="text-sm font-medium text-gray-900 block mb-2 mt-4"
      >Bio</label
    >
    <textarea
      bind:value={bio}
      name="bio"
      id="bio"
      class="bg-gray-100 focus:bg-white border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 "
      placeholder="i am ..."
    />
  </div>

  <!-- <div>
    <label
      for="profile"
      class="text-sm font-medium text-gray-900 block mb-2 mt-4">Profile</label
    >
    <input
      type="file"
      name="profile"
      id="profile"
      class="bg-gray-100 focus:bg-white border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 mb-2"
    />
  </div> -->

  <button
    on:click={applyNext}
    class="w-full px-4 py-2 tracking-wide text-white transition-colors duration-200 transform bg-blue-700 rounded hover:bg-blue-400 mt-4"
  >
    Next
  </button>
</div>
