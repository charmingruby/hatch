package rest

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendOKResponse(c *gin.Context, msg string, data any) {
	sendResponse(c, http.StatusOK, msg, data, nil)
}

func SendCreatedResponse(c *gin.Context, msg, id, resource string) {
	if msg == "" {
		msg = fmt.Sprintf("%s created successfully", resource)
	}

	sendResponse(
		c,
		http.StatusCreated,
		msg,
		gin.H{"id": id},
		nil,
	)
}

func SendBadRequestErrorResponse(c *gin.Context, msg string) {
	sendResponse(c, http.StatusBadRequest, "", nil, msg)
}

func SendConflictErrorResponse(c *gin.Context, msg string) {
	sendResponse(c, http.StatusConflict, "", nil, msg)
}

func SendUncaughtErrorResponse(c *gin.Context) {
	sendResponse(c, http.StatusInternalServerError, "", nil, "Internal Server Error")
}

func sendResponse(c *gin.Context, status int, msg string, data any, err any) {
	resp := gin.H{}

	if msg != "" {
		resp["message"] = msg
	}

	if data != nil {
		resp["data"] = data
	}

	if err != nil {
		resp["error"] = err
	}

	c.JSON(status, resp)
}
