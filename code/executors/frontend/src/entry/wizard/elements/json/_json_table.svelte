<script lang="ts">
  export let datas: any[];
  export let options = {};
  export let key: string;

  export let onRemove = (data: any) => {};
  export let onEdit = (data: any) => {};

  let key_names = options["column_names"] || [];
  let color = options["color_columns"] || [];
  let image_column = options["image_column"];

  const hashCode = (str) => {
    let hash = 77;
    for (var i = 0; i < str.length; i++) {
      hash = str.charCodeAt(i) + ((hash << 5) - hash);
    }
    return hash;
  };

  const color_it = (key, str) => {
    console.log(key, str);
    if (!color.includes(key)) {
      return "";
    }
    return `background: hsl(${hashCode(str) % 360}, 100%, 80%)`;
  };
</script>

<div class="overflow-x-auto">
  <table class="table-auto border w-full">
    <thead>
      <tr
        class="rounded-lg text-sm font-medium text-gray-700 text-left"
        style="font-size: 0.9674rem"
      >
        {#if image_column}
          <th class="px-2 py-1" style="background-color:#f8f8f8" />
        {/if}

        {#each key_names as key_name}
          {#if key_name !== image_column}
            <th class="px-2 py-1" style="background-color:#f8f8f8"
              >{key_name}</th
            >
          {/if}
        {/each}
        <th class="px-2 py-1" style="background-color:#f8f8f8" />
      </tr>
    </thead>
    <tbody class="text-sm font-normal text-gray-700">
      {#each datas as d}
        <tr
          class="hover:bg-gray-100 border-b border-gray-200 py-10 text-gray-700"
        >
          {#if image_column}
            <td class="px-3 py-1">
              <img src="" alt="" class="p-1 border" />
            </td>
          {/if}

          {#each key_names as key_name}
            {#if key_name !== image_column}
              <td class="px-3 py-1">
                <span
                  class="p-1 rounded-lg"
                  style={color_it(key_name, d[key_name] || "")}
                  >{d[key_name] || ""}</span
                >
              </td>
            {/if}
          {/each}

          <td class="p-1 cursor-pointer flex gap-1">
            <button>
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-5 w-5"
                viewBox="0 0 20 20"
                fill="currentColor"
              >
                <path
                  d="M17.414 2.586a2 2 0 00-2.828 0L7 10.172V13h2.828l7.586-7.586a2 2 0 000-2.828z"
                />
                <path
                  fill-rule="evenodd"
                  d="M2 6a2 2 0 012-2h4a1 1 0 010 2H4v10h10v-4a1 1 0 112 0v4a2 2 0 01-2 2H4a2 2 0 01-2-2V6z"
                  clip-rule="evenodd"
                />
              </svg>
            </button>

            <button on:click={() => onRemove(d)}>
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-5 w-5"
                viewBox="0 0 20 20"
                fill="currentColor"
              >
                <path
                  fill-rule="evenodd"
                  d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z"
                  clip-rule="evenodd"
                />
              </svg>
            </button>
          </td>
        </tr>
      {/each}
    </tbody>
  </table>
</div>
