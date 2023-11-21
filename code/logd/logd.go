package logd

import (
	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
	"github.com/temphia/temphia/code/backend/app/log/lreader"
	"github.com/temphia/temphia/code/backend/libx/xutils/kosher"
	"github.com/temphia/temphia/code/backend/xtypes/logx"
	"github.com/temphia/temphia/code/backend/xtypes/xserver/xnotz/httpx"
)

type Logd struct {
	listenpath string
	logpath    string
}

func New(lispath, logpath string) *Logd {
	return &Logd{
		listenpath: lispath,
		logpath:    logpath,
	}

}

func (l *Logd) Start() {

	sl := lreader.New(l.logpath)

	e := gin.New()

	e.GET("/log-query", func(ctx *gin.Context) {

		req := logx.QueryRequest{}
		err := ctx.BindJSON(&req)
		if err != nil {
			httpx.WriteErr(ctx, err)
			return
		}

		resp, err := sl.Query(ctx.Query("tenant_id"), req)
		if err != nil {
			httpx.WriteErr(ctx, err)
			return
		}

		for _, line := range resp {
			_, err = ctx.Writer.Write(kosher.Byte(string(line)))
			if err != nil {
				pp.Println("@err_log_query_write", err.Error())
				return
			}
		}

	})

	e.RunUnix(l.listenpath)

}
