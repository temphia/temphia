<script lang="ts">
    import { getContext } from "svelte";
    import { AutoForm, LoadingSpinner, PortalService } from "$lib/core";
    import { params } from "$lib/params";

    export let pid = $params["pid"];
    export let aid = $params["aid"];
    export let lid = $params["lid"];

    const app = getContext("__app__") as PortalService;
    const api = app.api_manager.get_admin_plug_api();

    let message = "";
    let data = {};
    let loading = true;

    const load = async () => {
        const resp = await api.get_agent_link(pid, aid, lid);
        if (!resp.ok) {
            return;
        }

        data = resp.data;
        loading = false;
    };

    load();

    const save = async (_data) => {
        const resp = await api.update_agent_link(pid, aid, lid, _data);
        if (!resp.ok) {
            message = resp.data;
            return;
        }
        app.nav.admin_agents(pid);
    };
</script>

{#if loading}
    <LoadingSpinner />
{:else}
    <AutoForm
        {message}
        schema={{
            fields: [
                {
                    name: "Name",
                    ftype: "TEXT",
                    key_name: "name",
                },
                {
                    name: "From Plug",
                    ftype: "TEXT",
                    key_name: "from_plug_id",
                    disabled: true,
                },
                {
                    name: "From Agent",
                    ftype: "TEXT",
                    key_name: "from_agent_id",
                    disabled: true,
                },

                {
                    name: "To Plug",
                    ftype: "TEXT",
                    key_name: "to_plug_id",
                },
                {
                    name: "To Agent",
                    ftype: "TEXT",
                    key_name: "to_agent_id",
                },

                {
                    name: "To Handler",
                    ftype: "TEXT",
                    key_name: "to_handler",
                },
                {
                    name: "Extra Meta",
                    ftype: "KEY_VALUE_TEXT",
                    key_name: "extra_meta",
                },
            ],
            name: "Update Link",
            required_fields: [],
        }}
        onSave={save}
        {data}
    />
{/if}
