
window.__register_factory__("plug.factory", "a.main", (ctx) => {
    console.log("@a.main", ctx)
})

window.__register_factory__("plug.factory", "b.main", (ctx) => {
    console.log("@b.main", ctx)
})