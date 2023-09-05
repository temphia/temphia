<script lang="ts">
    import { getContext } from "svelte";
    import { AutoTable, ActionAddButton, PortalService } from "$lib/core";
    import TopActions from "$lib/core/top_actions.svelte";

    const app: PortalService = getContext("__app__");
    const tapi = app.api_manager.get_admin_tenant_api();

    let domains = [];

    let tenant = {};
    let loaded = false;

    const load = async () => {
        const resp2 = await tapi.get_domains();
        if (!resp2.ok) {
            return;
        }

        domains = resp2.data;
        loaded = true;
    };

    load();
</script>

<div class="h-full w-full overflow-auto">
    <TopActions
        actions={{
            "System KV": () => app.nav.admin_tenant_system_kvs(),
            "System Event": () => app.nav.admin_tenant_system_events(),
        }}
    />
    <div class="md:p-8 bg-indigo-100 flex flex-row flex-wrap">
        <div
            class="md:w-1/2-screen m-0 p-5 bg-white w-full tw-h-full shadow md:rounded-lg relative"
        >
            <div class="text-2xl text-indigo-900">Organization</div>
            {#if loaded}
                <div class="flex-col flex py-3">
                    <label class="pb-2 text-gray-700 font-semibold">Name</label>
                    <input
                        type="text"
                        disabled
                        value={tenant["name"] || ""}
                        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
                    />
                </div>

                <div class="flex-col flex py-3 relative">
                    <label class="pb-2 text-gray-700 font-semibold">Slug</label>
                    <input
                        type="text"
                        value={tenant["slug"]}
                        disabled
                        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
                    />
                </div>

                <div class="flex-col flex py-3 relative border rounded p-2">
                    <div class="absolute right-1">
                        <ActionAddButton
                            onClick={() => app.nav.admin_tenant_domain_new()}
                        />
                    </div>

                    <label class="pb-2 text-gray-700 font-semibold"
                        >Domains</label
                    >

                    <AutoTable
                        action_key="id"
                        show_drop={true}
                        actions={[
                            {
                                Name: "Adapter Editor",
                                Action: (id) =>
                                    app.nav.admin_tenant_domain_adapter_editor(
                                        id
                                    ),
                                icon: "lightning-bolt",
                            },

                            {
                                Name: "Edit",
                                Action: (id) =>
                                    app.nav.admin_tenant_domain_edit(id),
                                icon: "pencil",
                            },

                            {
                                Name: "Hooks",
                                Action: (id) => {},
                                drop: true,
                                icon: "hashtag",
                            },

                            {
                                Name: "Widgets",
                                Action: (id) => {},
                                drop: true,
                                icon: "hashtag",
                            },

                            {
                                Name: "Reset",
                                drop: true,
                                icon: "refresh",
                                Action: (id) => {
                                    tapi.domain_adapter_reset(id);
                                },
                            },

                            {
                                Name: "Delete",
                                drop: true,
                                icon: "trash",
                                Action: async (id) => {
                                    await tapi.delete_domain(id);
                                    load();
                                },
                            },
                        ]}
                        key_names={[
                            ["id", "ID"],
                            ["name", "Name"],
                            ["adapter_type", "Http Adapter"],
                            ["about", "About"],
                        ]}
                        datas={domains}
                        color={["adapter_type"]}
                    />
                </div>
            {/if}
        </div>
    </div>
</div>
