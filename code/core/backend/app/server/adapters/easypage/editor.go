package easypage

func (e *EasyPage) serveEditorFile(file string) ([]byte, error) {

	switch file {
	case "main.js":
		return e.dataBox.GetAsset("build", "adapter_editor_easypage.js")
	case "main.css":
		return e.dataBox.GetAsset("build", "adapter_editor_easypage.css")
	}

	return []byte(``), nil
}

func (e *EasyPage) preformEditorAction(name string, data []byte) (any, error) {
	return nil, nil
}
