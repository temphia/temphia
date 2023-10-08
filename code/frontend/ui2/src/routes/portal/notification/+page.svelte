<script lang="ts">
    import { getContext } from "svelte";
    import type { PortalService } from "$lib/services/portal/portal";
    import { LoadingSpinner } from "$lib/compo";

    const app = getContext("__app__") as PortalService;
    const spai = app.api_manager.get_self_api();
    let data = [
        {
            id: 1,
            title: "Welcome User",
            read: true,
            type: "system_message",
            contents:
                "this is temphia interglactic information highway system connection portal blah blah",
            user_id: "superuser",
            created_at: "2023-08-22T16:20:40.772652104+05:45",
        },
    ];
    let loading = true;

    const load = async () => {
        const resp = await spai.list_message();
        if (!resp.ok) {
            return;
        }

        data = resp.data;
        loading = false;
    };

    load();
</script>

{#if loading}
    <LoadingSpinner />
{:else}
    <div class="card p-2 h-full">
        <div class="card-header">
            <h2 class="h2 mb-2">Notifications</h2>
        </div>

        <div>
            <div class="table-container">
                <table class="table text-token table-interactive" role="grid">
                    <thead class="table-head">
                        <tr>
                            <th class="">Title</th>
                            <th class="">Message</th>
                            <th class="">Source</th>
                            <th class="">Date</th>
                            <th class="">Actions</th>
                        </tr>
                    </thead>

                    <tbody class="table-body">
                        {#each data as nt}
                            <tr>
                                <th class="">{nt.title}</th>
                                <td class="" role="gridcell">{nt.contents}</td>
                                <td class="" role="gridcell">{nt.type}</td>
                                <td class="" role="gridcell">{nt.created_at}</td
                                >
                                <td class="" role="gridcell" />
                            </tr>
                        {/each}
                    </tbody>
                </table>
            </div>
        </div>
    </div>
{/if}
