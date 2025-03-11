package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JSONRequest struct {
	JSON string `json:"json"`
}

type JSONResponse struct {
	Formatted string `json:"formatted,omitempty"`
	Error     string `json:"error,omitempty"`
}

func FormatJSON(c *gin.Context) {
	var req JSONRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, JSONResponse{
			Error: "Invalid request body",
		})
		return
	}

	formatted, err := PrettyJSON([]byte(req.JSON))
	if err != nil {
		c.JSON(http.StatusBadRequest, JSONResponse{
			Error: "Invalid JSON input",
		})
		return
	}

	c.JSON(http.StatusOK, JSONResponse{
		Formatted: string(formatted),
	})
}

func PrettyJSON(input []byte) ([]byte, error) {
	var temp interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, err
	}
	return json.MarshalIndent(temp, "", "  ")
} 