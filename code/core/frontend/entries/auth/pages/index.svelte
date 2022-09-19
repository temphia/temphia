<script>
  import { routes, Router, params } from "svelte-hash-router";

  import Start from "./start/index.svelte";
  import Login from "./login/login.svelte";
  import LoginNextStage from "./login/nextstage/index.svelte";
  import FirstStage from "./alt/firststage/index.svelte";
  import SecondStage from "./alt/secondstage/index.svelte";

  import Signup from "./signup/index.svelte";
  import SignupNextStage from "./signup/nextstage/index.svelte";

  import Reset from "./reset/index.svelte";
  import ResetFinal from "./reset/finish/index.svelte";

  import PreHook from "./prehook/index.svelte";
  import Final from "./final/index.svelte";
  import ErrorPage from "./error/index.svelte";

  import { AuthApp } from "../../lib/app/auth/auth";
  import { setContext } from "svelte";
  import Playground from "./common/playground.svelte";
  import Tailwind from "../common/_tailwind.svelte";

  const app = new AuthApp();

  setContext("_auth_app_", app);

  routes.set({
    "/": Start,
    "/login": {
      "/": Login,
      "/next_stage": LoginNextStage,
    },
    "/signup": {
      "/": Signup,
      "/next_stage": SignupNextStage,
    },
    "/final": Final,
    "/prehook": PreHook,
    "/reset": {
      "/": Reset,
      "/finish": ResetFinal,
    },
    "/alt": {
      "/first_stage": FirstStage,
      "/second_stage": SecondStage,
    },
    "/error": ErrorPage,
  });
</script>

<Router />

<Tailwind />
