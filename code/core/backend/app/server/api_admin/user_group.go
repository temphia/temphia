package apiadmin

import (
	"strconv"

	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
)

func (r *ApiAdmin) AddUserGroup(ctx httpx.Request) {
	group := &entities.UserGroup{}

	err := ctx.Http.BindJSON(group)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	group.TenantID = ctx.Session.TenentId

	r.rutil.WriteJSON(ctx.Http, nil, r.cAdmin.AddUserGroup(ctx.Session, group))
}

func (r *ApiAdmin) GetUserGroup(ctx httpx.Request) {
	resp, err := r.cAdmin.GetUserGroup(ctx.Session, ctx.Http.Param("user_group"))
	r.rutil.WriteJSON(ctx.Http, resp, err)
}

func (r *ApiAdmin) ListUserGroup(ctx httpx.Request) {
	resp, err := r.cAdmin.ListUserGroup(ctx.Session)
	r.rutil.WriteJSON(ctx.Http, resp, err)
}

func (r *ApiAdmin) UpdateUserGroup(ctx httpx.Request) {
	data := make(map[string]interface{})
	err := ctx.Http.BindJSON(&data)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	err = r.cAdmin.UpdateUserGroup(ctx.Session, ctx.Http.Param("user_group"), data)
	r.rutil.WriteJSON(ctx.Http, nil, err)
}

func (r *ApiAdmin) RemoveUserGroup(ctx httpx.Request) {
	r.rutil.WriteJSON(ctx.Http, nil, r.cAdmin.RemoveUserGroup(ctx.Session, ctx.Http.Param("user_group")))
}

// auth user group meta

func (r *ApiAdmin) AddUserGroupAuth(ctx httpx.Request) {
	data := &entities.UserGroupAuth{}
	err := ctx.Http.BindJSON(data)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	r.rutil.WriteJSON(
		ctx.Http,
		nil,
		r.cAdmin.AddUserGroupAuth(ctx.Session, ctx.Http.Param("ugroup"), data),
	)
}

