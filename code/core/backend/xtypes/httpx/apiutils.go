package httpx

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	MessageOk = []byte(`{"message": "success"}`)
)

func WriteJSON(c *gin.Context, resp interface{}, err error) {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if resp == nil {
		WriteOk(c)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func WriteFinal(c *gin.Context, err error) {
	if err != nil {
		WriteErr(c, err.Error())
		return
	}
	WriteOk(c)
}

func WriteErr(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{"message": msg})
}

func WriteOk(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(MessageOk)
}

func WriteBinary(c *gin.Context, data []byte) {
	w := c.Writer
	w.Write(data)
	w.Header().Del("Content-Type")
	w.Header().Add("Content-Type", "application/octet-stream")
}
