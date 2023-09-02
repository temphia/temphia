<script lang="ts">
    import { getContext } from "svelte";
    import Issuer from "./panels/issuer/issuer.svelte";
    import {
        AutoTable,
        LoadingSpinner,
        FloatingAdd,
        PortalService,
    } from "$lib/core";

    import PickNewBprint from "./panels/_pick_new_bprint.svelte";

    let datas = [];
    let loading = true;
    const app = getContext("__app__") as PortalService;

    const load = async () => {
        const api = app.api_manager.get_admin_bprint_api();
        const resp = await api.list();
        if (!resp.ok) {
            return;
        }

        datas = resp.data;
        loading = false;
    };

    load();

    // actions

    const action_instance = async (id: string) => {
        app.nav.admin_bprint_instancer(id);
    };

    const action_edit = (id: string) => app.nav.admin_bprint(id);
    const action_issue = (id: string) =>
        app.utils.small_modal_open(Issuer, { app, bid: id });

    const action_goto_files = (id: string) => app.nav.admin_bprint_files(id);

    const action_delete = async (id: string) => {
        const api = app.api_manager.get_admin_bprint_api();
        await api.delete(id);
        load();
    };
    const action_new = () => {
        app.utils.small_modal_open(PickNewBprint, { app });
    };


    const action_zipit = (id: string) => app.nav.admin_bprint_export_zip(id);
</script>

{#if loading}
    <LoadingSpinner />
{:else}
    <AutoTable
        show_drop={true}
        color={["type"]}
        action_key="id"
        actions={[
            {
                Name: "Instance",
                Class: "bg-blue-400",
                icon: "document-download",
                Action: action_instance,
            },
            {
                Name: "Edit",
                Action: action_edit,
                icon: "pencil-alt",
                drop: true,
            },

            {
                Name: "Files",
                Action: action_goto_files,
                Class: "bg-green-400",
                icon: "document-duplicate",
            },

            {
                Name: "Issue",
                Action: action_issue,
                drop: true,
                icon: "terminal",
            },

            {
                Name: "Zip it",
                Action: action_zipit,
                drop: true,
                icon: "archive",
            },

            {
                Name: "Delete",
                drop: true,
                icon: "trash",
                Action: action_delete,
            },
        ]}
        key_names={[
            ["id", "ID"],
            ["name", "Name"],
            ["slug", "Slug"],
            ["type", "Type"],
            ["sub_type", "Sub Type"],
        ]}
        {datas}
    />
{/if}

<FloatingAdd onClick={action_new} />
