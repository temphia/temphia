package httpx

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func ReadForm(ctx *gin.Context) ([]byte, error) {
	fh, err := ctx.FormFile("file")
	if err != nil {
		return nil, err
	}

	file, err := fh.Open()
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(file)
}
