<script>
    export let title;
    export let links = [];
</script>

<div>
    <div class="w-96 mx-auto my-20">
        <header class="card-header mb-2 border-b-2 border-primary-400">
            <h3 class="h3">{title}</h3>
        </header>

        <article>
            <ul>
                {#each links as link}
                    <li>
                        <a class="underline" href={link[1]}>{link[0]}</a>
                    </li>
                {/each}
            </ul>
        </article>
    </div>
</div>
