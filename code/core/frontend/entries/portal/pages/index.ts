// lauch
import LauncherStart from "./launcher/start.svelte";
import LauncherPlayer from "./launcher/player.svelte";

// data
import DataLoader from "./data/loader.svelte";
import DataGroups from "./data/groups.svelte";
import DataGroup from "./data/group.svelte";
import DataTable from "./data/table.svelte";
// cabinet
import CabLoader from "./cabinet/loader.svelte";
import CabFolders from "./cabinet/folders.svelte";
import CabFolder from "./cabinet/folder.svelte";
//repo
import RepoLoader from "./repo/loader.svelte";
import Repo from "./repo/repo.svelte";

import admin_pages from "./admin";
import Play from "./play/play.svelte";

export default {
  "/": LauncherStart,
  "/launch/:app": LauncherPlayer,

  "/data": {
    "/": DataLoader,
    "/:source": DataGroups,
    "/:source/:dgroup": DataGroup,
    "/:source/:dgroup/:dtable": DataTable,
  },

  "/cabinet": {
    "/": CabLoader,
    "/:source": CabFolders,
    "/:source/:folder": CabFolder,
    "/:source/:folder/:file": null,
  },
  "/repo": {
    "/": RepoLoader,
    "/:source": Repo,
  },
  admin: admin_pages,
  "/play": Play,
};
