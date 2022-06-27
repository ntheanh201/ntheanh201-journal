package v1

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type response struct {
	Error string `json:"error"`
}

func errorResponse(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, response{msg})
}

type APIError struct {
	Object  string `json:"object"`
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (err *APIError) Error() string {
	return fmt.Sprintf("%v (code: %v, status: %v)", err.Message, err.Code, err.Status)
}

func parseErrorResponse(res *http.Response) error {
	var apiErr APIError

	err := json.NewDecoder(res.Body).Decode(&apiErr)
	if err != nil {
		return fmt.Errorf("failed to parse error from HTTP response: %w", err)
	}

	return &apiErr
}
