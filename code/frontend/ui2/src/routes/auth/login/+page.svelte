<script lang="ts">
    import { LoadingSpinner } from "$lib/compo";
    import { LoginService } from "./login";
    const service = new LoginService();

    let user = "";
    let password = "";
    let loading = false;

    let message = "";
</script>

<div>
    <h3 class="h3">User Login</h3>

    <p class="text-red-500">{""}</p>

    <label class="label my-1">
        <span>User</span>
        <input
            class="input p-2"
            type="text"
            placeholder="User1"
            bind:value={user}
        />
    </label>

    <label class="label my-1">
        <span>Password</span>
        <input
            class="input p-2"
            title="Password"
            type="password"
            placeholder="password"
            bind:value={password}
        />
    </label>

    <button
        type="button"
        on:click={async () => {
            await service.init();
            const resp = await service.loginWithPassword(user, password);
            if (resp) {
                message = resp["message"];
            }
        }}
        class="btn variant-filled my-1"
    >
        {#if loading}
            <LoadingSpinner />
        {/if}

        Login</button
    >
</div>
