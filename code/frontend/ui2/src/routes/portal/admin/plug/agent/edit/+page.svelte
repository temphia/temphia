<script lang="ts">
    import { getContext } from "svelte";
    import { AutoForm, PortalService, LoadingSpinner } from "$lib/core";
    import { params  } from "$lib/params";

    let data = $params;

    let pid = data["pid"];
    let aid = data["aid"];

    const app = getContext("__app__") as PortalService;
    const api = app.api_manager.get_admin_plug_api();

    let message = "";
    let rdata = {};
    let loading = true;

    let executors = [];
    app.api_manager.self_api
        .list_executors()
        .then((resp) => {
            executors = resp.data || [];
        })
        .catch(() => {});

    const load = async () => {
        const resp = await api.get_agent(pid, aid);
        if (!resp.ok) {
            message = resp.data;
            return;
        }

        rdata = resp.data;
        loading = false;
    };

    const save = async (_data) => {
        const resp = await api.update_agent(pid, aid, _data);
        if (!resp.ok) {
            message = resp.data;
            return;
        }
        app.nav.admin_agents(pid);
    };

    load();
</script>

{#if loading}
    <LoadingSpinner />
{:else}
    <AutoForm
        {message}
        schema={{
            fields: [
                {
                    name: "Id",
                    ftype: "TEXT_SLUG",
                    key_name: "id",
                },
                {
                    name: "Renderer",
                    ftype: "TEXT",
                    key_name: "renderer",
                },
                {
                    name: "Executor",
                    ftype: "SELECT",
                    key_name: "executor",
                    options: executors,
                },
                {
                    name: "Entry File",
                    ftype: "TEXT",
                    key_name: "entry_file",
                },
                {
                    name: "Web Files",
                    ftype: "KEY_VALUE_TEXT",
                    key_name: "web_files",
                },
                {
                    name: "Extra Meta",
                    ftype: "KEY_VALUE_TEXT",
                    key_name: "extra_meta",
                },
            ],
            name: "Edit Agent",
            required_fields: [],
        }}
        onSave={save}
        data={rdata}
    />
{/if}
