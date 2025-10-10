package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendOKResponse(c *gin.Context, msg string, data any) {
	sendResponse(c, http.StatusOK, msg, data, nil)
}

func SendCreatedResponse(c *gin.Context, id, resource string) {
	msg := fmt.Sprintf("%s created successfully", resource)

	sendResponse(
		c,
		http.StatusCreated,
		msg,
		gin.H{"id": id},
		nil,
	)
}

func SendBadRequestResponse(c *gin.Context, msg string) {
	sendResponse(c, http.StatusBadRequest, "", nil, msg)
}

func SendNotFoundResponse(c *gin.Context, msg string) {
	sendResponse(c, http.StatusNotFound, "", nil, msg)
}

func SendServiceUnavailableResponse(c *gin.Context, msg string) {
	sendResponse(c, http.StatusServiceUnavailable, "", nil, msg)
}

func SendConflictResponse(c *gin.Context, msg string) {
	sendResponse(c, http.StatusConflict, "", nil, msg)
}

func SendInternalServerErrorResponse(c *gin.Context) {
	sendResponse(c, http.StatusInternalServerError, "", nil, "Internal Server Error")
}

func SendEmptyResponse(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func sendResponse(c *gin.Context, status int, msg string, data any, err any) {
	isRespEmpty := true

	resp := gin.H{}

	if msg != "" {
		isRespEmpty = false
		resp["message"] = msg
	}

	if data != nil {
		isRespEmpty = false
		resp["data"] = data
	}

	if err != nil {
		isRespEmpty = false
		resp["error"] = err
	}

	if isRespEmpty {
		c.Status(status)
		return
	}

	c.JSON(status, resp)
}
