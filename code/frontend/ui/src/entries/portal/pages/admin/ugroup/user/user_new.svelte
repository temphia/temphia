<script lang="ts">
  import { getContext } from "svelte";
  import { validateEmail, validateUserId } from "../../../../../../lib/utils";
  import { AutoForm, PortalService } from "../../core";
  import { params } from "svelte-hash-router";

  const app = getContext("__app__") as PortalService;

  export let ugroup = $params.ugroup;

  let message = "";
  let data = {};

  if (app.nav.options) {
    const user_group = app.nav.options["new_user_user_group"];
    if (user_group) {
      data["group_id"] = user_group;
    }
  }

  const save = async (_data) => {
    if (!validateEmail(_data["email"])) {
      message = "Invalid Email";
      return;
    }

    const userid = _data["user_id"] || "";
    if (!userid || !validateUserId(userid)) {
      message = "Invalid user id only use a-z,1-9 and _ chars";
      return;
    }

    if (userid.length > 16 || userid.length < 4) {
      message = "user id should be shorter than 16 and atleast 4 chars";
    }

    message = "";
    const api = await app.api_manager.get_ugroup_tkt_api(ugroup);
    if (!api) {
      console.log("ugroup tkt api not found");
      return;
    }

    const resp = await api.new({...data, ..._data});
    if (!resp.ok) {
      console.log("Err", resp);
      message = resp.data;
      return;
    }

    app.nav.admin_ugroup_users(ugroup);
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
  {data}
/>
