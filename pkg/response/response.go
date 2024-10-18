package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	ResponseCode    string      `json:"responseCode"`
	ResponseMessage string      `json:"responseMessage"`
	Data            interface{} `json:"data,omitempty"`
}

func JSON(c *gin.Context, statusCode int, responseCode string, responseMessage string, data interface{}) {
	res := Response{
		ResponseCode:    responseCode,
		ResponseMessage: responseMessage,
		Data:            data,
	}
	c.JSON(statusCode, res)
}

func SuccessResponse(c *gin.Context, data interface{}) {
	JSON(c, http.StatusOK, "200", "Success", data)
}

func ErrorResponse(c *gin.Context, responseCode, responseMessage string) {
	JSON(c, http.StatusBadRequest, responseCode, responseMessage, nil)
}
