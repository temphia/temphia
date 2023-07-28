package apiadmin

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/controllers/admin"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

func (a *ApiAdmin) userAPI(rg *gin.RouterGroup) {

	rg.GET("/", a.X(a.ListUsers))
	rg.POST("/", a.X(a.AddUser))
	rg.GET("/:user_id", a.X(a.GetUserByID))
	rg.POST("/:user_id", a.X(a.UpdateUser))
	rg.DELETE("/:user_id", a.X(a.RemoveUser))

	rg.GET("/:user_id/device", a.X(a.ListUserDevices))
	rg.GET("/:user_id/device/:id", a.X(a.GetUserDevice))
	rg.POST("/:user_id/device/:id", a.X(a.UpdateUserDevice))
	rg.DELETE("/:user_id/device/:id", a.X(a.RemoveUserDevices))
	rg.POST("/:user_id/device", a.X(a.AddUserDevices))

}

func (r *ApiAdmin) AddUser(ctx httpx.Request) {
	usr := &entities.User{}

	err := ctx.Http.BindJSON(usr)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}
	usr.TenantID = ctx.Session.TenantId

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

func (r *ApiAdmin) ListUserDevices(ctx httpx.Request) {
	resp, err := r.cAdmin.ListUserDevice(ctx.Session, ctx.MustParam("user_id"))
	r.rutil.WriteJSON(ctx.Http, resp, err)
}

func (r *ApiAdmin) GetUserDevice(ctx httpx.Request) {
	resp, err := r.cAdmin.GetUserDevice(ctx.Session, ctx.MustParam("user_id"), ctx.MustParamInt("id"))
	r.rutil.WriteJSON(ctx.Http, resp, err)
}

func (r *ApiAdmin) UpdateUserDevice(ctx httpx.Request) {
	data := make(map[string]interface{})
	err := ctx.Http.BindJSON(&data)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	err = r.cAdmin.UpdateUserDevice(ctx.Session, ctx.MustParam("user_id"), ctx.MustParamInt("id"), data)
	r.rutil.WriteJSON(ctx.Http, nil, err)
}

func (r *ApiAdmin) RemoveUserDevices(ctx httpx.Request) {
	err := r.cAdmin.RemoveUserDevice(ctx.Session, ctx.MustParam("user_id"), ctx.MustParamInt("id"))
	r.rutil.WriteJSON(ctx.Http, nil, err)

}

func (r *ApiAdmin) AddUserDevices(ctx httpx.Request) {

	data := admin.NewUserDevice{}
	err := ctx.Http.BindJSON(&data)
	if err != nil {
		r.rutil.WriteErr(ctx.Http, err.Error())
		return
	}

	err = r.cAdmin.AddUserDevice(ctx.Session, ctx.MustParam("user_id"), &data)
	r.rutil.WriteJSON(ctx.Http, nil, err)
}
