package rest

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SendOKResponse sends a 200 code response.
//
// Parameters:
//   - *gin.Context: gin context is used to manage the response with the server.
//   - string: a message to indicates what happened.
//   - any: any relevant data for the consumer, will ne ommited if empty.
//
// Example:
//
//	{
//		"message": "example",
//		"data": {
//			"user_id": "123",
//			"organization": "abc"
//		}
//	}
func SendOKResponse(c *gin.Context, msg string, data any) {
	sendResponse(c, http.StatusOK, msg, data, nil)
}

// SendCreatedResponse sends a 201 code response.
//
// Parameters:
//   - *gin.Context: gin context is used to manage the response with the server.
//   - string: a message to return to the client, if is empty, will create a default creation message: "id created successfully".
//   - string: created resource id.
//   - string: the resource created (e.g.: "device", "user").
//
// Example:
//
//	{
//		"message": "id created successfully",
//		"data": {
//			"id": "123"
//		}
//	}
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

// SendBadRequestErrorResponse sends a 400 code response.
//
// Mostly used for payload validation.
//
// Parameters:
//   - *gin.Context: gin context is used to manage the response with the server.
//   - string: any error message.
//
// Example:
//
//	{
//		"error": "invalid email format",
//	}
func SendBadRequestErrorResponse(c *gin.Context, msg string) {
	sendResponse(c, http.StatusBadRequest, "", nil, msg)
}

// SendConflictErrorResponse sends a 409 code response.
//
// Parameters:
//   - *gin.Context: gin context is used to manage the response with the server.
//   - string: any error message.
//
// Example:
//
//	{
//		"error": "device already exists",
//	}
func SendConflictErrorResponse(c *gin.Context, msg string) {
	sendResponse(c, http.StatusConflict, "", nil, msg)
}

// SendUncaughtErrorResponse sends a 500 code response. Error memssage is alawys: "Internal Server Error".
//
// Mostly used for when error is not validdated.
//
// Parameters:
//   - *gin.Context: gin context is used to manage the response with the server.
//   - string: any error message.
//
// Example:
//
//	{
//		"error": "Internal  Server Error",
//	}
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
