// lauch
import Start from "./start/start.svelte";

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
import RepoItem from "./repo/repo_item.svelte";

// profile
import SelfProfile from "./profile/self.svelte";
import UserProfile from "./profile/user.svelte";
import SelfDevices from "./profile/device/devices.svelte";
import SelfDeviceNew from "./profile/device/device_new.svelte";

import admin_pages from "./admin/admin_routes";
import Play from "./play/play.svelte";
import NotFound from "./notfound.svelte";

import LaunchTargetPage from "../launcher/target_page.svelte";
import Launcher from "../launcher/launcher.svelte";
import Breath from "./play/breath.svelte";

export default {
  "/": Start,
  "/admin": admin_pages,
  "/data": {
    "/": DataLoader,
    "/:source": DataGroups,
    "/:source/:dgroup": DataGroup,
    "/:source/:dgroup/:dtable": DataTable,
    "/:source/:dgroup/:dtable/:layout": DataTable,
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
    "/:source/:group/*": RepoItem,
  },

  "/play": Play,
  "/play/breath": Breath,
  "/profile": {
    "/self": SelfProfile,
    "/user/:id": UserProfile,
    "/device": SelfDevices,
    "/device/new": SelfDeviceNew,
  },

  "/launch": {
    "/": Launcher,
    "/:target": LaunchTargetPage,
    "/:target/*": LaunchTargetPage,
  },

  "*": NotFound,
};
