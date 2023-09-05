<script lang="ts">
    import { getContext } from "svelte";
    import { AutoForm, PortalService } from "$lib/core";

    const app = getContext("__app__") as PortalService;
    const api = app.api_manager.get_admin_resource_api();

    let message = "";
    let data = {};
    let modified = false;

    if (app.nav.options) {
        modified = true;
        data = { ...app.nav.options };
    }

    const save = async (_data) => {
        const resp = await api.new({ ...data, ..._data });
        if (!resp.ok) {
            message = resp.data;
            return;
        }
        app.nav.admin_resources();
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
                name: "Sub Type",
                ftype: "TEXT",
                key_name: "sub_type",
            },

            {
                name: "Target",
                ftype: "TEXT",
                key_name: "target",
            },

            {
                name: "Policy",
                ftype: "TEXT_POLICY",
                key_name: "policy",
            },
            {
                name: "Payload",
                ftype: "LONG_TEXT",
                key_name: "payload",
            },

            {
                name: "Bprint Id",
                ftype: "TEXT",
                key_name: "bprint_id",
            },

            {
                name: "Extra Meta",
                ftype: "KEY_VALUE_TEXT",
                key_name: "extra_meta",
            },
        ],
        name: "New Resource",
        required_fields: [],
    }}
    onSave={save}
    {data}
    {modified}
/>
