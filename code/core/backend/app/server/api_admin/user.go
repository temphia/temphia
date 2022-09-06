package apiadmin

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
)

func (a *ApiAdmin) userAPI(rg *gin.RouterGroup) {

}

func (r *ApiAdmin) AddUser(ctx httpx.Request) {
	usr := &entities.User{}

	err := ctx.Http.BindJSON(usr)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}
	usr.TenantID = ctx.Session.TenentId

	r.rutil.WriteJSON(ctx.Http, nil, r.cAdmin.AddUser(ctx.Session, usr))
}

func (r *ApiAdmin) UpdateUser(ctx httpx.Request) {
	data := make(map[string]interface{})

	err := ctx.Http.BindJSON(&data)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	r.rutil.WriteJSON(ctx.Http, nil, r.cAdmin.UpdateUser(ctx.Session, data))
}

func (r *ApiAdmin) RemoveUser(ctx httpx.Request) {
	r.rutil.WriteJSON(ctx.Http, nil, r.cAdmin.RemoveUser(ctx.Session, ctx.Http.Param("user_id")))
}

func (r *ApiAdmin) GetUserByID(ctx httpx.Request) {
	usr, err := r.cAdmin.GetUserByID(ctx.Session, ctx.Http.Param("user_id"))
	if usr != nil {
		usr.Password = ""
	}

	r.rutil.WriteJSON(ctx.Http, usr, err)
}

func (r *ApiAdmin) ListUsers(ctx httpx.Request) {
	ugroup := ctx.Http.Query("user_group")
	var usrs []*entities.User
	var err error

	if ugroup != "" {
		usrs, err = r.cAdmin.ListUsersByGroup(ctx.Session, ugroup)
	} else {
		usrs, err = r.cAdmin.ListUsers(ctx.Session)
	}

	r.rutil.WriteJSON(ctx.Http, usrs, err)
}

// perm placeholder stuff

func (r *ApiAdmin) AddPerm(ctx httpx.Request) {

}
func (r *ApiAdmin) UpdatePerm(ctx httpx.Request) {

}
func (r *ApiAdmin) GetPerm(ctx httpx.Request) {

}
func (r *ApiAdmin) RemovePerm(ctx httpx.Request) {

}
func (r *ApiAdmin) AddRole(ctx httpx.Request) {

}
func (r *ApiAdmin) GetRole(ctx httpx.Request) {

}
func (r *ApiAdmin) UpdateRole(ctx httpx.Request) {

}
func (r *ApiAdmin) RemoveRole(ctx httpx.Request) {

}
func (r *ApiAdmin) AddUserRole(ctx httpx.Request) {

}
func (r *ApiAdmin) RemoveUserRole(ctx httpx.Request) {

}
func (r *ApiAdmin) ListAllPerm(ctx httpx.Request) {

}
func (r *ApiAdmin) ListAllRole(ctx httpx.Request) {

}
func (r *ApiAdmin) ListAllUserRole(ctx httpx.Request) {

}
func (r *ApiAdmin) ListAllUserPerm(ctx httpx.Request) {

}
func (r *ApiAdmin) ListUserPerm(ctx httpx.Request) {

}
