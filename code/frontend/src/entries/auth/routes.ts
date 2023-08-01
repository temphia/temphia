import Start from "./pages/start/index.svelte";
import Login from "./pages/login/login.svelte";
import LoginNextStage from "./pages/login/nextstage/index.svelte";
import FirstStage from "./pages/alt/firststage/index.svelte";
import SecondStage from "./pages/alt/secondstage/index.svelte";

import Signup from "./pages/signup/index.svelte";
import SignupNextStage from "./pages/signup/nextstage/index.svelte";

import Reset from "./pages/reset/index.svelte";
import ResetFinal from "./pages/reset/finish/index.svelte";

import PreHook from "./pages/prehook/index.svelte";
import Final from "./pages/final/index.svelte";
import ErrorPage from "./pages/error/index.svelte";

export default {
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
};
