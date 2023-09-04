<script lang="ts">
    import { sleep } from "yootils";

    let message = "Waiting";
    let loading = true;

    const handleMessage = async (ev) => {
        console.log("@handle_message");

        const data = JSON.parse(ev.data);
        if (data["mtype"] === "exec_data") {
            window["__exec_data__"] = data["data"];

            const params = new URLSearchParams(location.search);
            const redirrect_url = params.get("redirrect_url");
            if (redirrect_url) {
                message = "Data received, Redirrecting";
                sleep(1000);
                window.location.pathname = redirrect_url;
            } else {
                message = "Data received, Go home ?";
            }

            loading = false;
        }
    };

    window.addEventListener(
        "message",
        (ev) => {
            if (ev.data !== "transfer_port") {
                console.log("wrong message", ev);
                return;
            }

            const port = ev.ports[0];
            port.onmessage = handleMessage;
            window["__parent_port__"] = port;

            port.postMessage(
                JSON.stringify({
                    mtype: "get_exec_data",
                })
            );

            message = "Waiting for data";
        },
        false
    );
</script>

<div class="flex items-center justify-center min-h-screen">
    <button
        class="bg-indigo-400 h-max w-max rounded-lg text-white font-bold hover:bg-indigo-300 duration-[500ms,800ms]{ loading ? 'hover:cursor-not-allowed': ''} "
        disabled={loading}
        on:click={() => {
            window.location.pathname = "/";
        }}
    >
        <div class="flex items-center justify-center m-[10px]">
            {#if loading}
                <div
                    class="h-5 w-5 border-t-transparent border-solid animate-spin rounded-full border-white border-4"
                />
            {/if}

            <div class="ml-2">
                {message}...
                <div />
            </div>
        </div></button
    >
</div>
