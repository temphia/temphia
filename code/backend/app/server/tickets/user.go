package tickets

import (
	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/backend/xtypes/httpx"
	"github.com/temphia/temphia/code/backend/xtypes/models/claim"
	"github.com/temphia/temphia/code/backend/xtypes/models/entities"
)

func (a *TicketAPI) User(rg *gin.RouterGroup) {

	rg.GET("/", a.middleware.UX(a.ugListUser))
	rg.POST("/", a.middleware.UX(a.ugAddUser))
	rg.GET("/:user_id", a.middleware.UX(a.ugGetUser))
	rg.POST("/:user_id", a.middleware.UX(a.ugUpdateUser))
	rg.DELETE("/:user_id", a.middleware.UX(a.ugDeleteUser))
}

func (a *TicketAPI) ugListUser(uclaim *claim.UserMgmtTkt, http *gin.Context) {
	resp, err := a.cAdmin.UgroupListUsersByGroup(uclaim)

	httpx.WriteJSON(http, resp, err)
}

func (a *TicketAPI) ugAddUser(uclaim *claim.UserMgmtTkt, http *gin.Context) {
	user := &entities.User{}
	err := http.BindJSON(user)
	if err != nil {
		httpx.WriteErr(http, err)
		return
	}

	err = a.cAdmin.UgroupAddUser(uclaim, user)
	httpx.WriteFinal(http, err)
}

func (a *TicketAPI) ugGetUser(uclaim *claim.UserMgmtTkt, http *gin.Context) {
	resp, err := a.cAdmin.UgroupGetUserByID(uclaim, http.Param("user_id"))
	httpx.WriteJSON(http, resp, err)
}

func (a *TicketAPI) ugUpdateUser(uclaim *claim.UserMgmtTkt, http *gin.Context) {
	data := make(map[string]any)
	err := http.BindJSON(&data)
	if err != nil {
		httpx.WriteErr(http, err)
		return
	}

	err = a.cAdmin.UgroupUpdateUser(uclaim, http.Param("user_id"), data)
	httpx.WriteFinal(http, err)
}

func (a *TicketAPI) ugDeleteUser(uclaim *claim.UserMgmtTkt, http *gin.Context) {
	err := a.cAdmin.UgroupDeleteUser(uclaim, http.Param("user_id"))
	httpx.WriteFinal(http, err)
}
