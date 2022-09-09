package apiadmin

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/entities"
)

func (a *ApiAdmin) userAPI(rg *gin.RouterGroup) {

	rg.GET("/", a.X(a.ListUsers))
	rg.POST("/", a.X(a.AddUser))
	rg.GET("/:user_id", a.X(a.GetUserByID))
	rg.POST("/:user_id", a.X(a.UpdateUser))
	rg.DELETE("/:user_id", a.X(a.RemoveUser))

	rg.GET("/perm", a.X(a.ListAllPerm))
	rg.POST("/perm", a.X(a.AddPerm))
	rg.GET("/perm/:perm", a.X(a.GetPerm))
	rg.POST("/perm/:perm", a.X(a.UpdatePerm))
	rg.DELETE("/perm/:perm", a.X(a.RemovePerm))

	rg.GET("/role", a.X(a.ListAllRole))
	rg.POST("/role", a.X(a.AddRole))
	rg.GET("/role/:role", a.X(a.GetRole))
	rg.POST("/role/:role", a.X(a.UpdateRole))
	rg.DELETE("/role/:role", a.X(a.RemoveRole))

	rg.GET("/user_role", a.X(a.ListAllUserRole))
	rg.POST("/user_role", a.X(a.AddUserRole))
	rg.DELETE("/user_role", a.X(a.RemoveUserRole))

	// rg.GET("/user_perm", r.Authed(r.ListUserPerm)) // user query
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
