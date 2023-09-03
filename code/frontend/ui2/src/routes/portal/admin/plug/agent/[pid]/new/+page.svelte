<script lang="ts">
    import { getContext } from "svelte";
    import { AutoForm, PortalService } from "$lib/core";

    export let data;

    let pid = data["pid"];

    const app = getContext("__app__") as PortalService;
    const api = app.api_manager.get_admin_plug_api();

    let executors = [];
    app.api_manager.self_api
        .list_executors()
        .then((resp) => {
            executors = resp.data || [];
        })
        .catch(() => {});

    let message = "";

    const save = async (_data) => {
        const resp = await api.new_agent(pid, _data);
        if (!resp.ok) {
            message = resp.data;
            return;
        }
        app.nav.admin_agents(pid);
    };
</script>

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
                name: "Name",
                ftype: "TEXT",
                key_name: "name",
            },
            {
                name: "Type",
                ftype: "TEXT",
                key_name: "type",
            },
            {
                name: "Executor",
                ftype: "SELECT",
                key_name: "executor",
                options: executors,
            },
            {
                name: "Interface File",
                ftype: "TEXT",
                key_name: "iface_file",
            },
            {
                name: "Entry File",
                ftype: "TEXT",
                key_name: "entry_file",
            },
            {
                name: "Web Entry",
                ftype: "TEXT",
                key_name: "web_entry",
            },
            {
                name: "Web Script",
                ftype: "TEXT",
                key_name: "web_script",
            },
            {
                name: "Web Style",
                ftype: "TEXT",
                key_name: "web_style",
            },
            {
                name: "Web Loader",
                ftype: "TEXT",
                key_name: "web_loader",
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
        name: "New Agent",
        required_fields: [],
    }}
    onSave={save}
    data={{}}
/>
