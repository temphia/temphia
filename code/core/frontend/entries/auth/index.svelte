<script>
  import { routes, Router, params } from "svelte-hash-router";
  import { setContext } from "svelte";
  import auth_routes from "./routes";
  import Tailwind from "../xcompo/common/_tailwind.svelte";

  import { AuthApp } from "./app";

  const app = new AuthApp();

  setContext("_auth_app_", app);

  routes.set(auth_routes);
</script>

<Router />

<Tailwind />
