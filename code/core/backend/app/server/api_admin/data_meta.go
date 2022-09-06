package apiadmin

/*


// dyn_table_meta

func (r *R) AddIndex(req request.Ctx) {

}


func (r *R) AddUniqueIndex(req request.Ctx) {

}
func (r *R) AddFTSIndex(req request.Ctx) {

}
func (r *R) AddColumnFRef(req request.Ctx) {

}
func (r *R) ListIndex(req request.Ctx) {

}
func (r *R) RemoveIndex(req request.Ctx) {

}


// view stuff

func (r *R) NewView(req request.Ctx) {
	view := entities.DataView{}
	err := req.GinCtx.BindJSON(&view)
	if err != nil {
		r.rutil.WriteErr(req.GinCtx, err.Error())
		return
	}
	err = r.cDtable.NewView(req.Session, req.GinCtx.Param("table_id"), &view)
	r.rutil.WriteFinal(req.GinCtx, err)
}

func (r *R) ModifyView(req request.Ctx) {
	view := make(map[string]interface{})

	err := req.GinCtx.BindJSON(&view)
	if err != nil {
		r.rutil.WriteErr(req.GinCtx, err.Error())
		return
	}

	id, err := strconv.ParseInt(req.GinCtx.Param("id"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(req.GinCtx, err.Error())
		return
	}

	err = r.cDtable.ModifyView(req.Session, req.GinCtx.Param("table_id"), id, view)
	r.rutil.WriteFinal(req.GinCtx, err)

}

func (r *R) GetView(req request.Ctx) {
	id, err := strconv.ParseInt(req.GinCtx.Param("id"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(req.GinCtx, err.Error())
		return
	}

	resp, err := r.cDtable.GetView(req.Session, req.GinCtx.Param("table_id"), id)
	r.rutil.WriteJSON(req.GinCtx, resp, err)
}

func (r *R) ListView(req request.Ctx) {
	resp, err := r.cDtable.ListView(req.Session, req.GinCtx.Param("table_id"))
	if err != nil {
		r.rutil.WriteErr(req.GinCtx, err.Error())
		return
	}

	r.rutil.WriteJSON(req.GinCtx, resp, err)
}

func (r *R) DelView(req request.Ctx) {
	id, err := strconv.ParseInt(req.GinCtx.Param("id"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(req.GinCtx, err.Error())
		return
	}

	err = r.cDtable.DelView(req.Session, req.GinCtx.Param("table_id"), id)
	r.rutil.WriteFinal(req.GinCtx, err)
}

// hooks

func (r *R) NewHook(req request.Ctx) {
	hook := entities.DataHook{}
	err := req.GinCtx.BindJSON(&hook)
	if err != nil {
		r.rutil.WriteErr(req.GinCtx, err.Error())
		return
	}
	err = r.cDtable.NewHook(req.Session, req.GinCtx.Param("table_id"), &hook)
	r.rutil.WriteFinal(req.GinCtx, err)
}

func (r *R) ModifyHook(req request.Ctx) {
	data := make(map[string]interface{})

	err := req.GinCtx.BindJSON(&data)
	if err != nil {
		r.rutil.WriteErr(req.GinCtx, err.Error())
		return
	}

	id, err := strconv.ParseInt(req.GinCtx.Param("id"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(req.GinCtx, err.Error())
		return
	}

	err = r.cDtable.ModifyHook(req.Session, req.GinCtx.Param("table_id"), id, data)
	r.rutil.WriteFinal(req.GinCtx, err)

}

func (r *R) GetHook(req request.Ctx) {
	id, err := strconv.ParseInt(req.GinCtx.Param("id"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(req.GinCtx, err.Error())
		return
	}

	resp, err := r.cDtable.GetHook(req.Session, req.GinCtx.Param("table_id"), id)
	r.rutil.WriteJSON(req.GinCtx, resp, err)

}

func (r *R) ListHook(req request.Ctx) {
	resp, err := r.cDtable.ListHook(req.Session, req.GinCtx.Param("table_id"))
	if err != nil {
		r.rutil.WriteErr(req.GinCtx, err.Error())
		return
	}

	r.rutil.WriteJSON(req.GinCtx, resp, err)
}

func (r *R) DelHook(req request.Ctx) {
	id, err := strconv.ParseInt(req.GinCtx.Param("id"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(req.GinCtx, err.Error())
		return
	}

	err = r.cDtable.DelHook(req.Session, req.GinCtx.Param("table_id"), id)
	r.rutil.WriteFinal(req.GinCtx, err)
}


*/
