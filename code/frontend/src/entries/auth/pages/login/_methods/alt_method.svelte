<script lang="ts">
  import type { AuthService } from "../../../../lib/app/auth/auth";
  import Oauth from "./_oauth.svelte";

  export let app: AuthService;

  export let method;
  export let data;
</script>

<div class="p-1 flex flex-col">
  {#if method["type"] === "oauth"}
    <Oauth {app} {data} {method} />
  {/if}
</div>
