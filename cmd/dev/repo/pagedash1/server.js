function onload(opts) {
    core.log("@onload")
}

function onbuild(opts) {
    core.log("@onbuild")
}

function load_data1(params) {
    set_data_value("dyamic_data1", {value: 42})
}
