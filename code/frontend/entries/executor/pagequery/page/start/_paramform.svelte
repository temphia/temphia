<script>
  export let form = {
    Test: "text",
    Test2: "longtext",
  };
  export let title = "Form 1";

  let values = {};

  export const getParamData = () => {
    return values;
  };
</script>

<div class="w-full h-full flex flex-col p-2 border my-2 text-gray-700">
  <h1
    class="text-gray-800 font-lg font-bold tracking-normal leading-tight mb-4"
  >
    {title}
  </h1>

  {#each Object.entries(form) as [elemkey, elemtype]}
    <label for="name" class="pb-2 text-gray-700 flex uppercase">{elemkey}</label
    >
    <div class="flex flex-col w-full">
      {#if elemtype === "text"}
        <input
          type="text"
          class="p-2 border w-full rounded-lg bg-gray-50 outline-none focus:bg-gray-100 mr-1"
        />
      {:else if elemtype === "longtext"}
        <textarea
          class="p-2 border w-full rounded-lg bg-gray-50 outline-none focus:bg-gray-100 mr-1"
        />
      {/if}

      <p class="text-sm italic mb-4" />
    </div>
  {/each}
</div>
