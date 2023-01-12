function list_items() {
    const [resp, err] = plugkv.query({
        tag1s: ["todoitems"],
    });
    if (err) {
        if (utils.is_db_not_found(err)) {
            core.log("db not found");
            return utils.ok_response([]);
        }
        return utils.err_response(err);
    }

    core.log("@====>" + JSON.stringify(resp));

    const fresp = (resp || []).map((elem) => {
        core.log("@what the hell happended" + typeof elem.value);
        let pval = {};
        try {
            pval = JSON.parse(elem.value);
        } catch (e) {
            core.log("@err parsing" + e + "<=====>" + typeof elem.value + elem.value);
        }

        return { id: elem.key, value: pval };
    });

    core.log("@====> BEFORE RETURNING RESP");
    return utils.ok_response(fresp);
}

function action_list_items(params) {
    core.log("LIST_ITEMS " + Math.random());
    return list_items();
}

function action_add_item(params) {
    const data = utils.ab2str(params["data"]);
    const id = utils.generate_str_id();

    const err = plugkv.set(id, data, {
        tag1: "todoitems",
    });
    if (err) {
        core.log("LIST ERR =>" + err);
        return utils.err_response(err);
    }

    return list_items();
}

function action_update_item(params) { }
