<script lang="ts">
  import type { Panel } from "../service";
  export let panel: Panel;
  export let data: any[];

  let options = panel.options || {};

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

<h4 class="text-xl leading-none font-bold text-gray-600 mb-5">
  {panel.name || ""}
</h4>

<div class="overflow-x-auto">
  <table class="table-auto border-collapse w-full">
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
      </tr>
    </thead>
    <tbody class="text-sm font-normal text-gray-700">
      {#each data as d}
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
        </tr>
      {/each}
    </tbody>
  </table>
</div>