func (r *ApiAdmin) UpdateUserGroupAuth(ctx httpx.Request) {
	data := make(map[string]interface{})
	err := ctx.Http.BindJSON(&data)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	id, err := strconv.ParseInt(ctx.Http.Param("id"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	r.rutil.WriteJSON(
		ctx.Http,
		nil,
		r.cAdmin.UpdateUserGroupAuth(ctx.Session, ctx.Http.Param("ugroup"), id, data),
	)
}

func (r *ApiAdmin) ListUserGroupAuth(ctx httpx.Request) {
	data, err := r.cAdmin.ListUserGroupAuth(ctx.Session, ctx.Http.Param("ugroup"))
	r.rutil.WriteJSON(ctx.Http, data, err)
}

func (r *ApiAdmin) GetUserGroupAuth(ctx httpx.Request) {

	id, err := strconv.ParseInt(ctx.Http.Param("id"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	resp, err := r.cAdmin.GetUserGroupAuth(ctx.Session, ctx.Http.Param("ugroup"), id)
	r.rutil.WriteJSON(ctx.Http, resp, err)
}
func (r *ApiAdmin) RemoveUserGroupAuth(ctx httpx.Request) {
	id, err := strconv.ParseInt(ctx.Http.Param("id"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	err = r.cAdmin.RemoveUserGroupAuth(ctx.Session, ctx.Http.Param("ugroup"), id)
	r.rutil.WriteJSON(ctx.Http, nil, err)
}

// hook

func (r *ApiAdmin) AddUserGroupHook(ctx httpx.Request) {
	data := &entities.UserGroupHook{}
	err := ctx.Http.BindJSON(data)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	r.rutil.WriteJSON(
		ctx.Http,
		nil,
		r.cAdmin.AddUserGroupHook(ctx.Session, ctx.Http.Param("ugroup"), data),
	)
}

func (r *ApiAdmin) UpdateUserGroupHook(ctx httpx.Request) {
	data := make(map[string]interface{})
	err := ctx.Http.BindJSON(&data)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	id, err := strconv.ParseInt(ctx.Http.Param("id"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	r.rutil.WriteJSON(
		ctx.Http,
		nil,
		r.cAdmin.UpdateUserGroupHook(ctx.Session, ctx.Http.Param("ugroup"), id, data),
	)
}

func (r *ApiAdmin) ListUserGroupHook(ctx httpx.Request) {
	resp, err := r.cAdmin.ListUserGroupHook(ctx.Session, ctx.Http.Param("ugroup"))
	r.rutil.WriteJSON(ctx.Http, resp, err)
}

func (r *ApiAdmin) GetUserGroupHook(ctx httpx.Request) {
	id, err := strconv.ParseInt(ctx.Http.Param("id"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	resp, err := r.cAdmin.GetUserGroupHook(ctx.Session, ctx.Http.Param("ugroup"), id)
	r.rutil.WriteJSON(ctx.Http, resp, err)
}

func (r *ApiAdmin) RemoveUserGroupHook(ctx httpx.Request) {
	id, err := strconv.ParseInt(ctx.Http.Param("id"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	r.cAdmin.RemoveUserGroupHook(ctx.Session, ctx.Http.Param("ugroup"), id)
}

// plug

func (r *ApiAdmin) AddUserGroupPlug(ctx httpx.Request) {
	data := &entities.UserGroupPlug{}
	err := ctx.Http.BindJSON(data)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	r.rutil.WriteJSON(
		ctx.Http,
		nil,
		r.cAdmin.AddUserGroupPlug(ctx.Session, ctx.Http.Param("ugroup"), data),
	)
}

func (r *ApiAdmin) UpdateUserGroupPlug(ctx httpx.Request) {
	data := make(map[string]interface{})
	err := ctx.Http.BindJSON(&data)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	id, err := strconv.ParseInt(ctx.Http.Param("id"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	err = r.cAdmin.UpdateUserGroupPlug(ctx.Session, ctx.Http.Param("ugroup"), id, data)
	r.rutil.WriteJSON(ctx.Http, nil, err)
}

func (r *ApiAdmin) ListUserGroupPlug(ctx httpx.Request) {
	resp, err := r.cAdmin.ListUserGroupPlug(ctx.Session, ctx.Http.Param("ugroup"))
	r.rutil.WriteJSON(ctx.Http, resp, err)
}

func (r *ApiAdmin) GetUserGroupPlug(ctx httpx.Request) {
	id, err := strconv.ParseInt(ctx.Http.Param("id"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	resp, err := r.cAdmin.GetUserGroupPlug(ctx.Session, ctx.Http.Param("ugroup"), id)
	r.rutil.WriteJSON(ctx.Http, resp, err)
}

func (r *ApiAdmin) RemoveUserGroupPlug(ctx httpx.Request) {
	id, err := strconv.ParseInt(ctx.Http.Param("id"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	r.cAdmin.RemoveUserGroupPlug(ctx.Session, ctx.Http.Param("ugroup"), id)
}

// data

func (r *ApiAdmin) AddUserGroupData(ctx httpx.Request) {
	data := &entities.UserGroupData{}
	err := ctx.Http.BindJSON(data)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	r.rutil.WriteJSON(
		ctx.Http,
		nil,
		r.cAdmin.AddUserGroupData(ctx.Session, ctx.Http.Param("ugroup"), data),
	)
}

func (r *ApiAdmin) UpdateUserGroupData(ctx httpx.Request) {
	data := make(map[string]interface{})
	err := ctx.Http.BindJSON(&data)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	id, err := strconv.ParseInt(ctx.Http.Param("id"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	err = r.cAdmin.UpdateUserGroupData(ctx.Session, ctx.Http.Param("ugroup"), id, data)
	r.rutil.WriteJSON(ctx.Http, nil, err)
}

func (r *ApiAdmin) ListUserGroupData(ctx httpx.Request) {
	resp, err := r.cAdmin.ListUserGroupData(ctx.Session, ctx.Http.Param("ugroup"))
	r.rutil.WriteJSON(ctx.Http, resp, err)
}

func (r *ApiAdmin) GetUserGroupData(ctx httpx.Request) {
	id, err := strconv.ParseInt(ctx.Http.Param("id"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	resp, err := r.cAdmin.GetUserGroupData(ctx.Session, ctx.Http.Param("ugroup"), id)
	r.rutil.WriteJSON(ctx.Http, resp, err)
}

func (r *ApiAdmin) RemoveUserGroupData(ctx httpx.Request) {
	id, err := strconv.ParseInt(ctx.Http.Param("id"), 10, 64)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	r.cAdmin.RemoveUserGroupData(ctx.Session, ctx.Http.Param("ugroup"), id)
}
