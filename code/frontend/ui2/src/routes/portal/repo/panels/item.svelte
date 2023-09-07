<script>
  import Icon from "./icon.svelte";
  // import SvelteMarkdown from "svelte-markdown";
  import {PrimaryButton} from "$lib/compo/common";

  export let data = {
    id: "",
    name: "",
    slug: "",
    group: "",
    subgroup: "",
    description: ``,
    screenshots: [],
    icon: "",
    schema: "",
    meta: ``,
    source: "",
    tenant_id: "",
    tags: [],
    files: {},
  };

  export let readme = ``;
  export let importFunc;
</script>

<div class="h-full w-full bg-gray-200 overflow-auto">
  <div class="w-full mt-10">
    <div
      class="m-auto flex flex-col p-4 bg-white shadow rounded"
      style="max-width: 50rem;"
    >
      <div class="flex flex-wrap  border-b">
        <div class="w-1/4">
          <h2
            class="m-1 text-xs font-semibold text-lime-500 uppercase title-font"
          >
            BLUEPRINT
          </h2>
          <div class="w-32 p-2 bg-gray-100">
            <Icon src={data.icon} />
          </div>
        </div>
        <div class="w-3/4 flex flex-col ml-5 mb-4">
          <h1 class="text-gray-700 text-3xl font-medium">{data.name}</h1>
          <p
            class="mb-8 text-base leading-relaxed text-left text-blueGray-600 "
          >
            {data.description}
          </p>
        </div>
      </div>

      <div class="m-auto">
        <!-- <SvelteMarkdown source={readme} /> -->
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
        <!-- <div
          class="md:grid md:grid-cols-2 hover:bg-gray-50 md:space-y-0 space-y-1 p-4"
        >
          <p class="text-gray-600">Files</p>
          <div class="space-y-2">
            {#each data.files || [] as file}
              <div
                class="border-2 flex items-center p-2 rounded justify-between space-x-2"
              >
                <div class="space-x-2 truncate">
                  <svg
                    class="fill-current inline text-gray-500"
                    width="24"
                    height="24"
                    viewBox="0 0 24 24"
                    ><path
                      d="M17 5v12c0 2.757-2.243 5-5 5s-5-2.243-5-5v-12c0-1.654 1.346-3 3-3s3 1.346 3 3v9c0 .551-.449 1-1 1s-1-.449-1-1v-8h-2v8c0 1.657 1.343 3 3 3s3-1.343 3-3v-9c0-2.761-2.239-5-5-5s-5 2.239-5 5v12c0 3.866 3.134 7 7 7s7-3.134 7-7v-12h-2z"
                    /></svg
                  >
                  <span> {file} </span>
                </div>
                <a href="#" class="text-purple-700 hover:underline">
                  Download
                </a>
              </div>
            {/each}
          </div>
        </div> -->
      </div>

      <div class="flex justify-end items-center p-1">
        <PrimaryButton
          icon="arrow-circle-down"
          label="IMPORT"
          onClick={importFunc}
        />
      </div>
    </div>
  </div>

  <div class="w-full m-5 " />
</div>
