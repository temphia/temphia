(function () {
    console.log("noop executor loaded")
    const pf = window["__register_factory__"];
    if (!pf) {
        console.warn("factory registry func not found");
        return
    }

    pf("loader.factory", "noop.main", (ctx) => {
        console.log("@ctx", ctx)
    });

})()
