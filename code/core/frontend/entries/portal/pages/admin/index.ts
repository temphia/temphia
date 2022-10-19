import AdminBprint from "./bprint/bprint.svelte";
import AdminBprints from "./bprint/bprints.svelte";
import AdminBprintEditor from "./bprint/editor/editor.svelte";
import AdminBprintInstancer from "./bprint/instancer/instancer.svelte";

// repo
import AdminRepo from "./repo/repo.svelte";
import AdminRepos from "./repo/repos.svelte";

// data
import AdminDataLoader from "./data/loader.svelte";
import AdminDataGroup from "./data/group/group.svelte";
import AdminDataGroups from "./data/group/groups.svelte";
import AdminDataTable from "./data/table/table.svelte";
import AdminDataTables from "./data/table/tables.svelte";
import AdminDataColumns from "./data/column/columns.svelte";
import AdminDataColumn from "./data/column/column.svelte";
import AdminDataHook from "./data/hook/hook.svelte";
import AdminDataHooks from "./data/hook/hooks.svelte";
import AdminDataView from "./data/view/view.svelte";
import AdminDataViews from "./data/view/views.svelte";
import Admin from "./admin.svelte";

export default {
  $$component: Admin,
  "/bprint": {
    "/": AdminBprints,
    "/:bid": AdminBprint,
    "/:bid/editor": AdminBprintEditor,
    "/:bid/instancer": AdminBprintInstancer,
  },
  repo: {
    "/": AdminRepos,
    "/:rid": AdminRepo,
  },

  data: {
    "/": AdminDataLoader,
    "/group": AdminDataGroups,
    "/group/:group": AdminDataGroup,
    "/table/:group": AdminDataTables,
    "/table/:group/:table": AdminDataTable,
    "/column/:group/:table": AdminDataColumns,
    "/column/:group/:table/:column": AdminDataColumn,
    "/hook/:group/:table": AdminDataHooks,
    "/hook/:group/:table/:id": AdminDataHook,
    "/view/:group/:table": AdminDataViews,
    "/view/:group/:table/:id": AdminDataView,
  },

  resource: {},
  user: {},
  ugroup: {},
  lens: {},
  tenant: {},
};
