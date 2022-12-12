<script>
  let data = {
    sections: [
      {
        layout_type: "flex",
        bg_color: "#ff721",
        panels: [
          {
            size: "",
            name: "test",
            type: "bold_value",
            value: "77",
          },
          {
            size: "",
            name: "test",
            type: "bold_value",
            value: "77",
          },
          {
            size: "",
            name: "test",
            type: "bold_value",
            value: "77",
          },
          {
            size: "",
            name: "test",
            type: "bold_value",
            value: "77",
          },
          {
            size: "",
            name: "test",
            type: "bold_value",
            value: "77",
          },
          {
            size: "",
            name: "test",
            type: "bold_value",
            value: "77",
          },
        ],
      },
      {
        layout_type: "flex",
        items: 9,
        bg_color: "#ff721",
        panels: [
          {
            name: "test",
            type: "bold_value",
            value: "77",
          },
        ],
      },
    ],
  };

  const classIt = (mode) => {
    switch (mode) {
      case "xm":
        return "w-28";
      case "sm":
        return "w-32";
      case "md":
        return "w-48";
      case "lg":
        return "w-64";
      case "xl":
        return "w-72";
      case "2xl":
        return "w-96";
      default:
        return "w-64";
    }
  };
</script>

{#each data.sections as section}
  <div class="flex flex-wrap justify-between gap-2 py-2">
    {#each section.panels as panel}
      <div class={classIt(panel.class)}>
        <div class="flex items-center px-5 py-6 shadow-sm rounded-md bg-white">
          <div class="p-3 rounded-full bg-indigo-600 bg-opacity-75" />

          <div class="mx-5">
            <h4 class="text-2xl font-semibold text-gray-700">17</h4>
            <div class="text-gray-500">Users</div>
          </div>
        </div>
      </div>
    {/each}
  </div>
{/each}
