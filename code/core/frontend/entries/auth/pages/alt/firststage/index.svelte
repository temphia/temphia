<script lang="ts">
  import type { AuthApp } from "../../../app";

  import { getContext } from "svelte";
  import NewUserInfo from "../../common/new_user_info.svelte";
  import Layout from "../../common/layout.svelte";
  const app: AuthApp = getContext("_auth_app_");

  const opts: {
    message: string;
    ok: boolean;
    first_token: string;
    new_user: boolean;
    email: string;
    user_id_hints: string[];
  } = app.nav.nav_options;

  if (!opts) {
    app.nav.goto_login_page();
  }

  (async () => {
    if (opts.new_user) {
      return;
    }

    const resp = await app.alt_next_second(opts.first_token);
    if (resp.status !== 200) {
      app.nav.goto_error_page(resp.data);
      return;
    }
    app.nav.goto_alt_second_stage({ email: opts.email, ...resp.data });
  })();
</script>

{#if opts.new_user}
  <Layout>
    <NewUserInfo
      user_id_hints={opts.user_id_hints}
      email={opts.email}
      onNext={async (data) => {
        const resp = await app.alt_next_second(opts.first_token, data);
        if (resp.status !== 200) {
          app.nav.goto_error_page(resp.data);
          return;
        }
        app.nav.goto_alt_second_stage({ email: opts.email, ...resp.data });
      }}
    />
  </Layout>
{/if}
