function OnLoadForm(mode, stage) {
    core.log("load_form " + mode + " " + stage)
}


function OnStartSubmit(mode, stage) {
    core.log("start_submit " + mode + " " + stage)
}

function OnStartGenerate(mode, stage) {
    core.log("start_generate " + mode + " " + stage)
}


function OnSecondGenerate(mode, stage) {
    core.log("@second_generate")
    set_message("This could be great", true)
}

 