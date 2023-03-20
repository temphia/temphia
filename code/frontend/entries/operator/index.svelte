<script lang="ts">
  import Modal from "svelte-simple-modal";
  import { Router, Route } from "svelte-routing";
  import BigModal from "../xcompo/modal/big.svelte";
  import Magicdial from "./pages/magicdial/magicdial.svelte";
  import Start from "./pages/start/start.svelte";
  import Tenants from "./pages/tenant/tenants.svelte";
  import Tenant from "./pages/tenant/tenant.svelte";
  import NewTenant from "./pages/tenant/tenant_new.svelte";

  import { loadOperatorData } from "./service";
  import Querytool from "./pages/query/querytool.svelte";
  import Logs from "./pages/logs/logs.svelte";
  import Layout from "./layout.svelte";

  let url = "/z/operator";

  loadOperatorData();
</script>

<Layout>
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
</Layout>
