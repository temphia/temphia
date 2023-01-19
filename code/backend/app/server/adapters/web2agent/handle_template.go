package web2agent

func (w *WATarget) serveTemplate(file string) {
	tpl := w.adapter.state.template
	if tpl == nil {
		return
	}

	tpl.Execute(w.http.Writer, map[string]any{})
}
