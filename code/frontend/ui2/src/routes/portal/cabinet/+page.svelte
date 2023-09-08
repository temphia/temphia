<script>
    import Folder from "./_panels/Icons/Folder.svelte";
    import Icon from "./_panels/Icons/Icon.svelte";

    const tableArr = [
        {
            name: "public",
            modified: "",
            size: "",
            is_dir: true,
        },
        {
            name: "test.mp4",
            modified: "2022/03/45",
            size: "23M",
        },

        {
            name: "test.txt",
            modified: "2022/03/45",
            size: "1K",
        },

        {
            name: "test.zip",
            modified: "2022/03/45",
            size: "1K",
        },

        {
            name: "test.png",
            modified: "2022/03/45",
            size: "1K",
        },

        {
            name: "xyz.mp3",
            modified: "2022/03/45",
            size: "1K",
        },

        {
            name: "test.txt",
            modified: "2022/03/45",
            size: "1K",
        },
    ];

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
            {#each tableArr as row, i}
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
                    <td>{row.modified}</td>
                    <td>{row.size}</td>
                    <td>A</td>
                </tr>
            {/each}
        </tbody>
    </table>
</div>
