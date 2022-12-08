console.log("registering todo.helloworld.main ....")


window.__register_factory__("plug.factory", "todo.helloworld.main", (ctx) => {
    console.log("FACTORY CALLED WITH CTX", ctx)

    const Qs = (qs) => ctx.target.querySelector(qs);
    const api = new API(ctx.env);

    let list = [{ id: "xyz", value: { done: false, data: "buy milk" } }]


    ctx.target.innerHTML = `
        <h1>Todo App</h1>
        <div class="list">
        </div>
        <div>
            <textarea class="new-text"> </textarea>
            <button class="add-btn">Add</button>
        </div>
        <button class="refresh">Refresh</button>
    `



    const listElem = Qs("div .list")
    const refreshElem = Qs(".refresh")
    const textElem = Qs("div .new-text")
    const addElem = Qs("div .add-btn")


    const redraw = () => {
        const ul = document.createElement("ul")
        list.forEach((val) => {
            const li = document.createElement("li")
            li.innerHTML = `<input type="checkbox"></input> ${val.value.data}`
            ul.appendChild(li)
        })
        listElem.replaceChildren(ul)
    }

    const refresh = async () => {
        const rlist = await api.list_items()
        if (rlist) {
            list = rlist
        }
        redraw()
    }

    const add = async () => {
        const rlist = await api.add_item(textElem.value)
        if (rlist) {
            list = rlist
        }
        redraw()
    }

    refreshElem.addEventListener("click", refresh)
    addElem.addEventListener("click", add)
    refresh()
})


class API {
    constructor(env) {
        this.env = env
    }

    list_items = async () => {
        const resp = await this.env.PreformAction("list_items", {})
        return resp.data.data;
    }

    add_item = async (item) => {
        const resp = await this.env.PreformAction("add_item", {
            done: false,
            data: item
        })
        return resp.data.data;
    }

}