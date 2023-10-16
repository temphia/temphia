<script>
    export let file;
    export let filename;
    const imageTypes = ["jpg", "png", "jpeg"];
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

{#key file}
    {#if !file}
        <div
            class="h-full w-full text-center flex flex-col items-center justify-center cursor-pointer"
        >
            <p class="pointer-none text-gray-500">
                <span class="text-sm">Drag and drop</span>
                files here <br />
                or
                <span class="text-blue-600 hover:underline">select a file</span>
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
