<script lang="ts">
  import Tailwind from "../../src/entries/common/_tailwind.svelte";
  import { doLogin, gotoLoginPage } from "./root";

  const LOGIN = "login";
  const TICKET = "ticket";
  const OPERATOR = "operator";

  $: _mode = LOGIN;
  let message = "";

  let op_username = "";
  let op_password = "";
  let tenant_id = "";
  let group = "";

  const letsGo = async () => {
    switch (_mode) {
      case LOGIN:
        if (!tenant_id || tenant_id === "") {
          message = "Enter valid Tenant Id";
          return;
        }

        gotoLoginPage(tenant_id, group);

        break;
      case OPERATOR:
        message = "";
        message = await doLogin(op_username, op_password);
      default:
        break;
    }
  };
</script>

<section class="flex justify-center items-center h-screen bg-gray-100">
  <div class="max-w-md w-full bg-white rounded p-6 space-y-4 mt-4">
    <div class="mb-4">
      <h2 class="text-xl font-bold">Welcome to Temphia Home!</h2>
    </div>
    <select class="p-2 w-full" bind:value={_mode}>
      <option value={LOGIN}>Goto Login Page</option>
      <option value={TICKET}>Use Ticket</option>
      <option value={OPERATOR}>Operator Login</option>
    </select>

    {#if _mode === LOGIN}
      <div>
        <input
          bind:value={tenant_id}
          class="w-full p-4 text-sm bg-gray-50 focus:outline-none border border-gray-200 rounded text-gray-600"
          type="text"
          placeholder="tenant id"
        />
      </div>

      <div>
        <input
          bind:value={group}
          class="w-full p-4 text-sm bg-gray-50 focus:outline-none border border-gray-200 rounded text-gray-600"
          type="text"
          placeholder="user group (optional)"
        />
      </div>
      <span class="font-sans text-sm text-red-500 italic">{message}</span>
    {:else if _mode === TICKET}
      <div>
        <input
          class="w-full p-4 text-sm bg-gray-50 focus:outline-none border border-gray-200 rounded text-gray-600"
          type="text"
          placeholder="ticket"
        />
      </div>
    {:else if _mode === OPERATOR}
      <div>
        <input
          class="w-full p-4 text-sm bg-gray-50 focus:outline-none border border-gray-200 rounded text-gray-600"
          type="text"
          bind:value={op_username}
          placeholder="username"
        />
      </div>

      <div>
        <input
          class="w-full p-4 text-sm bg-gray-50 focus:outline-none border border-gray-200 rounded text-gray-600"
          type="password"
          bind:value={op_password}
          placeholder="password"
        />
      </div>
    {/if}

    <div>
      <button
        on:click={letsGo}
        class="w-full py-4 bg-blue-600 hover:bg-blue-700 rounded text-sm font-bold text-gray-50 transition duration-200"
        >Go</button
      >
    </div>
  </div>
</section>

<Tailwind />
