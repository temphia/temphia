// lauch
import Start from "./start/start.svelte";

// data
import DataLoader from "./data/loader.svelte";
import DataGroups from "./data/groups.svelte";

import DataCustomLoader from "./data/custom/loader.svelte"
import DataTableLoader from "./data/table/loader.svelte"
import DataTable from "./data/table/table.svelte";
import DataSheetLoader from "./data/sheet/loader.svelte";
import DataSheet from "./data/sheet/sheet.svelte";
import DataSheetsNew from "./data/sheet/new/sheets_new.svelte";

// cabinet
import CabLoader from "./cabinet/loader.svelte";
import CabFolders from "./cabinet/folders.svelte";
import CabFolder from "./cabinet/folder.svelte";
import CabFile from "./cabinet/file.svelte";
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
    "/:source/:dgroup/custom": DataCustomLoader,
    "/:source/:dgroup/sheet": DataSheetLoader, 
    "/:source/:dgroup/sheet/:sheet": DataSheet,
    "/:source/:dgroup/table": DataTableLoader,
    "/:source/:dgroup/table/:dtable": DataTable,
    "/:source/:dgroup/table/:dtable/:layout": DataTable,
  },
  "/data_sheets/new": DataSheetsNew,

  "/cabinet": {
    "/": CabLoader,
    "/:source": CabFolders,
    "/:source/:folder": CabFolder,
    "/:source/:folder/*": CabFile,
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
