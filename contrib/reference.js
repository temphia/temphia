function test(params) {
    const invoker = temphia.invoker()
    if (invoker !== "target_socket") {
        throw "this agent is expected to be executed invoked by target_socket"
    }

    const [lockmod, err] = temphia.NewModule("xplane_lock", {
        "nested": false,
    })
    if (err) {
        return `could not create module lock: ${err}`
    }

    const [_, err1] = lockmod.Execute("grab", { "timeout": "10m" });
    if (err) {
        return `could not grab lock: ${err1}`
    }


    const [wsmod, err2] = temphia.NewModule("wspool", { "url": "wss://mno.xyz/ws" })
    if (err2) {
        throw `could not create wspool ${err2}`
    }

    while (true) {
        const [messages, err3] = wsmod.Execute("poll", {
            "max_time_out": "10s",
        })

        if (err3) {
            temphia.Log(err3)
            continue
        }

        messages.forEach((msg) => {
            const perr = process_message(msg);
            if (perr) {
                temphia.LazyLog(perr)
            }
        })

        temphia.SyncLog()
    }
}