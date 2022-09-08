<script lang="ts">
  import Modal from "svelte-simple-modal";
  import { Router, Route } from "svelte-routing";

  import BigModal from "../../components/modal/big.svelte";
  import Tailwind from "../common/_tailwind.svelte";
  import Magicdial from "./pages/magicdial/magicdial.svelte";
  import Start from "./pages/start/start.svelte";
  import Tenants from "./pages/tenant/tenants.svelte";
  import Tenant from "./pages/tenant/tenant.svelte";
  import NewTenant from "./pages/tenant/tenant_new.svelte";

  import { goto, loadOperatorData } from "./operator";
  import Querytool from "./pages/query/querytool.svelte";
  import Logs from "./pages/logs/logs.svelte";

  let url = "/z/operator";

  loadOperatorData();
</script>

<Tailwind />

<div
  class="h-screen flex flex-row flex-auto flex-shrink-0 antialiased bg-gray-50 text-gray-800"
>
  <div class="flex flex-col top-0 left-0 w-64 bg-gray-900 h-full shadow-lg">
    <div class="flex items-center pl-6 h-20 border-b border-gray-800">
      <div class="ml-1">
        <p
          class="ml-1 text-md font-medium tracking-wide truncate text-gray-100 font-sans"
        >
          Operator Console
        </p>
      </div>
    </div>
    <div class="overflow-y-auto overflow-x-hidden flex-grow">
      <ul class="flex flex-col py-6 space-y-1">
        <li on:click={goto("/z/operator")}>
          <a
            href="#"
            class="relative flex flex-row items-center h-11 focus:outline-none hover:bg-gray-700 text-gray-500 hover:text-gray-200 border-l-4 border-transparent hover:border-blue-500 pr-6"
          >
            <span class="inline-flex justify-center items-center ml-4">
              <svg
                class="w-5 h-5"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
                xmlns="http://www.w3.org/2000/svg"
                ><path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"
                /></svg
              >
            </span>
            <span
              class="ml-2 font-semibold text-sm tracking-wide truncate font-sans"
              >Home</span
            >
          </a>
        </li>
        <li on:click={goto("/z/operator/tenants")}>
          <a
            href="#"
            class="relative flex flex-row items-center h-11 focus:outline-none hover:bg-gray-700 text-gray-500 hover:text-gray-200 border-l-4 border-transparent hover:border-blue-500 pr-6"
          >
            <span class="inline-flex justify-center items-center ml-4">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-6 w-6"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
                ><path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4"
                /></svg
              >
            </span>
            <span
              class="ml-2 font-semibold text-sm tracking-wide truncate font-sans"
              >Tenants</span
            >
          </a>
        </li>

        <li on:click={goto("/z/operator/magicdial")}>
          <a
            href="#"
            class="relative flex flex-row items-center h-11 focus:outline-none hover:bg-gray-700 text-gray-500 hover:text-gray-200 border-l-4 border-transparent hover:border-blue-500 pr-6"
          >
            <span class="inline-flex justify-center items-center ml-4">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-5 w-5"
                viewBox="0 0 20 20"
                fill="currentColor"
              >
                <path
                  d="M5 4a1 1 0 00-2 0v7.268a2 2 0 000 3.464V16a1 1 0 102 0v-1.268a2 2 0 000-3.464V4zM11 4a1 1 0 10-2 0v1.268a2 2 0 000 3.464V16a1 1 0 102 0V8.732a2 2 0 000-3.464V4zM16 3a1 1 0 011 1v7.268a2 2 0 010 3.464V16a1 1 0 11-2 0v-1.268a2 2 0 010-3.464V4a1 1 0 011-1z"
                />
              </svg>
            </span>
            <span
              class="ml-2 font-semibold text-sm tracking-wide truncate font-sans"
              >Magic Dial</span
            >
          </a>
        </li>

        <li on:click={goto("/z/operator/logs")}>
          <a
            href="#"
            class="relative flex flex-row items-center h-11 focus:outline-none hover:bg-gray-700 text-gray-500 hover:text-gray-200 border-l-4 border-transparent hover:border-blue-500 pr-6"
          >
            <span class="inline-flex justify-center items-center ml-4">
              <svg
                class="h-6 w-6"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
                stroke-width="2"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  d="M9 17v-2m3 2v-4m3 4v-6m2 10H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"
                />
              </svg>
            </span>
            <span
              class="ml-2 font-semibold text-sm tracking-wide truncate font-sans"
              >Logs</span
            >
          </a>
        </li>

        <li on:click={goto("/z/operator/querytool")}>
          <a
            href="#"
            class="relative flex flex-row items-center h-11 focus:outline-none hover:bg-gray-700 text-gray-500 hover:text-gray-200 border-l-4 border-transparent hover:border-blue-500 pr-6"
          >
            <span class="inline-flex justify-center items-center ml-4">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-6 w-6"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M4 7v10c0 2.21 3.582 4 8 4s8-1.79 8-4V7M4 7c0 2.21 3.582 4 8 4s8-1.79 8-4M4 7c0-2.21 3.582-4 8-4s8 1.79 8 4m0 5c0 2.21-3.582 4-8 4s-8-1.79-8-4"
                />
              </svg>
            </span>
            <span
              class="ml-2 font-semibold text-sm tracking-wide truncate font-sans"
              >Query Tool</span
            >
          </a>
        </li>

        <li>
          <a
            href="#"
            class="relative flex flex-row items-center h-11 focus:outline-none hover:bg-gray-700 text-gray-500 hover:text-gray-200 border-l-4 border-transparent hover:border-red-500 pr-6"
          >
            <span
              class="inline-flex justify-center items-center ml-4 text-red-400"
            >
              <svg
                class="w-5 h-5"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
                xmlns="http://www.w3.org/2000/svg"
                ><path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"
                /></svg
              >
            </span>
            <span
              class="ml-2 font-semibold text-sm tracking-wide truncate font-sans"
              >Logout</span
            >
          </a>
        </li>
      </ul>
    </div>
  </div>

  <div class="w-full h-screen overflow-auto">
    <Router {url}>
      <BigModal />
      <Modal>
        <Route path="/z/operator" component={Start} />
        <Route path="/z/operator/tenants" component={Tenants} />

        <Route path="/z/operator/tenants/:id" let:params>
          <Tenant id={params.id} />
        </Route>
        <Route path="/z/operator/tenant_new">
          <NewTenant />
        </Route>

        <Route path="/z/operator/magicdial" let:params>
          <Magicdial />
        </Route>

        <Route path="/z/operator/logs" let:params>
          <Logs />
        </Route>

        <Route path="/z/operator/querytool" let:params>
          <Querytool />
        </Route>
      </Modal>
    </Router>
  </div>
</div>
