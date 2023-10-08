<script lang="ts">
    import { getContext, onDestroy, onMount } from "svelte";
    import type { PortalService } from "$lib/services/portal/portal";

    import { params } from "$lib/params";

    const app = getContext("__app__") as PortalService;
    const launcher = app.launcher;

    onMount(() => {
        let lopts = app.nav.options || {};

        let instance = launcher.target_index[$params["target"]];
        if (instance) {
            launcher.instance_change(instance);
        } else {
            const name = $params["name"] ? window.atob($params["name"]) : "";
            instance = launcher.instance_by_target({
                invoker_name: "user_app",
                target_id: $params["target"],
                target_name: name,
                target_type: lopts["target_type"] || "",
                startup_payload: {},
                invoker: null,
            });
        }

        launcher.plane_show();
    });

    onDestroy(() => {
        launcher.plane_hide();
    });
</script>
