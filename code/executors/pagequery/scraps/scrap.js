(function () {
    core.Log("@total_amount")

    // consts
    const totalcolumnid = 2;
    const expectedcolid = 3;

    const sheet_id = get_execdata_item("sheet_id")
    const data_group = get_execdata_item("data_group")
    const source = get_execdata_item("source")

    // const rows = get_execdata_item("rows")


    if (sheet_id !== expectedcolid) {
        throw new Error("Wrong sheetid")
    }

    const [resp, err] = query_sheet({
        source,
        group: data_group,
        sheetid: sheet_id,
    })
    if (err) {
        throw new Error(`Query error: ${err}`)
    }

    // process here
    let acctotal = 0;
    (resp["cells"] || []).forEach((cell) => {
        if (cell["colid"] === totalcolumnid) {
            acctotal = acctotal + cell["numval"] || 0
        }
    })

    return [{
        data: { "total": acctotal },
        elements: {
            "Total": {
                "type": "dump",
                "source": "total"
            }
        }
    }, nil]
}());



(function () {
    core.log("This is test 111");

    return {
        data: { "total": { "hey": 100 } },
        elements: [
            {
                "name": "Total",
                "type": "dump",
                "source": "data/total"
            }
        ]
    }
})();



// /default/sheetsmhsj7/sheet/1


(function () {
    core.log("Calculating total amount.");

    if (get_execdata_item("sheet_name") !== "todo") {
        throw "wrong sheet"
    }

    const columnId = 5

    const [resp, err] = query_sheet({
        source: get_execdata_item("source"),
        group: get_execdata_item("data_group"),
        sheet_id: get_execdata_item("sheet_id"),
    })
    if (err) {
        throw err
    }

    let acctotal = 0;
    (resp["cells"] || []).forEach((cell) => {
        if (cell["colid"] === columnId) {
            acctotal = acctotal + Number(cell["numval"] || 0)  
        }
    })



    return {
        data: { "total": acctotal },
        elements: [
            {
                "name": "Total",
                "type": "dump",
                "source": "data/total"
            }
        ]
    }
})();