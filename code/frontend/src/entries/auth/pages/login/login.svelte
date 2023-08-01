<script lang="ts">
  import Inner from "./_inner.svelte";
  import AuthLayout from "../common/layout.svelte";

  import { getContext } from "svelte";
  import type { AuthService } from "../../services";

  const app: AuthService = getContext("_auth_app_");

  let loaded = false;

  let alt_methods = [];
  let password = false;
  let opensignup = false;

  (async () => {
    const resp = await app.list_methods();
    alt_methods = resp.alt_auth_method;
    password = resp.pass_auth;
    opensignup = resp.open_signup;

    loaded = true;
  })();
</script>

<AuthLayout>
  {#if loaded}
    <Inner {app} {alt_methods} {opensignup} {password} />
  {:else}
    <div>Loading..</div>
  {/if}
</AuthLayout>
