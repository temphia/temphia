<script lang="ts">
    import { getContext } from "svelte";
    import { AutoForm, LoadingSpinner, PortalService } from "$lib/core";

    export let data;

    let bid = data["slug"];

    const app = getContext("__app__") as PortalService;
    const api = app.api_manager.get_admin_bprint_api();

    let message = "";
    let rdata = {};
    let loading = true;

    const load = async () => {
        const resp = await api.get(bid);
        if (!resp.ok) {
            message = resp.data;
            return;
        }

        rdata = resp.data;
        loading = false;
    };

    const save = async (_data) => {
        console.log("@@data", _data);

        const resp = await api.update(bid, _data);
        if (!resp.ok) {
            message = resp.data;
            return;
        }
        app.nav.admin_bprints();
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
                    disabled: true,
                },
                {
                    name: "Name",
                    ftype: "TEXT",
                    key_name: "name",
                },
                {
                    name: "Group",
                    ftype: "TEXT",
                    key_name: "group",
                    options: ["plug", "tschema"],
                },

                {
                    name: "Sub Group",
                    ftype: "TEXT",
                    key_name: "sub_group",
                },
                {
                    name: "Source",
                    ftype: "TEXT",
                    key_name: "source",
                },
                {
                    name: "Extra Meta",
                    ftype: "KEY_VALUE_TEXT",
                    key_name: "extra_meta",
                },
            ],
            name: "Edit Blueprint",
            required_fields: [],
        }}
        onSave={save}
        data={rdata}
    />
{/if}
