<script>
    const tasks = [
        {
            group: "Tasks Group Name",
            items: [
                {
                    name: "task 1",
                },
                {
                    name: "task 2",
                },
                {
                    name: "task 3",
                },
                {
                    name: "task 4",
                },
                {
                    name: "task 5",
                },
            ],
        },

        {
            group: "Tasks Group Name",
            items: [
                {
                    name: "task 1",
                },
                {
                    name: "task 2",
                },
                {
                    name: "task 3",
                },
                {
                    name: "task 4",
                },
                {
                    name: "task 5",
                },
            ],
        },

        {
            group: "Tasks Group Name",
            items: [
                {
                    name: "task 1",
                },
                {
                    name: "task 2",
                },
                {
                    name: "task 3",
                },
                {
                    name: "task 4",
                },
                {
                    name: "task 5",
                },
            ],
        },

        {
            group: "Tasks Group Name",
            items: [
                {
                    name: "task 1",
                },
                {
                    name: "task 2",
                },
                {
                    name: "task 3",
                },
                {
                    name: "task 4",
                },
                {
                    name: "task 5",
                },
            ],
        },

        {
            group: "Tasks Group Name",
            items: [
                {
                    name: "task 1",
                },
                {
                    name: "task 2",
                },
                {
                    name: "task 3",
                },
                {
                    name: "task 4",
                },
                {
                    name: "task 5",
                },
            ],
        },
    ];
</script>

<div class="bg-blue-500 w-full h-full font-sans">
    <div class="flex px-4 pb-8 items-start overflow-x-scroll">
        {#each tasks as task}
            <div class="rounded bg-gray-light  flex-no-shrink w-64 p-2 mr-3">
                <div class="flex justify-between py-1">
                    <h3 class="text-sm">New landing page</h3>
                    <svg
                        class="h-4 fill-current text-gray-dark cursor-pointer"
                        xmlns="http://www.w3.org/2000/svg"
                        viewBox="0 0 24 24"
                        ><path
                            d="M5 10a1.999 1.999 0 1 0 0 4 1.999 1.999 0 1 0 0-4zm7 0a1.999 1.999 0 1 0 0 4 1.999 1.999 0 1 0 0-4zm7 0a1.999 1.999 0 1 0 0 4 1.999 1.999 0 1 0 0-4z"
                        /></svg
                    >
                </div>
                <div class="text-sm mt-2">
                    {#each task.items as item}
                        <div
                            class="bg-white p-2 rounded mt-1 border-b border-gray cursor-pointer hover:bg-gray-lighter"
                        >
                            {item.name}
                        </div>
                    {/each}

                    <p class="mt-3 text-gray-dark">Add a card...</p>
                </div>
            </div>
        {/each}
    </div>
</div>

<!-- 


    <div
                        class="bg-white p-2 rounded mt-1 border-b border-gray cursor-pointer hover:bg-gray-lighter"
                    >
                        Check the responsive layout on all devices
                        <div
                            class="text-gray-darker mt-2 ml-2 flex justify-between items-start"
                        >
                            <span class="text-xs flex items-center">
                                <svg
                                    class="h-4 fill-current mr-1"
                                    xmlns="http://www.w3.org/2000/svg"
                                    viewBox="0 0 50 50"
                                    ><path
                                        d="M11 4c-3.855 0-7 3.145-7 7v28c0 3.855 3.145 7 7 7h28c3.855 0 7-3.145 7-7V11c0-3.855-3.145-7-7-7zm0 2h28c2.773 0 5 2.227 5 5v28c0 2.773-2.227 5-5 5H11c-2.773 0-5-2.227-5-5V11c0-2.773 2.227-5 5-5zm25.234 9.832l-13.32 15.723-8.133-7.586-1.363 1.465 9.664 9.015 14.684-17.324z"
                                    /></svg
                                >
                                3/5
                            </span>
                            <img
                                src="https://i.imgur.com/OZaT7jl.png"
                                class="rounded-full"
                            />
                        </div>
                    </div>


 -->
