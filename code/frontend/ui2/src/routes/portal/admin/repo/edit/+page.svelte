<script lang="ts">
    import { getContext } from "svelte";
    import type { PortalService } from "$lib/core";
    import { AutoForm, LoadingSpinner } from "$lib/core";

    const app = getContext("__app__") as PortalService;
    const rapi = app.api_manager.get_admin_repo_api();

    export let data;

    let rid = data["slug"];

    let rdata = {};
    let message = "";

    let loading = true;

    const load = async () => {
        const resp = await rapi.get(rid);
        if (resp.status !== 200) {
            console.log("Err", resp);
            return;
        }

        rdata = resp.data;
        loading = false;
    };

    load();

    const saveHandle = async (_data) => {
        const resp = await rapi.update(rid, _data);
        if (!resp.ok) {
            console.log("Err", resp);
            message = resp.data;
            return;
        }

        app.nav.admin_repos();
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
                    name: "Provider",
                    ftype: "TEXT",
                    key_name: "provider",
                },

                {
                    name: "URL",
                    ftype: "TEXT",
                    key_name: "url",
                },

                {
                    name: "Extra Meta",
                    ftype: "KEY_VALUE_TEXT",
                    key_name: "extra_meta",
                },
            ],
            name: "Edit Repo",
            required_fields: [],
        }}
        onSave={saveHandle}
        data={rdata}
    />
{/if}
