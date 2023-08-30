<script>
    import { items } from "./admin";
</script>

<h1>Welcome to Admin</h1>

<ul>
    {#each items as item }
    <li>
        <a href="/z/pages/portal/admin{item.path}">{item.name}</a>
    </li>        
    {/each}
</ul>


