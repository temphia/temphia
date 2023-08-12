<script lang="ts">
  import { AuthService } from "../../../services";
  export let app: AuthService;

  let emailuser;
  let password;
  let message;

  const submit = async () => {
    const resp = await app.login_next(emailuser, password);
    if (resp.status !== 200) {
      return;
    }

    if (resp.data["ok"]) {
      app.nav.goto_login_next_stage(resp.data);
    } else {
      message = resp.data["message"];
    }
  };
</script>

{#if message}
  <p class="text-red-500">{message}</p>
{/if}

<div class="mt-4">
  <label
    class="block mb-2 text-sm font-medium text-gray-600"
    for="LoggingEmailAddress">Email Address / Username</label
  >
  <input
    id="LoggingEmailAddress"
    class="block w-full px-4 py-2 text-gray-700 bg-white border border-gray-300 rounded-md focus:border-blue-500 focus:outline-none focus:ring"
    bind:value={emailuser}
    type="email"
  />
</div>

<div class="mt-4">
  <div class="flex justify-between">
    <label
      class="block mb-2 text-sm font-medium text-gray-600"
      for="loggingPassword">Password</label
    >
    <a href="#" class="text-xs text-gray-500 hover:underline"
      >Forget Password?</a
    >
  </div>

  <input
    class="block w-full px-4 py-2 text-gray-700 bg-white border border-gray-300 rounded-md focus:border-blue-500 focus:outline-none focus:ring"
    type="password"
    bind:value={password}
  />
</div>

<div class="mt-8">
  <button
    on:click={submit}
    class="w-full px-4 py-2 tracking-wide text-white font-semibold transition-colors duration-200 transform bg-blue-700 rounded hover:bg-blue-400"
  >
    Login
  </button>
</div>
