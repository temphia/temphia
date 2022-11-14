<script lang="ts">
  /*
  import { getContext } from "svelte";
  import { params } from "svelte-hash-router";
  import { LoadingSpinner } from "../../../../../../xcompo";
  import type { PortalService } from "../../../core";
  import Inner from "./inner.svelte";
  import ShellInner from "./shell_inner.svelte";

  export let pid = $params.pid;
  export let aid = $params.aid;

  const app: PortalService = getContext("__app__");

  let message = "";
  let loading = true;
  let tkt = "";
  let interface_loaded = false;
  let iface_data = {};

  const load = async () => {
    // const papi = await app.get_apm().get_plug_api();
    // const resp = await papi.get_plug(pid);
    // if (resp.status !== 200) {
    //   message = resp.data;
    //   loading = false;
    //   return;
    // }

    // loadInterface(papi, resp.data);

    // const bapi = await app.get_apm().get_bprint_api();
    // const resp1 = await bapi.issue_tkt(resp.data["bprint_id"], true, false);
    // if (resp1.status !== 200) {
    //   message = resp1.data;
    //   loading = false;
    //   return;
    // }

    // tkt = resp1.data["dev_ticket"];
    // loading = false;
  };

  const loadInterface = async (papi: PlugAPI, plug: object) => {
    // const resp = await papi.get_agent(pid, aid);
    // if (resp.status !== 200) {
    //   loading = false;
    //   message = "Error loading agent.";
    //   return;
    // }

    // const ifaceUrl: string = resp.data["iface_file"];
    // if (!ifaceUrl) {
    //   console.log("Empty openrpc iface file");
    //   return;
    // }

    // const bapi = await app.get_apm().get_bprint_api();
    // if (ifaceUrl.startsWith("http://") || ifaceUrl.startsWith("https://")) {
    //   const bresp = await fetch(ifaceUrl);
    //   iface_data = await bresp.json();
    //   interface_loaded = true;
    // } else if (ifaceUrl.startsWith("@bprint/")) {
    //   const bpath = ifaceUrl.replace("@bprint/", "").split("/");
    //   const bresp = await bapi.bprint_get_file(bpath[0], bpath[1]);
    //   if (bresp.status !== 200) {
    //     console.log("connot retrive", bresp);
    //     return;
    //   }

    //   iface_data =
    //     typeof bresp.data === "object"
    //       ? bresp.data
    //       : JSON.stringify(bresp.data);

    //   interface_loaded = true;
    // } else {
    //   const bresp = await bapi.bprint_get_file(plug["bprint_id"], ifaceUrl);
    //   if (bresp.status !== 200) {
    //     console.log("connot retrive", bresp);
    //     return;
    //   }

    //   iface_data =
    //     typeof bresp.data === "object"
    //       ? bresp.data
    //       : JSON.stringify(bresp.data);
    //   interface_loaded = true;
    // }
  };

  $: console.log("@interface =>", interface_loaded, iface_data);

  load();
  */
</script>

<div>Shell</div>

<!-- {#if loading}
  <LoadingSpinner />
{:else}
  <Inner {aid} {pid} {app} tkt="" />
{/if}

{#if !message}
  <ShellInner {aid} {app} {pid} {tkt} />
{:else}
  <p class="bg-red-700">{message}</p>
{/if}
 -->
