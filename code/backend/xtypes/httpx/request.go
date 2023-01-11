package httpx

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/temphia/temphia/code/core/backend/xtypes/models/claim"
)

type Request struct {
	Id      int64
	Http    *gin.Context
	Session *claim.Session
}

func (r Request) MustParamInt(name string) int64 {
	i, err := strconv.ParseInt(r.Http.Param(name), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("could not parse param %s", name))
	}
	return i
}

func (r Request) MustParam(name string) string {
	val := r.Http.Param(name)
	if val == "" {
		panic(fmt.Sprintf("could not extract param %s", name))
	}

	return val
}

func (r Request) MustQueryInt(name string) int64 {
	i, err := strconv.ParseInt(r.Http.Query(name), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("could not parse param %s", name))
	}
	return i
}

func (r Request) MustQuery(name string) string {
	val := r.Http.Query(name)
	if val == "" {
		panic(fmt.Sprintf("could not extract query %s", name))
	}

	return val
}
