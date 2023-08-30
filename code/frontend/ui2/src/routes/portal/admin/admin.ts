import {
    BluprintIcon,
    EventIcon,
    GroupIcon,
    UserIcon,
    PlugIcon,
    CodeIcon,
    OrgIcon,
    UserGroupIcon,
    ResourceIcon,
    StoreIcon,
    SearchIcon,
  } from "$lib/compo/svg";


export const items = [
    {
      id: "repo",
      name: "Repos",
      icon: StoreIcon,
      path: "/repo",
    },

    {
      id: "bprint",
      name: "Bluprints",
      icon: BluprintIcon,
      path: "/bprint",
    },
    {
      id: "plug",
      name: "Plugs",
      icon: PlugIcon,
      path: "/plug",
    },
    {
      id: "resource",
      name: "Resources",
      icon: ResourceIcon,
      path: "/resource",
    },
    {
      id: "target",
      name: "Target Apps and Hooks",
      icon: CodeIcon,
      path: "/target/app/",
    },
    {
      id: "lens",
      name: "lens",
      icon: SearchIcon,
      path: "/lens/logs",
    },

    {
      id: "ugroup",
      name: "Users and Groups",
      icon: UserGroupIcon,
      path: "/ugroup",
    },
    {
      id: "tenant",
      name: "Organization",
      icon: OrgIcon,
      path: "/tenant",
    },

    {
      id: "data",
      name: "Data Tables",
      icon: GroupIcon,
      path: "/data",
    },
  ];