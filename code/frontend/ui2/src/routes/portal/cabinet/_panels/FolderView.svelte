<script>
    import Folder from "./Icons/Folder.svelte";
    import Icon from "./Icons/Icon.svelte";
    import HeroIcon from "@krowten/svelte-heroicons/Icon.svelte";

    export let files = [
        {
            name: "public",
            last_modified: "",
            size: "",
            is_dir: true,
        },
        {
            name: "test.mp4",
            last_modified: "2022/03/45",
            size: "23M",
        },

        {
            name: "test.zip",
            last_modified: "2022/03/45",
            size: "1K",
        },

        {
            name: "test.png",
            last_modified: "2022/03/45",
            size: "1K",
        },

        {
            name: "xyz.mp3",
            last_modified: "2022/03/45",
            size: "1K",
        },

        {
            name: "test.txt",
            last_modified: "2022/03/45",
            size: "1K",
        },
    ]

    

    let size = "32";


</script>

<div class="table-container">
    <table class="table table-hover">
        <thead>
            <tr>
                <th>Name</th>
                <th>Last Modified</th>
                <th>Size</th>
                <th>Action</th>
            </tr>
        </thead>
        <tbody>
            {#each files as row, i}
                <tr>
                    <td>
                        <span class="mr-1 text-indigo-500">
                            {#if row.is_dir}
                                <Folder {size} />
                            {:else}
                                <Icon {size} name={row.name} />
                            {/if}
                        </span>

                        {row.name}
                    </td>
                    <td>{row.last_modified || ""}</td>
                    <td>{row.size || ""}</td>
                    <td>
                        <button
                            type="button"
                            class="btn btn-sm variant-filled-primary"
                        >
                            <HeroIcon name="eye" class="w-4 h-4" />
                        </button>
                    </td>
                </tr>
            {/each}
        </tbody>
    </table>
</div>
