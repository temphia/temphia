<script lang="ts">
  const datas = [
    {
      id: 1,
      type: "xyz",
      row_id: 11,
      row_version: 0,
      init_sign: "",
      user_id: "mno",
      user_sign: "xyz",
      payload: "{}",
      message: "",
      created_at: "",
    },
    {
      id: 1,
      type: "xyz",
      row_id: 11,
      row_version: 0,
      init_sign: "",
      user_id: "mno",
      user_sign: "xyz",
      payload: "{}",
      message: "",
      created_at: "",
    },
    {
      id: 1,
      type: "xyz",
      row_id: 11,
      row_version: 0,
      init_sign: "",
      user_id: "mno",
      user_sign: "xyz",
      payload: "{}",
      message: "",
      created_at: "",
    },
   
  ];
</script>

<div class="p-5">
  <div
    class="overflow-x-auto bg-white rounded-lg shadow overflow-y-auto relative "
  >
    <table
      class="border-collapse table-auto w-full whitespace-no-wrap bg-white table-striped relative"
    >
      <thead>
        <tr class="text-left">
          <th
            class="bg-gray-100 sticky top-0 border-b border-gray-200 px-6 py-2 text-gray-600 font-bold tracking-wider uppercase text-xs userId"
            >ID</th
          >

          <th
            class="bg-gray-100 sticky top-0 border-b border-gray-200 px-6 py-2 text-gray-600 font-bold tracking-wider uppercase text-xs firstName"
            >Type</th
          >

          <th
            class="bg-gray-100 sticky top-0 border-b border-gray-200 px-6 py-2 text-gray-600 font-bold tracking-wider uppercase text-xs lastName"
            >Row ID</th
          >

          <th
            class="bg-gray-100 sticky top-0 border-b border-gray-200 px-6 py-2 text-gray-600 font-bold tracking-wider uppercase text-xs emailAddress"
            >Row Version</th
          >

          <th
            class="bg-gray-100 sticky top-0 border-b border-gray-200 px-6 py-2 text-gray-600 font-bold tracking-wider uppercase text-xs gender"
            >User ID</th
          >

          <th
            class="bg-gray-100 sticky top-0 border-b border-gray-200 px-6 py-2 text-gray-600 font-bold tracking-wider uppercase text-xs phoneNumber"
            >User Sign</th
          >

          <th
            class="bg-gray-100 sticky top-0 border-b border-gray-200 px-6 py-2 text-gray-600 font-bold tracking-wider uppercase text-xs phoneNumber"
            >Init Sign</th
          >

          <th
            class="bg-gray-100 sticky top-0 border-b border-gray-200 px-6 py-2 text-gray-600 font-bold tracking-wider uppercase text-xs phoneNumber"
            >Payload</th
          >

          <th
            class="bg-gray-100 sticky top-0 border-b border-gray-200 px-6 py-2 text-gray-600 font-bold tracking-wider uppercase text-xs phoneNumber"
            >Message</th
          >
        </tr>
      </thead>
      <tbody>
        {#each datas as data}
          <tr>
            <td class="border-dashed border-t border-gray-200 userId">
              <span class="text-gray-700 px-6 py-3 flex items-center"
                >{data.id}</span
              >
            </td>
            <td class="border-dashed border-t border-gray-200 firstName">
              <span class="text-gray-700 px-6 py-3 flex items-center"
                >{data.type}</span
              >
            </td>
            <td class="border-dashed border-t border-gray-200 lastName">
              <span class="text-gray-700 px-6 py-3 flex items-center"
                >{data.row_id}</span
              >
            </td>
            <td class="border-dashed border-t border-gray-200 emailAddress">
              <span class="text-gray-700 px-6 py-3 flex items-center"
                >{data.row_version}</span
              >
            </td>
            <td class="border-dashed border-t border-gray-200 gender">
              <span class="text-gray-700 px-6 py-3 flex items-center"
                >{data.user_id}</span
              >
            </td>
            <td class="border-dashed border-t border-gray-200 phoneNumber">
              <span class="text-gray-700 px-6 py-3 flex items-center"
                >{data.user_sign}</span
              >
            </td>

            <td class="border-dashed border-t border-gray-200 phoneNumber">
              <span class="text-gray-700 px-6 py-3 flex items-center"
                >{data.init_sign}</span
              >
            </td>

            <td class="border-dashed border-t border-gray-200 phoneNumber">
              <span class="text-gray-700 px-6 py-3 flex items-center"
                >{data.payload}</span
              >
            </td>

            <td class="border-dashed border-t border-gray-200 phoneNumber">
              <span class="text-gray-700 px-6 py-3 flex items-center"
                >{data.message}</span
              >
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  </div>
</div>
