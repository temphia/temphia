<script lang="ts">
  import { getContext } from "svelte";
  import type { AuthService } from "../../services";

  const app: AuthService = getContext("_auth_app_");

  (async () => {
    if (!app.site_manager.isLogged()) {
      app.nav.goto_login_page();
      return;
    }
    app.nav.goto_final_page();
    return;
  })();
</script>
