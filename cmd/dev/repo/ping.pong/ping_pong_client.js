console.log("registering ping.pong.main ....")

window.__register_factory__("plug.factory", "ping.pong.main", (ctx) => {
    console.log("FACTORY CALLED WITH CTX", ctx)

    let value = 0

    ctx.target.innerHTML = `
    <label>name</label>
    <input id="name-input" />
    <button id="ping-button">Ping</button>
    <p id="message"></p>
    `

    const nameInputElem = ctx.target.getElementById("name-input");
    const pingButtonElem = ctx.target.getElementById("ping-button");
    const messageElem = ctx.target.getElementById("message")
    
    pingButtonElem.addEventListener("click", (evt) => {
        console.log("@value", nameInputElem.value);
        console.log("@event", evt);
        messageElem.innerText = `fixme`; 
    })
})