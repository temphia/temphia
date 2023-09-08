<script lang="ts">
    import { getContext } from "svelte";
    import { validateEmail, validateSlug } from "$lib/utils";
    import { AutoForm, PortalService } from "$lib/core";

    const app = getContext("__app__") as PortalService;
    const api = app.api_manager.get_admin_user_api();

    let message = "";

    const save = async (_data) => {
        if (!validateEmail(_data["email"])) {
            message = "Invalid Email";
            return;
        }

        const userid = _data["user_id"] || "";
        if (!userid || !validateSlug(userid)) {
            message = "Invalid user id";
            return;
        }

        message = "";

        const resp = await api.new(_data);
        if (!resp.ok) {
            message = resp.data;
            return;
        }

        app.nav.admin_users();
    };
</script>

<AutoForm
    {message}
    schema={{
        fields: [
            {
                name: "Full Name",
                ftype: "TEXT",
                key_name: "full_name",
            },
            {
                name: "User Id",
                ftype: "TEXT",
                key_name: "user_id",
            },

            {
                name: "Email",
                ftype: "TEXT",
                key_name: "email",
            },

            {
                name: "Group",
                ftype: "TEXT",
                key_name: "group_id",
            },

            {
                name: "Bio",
                ftype: "LONG_TEXT",
                key_name: "bio",
            },

            {
                name: "Password",
                ftype: "TEXT",
                key_name: "password",
            },

            {
                name: "Public Key",
                ftype: "LONG_TEXT",
                key_name: "pub_key",
            },

            {
                name: "Active",
                ftype: "BOOL",
                key_name: "active",
            },
        ],
        name: "New User",
        required_fields: [],
    }}
    onSave={save}
    data={{}}
/>
