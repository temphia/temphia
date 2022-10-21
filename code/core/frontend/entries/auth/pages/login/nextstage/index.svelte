<script lang="ts">
  import Layout from "../../common/layout.svelte";
  import { getContext } from "svelte";
  import type { AuthApp } from "../../../app";

  const app: AuthApp = getContext("_auth_app_");

  const opts: {
    ok: boolean;
    message: string;
    next_token: string;
    user_id: string;
    email_verify: boolean;
    pass_change: boolean;
  } = app.nav.nav_options;

  let show = false;
  let message = "";

  (async () => {
    if (!opts.email_verify && !opts.pass_change) {
      const resp = await app.login_submit(opts.next_token);
      if (resp.status !== 200) {
        console.log("Err =>", resp);
        return;
      }

      if (!resp.data["ok"]) {
        message = resp.data["message"];
        return;
      }

      app.save_preauthed_data(resp.data);
      app.nav.goto_prehook_page(resp.data);
    } else {
      show = true;
    }
  })();
</script>

<Layout>
  <p>{message}</p>

  {#if show}
    <div>fixme</div>
  {/if}
</Layout>
