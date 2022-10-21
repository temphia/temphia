<script lang="ts">

  import { getContext } from "svelte";
  import type { AuthApp } from "../../../app";
  import Layout from "../../common/layout.svelte";
  const app: AuthApp = getContext("_auth_app_");

  const opts: {
    ok: number;
    message: string;
    next_token: string;
    user_id: string;
    email_verify: boolean;
    email: string;
  } = app.nav.nav_options;

  if (!opts) {
    app.nav.goto_login_page();
  }

  let code = "";
  (async () => {
    if (!opts.email_verify) {
      const resp = await app.submit_alt_auth(opts.next_token);
      if (resp.status !== 200) {
        app.nav.goto_error_page(resp.data);
        return;
      }
      app.save_preauthed_data(resp.data)
      app.nav.goto_prehook_page(resp.data);
    }
  })();
</script>

{#if opts.email_verify}
  <Layout>
    <div>
      <label
        for="code"
        class="text-sm font-medium text-gray-900 block mb-2 mt-4 "
        >Verify your email: <span class="text-blue-700">{opts.email}</span>
      </label>
      <input
        bind:value={code}
        type="text"
        name="code"
        id="code"
        class="bg-gray-100 h-20 focus:bg-white border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 "
      />
    </div>

    <button
      class="w-full px-4 py-2 tracking-wide text-white transition-colors duration-200 transform bg-blue-700 rounded hover:bg-blue-400 mt-4"
    >
      Verify
    </button>
  </Layout>
{/if}
