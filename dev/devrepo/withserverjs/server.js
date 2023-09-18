const rePort = Bun.env["TEMPHIA_RE_PORT"]
const reToken = Bun.env["TEMPHIA_RE_TOKEN"]

console.log("@starting_server", {
    rePort,
    reToken
})

const handleRPX = (req) => (new Response("Hello from RPX!"))

Bun.serve({
    port: 8080,
    fetch(req) {

        if (url.pathname.startsWith("/rpx")) {
            return handleRPX(req)
        }

        return new Response("Hello");
    },
});