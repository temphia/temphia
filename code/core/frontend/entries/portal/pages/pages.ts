import StartPage from "./start/start.svelte";
import admin_pages from "./admin";
import Loader from "./data/loader.svelte";

export default {
  "/": StartPage,
  "/data": {
    "/": Loader,
    
  },

  "/cabinet": {
    "/": null,
  },
  "/play": null,

  ...admin_pages,
};


/*

      <!-- store stuff -->
      <Route path="/z/portal/launcher/:plugid/:agentid" let:params>
        <Launcher plugid={params.plugid} agentid={params.agentid} />
      </Route>

      <Route
        path="/z/portal/apps_launcher"
        let:params
        component={AppsLauncher}
      />

     <Route path="/z/portal/admin/bprints/:id/instance/plug/:file" let:params>
      <InstancePlug bid={params.id} file={params.file} />
    </Route>

    <Route path="/z/portal/admin/bprints/:id/instance/data_group/:file" let:params>
      <InstanceDataGroup bid={params.id} file={params.file} />
    </Route>

    <Route path="/z/portal/admin/bprints/:id/instance/data_table/:file" let:params>
      <InstanceDataTable bid={params.id} file={params.file} />
    </Route>

*/