package apiadmin

/*

func (r *R) ListDtableSources(req request.Ctx) {
	sources, err := r.cBasic.ListDyndbSources(req.Session)
	r.rutil.WriteJSON(req.GinCtx, sources, err)
}

// dyn_table_group

func (r *R) NewGroup(req request.Ctx) {
	tg := &bprints.NewTableGroup{}
	err := req.GinCtx.BindJSON(tg)

	if err != nil {
		r.rutil.WriteErr(req.GinCtx, err.Error())
		return
	}
	err = r.cDtable.NewGroup(req.Session, req.GinCtx.Param("source"), tg)
	r.rutil.WriteFinal(req.GinCtx, err)

}
func (r *R) EditGroup(req request.Ctx) {
	tg := &entities.TableGroupPartial{}
	err := req.GinCtx.BindJSON(tg)
	if err != nil {
		r.rutil.WriteErr(req.GinCtx, err.Error())
		return
	}
	err = r.cDtable.EditGroup(req.Session, req.GinCtx.Param("source"), req.GinCtx.Param("group_id"), tg)
	r.rutil.WriteFinal(req.GinCtx, err)
}

func (r *R) GetGroup(req request.Ctx) {
	resp, err := r.cDtable.GetGroup(req.Session, req.GinCtx.Param("source"), req.GinCtx.Param("group_id"))
	r.rutil.WriteJSON(req.GinCtx, resp, err)
}

func (r *R) ListGroup(req request.Ctx) {
	gr, err := r.cDtable.ListGroup(req.Session, req.GinCtx.Param("source"))
	if err != nil {
		r.rutil.WriteErr(req.GinCtx, err.Error())
		return
	}
	r.rutil.WriteJSON(req.GinCtx, gr, err)
}

func (r *R) LoadGroup(req request.Ctx) {
	gr, err := r.cDtable.LoadGroup(req.Session)
	if err != nil {
		r.rutil.WriteErr(req.GinCtx, err.Error())
		return
	}

	r.rutil.WriteJSON(req.GinCtx, gr, err)
}

func (r *R) DeleteGroup(req request.Ctx) {
	err := r.cDtable.DeleteGroup(req.Session, req.GinCtx.Param("source"), req.GinCtx.Param("group_id"))
	r.rutil.WriteFinal(req.GinCtx, err)

}

// dyn_table

func (r *R) AddTable(req request.Ctx) {
	t := &bprints.NewTable{}
	err := req.GinCtx.BindJSON(t)
	if err != nil {
		r.rutil.WriteErr(req.GinCtx, err.Error())
		return
	}
	err = r.cDtable.AddTable(req.Session, t)
	r.rutil.WriteFinal(req.GinCtx, err)

}
func (r *R) EditTable(req request.Ctx) {
	tp := &entities.TablePartial{}
	err := req.GinCtx.BindJSON(tp)
	if err != nil {
		r.rutil.WriteErr(req.GinCtx, err.Error())
		return
	}
	err = r.cDtable.EditTable(req.Session, req.GinCtx.Param("table_id"), tp)
	r.rutil.WriteFinal(req.GinCtx, err)

}

func (r *R) GetTable(req request.Ctx) {
	tbl, err := r.cDtable.GetTable(req.Session, req.GinCtx.Param("table_id"))
	r.rutil.WriteJSON(req.GinCtx, tbl, err)
}

func (r *R) ListTables(req request.Ctx) {
	tbls, err := r.cDtable.ListTables(req.Session)
	r.rutil.WriteJSON(req.GinCtx, tbls, err)

}
func (r *R) DeleteTable(req request.Ctx) {
	err := r.cDtable.DeleteTable(req.Session, req.GinCtx.Param("table_id"))
	r.rutil.WriteFinal(req.GinCtx, err)

}

// dyn_table_column

func (r *R) AddColumn(req request.Ctx) {
	nc := &bprints.NewColumn{}
	err := req.GinCtx.BindJSON(nc)
	if err != nil {
		r.rutil.WriteErr(req.GinCtx, err.Error())
		return
	}

	err = r.cDtable.AddColumn(req.Session, req.GinCtx.Param("table_id"), nc)
	r.rutil.WriteFinal(req.GinCtx, err)

}
func (r *R) EditColumn(req request.Ctx) {
	cp := &entities.ColumnPartial{}
	err := req.GinCtx.BindJSON(cp)
	if err != nil {
		r.rutil.WriteErr(req.GinCtx, err.Error())
		return
	}
	err = r.cDtable.EditColumn(req.Session, req.GinCtx.Param("table_id"), req.GinCtx.Param("column_id"), cp)
	r.rutil.WriteFinal(req.GinCtx, err)

}

func (r *R) GetColumn(req request.Ctx) {
	resp, err := r.cDtable.GetColumn(req.Session, req.GinCtx.Param("table_id"), req.GinCtx.Param("column_id"))
	r.rutil.WriteJSON(req.GinCtx, resp, err)
}

func (r *R) ListColumns(req request.Ctx) {
	cols, err := r.cDtable.ListColumns(req.Session, req.GinCtx.Param("table_id"))
	r.rutil.WriteJSON(req.GinCtx, cols, err)

}
func (r *R) DeleteColumn(req request.Ctx) {

	err := r.cDtable.DeleteColumn(req.Session, req.GinCtx.Param("table_id"), req.GinCtx.Param("column_id"))
	r.rutil.WriteFinal(req.GinCtx, err)

}


*/
