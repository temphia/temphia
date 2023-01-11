<script lang="ts">
  import type { AuthService } from "../../../../lib/app/auth/auth";

  export let app: AuthService;
  export let method: object;
  export let data: object;

  const handleClick = (ev: any) => {
    const redirect_uri = `${window.origin}/z/auth/oauth_redirect`;

    let up = new URLSearchParams({
      client_id: data["client_id"],
      response_type: "code",
      redirect_uri,
      state: data["state_token"],
    });

    data["scopes"].forEach((s) => {
      up.set("scope", s);
    });

    const tabWindow = window.open(
      `${data["auth_url"]}?${up.toString()}`,
      "_blank"
    );

    console.log("@tabwindow =>", tabWindow);

    const i = setInterval(() => {
      console.log("@checking .....");
      const { location } = tabWindow;

      try {
        if (location.href.indexOf(redirect_uri) !== 0) return;
        parseAndExtract(location.search);
      } catch (error) {
        if (
          error instanceof DOMException ||
          error.message === "Permission denied"
        ) {
          return;
        }
        if (!tabWindow.closed) tabWindow.close();
      }

      tabWindow.close();
      clearInterval(i);
    }, 250);
  };

  const parseAndExtract = async (qstr: string) => {
    let up = new URLSearchParams(qstr);
    const resp = await app.alt_next_first(up.get("code"), up.get("state"));
    if (resp.status !== 200) {
      console.log("Err", resp);
      return;
    }
    app.nav.goto_alt_first_stage(resp.data);
  };
</script>

<button
  class="p-2 text-white text-lg rounded bg-gray-600 flex justify-center"
  on:click={handleClick}
>
  <svg
    xmlns="http://www.w3.org/2000/svg"
    class="h-5 w-5"
    viewBox="0 0 20 20"
    fill="currentColor"
  >
    <path
      d="M11 3a1 1 0 100 2h2.586l-6.293 6.293a1 1 0 101.414 1.414L15 6.414V9a1 1 0 102 0V4a1 1 0 00-1-1h-5z"
    />
    <path
      d="M5 5a2 2 0 00-2 2v8a2 2 0 002 2h8a2 2 0 002-2v-3a1 1 0 10-2 0v3H5V7h3a1 1 0 000-2H5z"
    />
  </svg>

  Open {method["name"]}
</button>
