<script>
  import BackgroundImage from "./_bg_image.svelte";
  export let uploadFile;

  const imageTypes = ["jpg", "png", "jpeg"];

  let filename = "";
  let file;

  const fileSelect = (ev) => {
    console.log(ev);
    file = ev.target.files[0];
    filename = file.name;
    console.log(file);
  };

  const upload = async () => {
    const formdata = new FormData();
    formdata.append("file", file);
    uploadFile(filename, formdata);
  };

  const canPreview = () => {
    const fparts = filename.split(".");
    if (fparts.length <= 1) {
      return false;
    }

    if (imageTypes.includes(fparts.pop())) {
      return true;
    }

    return false;
  };
</script>

<div class="w-full h-full">
  <div class="text-center">
    <h2 class="text-3xl font-bold text-gray-900">File Upload!</h2>
  </div>
  <form class="mt-4 space-y-3" on:submit|preventDefault={upload}>
    <div class="grid grid-cols-1 space-y-2">
      <label class="text-sm font-bold text-gray-500 tracking-wide">Title</label>
      <input
        class="text-base p-2 border border-gray-300 rounded-lg focus:outline-none focus:border-indigo-500"
        type="text"
        bind:value={filename}
        placeholder="file.txt"
      />
    </div>
    <div class="grid grid-cols-1 space-y-2">
      <span class="text-sm font-bold text-gray-500 tracking-wide"
        >Attach Document</span
      >
      <div class="flex items-center justify-center w-full">
        <label
          class="flex flex-col rounded-lg border-4 border-dashed w-full h-60 p-5 group text-center"
        >
          {#key file}
            {#if !file}
              <div
                class="h-full w-full text-center flex flex-col items-center justify-center items-center  cursor-pointer"
              >
                <div class="flex flex-auto max-h-48 w-2/5 mx-auto">
                  <div class="h-36 p-4">
                    <BackgroundImage />
                  </div>
                </div>

                <p class="pointer-none text-gray-500 ">
                  <span class="text-sm">Drag and drop</span> files here <br />
                  or
                  <span class="text-blue-600 hover:underline"
                    >select a file</span
                  >
                  from your computer
                </p>
              </div>
            {:else if canPreview()}
              <div class="flex justify-center">
                <img
                  class="h-48 w-auto p-2 rounded border"
                  src={URL.createObjectURL(file)}
                  alt=""
                />
              </div>
            {:else}
              <div class="flex justify-center">
                <div class="bg-text-white rounded text-gray-400 p-1 border">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    class="h-48 w-auto"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M7 21h10a2 2 0 002-2V9.414a1 1 0 00-.293-.707l-5.414-5.414A1 1 0 0012.586 3H7a2 2 0 00-2 2v14a2 2 0 002 2z"
                    />
                  </svg>
                </div>
              </div>
            {/if}
          {/key}

          <input type="file" class="hidden" on:change={fileSelect} />
        </label>
      </div>
    </div>
    <p class="text-sm text-gray-300">
      <span>File type: any</span>
    </p>
    <div>
      {#if file}
        <button
          type="submit"
          class="my-5 w-full flex justify-center bg-blue-500 text-gray-100 p-4  rounded-full tracking-wide font-semibold  focus:outline-none focus:shadow-outline hover:bg-blue-800 shadow-lg cursor-pointer transition ease-in duration-300"
        >
          Upload
        </button>
      {/if}
    </div>
  </form>
</div>
