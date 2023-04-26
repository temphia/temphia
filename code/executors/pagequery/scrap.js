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
        core.Log("@wrong_sheetid")
        return
    }

    const [resp, err] = query_sheet({
        source,
        group: data_group,
        sheetid: sheet_id,
    })
    if (err != nil) {
        core.Log("@query_error" + err)
        return [nil, nil]
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
                "type": "paragraph",
                "source": "total"
            }
        }
    }, nil]
}());
