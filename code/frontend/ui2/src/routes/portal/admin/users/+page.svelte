<script lang="ts">
    import { getContext } from "svelte";
    import {
        AutoTable,
        LoadingSpinner,
        FloatingAdd,
        PortalService,
    } from "$lib/core";
    import TopActions from "$lib/core/top_actions.svelte";
    import ChangeEmail from "./actions/change_email.svelte";
    import ResetPassword from "./actions/reset_password.svelte";

    const app = getContext("__app__") as PortalService;

    let datas = [];
    let loading = true;
    const api = app.api_manager.get_admin_user_api();

    const load = async () => {
        loading = true;
        const resp = await api.list();
        if (!resp.ok) {
            return;
        }

        datas = resp.data;
        loading = false;
    };

    load();

    // actions

    const action_edit = (id: string) => app.nav.admin_user_edit(id);
    const action_profile = (id: string) => app.nav.user_profile(id);
    const action_delete = async (id: string) => {
        await api.delete(id);
        load();
    };

    // syncme => ../ugroup/users.svelte

    const action_reset_password = (id: string) => {
        app.utils.small_modal_open(ResetPassword, {
            uid: id,
            onComplete: (opts) => {
                console.log("RESET PASSWORD", opts);
                app.utils.small_modal_close();
            },
        });
    };
    const action_email_change = (id: string) => {
        app.utils.small_modal_open(ChangeEmail, {
            uid: id,
            onComplete: (opts) => {
                console.log("CHANGE EMAIL", opts);
                app.utils.small_modal_close();
            },
        });
    };

    const action_roles = (id: string) => {};

    const action_new = () => app.nav.admin_user_new();
</script>

<TopActions
    actions={{
        "User Groups": () => app.nav.admin_ugroups(),
    }}
/>

{#if loading}
    <LoadingSpinner />
{:else}
    <AutoTable
        action_key="user_id"
        show_drop={true}
        actions={[
            {
                Name: "Profile",
                Action: action_profile,
                icon: "user-circle",
            },

            {
                Name: "Edit",
                Action: action_edit,
                icon: "pencil-alt",
            },

            {
                Name: "Disable",
                Action: (id) => {},
                icon: "user-remove",
                drop: true,
            },

            {
                Name: "Reset Password",
                Action: action_reset_password,
                icon: "lock-open",
                drop: true,
            },

            {
                Name: "Change Email",
                Action: action_email_change,
                icon: "at-symbol",
                drop: true,
            },

            {
                Name: "Roles",
                Action: action_roles,
                icon: "identification",
                drop: true,
            },

            {
                Name: "Delete",
                Action: action_delete,
                icon: "trash",
                drop: true,
            },
        ]}
        key_names={[
            ["user_id", "User Id"],
            ["full_name", "Full Name"],
            ["group_id", "Group"],
            ["created_at", "Created At"],
            ["active", "Active"],
        ]}
        color={["group_id", "active"]}
        {datas}
    />
{/if}

<FloatingAdd onClick={action_new} />
