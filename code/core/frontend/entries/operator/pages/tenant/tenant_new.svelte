<script lang="ts">
  import { createTenant } from "../../service";

  let name = "";
  let slug = "";
  let super_password = "";
  let super_email = "";
</script>

<div class="h-full w-full bg-indigo-100 overflow-auto">
  <div class="md:p-12 flex flex-row flex-wrap">
    <div
      class="md:w-1/2-screen m-0 p-5 bg-white w-full tw-h-full shadow md:rounded-lg"
    >
      <div class="text-2xl text-indigo-900">New Tenant</div>

      <div class="flex-col flex py-3">
        <label class="pb-2 text-gray-700 font-semibold">Name</label>
        <input
          type="text"
          bind:value={name}
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          placeholder="Tenant Name"
        />
      </div>

      <div class="flex-col flex py-3 relative">
        <label class="pb-2 text-gray-700 font-semibold">Slug</label>
        <input
          type="text"
          bind:value={slug}
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          placeholder="tenantslug"
        />
      </div>

      <div class="flex-col flex py-3 relative">
        <label class="pb-2 text-gray-700 font-semibold"
          >Super User Email
        </label>

        <input
          type="text"
          bind:value={super_email}
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          placeholder="admin@example.com"
        />
      </div>

      <div class="flex-col flex py-3 relative">
        <label class="pb-2 text-gray-700 font-semibold"
          >Super User Passowrd
        </label>

        <input
          type="password"
          bind:value={super_password}
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex py-3">
        <button
          on:click={() =>
            createTenant({ name, slug, super_password, super_email })}
          class="p-2 bg-blue-400 m-1 w-20 text-white rounded hover:bg-blue-600"
          >Save</button
        >
      </div>
    </div>
  </div>
</div>
