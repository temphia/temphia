console.log("registering counter.helloworld.main ....")

window.__register_factory__("plug.factory", "counter.helloworld.main", (ctx) => {
    console.log("FACTORY CALLED WITH CTX", ctx)

    let value = 0

    ctx.target.innerHTML = `
        <h1>Counter is <span>${value}<span></h1>
        <button class="add-btn">Add</button>    
        <button class="sub-btn">Subtract</button>
    `


    const Qs = (qs) => ctx.target.querySelector(qs);
    const counter = Qs("h1 span")
    const add = Qs(".add-btn")
    const sub = Qs(".sub-btn")

    const redraw = () => {
        counter.innerHTML = value + "";
    }

    add.addEventListener("click", () => {
        value = value + 1;
        redraw()
    })

    sub.addEventListener("click", () => {
        value = value - 1;
        redraw()
    })

})