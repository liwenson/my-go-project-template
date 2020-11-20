package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Ret struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type ResponseMsg struct {
	Ret Ret `json:"ret"`
}

type ResponseData struct {
	Ret  Ret         `json:"ret"`
	Data interface{} `json:"data"`
}

func OkData(c *gin.Context, ret Ret, data interface{}) {
	//res := response{
	//	Ret: ret,
	//	Data: map[string]interface{}{},
	//}
	if data == nil {
		c.JSON(http.StatusOK, ResponseMsg{
			Ret: ret,
		})
		return
	} else {
		c.JSON(http.StatusOK, ResponseData{
			Ret:  ret,
			Data: data,
		})
		return
	}
}
