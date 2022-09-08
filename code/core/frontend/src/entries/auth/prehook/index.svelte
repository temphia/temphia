<script lang="ts">
  import type { AuthApp } from "../../../lib/app/auth/auth";
  import { getContext } from "svelte";

  const app: AuthApp = getContext("_auth_app_");

  const opts: {
    preauthed_token: string;
    has_exec_hook: boolean;
    hook_plug_id: string;
    hook_agent_id: string;
    hook_exec_token: string;
  } = app.get_preauthed_data();

  console.log("@opts =>", opts);

  (async () => {
    if (!opts.has_exec_hook) {
      const resp = await app.login_finish(opts.preauthed_token);
      if (resp.status !== 200) {
        console.log("Err =>", resp);
        return;
      }
      app.save_authed_data(resp.data["user_token"]);
      app.nav.goto_final_page();
      return;
    }

    console.log("TODO RUN HOOK");
  })();
</script>

<div>Pre Hook Page</div>
