<script lang="ts">
    import { getContext } from "svelte";
    import {
        AutoTable,
        LoadingSpinner,
        FloatingAdd,
        PortalService,
    } from "$lib/core";

    import { params  } from "$lib/params";

    let data = $params;

    export let pid = data["pid"];

    let datas = [];
    let loading = true;
    const app = getContext("__app__") as PortalService;

    const load = async () => {
        const api = app.api_manager.get_admin_plug_api();
        const resp = await api.list_agent(pid);
        if (!resp.ok) {
            return;
        }

        datas = resp.data;
        loading = false;
    };

    load();

    // actions

    const action_agent_edit = (id: string) => app.nav.admin_agent_edit(pid, id);

    const action_execute = (id: string) =>
        app.nav.admin_plug_dev_execute(pid, id);
    const action_agent_links = (id: string) =>
        app.nav.admin_agent_links(pid, id);
    const action_agent_exts = (id: string) => app.nav.admin_agent_ext(pid, id);
    const action_agent_resources = (id: string) =>
        app.nav.admin_agent_res(pid, id);

    const action_live_shell = (id: string) =>
        app.nav.admin_plug_dev_live_shell(pid, id);
    const action_logs = (id: string) => {
        app.nav.admin_lens_logs({
            filters: {
                plug_id: pid,
                agent_id: id,
            },
        });
    };

    const action_dev_reset = async (id: string) => {
        const eapi = app.api_manager.get_engine_api();
        loading = true;
        await eapi.reset(pid, id);
        loading = false;
    };

    const action_dev_docs = (id: string) =>
        app.nav.admin_plug_dev_docs(pid, id);
    const action_delete = async (id: string) => {
        const api = app.api_manager.get_admin_plug_api();
        await api.delete_agent(pid, id);
        load();
    };

    const action_new = () => app.nav.admin_agent_new(pid);
</script>

{#if loading}
    <LoadingSpinner />
{:else}
    <AutoTable
        action_key="id"
        actions={[
            {
                Name: "Execute",
                Class: "bg-blue-400",

                Action: action_execute,
                icon: "lightning-bolt",
            },
            {
                Name: "Edit",
                Action: action_agent_edit,
                drop: true,
                icon: "pencil-alt",
            },
            {
                Name: "Extern Execute",
                Action: (agent_id) => {},
                drop: true,
                icon: "lightning-bolt",
            },

            {
                Name: "Duplicate",
                Action: (id) => {},
                drop: true,
                icon: "duplicate",
            },

            {
                Name: "Docs",
                drop: true,
                icon: "book-open",
                Action: action_dev_docs,
            },

            {
                Name: "Links",
                drop: true,
                icon: "link",
                Action: action_agent_links,
            },

            {
                Name: "Extensions",
                drop: true,
                icon: "puzzle",
                Action: action_agent_exts,
            },

            {
                Name: "Resources",
                drop: true,
                icon: "paper-clip",
                Action: action_agent_resources,
            },

            {
                Name: "Live Shell",
                drop: true,
                icon: "eye",
                Action: action_live_shell,
            },

            {
                Name: "Reset",
                Action: action_dev_reset,
                drop: true,
                icon: "refresh",
            },

            {
                Name: "Logs",
                drop: true,
                icon: "document-search",
                Action: action_logs,
            },

            {
                Name: "Delete",
                Action: action_delete,
                drop: true,
                icon: "trash",
            },
        ]}
        key_names={[
            ["id", "ID"],
            ["name", "Name"],
            ["type", "Type"],
            ["executor", "Executor"],
            ["plug_id", "Plug Id"],
        ]}
        {datas}
        show_drop={true}
        color={["type", "executor"]}
    />
{/if}

<FloatingAdd onClick={action_new} />
