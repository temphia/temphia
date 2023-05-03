<script lang="ts">
  import Icon from "./_icon.svelte";

  export let data: {
    name?: string;
    slug?: string;
    type?: string;
    subtype?: string;
    description?: string;
    icon?: string;
    tags?: string[];
    versions?: string[];
    files?: {};
  };

  let versions = data.versions || [];
  export let version = versions[versions.length - 1] || "current";

  export let readme = ``;
</script>

<div class="flex flex-wrap border-b">
  <div class="w-1/4">
    <h2 class="m-1 text-xs font-semibold text-lime-500 uppercase title-font">
      BLUEPRINT
    </h2>
    <div class="w-32 p-2 bg-gray-100">
      <Icon src={data.icon} />
    </div>
  </div>
  <div class="w-3/4 flex flex-col ml-5 mb-4">
    <h1 class="text-gray-700 text-3xl font-medium">{data.name}</h1>
    <p class="mb-8 text-base leading-relaxed text-left text-blueGray-600">
      {data.description}
    </p>
  </div>
</div>

<div class="m-auto">
  <p>{readme}</p>
</div>

<div class="w-full border bg-white shadow">
  <div
    class="md:grid md:grid-cols-2 hover:bg-gray-50 md:space-y-0 space-y-1 p-4 border-b"
  >
    <p class="text-gray-600">Name</p>
    <p>{data.name || ""}</p>
  </div>
  <div
    class="md:grid md:grid-cols-2 hover:bg-gray-50 md:space-y-0 space-y-1 p-4 border-b"
  >
    <p class="text-gray-600">Slug</p>
    <p>{data.slug || ""}</p>
  </div>
  <div
    class="md:grid md:grid-cols-2 hover:bg-gray-50 md:space-y-0 space-y-1 p-4 border-b"
  >
    <p class="text-gray-600">Type</p>
    <p>{data.type || ""}</p>
  </div>
  <div
    class="md:grid md:grid-cols-2 hover:bg-gray-50 md:space-y-0 space-y-1 p-4 border-b"
  >
    <p class="text-gray-600">Sub Type</p>
    <p>{data.subtype || ""}</p>
  </div>

  <div
    class="md:grid md:grid-cols-2 hover:bg-gray-50 md:space-y-0 space-y-1 p-4 border-b"
  >
    <p class="text-gray-600">Version</p>

    <select bind:value={version} class="p-1 rounded">
      {#each versions as vv}
        <option value={vv}>{vv}</option>
      {/each}
    </select>
  </div>
</div>
